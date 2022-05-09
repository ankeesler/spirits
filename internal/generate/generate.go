package generate

import (
	"fmt"
	"math/rand"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
)

func Spirit(r *rand.Rand, spirit *spiritsinternal.Spirit) string {
	action := spiritsinternal.SpiritWellKnownActionAttack
	spirit.Spec = spiritsinternal.SpiritSpec{
		Attributes: spiritsinternal.SpiritAttributes{
			Stats: spiritsinternal.SpiritStats{
				Health:  generateStat(r) * 2,
				Power:   generateStat(r) / 2,
				Agility: generateStat(r),
				Armor:   generateStat(r) / 4,
			},
		},
		Action: spiritsinternal.SpiritAction{
			WellKnown: &action,
		},
	}
	return generateName(r)
}

func generateName(r *rand.Rand) string {
	words := randomWords()
	i, j := r.Intn(len(words)), r.Intn(len(words))
	return fmt.Sprintf("%s-%s", words[i], words[j])
}

func generateStat(r *rand.Rand) int64 {
	// These constants should make it so that 99.9% of values in [2, 18].
	const (
		desiredMean   = 10
		desiredStdDev = 2
	)
	normalNumber := r.NormFloat64()
	stat := normalNumber*desiredStdDev + desiredMean
	if stat < 0 {
		stat = 1
	}
	return int64(stat)
}

func randomWords() []string {
	return []string{
		"tofu",
		"mayonaise",
		"ice",
		"bulldozer",
		"kitchen",
		"picture",
		"sweatshirt",
		"glasses",
		"dark",
		"orange",
		"red",
		"blue",
		"pizza",
		"president",
		"leaf",
		"iron",
		"uranium",
		"sleepy",
		"banana",
		"sam",
		"ringer",
		"lantern",
		"mountain",
		"dessert",
		"forest",
		"tundra",
		"space",
		"dragon",
		"orc",
		"goblin",
		"bird",
		"snake",
		"panda",
		"tuna",
		"cat",
		"monkey",
		"mystic",
		"marshmallow",
		"redwood",
		"club",
		"carpet",
		"pillow",
		"dj",
		"squash",
		"shovel",
		"bunny",
		"dog",
		"notebook",
		"pencil",
		"farmer",
		"cliff",
		"pot",
		"ditch",
		"cup",
		"black",
		"dust",
	}
}
