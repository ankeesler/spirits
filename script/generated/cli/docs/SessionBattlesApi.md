# SessionBattlesApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionBattles**](SessionBattlesApi.md#createSessionBattles) | **POST** /sessions/{sessionName}/battles | 
[**deleteSessionBattles**](SessionBattlesApi.md#deleteSessionBattles) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
[**getSessionBattles**](SessionBattlesApi.md#getSessionBattles) | **GET** /sessions/{sessionName}/battles/{battleName} | 
[**listSessionBattles**](SessionBattlesApi.md#listSessionBattles) | **GET** /sessions/{sessionName}/battles | 



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


## listSessionBattles



List Battles

### Example

```bash
 listSessionBattles sessionName=value
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

