package clientcli

import (
	"flag"
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/ankeesler/spirits/pkg/webhook"
)

func runSpirit(c *config) error {
	var (
		name                          string
		health, power, armor, agility int64
		help                          bool
	)
	flags := flag.NewFlagSet(c.command, flag.ContinueOnError)
	flags.StringVar(&name, "name", "", "spirit name; if left empty, it will be generated")
	flags.Int64Var(&health, "health", 1, "spirit health")
	flags.Int64Var(&power, "power", 0, "spirit power")
	flags.Int64Var(&armor, "armor", 0, "spirit armor")
	flags.Int64Var(&agility, "agility", 0, "spirit agility")
	flags.BoolVar(&help, "help", false, "print usage")
	flags.Parse(c.args)

	if help {
		flags.Usage()
		return nil
	}

	spirit := spiritsv1alpha1.Spirit{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:    namespace,
			GenerateName: "spirit-",
			Annotations: map[string]string{
				webhook.GenerateSpiritAnnotation: "true",
			},
			Labels: map[string]string{
				createdBySpiritsClientLabelKey: createdBySpiritsClientLabelValue,
			},
		},
	}
	if len(name) > 0 {
		spirit.GenerateName = ""
		spirit.Name = name
	}

	if err := createOrPatch(c, &spirit, func() error {
		if health != 0 {
			spirit.Spec.Stats.Health = health
		}
		if power != 0 {
			spirit.Spec.Stats.Power = power
		}
		if armor != 0 {
			spirit.Spec.Stats.Armor = armor
		}
		if agility != 0 {
			spirit.Spec.Stats.Agility = agility
		}
		if !reflect.DeepEqual(spirit.Spec.Stats, reflect.Zero(reflect.TypeOf(spirit.Spec.Stats))) {
			delete(spirit.Annotations, webhook.GenerateSpiritAnnotation)
		}
		return nil
	}); err != nil {
		return err
	}

	printSpirit(c, &spirit, "")

	return nil
}

func printSpirit(c *config, spirit *spiritsv1alpha1.Spirit, indent string) {
	fmt.Fprintf(c.out, "%skind: %s\n", indent, reflect.ValueOf(spirit).Type().String())
	fmt.Fprintf(c.out, "%sname: %s\n", indent, spirit.Name)
	fmt.Fprintf(c.out, "%sstats\n", indent)
	fmt.Fprintf(c.out, "%s  health: %d\n", indent, spirit.Spec.Stats.Health)
	fmt.Fprintf(c.out, "%s  power: %d\n", indent, spirit.Spec.Stats.Power)
	fmt.Fprintf(c.out, "%s  armor: %d\n", indent, spirit.Spec.Stats.Armor)
	fmt.Fprintf(c.out, "%s  agility: %d\n", indent, spirit.Spec.Stats.Agility)
	fmt.Fprintf(c.out, "%sactions: %s\n", indent, spirit.Spec.Actions)
	fmt.Fprintf(c.out, "%sattributes: %s\n", indent, spirit.Spec.Attributes)
	fmt.Fprintf(c.out, "%sready: %t\n", indent, meta.IsStatusConditionTrue(spirit.Status.Conditions, "Ready"))
}
