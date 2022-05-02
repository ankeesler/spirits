package clientcli

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/ankeesler/spirits/pkg/webhook"
)

func runBattle(c *config) error {
	var (
		name        string
		spiritNames string
		help        bool
	)
	flags := flag.NewFlagSet(c.command, flag.ContinueOnError)
	flags.StringVar(&name, "name", "", "spirit name; if empty, it will be generated")
	flags.StringVar(&spiritNames, "spirits", "", "spirit names (comma-separated); if empty, they will be generated")
	flags.BoolVar(&help, "help", false, "print usage")
	flags.Parse(c.args)

	if help {
		flags.Usage()
		return nil
	}

	battle := spiritsv1alpha1.Battle{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:    namespace,
			GenerateName: "battle-",
			Labels: map[string]string{
				createdBySpiritsClientLabelKey: createdBySpiritsClientLabelValue,
			},
		},
	}
	if len(name) > 0 {
		battle.GenerateName = ""
		battle.Name = name
	} else {
		var err error
		battle.Spec.Spirits, err = generateSpirits(c)
		if err != nil {
			return fmt.Errorf("generate spirits: %w", err)
		}
	}

	if err := createOrPatch(c, &battle, func() error {
		if len(spiritNames) == 0 {
		} else {
			for _, spiritName := range strings.Split(spiritNames, ",") {
				battle.Spec.Spirits = append(battle.Spec.Spirits, corev1.LocalObjectReference{Name: spiritName})
			}
		}
		return nil
	}); err != nil {
		return err
	}

	if err := printBattle(c, &battle); err != nil {
		return fmt.Errorf("print battle: %w", err)
	}

	return nil
}

func generateSpirits(c *config) ([]corev1.LocalObjectReference, error) {
	var spiritRefs []corev1.LocalObjectReference
	for i := 0; i < 2; i++ {
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
		if err := c.client.Create(c.ctx, &spirit); err != nil {
			return nil, err
		}
		spiritRefs = append(spiritRefs, corev1.LocalObjectReference{Name: spirit.Name})
	}
	return spiritRefs, nil
}

func printBattle(c *config, battle *spiritsv1alpha1.Battle) error {
	fmt.Fprintln(c.out, "kind:", reflect.ValueOf(battle).Type().String())
	fmt.Fprintln(c.out, "name:", battle.Name)
	fmt.Fprintln(c.out, "spirits:")
	for _, spiritRef := range battle.Spec.Spirits {
		if err := printSpiritFromRef(c, spiritRef); err != nil {
			return err
		}
	}
	fmt.Fprintln(c.out, "progressing:", meta.IsStatusConditionTrue(battle.Status.Conditions, "Progressing"))
	fmt.Fprintln(c.out, "phase:", battle.Status.Phase)
	switch battle.Status.Phase {
	case spiritsv1alpha1.BattlePhaseError:
		fmt.Fprintln(c.out, "message:", battle.Status.Message)
	case spiritsv1alpha1.BattlePhaseAwaitingAction:
		fmt.Fprintln(c.out, "acting:", battle.Status.ActingSpirit)
	}
	return nil
}

func printSpiritFromRef(c *config, spiritRef corev1.LocalObjectReference) error {
	spirit := spiritsv1alpha1.Spirit{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      spiritRef.Name,
		},
	}
	if err := c.client.Get(c.ctx, client.ObjectKeyFromObject(&spirit), &spirit); err != nil {
		return fmt.Errorf("get spirit: %w", err)
	}
	printSpirit(c, &spirit, "  ")
	return nil
}
