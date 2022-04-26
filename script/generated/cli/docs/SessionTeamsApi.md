# SessionTeamsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionTeams**](SessionTeamsApi.md#createSessionTeams) | **POST** /sessions/{sessionName}/teams | 
[**deleteSessionTeams**](SessionTeamsApi.md#deleteSessionTeams) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
[**getSessionTeams**](SessionTeamsApi.md#getSessionTeams) | **GET** /sessions/{sessionName}/teams/{teamName} | 
[**listSessionTeams**](SessionTeamsApi.md#listSessionTeams) | **GET** /sessions/{sessionName}/teams | 
[**updateSessionTeams**](SessionTeamsApi.md#updateSessionTeams) | **PUT** /sessions/{sessionName}/teams/{teamName} | 



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

