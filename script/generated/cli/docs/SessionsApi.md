# SessionsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionBattleSpiritActions**](SessionsApi.md#createSessionBattleSpiritActions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**createSessionBattles**](SessionsApi.md#createSessionBattles) | **POST** /sessions/{sessionName}/battles | 
[**createSessionTeamSpirits**](SessionsApi.md#createSessionTeamSpirits) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**createSessionTeams**](SessionsApi.md#createSessionTeams) | **POST** /sessions/{sessionName}/teams | 
[**createSessions**](SessionsApi.md#createSessions) | **POST** /sessions | 
[**deleteSessionBattles**](SessionsApi.md#deleteSessionBattles) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
[**deleteSessionTeamSpirits**](SessionsApi.md#deleteSessionTeamSpirits) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**deleteSessionTeams**](SessionsApi.md#deleteSessionTeams) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
[**deleteSessions**](SessionsApi.md#deleteSessions) | **DELETE** /sessions/{sessionName} | 
[**getSessionBattleSpirits**](SessionsApi.md#getSessionBattleSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**getSessionBattles**](SessionsApi.md#getSessionBattles) | **GET** /sessions/{sessionName}/battles/{battleName} | 
[**getSessionTeamSpirits**](SessionsApi.md#getSessionTeamSpirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**getSessionTeams**](SessionsApi.md#getSessionTeams) | **GET** /sessions/{sessionName}/teams/{teamName} | 
[**getSessions**](SessionsApi.md#getSessions) | **GET** /sessions/{sessionName} | 
[**listSessions**](SessionsApi.md#listSessions) | **GET** /sessions | 
[**listSessionsBattles**](SessionsApi.md#listSessionsBattles) | **GET** /sessions/{sessionName}/battles | 
[**listSessionsBattlesSpirits**](SessionsApi.md#listSessionsBattlesSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 
[**listSessionsTeams**](SessionsApi.md#listSessionsTeams) | **GET** /sessions/{sessionName}/teams | 
[**listSessionsTeamsSpirits**](SessionsApi.md#listSessionsTeamsSpirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**updateSessionTeamSpirits**](SessionsApi.md#updateSessionTeamSpirits) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**updateSessionTeams**](SessionsApi.md#updateSessionTeams) | **PUT** /sessions/{sessionName}/teams/{teamName} | 
[**updateSessions**](SessionsApi.md#updateSessions) | **PUT** /sessions/{sessionName} | 



## createSessionBattleSpiritActions



Create a Action

### Example

```bash
 createSessionBattleSpiritActions sessionName=value battleName=value spiritName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Action name | [default to null]
 **battleName** | **string** | Action name | [default to null]
 **spiritName** | **string** | Action name | [default to null]
 **action** | [**Action**](Action.md) | Action to create | [optional]

### Return type

[**Action**](Action.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## createSessionBattles



Create a Battle

### Example

```bash
 createSessionBattles sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Battle name | [default to null]
 **battle** | [**Battle**](Battle.md) | Battle to create | [optional]

### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## createSessionTeamSpirits



Create a Spirit

### Example

```bash
 createSessionTeamSpirits sessionName=value teamName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Spirit name | [default to null]
 **teamName** | **string** | Spirit name | [default to null]
 **spirit** | [**Spirit**](Spirit.md) | Spirit to create | [optional]

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## createSessionTeams



Create a Team

### Example

```bash
 createSessionTeams sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Team name | [default to null]
 **team** | [**Team**](Team.md) | Team to create | [optional]

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## createSessions



Create a Session

### Example

```bash
 createSessions
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session** | [**Session**](Session.md) | Session to create | [optional]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## deleteSessionBattles



Watch Battle

### Example

```bash
 deleteSessionBattles sessionName=value battleName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Battle name | [default to null]
 **battleName** | **string** | Battle name | [default to null]

### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## deleteSessionTeamSpirits



Watch Spirit

### Example

```bash
 deleteSessionTeamSpirits sessionName=value teamName=value spiritName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Spirit name | [default to null]
 **teamName** | **string** | Spirit name | [default to null]
 **spiritName** | **string** | Spirit name | [default to null]

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## deleteSessionTeams



Watch Team

### Example

```bash
 deleteSessionTeams sessionName=value teamName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Team name | [default to null]
 **teamName** | **string** | Team name | [default to null]

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## deleteSessions



Watch Session

### Example

```bash
 deleteSessions sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Session name | [default to null]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## getSessionBattleSpirits



Get Spirit

### Example

```bash
 getSessionBattleSpirits sessionName=value battleName=value spiritName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Spirit name | [default to null]
 **battleName** | **string** | Spirit name | [default to null]
 **spiritName** | **string** | Spirit name | [default to null]

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## getSessionBattles



Get Battle

### Example

```bash
 getSessionBattles sessionName=value battleName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Battle name | [default to null]
 **battleName** | **string** | Battle name | [default to null]

### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## getSessionTeamSpirits



Get Spirit

### Example

```bash
 getSessionTeamSpirits sessionName=value teamName=value spiritName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Spirit name | [default to null]
 **teamName** | **string** | Spirit name | [default to null]
 **spiritName** | **string** | Spirit name | [default to null]

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## getSessionTeams



Get Team

### Example

```bash
 getSessionTeams sessionName=value teamName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Team name | [default to null]
 **teamName** | **string** | Team name | [default to null]

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## getSessions



Get Session

### Example

```bash
 getSessions sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Session name | [default to null]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## listSessions



List Sessions

### Example

```bash
 listSessions
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## listSessionsBattles



List Battles

### Example

```bash
 listSessionsBattles sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Battle name | [default to null]

### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## listSessionsBattlesSpirits



List Spirits

### Example

```bash
 listSessionsBattlesSpirits sessionName=value battleName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Spirit name | [default to null]
 **battleName** | **string** | Spirit name | [default to null]

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## listSessionsTeams



List Teams

### Example

```bash
 listSessionsTeams sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Team name | [default to null]

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## listSessionsTeamsSpirits



List Spirits

### Example

```bash
 listSessionsTeamsSpirits sessionName=value teamName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Spirit name | [default to null]
 **teamName** | **string** | Spirit name | [default to null]

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## updateSessionTeamSpirits



Update Spirit

### Example

```bash
 updateSessionTeamSpirits sessionName=value teamName=value spiritName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Spirit name | [default to null]
 **teamName** | **string** | Spirit name | [default to null]
 **spiritName** | **string** | Spirit name | [default to null]
 **spirit** | [**Spirit**](Spirit.md) | Spirit to update | [optional]

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## updateSessionTeams



Update Team

### Example

```bash
 updateSessionTeams sessionName=value teamName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Team name | [default to null]
 **teamName** | **string** | Team name | [default to null]
 **team** | [**Team**](Team.md) | Team to update | [optional]

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## updateSessions



Update Session

### Example

```bash
 updateSessions sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Session name | [default to null]
 **session** | [**Session**](Session.md) | Session to update | [optional]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

