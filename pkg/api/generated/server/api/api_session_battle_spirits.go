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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// SessionBattleSpiritsApiController binds http requests to an api service and writes the service results to the http response
type SessionBattleSpiritsApiController struct {
	service      SessionBattleSpiritsApiServicer
	errorHandler ErrorHandler
}

// SessionBattleSpiritsApiOption for how the controller is set up.
type SessionBattleSpiritsApiOption func(*SessionBattleSpiritsApiController)

// WithSessionBattleSpiritsApiErrorHandler inject ErrorHandler into controller
func WithSessionBattleSpiritsApiErrorHandler(h ErrorHandler) SessionBattleSpiritsApiOption {
	return func(c *SessionBattleSpiritsApiController) {
		c.errorHandler = h
	}
}

// NewSessionBattleSpiritsApiController creates a default api controller
func NewSessionBattleSpiritsApiController(s SessionBattleSpiritsApiServicer, opts ...SessionBattleSpiritsApiOption) Router {
	controller := &SessionBattleSpiritsApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SessionBattleSpiritsApiController
func (c *SessionBattleSpiritsApiController) Routes() Routes {
	return Routes{
		{
			"GetSessionBattleSpirits",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}",
			c.GetSessionBattleSpirits,
		},
		{
			"ListSessionBattleSpirits",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/battles/{battleName}/spirits",
			c.ListSessionBattleSpirits,
		},
	}
}

// GetSessionBattleSpirits -
func (c *SessionBattleSpiritsApiController) GetSessionBattleSpirits(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	battleNameParam := params["battleName"]

	spiritNameParam := params["spiritName"]

	result, err := c.service.GetSessionBattleSpirits(r.Context(), sessionNameParam, battleNameParam, spiritNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListSessionBattleSpirits -
func (c *SessionBattleSpiritsApiController) ListSessionBattleSpirits(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	battleNameParam := params["battleName"]

	result, err := c.service.ListSessionBattleSpirits(r.Context(), sessionNameParam, battleNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
