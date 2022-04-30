// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	scheme "github.com/ankeesler/spirits/pkg/apis/clientset/versioned/scheme"
	v1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BattlesGetter has a method to return a BattleInterface.
// A group's client should implement this interface.
type BattlesGetter interface {
	Battles(namespace string) BattleInterface
}

// BattleInterface has methods to work with Battle resources.
type BattleInterface interface {
	Create(ctx context.Context, battle *v1alpha1.Battle, opts v1.CreateOptions) (*v1alpha1.Battle, error)
	Update(ctx context.Context, battle *v1alpha1.Battle, opts v1.UpdateOptions) (*v1alpha1.Battle, error)
	UpdateStatus(ctx context.Context, battle *v1alpha1.Battle, opts v1.UpdateOptions) (*v1alpha1.Battle, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Battle, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BattleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Battle, err error)
	BattleExpansion
}

// battles implements BattleInterface
type battles struct {
	client rest.Interface
	ns     string
}

// newBattles returns a Battles
func newBattles(c *AnkeeslerV1alpha1Client, namespace string) *battles {
	return &battles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the battle, and returns the corresponding battle object, and an error if there is any.
func (c *battles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Battle, err error) {
	result = &v1alpha1.Battle{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("battles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Battles that match those selectors.
func (c *battles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BattleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BattleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("battles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested battles.
func (c *battles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("battles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a battle and creates it.  Returns the server's representation of the battle, and an error, if there is any.
func (c *battles) Create(ctx context.Context, battle *v1alpha1.Battle, opts v1.CreateOptions) (result *v1alpha1.Battle, err error) {
	result = &v1alpha1.Battle{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("battles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(battle).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a battle and updates it. Returns the server's representation of the battle, and an error, if there is any.
func (c *battles) Update(ctx context.Context, battle *v1alpha1.Battle, opts v1.UpdateOptions) (result *v1alpha1.Battle, err error) {
	result = &v1alpha1.Battle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("battles").
		Name(battle.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(battle).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *battles) UpdateStatus(ctx context.Context, battle *v1alpha1.Battle, opts v1.UpdateOptions) (result *v1alpha1.Battle, err error) {
	result = &v1alpha1.Battle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("battles").
		Name(battle.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(battle).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the battle and deletes it. Returns an error if one occurs.
func (c *battles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("battles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *battles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("battles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched battle.
func (c *battles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Battle, err error) {
	result = &v1alpha1.Battle{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("battles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}