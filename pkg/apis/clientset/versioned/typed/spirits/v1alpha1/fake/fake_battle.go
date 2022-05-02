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

// FakeBattles implements BattleInterface
type FakeBattles struct {
	Fake *FakeSpiritsV1alpha1
	ns   string
}

var battlesResource = schema.GroupVersionResource{Group: "spirits.ankeesler.github.com", Version: "v1alpha1", Resource: "battles"}

var battlesKind = schema.GroupVersionKind{Group: "spirits.ankeesler.github.com", Version: "v1alpha1", Kind: "Battle"}

// Get takes name of the battle, and returns the corresponding battle object, and an error if there is any.
func (c *FakeBattles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Battle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(battlesResource, c.ns, name), &v1alpha1.Battle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Battle), err
}

// List takes label and field selectors, and returns the list of Battles that match those selectors.
func (c *FakeBattles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BattleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(battlesResource, battlesKind, c.ns, opts), &v1alpha1.BattleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BattleList{ListMeta: obj.(*v1alpha1.BattleList).ListMeta}
	for _, item := range obj.(*v1alpha1.BattleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested battles.
func (c *FakeBattles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(battlesResource, c.ns, opts))

}

// Create takes the representation of a battle and creates it.  Returns the server's representation of the battle, and an error, if there is any.
func (c *FakeBattles) Create(ctx context.Context, battle *v1alpha1.Battle, opts v1.CreateOptions) (result *v1alpha1.Battle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(battlesResource, c.ns, battle), &v1alpha1.Battle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Battle), err
}

// Update takes the representation of a battle and updates it. Returns the server's representation of the battle, and an error, if there is any.
func (c *FakeBattles) Update(ctx context.Context, battle *v1alpha1.Battle, opts v1.UpdateOptions) (result *v1alpha1.Battle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(battlesResource, c.ns, battle), &v1alpha1.Battle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Battle), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBattles) UpdateStatus(ctx context.Context, battle *v1alpha1.Battle, opts v1.UpdateOptions) (*v1alpha1.Battle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(battlesResource, "status", c.ns, battle), &v1alpha1.Battle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Battle), err
}

// Delete takes name of the battle and deletes it. Returns an error if one occurs.
func (c *FakeBattles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(battlesResource, c.ns, name, opts), &v1alpha1.Battle{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBattles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(battlesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BattleList{})
	return err
}

// Patch applies the patch and returns the patched battle.
func (c *FakeBattles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Battle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(battlesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Battle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Battle), err
}
