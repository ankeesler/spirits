# SessionTeamSpiritApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionTeamSpirit**](SessionTeamSpiritApi.md#createSessionTeamSpirit) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**deleteSessionTeamSpirit**](SessionTeamSpiritApi.md#deleteSessionTeamSpirit) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**getSessionTeamSpirit**](SessionTeamSpiritApi.md#getSessionTeamSpirit) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**listSessionTeamSpirits**](SessionTeamSpiritApi.md#listSessionTeamSpirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**updateSessionTeamSpirit**](SessionTeamSpiritApi.md#updateSessionTeamSpirit) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 



## createSessionTeamSpirit



Create a Spirit

### Example

```bash
 createSessionTeamSpirit sessionName=value teamName=value
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


## deleteSessionTeamSpirit



Watch Spirit

### Example

```bash
 deleteSessionTeamSpirit sessionName=value teamName=value spiritName=value
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


## getSessionTeamSpirit



Get Spirit

### Example

```bash
 getSessionTeamSpirit sessionName=value teamName=value spiritName=value
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


## listSessionTeamSpirits



List Spirits

### Example

```bash
 listSessionTeamSpirits sessionName=value teamName=value
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


## updateSessionTeamSpirit



Update Spirit

### Example

```bash
 updateSessionTeamSpirit sessionName=value teamName=value spiritName=value
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

