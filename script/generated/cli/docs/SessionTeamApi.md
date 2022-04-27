# SessionTeamApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionTeam**](SessionTeamApi.md#createSessionTeam) | **POST** /sessions/{sessionName}/teams | 
[**deleteSessionTeam**](SessionTeamApi.md#deleteSessionTeam) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
[**getSessionTeam**](SessionTeamApi.md#getSessionTeam) | **GET** /sessions/{sessionName}/teams/{teamName} | 
[**listSessionTeams**](SessionTeamApi.md#listSessionTeams) | **GET** /sessions/{sessionName}/teams | 
[**updateSessionTeam**](SessionTeamApi.md#updateSessionTeam) | **PUT** /sessions/{sessionName}/teams/{teamName} | 



## createSessionTeam



Create a Team

### Example

```bash
 createSessionTeam sessionName=value
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


## deleteSessionTeam



Watch Team

### Example

```bash
 deleteSessionTeam sessionName=value teamName=value
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


## getSessionTeam



Get Team

### Example

```bash
 getSessionTeam sessionName=value teamName=value
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


## listSessionTeams



List Teams

### Example

```bash
 listSessionTeams sessionName=value
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


## updateSessionTeam



Update Team

### Example

```bash
 updateSessionTeam sessionName=value teamName=value
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

