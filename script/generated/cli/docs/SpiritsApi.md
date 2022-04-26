# SpiritsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionBattleSpiritActions**](SpiritsApi.md#createSessionBattleSpiritActions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**createSessionTeamSpirits**](SpiritsApi.md#createSessionTeamSpirits) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**deleteSessionTeamSpirits**](SpiritsApi.md#deleteSessionTeamSpirits) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**getSessionBattleSpirits**](SpiritsApi.md#getSessionBattleSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**getSessionTeamSpirits**](SpiritsApi.md#getSessionTeamSpirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**listSessionsBattlesSpirits**](SpiritsApi.md#listSessionsBattlesSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 
[**listSessionsTeamsSpirits**](SpiritsApi.md#listSessionsTeamsSpirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**updateSessionTeamSpirits**](SpiritsApi.md#updateSessionTeamSpirits) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 



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

