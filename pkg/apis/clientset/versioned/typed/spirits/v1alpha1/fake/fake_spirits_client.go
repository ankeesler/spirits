// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/ankeesler/spirits/pkg/apis/clientset/versioned/typed/spirits/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSpiritsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSpiritsV1alpha1) Battles(namespace string) v1alpha1.BattleInterface {
	return &FakeBattles{c, namespace}
}

func (c *FakeSpiritsV1alpha1) Spirits(namespace string) v1alpha1.SpiritInterface {
	return &FakeSpirits{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSpiritsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
