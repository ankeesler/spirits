// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSpirits implements SpiritInterface
type FakeSpirits struct {
	Fake *FakeSpiritsV1alpha1
	ns   string
}

var spiritsResource = schema.GroupVersionResource{Group: "spirits.ankeesler.github.com", Version: "v1alpha1", Resource: "spirits"}

var spiritsKind = schema.GroupVersionKind{Group: "spirits.ankeesler.github.com", Version: "v1alpha1", Kind: "Spirit"}

// Get takes name of the spirit, and returns the corresponding spirit object, and an error if there is any.
func (c *FakeSpirits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Spirit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spiritsResource, c.ns, name), &v1alpha1.Spirit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spirit), err
}

// List takes label and field selectors, and returns the list of Spirits that match those selectors.
func (c *FakeSpirits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SpiritList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spiritsResource, spiritsKind, c.ns, opts), &v1alpha1.SpiritList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpiritList{ListMeta: obj.(*v1alpha1.SpiritList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpiritList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spirits.
func (c *FakeSpirits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spiritsResource, c.ns, opts))

}

// Create takes the representation of a spirit and creates it.  Returns the server's representation of the spirit, and an error, if there is any.
func (c *FakeSpirits) Create(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.CreateOptions) (result *v1alpha1.Spirit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spiritsResource, c.ns, spirit), &v1alpha1.Spirit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spirit), err
}

// Update takes the representation of a spirit and updates it. Returns the server's representation of the spirit, and an error, if there is any.
func (c *FakeSpirits) Update(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.UpdateOptions) (result *v1alpha1.Spirit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spiritsResource, c.ns, spirit), &v1alpha1.Spirit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spirit), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpirits) UpdateStatus(ctx context.Context, spirit *v1alpha1.Spirit, opts v1.UpdateOptions) (*v1alpha1.Spirit, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(spiritsResource, "status", c.ns, spirit), &v1alpha1.Spirit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spirit), err
}

// Delete takes name of the spirit and deletes it. Returns an error if one occurs.
func (c *FakeSpirits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(spiritsResource, c.ns, name, opts), &v1alpha1.Spirit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpirits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spiritsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpiritList{})
	return err
}

// Patch applies the patch and returns the patched spirit.
func (c *FakeSpirits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Spirit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spiritsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Spirit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spirit), err
}
