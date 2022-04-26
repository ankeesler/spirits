# TeamsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionTeamSpirits**](TeamsApi.md#createSessionTeamSpirits) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**createSessionTeams**](TeamsApi.md#createSessionTeams) | **POST** /sessions/{sessionName}/teams | 
[**deleteSessionTeamSpirits**](TeamsApi.md#deleteSessionTeamSpirits) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**deleteSessionTeams**](TeamsApi.md#deleteSessionTeams) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
[**getSessionTeamSpirits**](TeamsApi.md#getSessionTeamSpirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**getSessionTeams**](TeamsApi.md#getSessionTeams) | **GET** /sessions/{sessionName}/teams/{teamName} | 
[**listSessionsTeams**](TeamsApi.md#listSessionsTeams) | **GET** /sessions/{sessionName}/teams | 
[**listSessionsTeamsSpirits**](TeamsApi.md#listSessionsTeamsSpirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**updateSessionTeamSpirits**](TeamsApi.md#updateSessionTeamSpirits) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**updateSessionTeams**](TeamsApi.md#updateSessionTeams) | **PUT** /sessions/{sessionName}/teams/{teamName} | 



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

