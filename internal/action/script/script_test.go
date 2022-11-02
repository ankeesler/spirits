package script

import (
	"context"
	"testing"

	"github.com/ankeesler/spirits0/internal/api"
	"github.com/ankeesler/spirits0/internal/spirit"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestScript(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{
			name: "attack",
			source: `
def main():
  print('source:', action.source)
  for target in action.targets:
    print('target:', target)
    target.stats.set_health(target.stats.health() - (action.source.stats.physical_power() - target.stats.physical_constitution()))

main()
`,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			action, err := Compile(test.source)
			if err != nil {
				t.Fatal("compile:", err)
			}

			s := &api.Spirit{
				Stats: &api.SpiritStats{
					Health:               3,
					PhysicalPower:        2,
					PhysicalConstitution: 1,
				},
			}
			source := &spirit.Spirit{API: proto.Clone(s).(*api.Spirit)}
			target := &spirit.Spirit{API: proto.Clone(s).(*api.Spirit)}
			targets := []*spirit.Spirit{target}
			if err := action.Run(context.Background(), source, targets); err != nil {
				t.Fatal("action run:", err)
			}

			wantSource := &spirit.Spirit{API: proto.Clone(s).(*api.Spirit)}
			wantTarget := &spirit.Spirit{API: proto.Clone(s).(*api.Spirit)}
			wantTarget.API.Stats.Health -= 1
			wantTargets := []*spirit.Spirit{wantTarget}
			if diff := cmp.Diff(source, wantSource, protocmp.Transform()); len(diff) > 0 {
				t.Errorf("source mismatch: -got, +want:\n%s", diff)
			}
			if diff := cmp.Diff(targets, wantTargets, protocmp.Transform()); len(diff) > 0 {
				t.Errorf("targets mismatch: -got, +want:\n%s", diff)
			}
		})
	}
}
