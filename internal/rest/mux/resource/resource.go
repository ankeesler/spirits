package resource

import (
	"net/http"

	"github.com/ankeesler/spirits/internal/rest"
)

type Verb string

const (
	VerbCreate Verb = "create"
)

type Verbs map[Verb]rest.Handler

type Resource struct {
	name         string
	verbs        Verbs
	subresources map[string]*Resource
}

func New(name string, verbs Verbs) *Resource {
	return &Resource{
		name:  name,
		verbs: verbs,
	}
}

func (r *Resource) WithSubresource(sr *Resource) *Resource {
	r.subresources[sr.name] = sr
	return r
}

func (resource *Resource) Handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
