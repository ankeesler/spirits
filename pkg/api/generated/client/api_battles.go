/*
spirits

spirits is a turn-based battle royale game

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// BattlesApiService BattlesApi service
type BattlesApiService service

type ApiCreateSessionBattleSpiritActionsRequest struct {
	ctx context.Context
	ApiService *BattlesApiService
	sessionName string
	battleName string
	spiritName string
	action *Action
}

// Action to create
func (r ApiCreateSessionBattleSpiritActionsRequest) Action(action Action) ApiCreateSessionBattleSpiritActionsRequest {
	r.action = &action
	return r
}

func (r ApiCreateSessionBattleSpiritActionsRequest) Execute() (*Action, *http.Response, error) {
	return r.ApiService.CreateSessionBattleSpiritActionsExecute(r)
}

/*
CreateSessionBattleSpiritActions Method for CreateSessionBattleSpiritActions

Create a Action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionName Action name
 @param battleName Action name
 @param spiritName Action name
 @return ApiCreateSessionBattleSpiritActionsRequest
*/
func (a *BattlesApiService) CreateSessionBattleSpiritActions(ctx context.Context, sessionName string, battleName string, spiritName string) ApiCreateSessionBattleSpiritActionsRequest {
	return ApiCreateSessionBattleSpiritActionsRequest{
		ApiService: a,
		ctx: ctx,
		sessionName: sessionName,
		battleName: battleName,
		spiritName: spiritName,
	}
}

// Execute executes the request
//  @return Action
func (a *BattlesApiService) CreateSessionBattleSpiritActionsExecute(r ApiCreateSessionBattleSpiritActionsRequest) (*Action, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Action
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BattlesApiService.CreateSessionBattleSpiritActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", url.PathEscape(parameterToString(r.sessionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"battleName"+"}", url.PathEscape(parameterToString(r.battleName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"spiritName"+"}", url.PathEscape(parameterToString(r.spiritName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.sessionName) < 1 {
		return localVarReturnValue, nil, reportError("sessionName must have at least 1 elements")
	}
	if strlen(r.battleName) < 1 {
		return localVarReturnValue, nil, reportError("battleName must have at least 1 elements")
	}
	if strlen(r.spiritName) < 1 {
		return localVarReturnValue, nil, reportError("spiritName must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.action
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateSessionBattlesRequest struct {
	ctx context.Context
	ApiService *BattlesApiService
	sessionName string
	battle *Battle
}

// Battle to create
func (r ApiCreateSessionBattlesRequest) Battle(battle Battle) ApiCreateSessionBattlesRequest {
	r.battle = &battle
	return r
}

func (r ApiCreateSessionBattlesRequest) Execute() (*Battle, *http.Response, error) {
	return r.ApiService.CreateSessionBattlesExecute(r)
}

/*
CreateSessionBattles Method for CreateSessionBattles

Create a Battle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionName Battle name
 @return ApiCreateSessionBattlesRequest
*/
func (a *BattlesApiService) CreateSessionBattles(ctx context.Context, sessionName string) ApiCreateSessionBattlesRequest {
	return ApiCreateSessionBattlesRequest{
		ApiService: a,
		ctx: ctx,
		sessionName: sessionName,
	}
}

// Execute executes the request
//  @return Battle
func (a *BattlesApiService) CreateSessionBattlesExecute(r ApiCreateSessionBattlesRequest) (*Battle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Battle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BattlesApiService.CreateSessionBattles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionName}/battles"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", url.PathEscape(parameterToString(r.sessionName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.sessionName) < 1 {
		return localVarReturnValue, nil, reportError("sessionName must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.battle
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteSessionBattlesRequest struct {
	ctx context.Context
	ApiService *BattlesApiService
	sessionName string
	battleName string
}

func (r ApiDeleteSessionBattlesRequest) Execute() (*Battle, *http.Response, error) {
	return r.ApiService.DeleteSessionBattlesExecute(r)
}

/*
DeleteSessionBattles Method for DeleteSessionBattles

Watch Battle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionName Battle name
 @param battleName Battle name
 @return ApiDeleteSessionBattlesRequest
*/
func (a *BattlesApiService) DeleteSessionBattles(ctx context.Context, sessionName string, battleName string) ApiDeleteSessionBattlesRequest {
	return ApiDeleteSessionBattlesRequest{
		ApiService: a,
		ctx: ctx,
		sessionName: sessionName,
		battleName: battleName,
	}
}

// Execute executes the request
//  @return Battle
func (a *BattlesApiService) DeleteSessionBattlesExecute(r ApiDeleteSessionBattlesRequest) (*Battle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Battle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BattlesApiService.DeleteSessionBattles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionName}/battles/{battleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", url.PathEscape(parameterToString(r.sessionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"battleName"+"}", url.PathEscape(parameterToString(r.battleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.sessionName) < 1 {
		return localVarReturnValue, nil, reportError("sessionName must have at least 1 elements")
	}
	if strlen(r.battleName) < 1 {
		return localVarReturnValue, nil, reportError("battleName must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSessionBattleSpiritsRequest struct {
	ctx context.Context
	ApiService *BattlesApiService
	sessionName string
	battleName string
	spiritName string
}

func (r ApiGetSessionBattleSpiritsRequest) Execute() (*Spirit, *http.Response, error) {
	return r.ApiService.GetSessionBattleSpiritsExecute(r)
}

/*
GetSessionBattleSpirits Method for GetSessionBattleSpirits

Get Spirit

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionName Spirit name
 @param battleName Spirit name
 @param spiritName Spirit name
 @return ApiGetSessionBattleSpiritsRequest
*/
func (a *BattlesApiService) GetSessionBattleSpirits(ctx context.Context, sessionName string, battleName string, spiritName string) ApiGetSessionBattleSpiritsRequest {
	return ApiGetSessionBattleSpiritsRequest{
		ApiService: a,
		ctx: ctx,
		sessionName: sessionName,
		battleName: battleName,
		spiritName: spiritName,
	}
}

// Execute executes the request
//  @return Spirit
func (a *BattlesApiService) GetSessionBattleSpiritsExecute(r ApiGetSessionBattleSpiritsRequest) (*Spirit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Spirit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BattlesApiService.GetSessionBattleSpirits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", url.PathEscape(parameterToString(r.sessionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"battleName"+"}", url.PathEscape(parameterToString(r.battleName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"spiritName"+"}", url.PathEscape(parameterToString(r.spiritName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.sessionName) < 1 {
		return localVarReturnValue, nil, reportError("sessionName must have at least 1 elements")
	}
	if strlen(r.battleName) < 1 {
		return localVarReturnValue, nil, reportError("battleName must have at least 1 elements")
	}
	if strlen(r.spiritName) < 1 {
		return localVarReturnValue, nil, reportError("spiritName must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSessionBattlesRequest struct {
	ctx context.Context
	ApiService *BattlesApiService
	sessionName string
	battleName string
}

func (r ApiGetSessionBattlesRequest) Execute() (*Battle, *http.Response, error) {
	return r.ApiService.GetSessionBattlesExecute(r)
}

/*
GetSessionBattles Method for GetSessionBattles

Get Battle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionName Battle name
 @param battleName Battle name
 @return ApiGetSessionBattlesRequest
*/
func (a *BattlesApiService) GetSessionBattles(ctx context.Context, sessionName string, battleName string) ApiGetSessionBattlesRequest {
	return ApiGetSessionBattlesRequest{
		ApiService: a,
		ctx: ctx,
		sessionName: sessionName,
		battleName: battleName,
	}
}

// Execute executes the request
//  @return Battle
func (a *BattlesApiService) GetSessionBattlesExecute(r ApiGetSessionBattlesRequest) (*Battle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Battle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BattlesApiService.GetSessionBattles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionName}/battles/{battleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", url.PathEscape(parameterToString(r.sessionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"battleName"+"}", url.PathEscape(parameterToString(r.battleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.sessionName) < 1 {
		return localVarReturnValue, nil, reportError("sessionName must have at least 1 elements")
	}
	if strlen(r.battleName) < 1 {
		return localVarReturnValue, nil, reportError("battleName must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListSessionsBattlesRequest struct {
	ctx context.Context
	ApiService *BattlesApiService
	sessionName string
}

func (r ApiListSessionsBattlesRequest) Execute() (*Battle, *http.Response, error) {
	return r.ApiService.ListSessionsBattlesExecute(r)
}

/*
ListSessionsBattles Method for ListSessionsBattles

List Battles

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionName Battle name
 @return ApiListSessionsBattlesRequest
*/
func (a *BattlesApiService) ListSessionsBattles(ctx context.Context, sessionName string) ApiListSessionsBattlesRequest {
	return ApiListSessionsBattlesRequest{
		ApiService: a,
		ctx: ctx,
		sessionName: sessionName,
	}
}

// Execute executes the request
//  @return Battle
func (a *BattlesApiService) ListSessionsBattlesExecute(r ApiListSessionsBattlesRequest) (*Battle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Battle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BattlesApiService.ListSessionsBattles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionName}/battles"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", url.PathEscape(parameterToString(r.sessionName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.sessionName) < 1 {
		return localVarReturnValue, nil, reportError("sessionName must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListSessionsBattlesSpiritsRequest struct {
	ctx context.Context
	ApiService *BattlesApiService
	sessionName string
	battleName string
}

func (r ApiListSessionsBattlesSpiritsRequest) Execute() (*Spirit, *http.Response, error) {
	return r.ApiService.ListSessionsBattlesSpiritsExecute(r)
}

/*
ListSessionsBattlesSpirits Method for ListSessionsBattlesSpirits

List Spirits

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sessionName Spirit name
 @param battleName Spirit name
 @return ApiListSessionsBattlesSpiritsRequest
*/
func (a *BattlesApiService) ListSessionsBattlesSpirits(ctx context.Context, sessionName string, battleName string) ApiListSessionsBattlesSpiritsRequest {
	return ApiListSessionsBattlesSpiritsRequest{
		ApiService: a,
		ctx: ctx,
		sessionName: sessionName,
		battleName: battleName,
	}
}

// Execute executes the request
//  @return Spirit
func (a *BattlesApiService) ListSessionsBattlesSpiritsExecute(r ApiListSessionsBattlesSpiritsRequest) (*Spirit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Spirit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BattlesApiService.ListSessionsBattlesSpirits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionName}/battles/{battleName}/spirits"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", url.PathEscape(parameterToString(r.sessionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"battleName"+"}", url.PathEscape(parameterToString(r.battleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.sessionName) < 1 {
		return localVarReturnValue, nil, reportError("sessionName must have at least 1 elements")
	}
	if strlen(r.battleName) < 1 {
		return localVarReturnValue, nil, reportError("battleName must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}