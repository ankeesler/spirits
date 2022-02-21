package manifest_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/ankeesler/spirits/internal/action"
	"github.com/ankeesler/spirits/internal/manifest"
	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/team"
	testpkg "github.com/ankeesler/spirits/internal/test"
	api "github.com/ankeesler/spirits/pkg/v0"
)

func TestLoad(t *testing.T) {
	goodManifest := &api.Manifest{
		Data: &api.ManifestData{
			Teams: []*api.Team{
				{
					Name: "team 0",
					Spirits: []*api.Spirit{
						{Name: "spirit 0a", Health: 1, Action: &api.Action{Type: "attack", Target: 0xFF}},
						{Name: "spirit 0b", Power: 2, Action: &api.Action{Type: "heal", Target: 0x08}},
					},
				},
				{
					Name: "team 1",
					Spirits: []*api.Spirit{
						{Name: "spirit 1a", Armour: 3, Action: &api.Action{Type: "buf", Target: 0x18}},
						{Name: "spirit 1b", Agility: 4, Action: &api.Action{Type: "debuf", Target: 0x07}},
					},
				},
			},
		},
	}

	badActionTypeManifest := &api.Manifest{
		Data: &api.ManifestData{
			Teams: []*api.Team{
				{
					Name: "good team",
					Spirits: []*api.Spirit{
						{Name: "good spirit 0", Action: &api.Action{Type: "attack"}},
						{Name: "good spirit 1", Action: &api.Action{Type: "attack"}},
					},
				},
				{
					Name: "bad team",
					Spirits: []*api.Spirit{
						{Name: "bad spirit", Action: &api.Action{Type: "bad action type"}},
					},
				},
			},
		},
	}

	tests := []struct {
		name      string
		m         *api.Manifest
		wantTeams []*team.Team
		wantError string
	}{
		{
			name: "happy path",
			m:    goodManifest,
			wantTeams: []*team.Team{
				team.New(
					"team 0",
					spirit.New("spirit 0a", 1, 0, 0, 0, action.Attack(action.TargetAll)),
					spirit.New("spirit 0b", 0, 2, 0, 0, action.Attack(action.TargetMe)),
				),
				team.New(
					"team 1",
					spirit.New("spirit 1a", 0, 0, 3, 0, action.Attack(action.TargetAll)),
					spirit.New("spirit 1b", 0, 0, 0, 4, action.Attack(action.TargetMe)),
				),
			},
		},
		{
			name:      "unrecognized action type",
			m:         badActionTypeManifest,
			wantError: `cannot instantiate team "bad team": cannot instantiate spirit "bad spirit": cannot instantiate action: unrecognized action type: "bad action type"`,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			gotTeams, gotErr := manifest.Load(test.m)
			if test.wantError != "" {
				require.EqualError(t, gotErr, test.wantError)
				return
			}
			if diff := cmp.Diff(test.wantTeams, gotTeams, testpkg.Comparer()); diff != "" {
				t.Fatalf("-want, +got:\n%s", diff)
			}
		})
	}
}
