package action

func TestScript(t *testing.T) {
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

	tests := []struct{
		name string
		codec runtime.Codec
		source string
		from, to         *spiritsinternal.Spirit
		wantFrom, wantTo *spiritsinternal.Spirit
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			a, err := Script(test.codec, test.source)
			require.NoError(t, err)
			a.Run(
		})
	}
}

func deltaHealth(spirit *spiritsinternal.Spirit, less int64) *spiritsinternal.Spirit {
	spirit = spirit.DeepCopy()
	spirit.Spec.Attributes.Stats.Health -= less
	return spirit
}
