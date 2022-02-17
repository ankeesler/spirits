package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ankeesler/spirits/pkg/api"
	"github.com/ankeesler/spirits/pkg/battle"
	"github.com/ankeesler/spirits/pkg/manifest"
	"github.com/ankeesler/spirits/pkg/team"
	"github.com/ankeesler/spirits/pkg/ui"
	"gopkg.in/yaml.v2"
)

func runBattle() error {
	var (
		manifestPath      string
		battleTeamsString string
		help              bool
	)

	fs := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	fs.StringVar(&manifestPath, "manifest", "", "path to manifest")
	fs.StringVar(&battleTeamsString, "teams", "", "teams to battle (e.g., 'team1,team2,team3')")
	fs.BoolVar(&help, "help", false, "print this help message")
	fs.Parse(os.Args[2:])

	if help {
		fs.Usage()
		return nil
	}

	if manifestPath == "" {
		return fmt.Errorf("must pass -manifest")
	}

	if battleTeamsString == "" {
		return fmt.Errorf("must pass -teams")
	}

	manifestFile, err := os.Open(manifestPath)
	if err != nil {
		return fmt.Errorf("cannot open manifest %q: %w", manifestPath, err)
	}
	defer manifestFile.Close()

	var m api.Manifest
	if err := yaml.NewDecoder(manifestFile).Decode(&m); err != nil {
		return fmt.Errorf("cannot decode manifest %q: %w", manifestPath, err)
	}

	teams, err := manifest.Load(&m)
	if err != nil {
		return fmt.Errorf("cannot load manifest %q: %w", manifestPath, err)
	}

	battleTeams := []*team.Team{}
	for _, battleTeamString := range strings.Split(battleTeamsString, ",") {
		battleTeam := findTeam(battleTeamString, teams)
		if battleTeam == nil {
			return fmt.Errorf("cannot find team %q in manifest %q", battleTeamString, manifestPath)
		}
		battleTeams = append(battleTeams, battleTeam)
	}

	b := battle.New(battleTeams...)
	b.Callback = ui.Text(os.Stdout)
	b.Run()

	return nil
}

func findTeam(needle string, haystack []*team.Team) *team.Team {
	for _, hay := range haystack {
		if hay.Name == needle {
			return hay
		}
	}
	return nil
}
