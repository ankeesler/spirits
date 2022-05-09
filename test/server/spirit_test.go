package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestInvalidSpirits(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	// Read test fixtures
	for _, test := range []struct {
		path            string
		passesAdmission bool
	}{
		{
			path:            "spirit-bad-intelligence.yaml",
			passesAdmission: true,
		},
		{
			path:            "spirit-too-many-actions.yaml",
			passesAdmission: false,
		},
	} {
		test := test
		t.Run(test.path, func(t *testing.T) {
			spirit := readObject(t, test.path).(*spiritsv1alpha1.Spirit)
			var err error
			spirit, err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Create(ctx, spirit, metav1.CreateOptions{})
			if !test.passesAdmission {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// Assert spirit is errored
			requireEventuallyConsistent(t, func() (bool, error) {
				spirit, err := tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Get(ctx, spirit.Name, metav1.GetOptions{})
				if err != nil {
					return false, fmt.Errorf("get: %w", err)
				}
				t.Logf("got spirit %q conditions: %#v", spirit.Name, spirit.Status.Conditions)
				return meta.IsStatusConditionFalse(spirit.Status.Conditions, "Ready"), nil
			})
		})
	}
}
