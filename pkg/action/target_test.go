package action_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ankeesler/spirits/pkg/action"
)

func TestTargetString(t *testing.T) {
	tests := []struct {
		in  action.Target
		out string
	}{
		{
			in:  action.TargetAll,
			out: "all",
		},
		{
			in:  action.TargetThem,
			out: "them",
		},
		{
			in:  action.TargetRandomThem,
			out: "random-them",
		},
		{
			in:  action.TargetOneOfThem,
			out: "one-of-them",
		},
		{
			in:  action.TargetUs,
			out: "us",
		},
		{
			in:  action.Target(0x00),
			out: "???",
		},
		{
			in:  action.Target(0x00) | action.TargetPower,
			out: "???+power",
		},
		{
			in:  action.TargetAll | action.TargetPower,
			out: "all+power",
		},
		{
			in:  action.TargetAll | action.TargetPower | action.TargetAgility,
			out: "all+power+agility",
		},
		{
			in:  action.TargetMe | action.TargetPower | action.TargetAgility,
			out: "me+power+agility",
		},
		{
			in:  action.TargetThem | action.TargetPower | action.TargetArmour | action.TargetAgility,
			out: "them+power+armour+agility",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.out, func(t *testing.T) {
			require.Equal(t, test.out, test.in.String())
		})
	}
}
