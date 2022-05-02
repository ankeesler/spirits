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

// SpiritsGetter has a method to return a SpiritInterface.
// A group's client should implement this interface.
type SpiritsGetter interface {
	Spirits(namespace string) SpiritInterface
}

// SpiritInterface has methods to work with Spirit resources.
type SpiritInterface interface {
	Create(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.CreateOptions) (*v1alpha1.Spirit, error)
	Update(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.UpdateOptions) (*v1alpha1.Spirit, error)
	UpdateStatus(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.UpdateOptions) (*v1alpha1.Spirit, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Spirit, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SpiritList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Spirit, err error)
	SpiritExpansion
}

// spirits implements SpiritInterface
type spirits struct {
	client rest.Interface
	ns     string
}

// newSpirits returns a Spirits
func newSpirits(c *SpiritsV1alpha1Client, namespace string) *spirits {
	return &spirits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the spirit, and returns the corresponding spirit object, and an error if there is any.
func (c *spirits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Spirit, err error) {
	result = &v1alpha1.Spirit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spirits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Spirits that match those selectors.
func (c *spirits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SpiritList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpiritList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spirits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spirits.
func (c *spirits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("spirits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a spirit and creates it.  Returns the server's representation of the spirit, and an error, if there is any.
func (c *spirits) Create(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.CreateOptions) (result *v1alpha1.Spirit, err error) {
	result = &v1alpha1.Spirit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("spirits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spirit).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a spirit and updates it. Returns the server's representation of the spirit, and an error, if there is any.
func (c *spirits) Update(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.UpdateOptions) (result *v1alpha1.Spirit, err error) {
	result = &v1alpha1.Spirit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spirits").
		Name(spirit.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spirit).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *spirits) UpdateStatus(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.UpdateOptions) (result *v1alpha1.Spirit, err error) {
	result = &v1alpha1.Spirit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spirits").
		Name(spirit.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(spirit).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the spirit and deletes it. Returns an error if one occurs.
func (c *spirits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spirits").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spirits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spirits").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched spirit.
func (c *spirits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Spirit, err error) {
	result = &v1alpha1.Spirit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("spirits").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
