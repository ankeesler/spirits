/*
 * spirits
 *
 * spirits is a turn-based battle royale game
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// SessionTeamsApiController binds http requests to an api service and writes the service results to the http response
type SessionTeamsApiController struct {
	service      SessionTeamsApiServicer
	errorHandler ErrorHandler
}

// SessionTeamsApiOption for how the controller is set up.
type SessionTeamsApiOption func(*SessionTeamsApiController)

// WithSessionTeamsApiErrorHandler inject ErrorHandler into controller
func WithSessionTeamsApiErrorHandler(h ErrorHandler) SessionTeamsApiOption {
	return func(c *SessionTeamsApiController) {
		c.errorHandler = h
	}
}

// NewSessionTeamsApiController creates a default api controller
func NewSessionTeamsApiController(s SessionTeamsApiServicer, opts ...SessionTeamsApiOption) Router {
	controller := &SessionTeamsApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SessionTeamsApiController
func (c *SessionTeamsApiController) Routes() Routes {
	return Routes{
		{
			"CreateSessionTeams",
			strings.ToUpper("Post"),
			"/sessions/{sessionName}/teams",
			c.CreateSessionTeams,
		},
		{
			"DeleteSessionTeams",
			strings.ToUpper("Delete"),
			"/sessions/{sessionName}/teams/{teamName}",
			c.DeleteSessionTeams,
		},
		{
			"GetSessionTeams",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/teams/{teamName}",
			c.GetSessionTeams,
		},
		{
			"ListSessionTeams",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/teams",
			c.ListSessionTeams,
		},
		{
			"UpdateSessionTeams",
			strings.ToUpper("Put"),
			"/sessions/{sessionName}/teams/{teamName}",
			c.UpdateSessionTeams,
		},
	}
}

// CreateSessionTeams -
func (c *SessionTeamsApiController) CreateSessionTeams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	teamParam := Team{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&teamParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTeamRequired(teamParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateSessionTeams(r.Context(), sessionNameParam, teamParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteSessionTeams -
func (c *SessionTeamsApiController) DeleteSessionTeams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	teamNameParam := params["teamName"]

	result, err := c.service.DeleteSessionTeams(r.Context(), sessionNameParam, teamNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSessionTeams -
func (c *SessionTeamsApiController) GetSessionTeams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	teamNameParam := params["teamName"]

	result, err := c.service.GetSessionTeams(r.Context(), sessionNameParam, teamNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListSessionTeams -
func (c *SessionTeamsApiController) ListSessionTeams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	result, err := c.service.ListSessionTeams(r.Context(), sessionNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateSessionTeams -
func (c *SessionTeamsApiController) UpdateSessionTeams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	teamNameParam := params["teamName"]

	teamParam := Team{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&teamParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTeamRequired(teamParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateSessionTeams(r.Context(), sessionNameParam, teamNameParam, teamParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
