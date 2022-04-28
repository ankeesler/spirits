package resource

import (
	"testing"

	"github.com/ankeesler/spirits/internal/rest"
)

func TestResource(t *testing.T) {
	r := New("root", nil).
		WithSubresource(
			New("child-a", nil).WithSubresource(
				New("grandchild-0", nil),
			).WithSubresource(
				New("grandchild-1", nil),
			),
		).
		WithSubresource(
			New("child-b", nil).WithSubresource(
				New("grandchild-0", nil),
			),
		)
	rest.ServeHTTP(r).ServeHTTP(nil, nil)
}
