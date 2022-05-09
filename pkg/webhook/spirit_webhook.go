package webhook

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	"github.com/ankeesler/spirits/internal/generate"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
)

const (
	GenerateSpiritAnnotation     = "spirits.ankeesler.github.io/generate"
	GenerateSeedSpiritAnnotation = "spirits.ankeesler.github.io/generate-seed"

	GeneratedNicknameSpiritAnnotation = "spirits.ankeesler.github.io/generated-nickname"
)

// SpiritWebhook handles Spirit object requests
type SpiritWebhook struct {
	Scheme *runtime.Scheme
}

// SetupWithManager sets up the webhook with the Manager.
func (r *SpiritWebhook) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&spiritsv1alpha1.Spirit{}).
		WithDefaulter(r).
		Complete()
}

func (w *SpiritWebhook) Default(ctx context.Context, obj runtime.Object) error {
	log := log.FromContext(ctx)

	spirit, ok := obj.(*spiritsv1alpha1.Spirit)
	if !ok {
		return fmt.Errorf("expected object to be %T, got %T", &spiritsv1alpha1.Spirit{}, obj)
	}

	// Make sure health is non-zero
	if spirit.Spec.Attributes.Stats.Health == 0 {
		spirit.Spec.Attributes.Stats.Health = 1
	}

	if _, generate := spirit.Annotations[GenerateSpiritAnnotation]; !generate {
		log.V(1).Info("spirit %q will not be generated", client.ObjectKeyFromObject(spirit).String)
		return nil
	}

	seed := int(time.Now().Unix())
	if seedString, ok := spirit.Annotations[GenerateSeedSpiritAnnotation]; ok {
		var err error
		seed, err = strconv.Atoi(seedString)
		if err != nil {
			return fmt.Errorf("convert seed string to int: %w", err)
		}
	}

	r := rand.New(rand.NewSource(int64(seed)))
	if err := w.withInternalSpirit(spirit, func(internalSpirit *spiritsinternal.Spirit) {
		spirit.Annotations[GeneratedNicknameSpiritAnnotation] = generate.Spirit(r, internalSpirit)
	}); err != nil {
		return fmt.Errorf("with internal spirit: %w", err)
	}

	return nil
}

func (w *SpiritWebhook) withInternalSpirit(spirit *spiritsv1alpha1.Spirit, doFunc func(*spiritsinternal.Spirit)) error {
	var internalSpirit spiritsinternal.Spirit
	if err := w.Scheme.Convert(spirit, internalSpirit, nil); err != nil {
		return fmt.Errorf("convert external spirit to internal spirit: %w", err)
	}

	doFunc(&internalSpirit)

	if err := w.Scheme.Convert(internalSpirit, spirit, nil); err != nil {
		return fmt.Errorf("convert internal spirit to external spirit: %w", err)
	}

	return nil
}
