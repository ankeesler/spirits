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
	var err error
	for _, path := range []string{
		"spirit-bad-intelligence.yaml",
	} {
		t.Run(path, func(t *testing.T) {
			spirit := readObject(t, path).(*spiritsv1alpha1.Spirit)
			spirit, err = tc.spiritsClientset.SpiritsV1alpha1().Spirits(tc.namespace.Name).Create(ctx, spirit, metav1.CreateOptions{})
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
