package main

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// Logic user has to write:
//     authN
//     authz (when should this come in the handler chain? prob right after authn :/)
//     path mux
//     method mux
//     controllers
//     serializers
//     services
//     mappers

// so we'd want to write the library that does all this
//   internal/rest                               <- generic utilities (Err)
//   internal/rest/doc                           <- given a handler tree, generates documentation using reflection
//   internal/rest/auth                          <- generic authN/authZ stuff (OIDC, etc.)
//   internal/rest/mux                           <- http mux'ing (maybe path, method, etc.)
//   internal/rest/mux/resource                  <- helper package using the parent package to setup CRUD resource muxing (Resource{Name string, Verbs Verbs, Subresources []*Resource})
//                                               <- also, dynamically build handler tree from a YAML file of Resource's
//                                               <- also also, generate openapi schema :)
//   internal/rest/controller                    <- the generic controller (including watch impl)
//   internal/rest/controller/serializer         <- the generic serializer with builder pattern (WithJSON, WithYAML, WithXML, etc.)
//   internal/rest/controller/service            <- the generic service
//   internal/rest/controller/service/mapper     <- the generic mapper (WithReflection, etc.)
//   internal/rest/controller/service/repository <- the generic repository
//
// and then the rest of spirits looks like this
//   internal/generated/api                      <- the openapi generated client code (mainly to get API types)
//   internal/mapper                             <- mapper implementation (from/to external to/from internal)
//   internal/domain                             <- domain objects
//   internal/cli                                <- cli entry point, wiring code
//   internal/server                             <- server start code
//   internal/log                                <- log utility (this should definitely go away)
//
// and we'd want to build an openapi generator that generates stuff using this
// serializers could be brought in directly from the library
// serializer options (validations) could be injected into the serializers from schema stuff
// mappers would be the responsibility of the user
// ...actually no - i worry that openapi would be such a ridiculous undertaking...
// because they deal with stuff like path queries...and we just want REST
// yeah this is gonna be...interesting...because if the openapi spec says POST, how do we know it is a create?
// maybe we could just generate from schemas...? or use the openapi-generated schema types? yeah that would be a good start

// how does the controller know what Service to get...it is always the same...
// so could the service always get the same Repository...
// could we always pass it to a Repository...?
// and then we would have a Repository of objects across parent objects. so team a and team a would both end up in the repo...
// maybe the key for the resource in the repository could be the url path...
// so we could have this

// bonus: watch handler
// New[ExternalT any, InternalT any](serializer Serializer[ExternalT, InternalT], service Service[ExternalT, InternalT], next http.Handler) http.Handler
type WatchService[T any] interface {
	Watch(context.Context, string, chan *T) error
}

type Err struct {
	Code int
	Err  error
}

type restPathContextKey struct{}

func a(r *http.Request) {
	r = r.WithContext(WithPath(r.Context(), r.URL.Path))
}

func WithPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, restPathContextKey{}, path)
}

func Path(ctx context.Context) string {
	return ctx.Value(restPathContextKey{}).(string)
}

type Repository[T any] interface {
	Create(context.Context, string, *T) (*T, error)
	Update(context.Context, string, *T) (*T, error)
	List(context.Context, string) ([]*T, error)
	Get(context.Context, string) (*T, error)
	Watch(context.Context, string, chan *T) error // Watch everything under that path
	Delete(context.Context, string) (*T, error)
}

// that would be easier for the Repository implementation...but how would the service pass the ID...
// the domain types would need to implement GetID()
// so we'd probably want to pass that to the domain type...why don't we pass it through a context to the mapper and the mapper can take care of it

// but then how would we implement GET /sessions/:id/teams...
// the default implementation service implementation would list everything from the repository, unless we filtered it down to just objects with the prefix...
// so maybe the repository needs a "list subresources" with a prefix...

// what about how would we implement GET /sessions/:id/teams/:id
// now we are talking about an etcd store :eye-roll:
// because we would want to get by prefix
// but i guess we could just make sure that the internal type always has that ID

// how would we implement watch
// we'd need some sort of injection mechanism...in the method mux we could wrap the GET handler to check for websocket before calling the next handler

// actually, why don't we just have the service pass the path as the ID so we don't have to
// so the Repository is actually a map of ID -> object
// and the ID is special because it can namespace stuff
// so like kube :eye-roll:

func main() {
	// Service logic
	sessionService := Service[ExternalSession, InternalSession]{
		Mapper: nil,
	}
	sessionController := Controller[ExternalSession, InternalSession]{
		Serializer: nil,
		Service:    &sessionService,
	}

	// Path mux
	var h http.Handler = &PathMux{
		Name: "sessions",
		Handler: MethodMux{
			http.MethodPost: http.HandlerFunc(sessionController.Create),
			http.MethodGet:  http.HandlerFunc(sessionController.List),
		},
		Children: []*PathMux{
			{
				Name:      "sessionName",
				Collecter: sessionService.CollectStore,
				Handler: MethodMux{
					http.MethodPut: authZ(
						[]string{"spirits:sessions.write"}, http.HandlerFunc(sessionController.Update),
					),
					http.MethodGet: authZ(
						[]string{"spirits:sessions.write", "spirits:sessions.read"}, http.HandlerFunc(sessionController.List),
					),
					http.MethodDelete: http.HandlerFunc(sessionController.Delete),
				},
			},
		},
	}

	h = authN(h)

	panic(http.ListenAndServe(":80", h))
}

func authN(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: authenticate r
		// TODO: respond with 401 on failure
		// TODO: store authentication material in r.Context()
		next.ServeHTTP(w, r)
	})
}

func authZ(scope []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: get authentication material in r.Context().Value(authContextKey)
		// TODO: assert scope in authentication material
		// TODO: respond with 403 on failure
		// TODO: store authorization material (scopes needed, scopes used) in r.Context()
		next.ServeHTTP(w, r)
	})
}

type ExternalSession struct{}
type InternalSession struct{}

// TODO: pass context
type Mapper[A any, B any] interface {
	AToB(*A) (*B, error)
	BToA(*B) (*A, error)
}

type Store[T any] interface {
	Create(context.Context, *T) (*T, error)
	Update(context.Context, *T) (*T, error)
	List(context.Context) ([]*T, error)
	Get(context.Context, string) (*T, error)
	Delete(context.Context, string) (*T, error)
}

type Service[ExternalT any, InternalT any] struct {
	Mapper    Mapper[ExternalT, InternalT]
	StoreFunc func() (Store[InternalT], error)
}

func (c *Service[ExternalT, Internal]) Create(ctx context.Context, externalT *ExternalT) (*ExternalT, error) {
	internalT, err := c.Mapper.AToB(externalT)
	if err != nil {
		return nil, &HTTPError{ /* TODO: bad request */ }
	}

	store, err := c.StoreFunc()
	if err != nil {
		return nil, &HTTPError{ /* TODO: internal server error */ }
	}

	internalT, err = store.Create(ctx, internalT)
	if err != nil {
		// TODO: process store error
		return nil, err
	}

	externalT, err = c.Mapper.BToA(internalT)
	if err != nil {
		return nil, err
	}

	return externalT, nil
}

func (c *Service[ExternalT, Internal]) Update(ctx context.Context, externalT *ExternalT) (*ExternalT, error) {
	// TODO: convert
	// TODO: store
	// TODO: convert
	return nil, nil
}

func (c *Service[ExternalT, InternalT]) Get(ctx context.Context) (*ExternalT, error) {
	// TODO: load
	// TODO: convert
	return nil, nil
}

func (c *Service[ExternalT, InternalT]) List() ([]*ExternalT, error) {
	// TODO: load
	// TODO: convert
	return nil, nil
}

func (c *Service[ExternalT, InternalT]) Delete() (*ExternalT, error) {
	// TODO: store
	// TODO: convert
	return nil, nil
}

var serviceStoreKey struct{}

func (c *Service[ExternalT, InternalT]) Collect(r *http.Request, segment string) *http.Request {
	// TODO: c.domain.Sessions
	return r
}

func (c *Service[ExternalT, InternalT]) store(ctx context.Context) Store[InternalT] {
	return ctx.Value(serviceStoreKey).(Store[InternalT])
}

type Serializer[T any] interface {
	Write(http.ResponseWriter, *http.Request, *T) error
	Read(http.ResponseWriter, *http.Request, *T) error
}

type HTTPError struct{}

func (e *HTTPError) Error() string { return "" }

type Controller[ExternalT, InternalT any] struct {
	Serializer Serializer[ExternalT]
	Service    *Service[ExternalT, InternalT]
}

func (c *Controller[ExternalT, InternalT]) Create(w http.ResponseWriter, r *http.Request) {
	var t ExternalT
	if err := c.Serializer.From(r.Context(), r.Body, &t); err != nil {
		// TODO: bad request
	}

	t, err = c.Service.Create(r.Context(), t)
	if err != nil {
		var httpErr HTTPError
		if errors.As(err, &httpErr) {
			// TODO: http error
		}
		// TODO: bad request
	}

	if err := c.Serializer.To(w, t); err != nil {
		// TODO: internal server error
	}
}

func (c *Controller[ExternalT, InternalT]) Update(w http.ResponseWriter, r *http.Request) {
	t, err := c.Serializer.From(r.Body)
	if err != nil {
		// TODO: bad request
	}

	t, err = c.Service.Update(r.Context(), t)
	if err != nil {
		var httpErr HTTPError
		if errors.As(err, &httpErr) {
			// TODO: http error
		}
		// TODO: bad request
	}

	if err := c.Serializer.To(w, t); err != nil {
		// TODO: internal server error
	}
}

func (c *Controller[ExternalT, InternalT]) List(w http.ResponseWriter, r *http.Request) {
	ts, err := c.Service.List()
	if err != nil {
		var httpErr HTTPError
		if errors.As(err, &httpErr) {
			// TODO: http error
		}
		// TODO: bad request
	}

	// TODO: serialize
	_ = ts
}

func (c *Controller[ExternalT, InternalT]) Get(w http.ResponseWriter, r *http.Request) {
	t, err := c.Service.Get(r.Context())
	if err != nil {
		var httpErr HTTPError
		if errors.As(err, &httpErr) {
			// TODO: http error
		}
		// TODO: bad request
	}

	if err := c.Serializer.To(w, t); err != nil {
		// TODO: internal server error
	}
}

func (c *Controller[ExternalT, InternalT]) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: store
	// TODO: convert
}

// Path mux

type PathMuxHandler interface {
	Handle(http.ResponseWriter, *http.Request, []string, http.Handler)
}

type PathMux struct {
	Name     string
	Handler  http.Handler
	Children []*PathMux
}

func (n *PathMux) AddChild(child *PathMux) error {
	// TODO: add to children, but return error if name already exists
	return nil
}

func (n *PathMux) RemoveChild(name string) error {
	// TODO: remove from children, but return error if name does not exist
	return nil
}

func (n *PathMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n.serveHTTP(w, r, strings.Split(r.URL.Path, "/"))
}

func (n *PathMux) serveHTTP(w http.ResponseWriter, r *http.Request, segments []string) {
	// If this is the last segment, we are done!
	if len(segments) == 0 {
		return
	}

	// Call n.Handler, if set
	if n.Handler != nil {
		n.Handler.ServeHTTP(w, r)
	}

	// Look for the next child
	for _, child := range n.Children {
		if child.Name == segments[0] {
			child.serveHTTP(w, r, segments[1:])
			return
		}
	}

	httpError(w, http.StatusNotFound)
}

// Method mux

type MethodMux map[string]http.Handler

func (m MethodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := m[r.Method]
	if !ok {
		httpError(w, http.StatusMethodNotAllowed)
	}
	handler.ServeHTTP(w, r)
}

func httpError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
