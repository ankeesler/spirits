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
	"context"
	"errors"
	"net/http"
)

// SessionBattleSpiritActionsApiService is a service that implements the logic for the SessionBattleSpiritActionsApiServicer
// This service should implement the business logic for every endpoint for the SessionBattleSpiritActionsApi API.
// Include any external packages or services that will be required by this service.
type SessionBattleSpiritActionsApiService struct {
}

// NewSessionBattleSpiritActionsApiService creates a default api service
func NewSessionBattleSpiritActionsApiService() SessionBattleSpiritActionsApiServicer {
	return &SessionBattleSpiritActionsApiService{}
}

// CreateSessionBattleSpiritActions -
func (s *SessionBattleSpiritActionsApiService) CreateSessionBattleSpiritActions(ctx context.Context, sessionName string, battleName string, spiritName string, action Action) (ImplResponse, error) {
	// TODO - update CreateSessionBattleSpiritActions with the required logic for this service method.
	// Add api_session_battle_spirit_actions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Action{}) or use other options such as http.Ok ...
	//return Response(201, Action{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateSessionBattleSpiritActions method not implemented")
}
