package controller

import (
	"context"
	"flag"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

func TestAction(t *testing.T) {
	flagset := flag.NewFlagSet("", flag.ContinueOnError)
	klog.InitFlags(flagset)
	flagset.Parse([]string{"--v=1", "--logtostderr"})

	attackAction := spiritsv1alpha1.SpiritWellKnownActionAttack
	spirit := &spiritsinternal.Spirit{
		Spec: spiritsinternal.SpiritSpec{
			Attributes: spiritsinternal.SpiritAttributes{
				Stats: spiritsinternal.SpiritStats{
					Health: 5,
					Power:  2,
					Armor:  1,
				},
			},
		},
	}
	tests := []struct {
		name                                   string
		action                                 spiritsv1alpha1.SpiritAction
		from, to                               *spiritsinternal.Spirit
		wantFrom, wantTo                       *spiritsinternal.Spirit
		wantGetActionError, wantRunActionError string
	}{
		{
			name: "wellknown",
			action: spiritsv1alpha1.SpiritAction{
				WellKnown: &attackAction,
			},
			from:     spirit,
			to:       spirit,
			wantFrom: spirit,
			wantTo:   deltaHealth(spirit, -(spirit.Spec.Attributes.Stats.Power - spirit.Spec.Attributes.Stats.Armor)),
		},
		{
			name: "good script",
			action: spiritsv1alpha1.SpiritAction{
				Script: &spiritsv1alpha1.Script{
					APIVersion: "plugin.spirits.ankeesler.github.io/v1alpha1",
					Text: `
print(apiVersion)
print(spec)
status = spec
`,
				},
			},
			from:     spirit,
			to:       spirit,
			wantFrom: deltaHealth(spirit, spirit.Spec.Attributes.Stats.Armor),
			wantTo:   deltaHealth(spirit, -spirit.Spec.Attributes.Stats.Power),
		},
		{
			name: "bad script api version",
			action: spiritsv1alpha1.SpiritAction{
				Script: &spiritsv1alpha1.Script{
					APIVersion: "plugin.spirits.ankeesler.github.io/v999alpha1",
				},
			},
			from:               spirit,
			to:                 spirit,
			wantGetActionError: `compile action script: get script predeclared symbols for compile: encode ActionRun to JSON: no kind "ActionRun" is registered for version "plugin.spirits.ankeesler.github.io/v999alpha1" in scheme "pkg/runtime/scheme.go:100"`,
		},
		{
			name: "bad script api group",
			action: spiritsv1alpha1.SpiritAction{
				Script: &spiritsv1alpha1.Script{
					APIVersion: "xxx/v1alpha1",
				},
			},
			from:               spirit,
			to:                 spirit,
			wantGetActionError: `compile action script: get script predeclared symbols for compile: encode ActionRun to JSON: plugin.ActionRun is not suitable for converting to "xxx/v1alpha1" in scheme "pkg/runtime/scheme.go:100"`,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actionChan := make(chan spiritsinternal.Action)
			action, err := getAction(&test.action, func(ctx context.Context) (spiritsinternal.Action, error) {
				return <-actionChan, nil
			}, scheme)
			if len(test.wantGetActionError) > 0 {
				require.EqualError(t, err, test.wantGetActionError)
				return
			}
			require.NoError(t, err)

			gotFrom, gotTo := test.from.DeepCopy(), test.to.DeepCopy()
			err = action.Run(context.Background(), gotFrom, gotTo)
			if len(test.wantRunActionError) > 0 {
				require.EqualError(t, err, test.wantRunActionError)
				return
			}
			require.NoError(t, err)

			require.Empty(t, cmp.Diff(test.wantFrom, gotFrom), "diff in 'from' spirit (-want, +got)")
			require.Empty(t, cmp.Diff(test.wantTo, gotTo), "diff in 'to' spirit (-want, +got)")
		})
	}
}

func deltaHealth(spirit *spiritsinternal.Spirit, delta int64) *spiritsinternal.Spirit {
	spirit = spirit.DeepCopy()
	spirit.Spec.Attributes.Stats.Health += delta
	return spirit
}
