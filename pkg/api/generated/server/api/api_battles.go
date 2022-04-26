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

// BattlesApiController binds http requests to an api service and writes the service results to the http response
type BattlesApiController struct {
	service      BattlesApiServicer
	errorHandler ErrorHandler
}

// BattlesApiOption for how the controller is set up.
type BattlesApiOption func(*BattlesApiController)

// WithBattlesApiErrorHandler inject ErrorHandler into controller
func WithBattlesApiErrorHandler(h ErrorHandler) BattlesApiOption {
	return func(c *BattlesApiController) {
		c.errorHandler = h
	}
}

// NewBattlesApiController creates a default api controller
func NewBattlesApiController(s BattlesApiServicer, opts ...BattlesApiOption) Router {
	controller := &BattlesApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the BattlesApiController
func (c *BattlesApiController) Routes() Routes {
	return Routes{
		{
			"CreateSessionBattleSpiritActions",
			strings.ToUpper("Post"),
			"/sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions",
			c.CreateSessionBattleSpiritActions,
		},
		{
			"CreateSessionBattles",
			strings.ToUpper("Post"),
			"/sessions/{sessionName}/battles",
			c.CreateSessionBattles,
		},
		{
			"DeleteSessionBattles",
			strings.ToUpper("Delete"),
			"/sessions/{sessionName}/battles/{battleName}",
			c.DeleteSessionBattles,
		},
		{
			"GetSessionBattleSpirits",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}",
			c.GetSessionBattleSpirits,
		},
		{
			"GetSessionBattles",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/battles/{battleName}",
			c.GetSessionBattles,
		},
		{
			"ListSessionsBattles",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/battles",
			c.ListSessionsBattles,
		},
		{
			"ListSessionsBattlesSpirits",
			strings.ToUpper("Get"),
			"/sessions/{sessionName}/battles/{battleName}/spirits",
			c.ListSessionsBattlesSpirits,
		},
	}
}

// CreateSessionBattleSpiritActions -
func (c *BattlesApiController) CreateSessionBattleSpiritActions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	battleNameParam := params["battleName"]

	spiritNameParam := params["spiritName"]

	actionParam := Action{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&actionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertActionRequired(actionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateSessionBattleSpiritActions(r.Context(), sessionNameParam, battleNameParam, spiritNameParam, actionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateSessionBattles -
func (c *BattlesApiController) CreateSessionBattles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	battleParam := Battle{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&battleParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBattleRequired(battleParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateSessionBattles(r.Context(), sessionNameParam, battleParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteSessionBattles -
func (c *BattlesApiController) DeleteSessionBattles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	battleNameParam := params["battleName"]

	result, err := c.service.DeleteSessionBattles(r.Context(), sessionNameParam, battleNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSessionBattleSpirits -
func (c *BattlesApiController) GetSessionBattleSpirits(w http.ResponseWriter, r *http.Request) {
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

// GetSessionBattles -
func (c *BattlesApiController) GetSessionBattles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	battleNameParam := params["battleName"]

	result, err := c.service.GetSessionBattles(r.Context(), sessionNameParam, battleNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListSessionsBattles -
func (c *BattlesApiController) ListSessionsBattles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	result, err := c.service.ListSessionsBattles(r.Context(), sessionNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListSessionsBattlesSpirits -
func (c *BattlesApiController) ListSessionsBattlesSpirits(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sessionNameParam := params["sessionName"]

	battleNameParam := params["battleName"]

	result, err := c.service.ListSessionsBattlesSpirits(r.Context(), sessionNameParam, battleNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}