# SessionBattleSpiritsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**getSessionBattleSpirits**](SessionBattleSpiritsApi.md#getSessionBattleSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**listSessionBattleSpirits**](SessionBattleSpiritsApi.md#listSessionBattleSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 



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


## listSessionBattleSpirits



List Spirits

### Example

```bash
 listSessionBattleSpirits sessionName=value battleName=value
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

