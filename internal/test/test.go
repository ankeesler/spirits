// Package test holds test utilities and integration tests for the spirits project.
package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ankeesler/spirits/internal/spirit"
	"github.com/ankeesler/spirits/internal/tbass"
	"github.com/ankeesler/spirits/internal/tbass/actor/action"
	"github.com/ankeesler/spirits/internal/tbass/team"
)

func newSpirit(t *testing.T, format string) *spirit.Spirit {
	t.Helper()

	// format: health.power.armour.speed
	// e.g., 5.1.2.1,
	split := strings.Split(format, ".")

	health, err := strconv.Atoi(split[0])
	require.NoError(t, err)

	power, err := strconv.Atoi(split[1])
	require.NoError(t, err)

	armour, err := strconv.Atoi(split[2])
	require.NoError(t, err)

	agility, err := strconv.Atoi(split[3])
	require.NoError(t, err)

	return spirit.New(
		fmt.Sprintf("s(%s)", format),
		tbass.StatValue(health),
		tbass.StatValue(power),
		tbass.StatValue(armour),
		tbass.StatValue(agility),
		action.Debuf(spirit.StatPower, spirit.StatHealth, false),
	)
}

func newTeam(t *testing.T, formats ...string) tbass.Team {
	t.Helper()

	var actors []tbass.Actor
	for _, format := range formats {
		actors = append(actors, newSpirit(t, format))
	}

	return team.New(fmt.Sprintf("t(%s)", strings.Join(formats, ",")), actors...)
}
