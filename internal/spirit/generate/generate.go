package generate

import (
	"fmt"
	"math/rand"

	"github.com/ankeesler/spirits/internal/spirit"
)

func Generate(seed int64) []*spirit.Spirit {
	r := rand.New(rand.NewSource(seed))
	return []*spirit.Spirit{generate(r), generate(r)}
}

func generate(r *rand.Rand) *spirit.Spirit {
	return &spirit.Spirit{
		Name:    generateName(r),
		Health:  generateStat(r) * 2,
		Power:   generateStat(r) / 2,
		Agility: generateStat(r),
		Armour:  generateStat(r) / 4,
	}
}

func generateName(r *rand.Rand) string {
	words := randomWords()
	i, j := r.Intn(len(words)), r.Intn(len(words))
	return fmt.Sprintf("%s-%s", words[i], words[j])
}

func generateStat(r *rand.Rand) int {
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
	return int(stat)
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
	}
}
