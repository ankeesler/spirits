package action

import (
	"context"
	"reflect"
	"testing"

	"github.com/ankeesler/spirits0/internal/spirit"
	"github.com/google/go-cmp/cmp"
)

var spiritExporter = cmp.Exporter(func(t reflect.Type) bool {
	spiritT := reflect.TypeOf(spirit.Spirit{})
	statsT := reflect.TypeOf(spirit.Stats{})
	return t == spiritT || t == statsT
})

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
			action, err := compile(test.source)
			if err != nil {
				t.Fatal("compile:", err)
			}

			stats := spirit.NewStats(3, 2, 1, 0, 0, 0)
			s := spirit.New("", "", stats, nil)

			source := s.Clone()
			target := s.Clone()
			targets := []*spirit.Spirit{target}
			if _, err := action.Run(context.Background(), source, targets); err != nil {
				t.Fatal("action run:", err)
			}

			wantSource := s.Clone()
			wantTarget := s.Clone()
			wantTarget.Stats().SetHealth(wantTarget.Stats().Health() - 1)
			wantTargets := []*spirit.Spirit{wantTarget}
			if diff := cmp.Diff(source, wantSource, spiritExporter); len(diff) > 0 {
				t.Errorf("source mismatch: -got, +want:\n%s", diff)
			}
			if diff := cmp.Diff(targets, wantTargets, spiritExporter); len(diff) > 0 {
				t.Errorf("targets mismatch: -got, +want:\n%s", diff)
			}
		})
	}
}
