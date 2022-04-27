# SessionBattleApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionBattle**](SessionBattleApi.md#createSessionBattle) | **POST** /sessions/{sessionName}/battles | 
[**deleteSessionBattle**](SessionBattleApi.md#deleteSessionBattle) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
[**getSessionBattle**](SessionBattleApi.md#getSessionBattle) | **GET** /sessions/{sessionName}/battles/{battleName} | 
[**listSessionBattles**](SessionBattleApi.md#listSessionBattles) | **GET** /sessions/{sessionName}/battles | 



## createSessionBattle



Create a Battle

### Example

```bash
 createSessionBattle sessionName=value
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


## deleteSessionBattle



Watch Battle

### Example

```bash
 deleteSessionBattle sessionName=value battleName=value
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


## getSessionBattle



Get Battle

### Example

```bash
 getSessionBattle sessionName=value battleName=value
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

