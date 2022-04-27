# SessionBattleSpiritApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**getSessionBattleSpirit**](SessionBattleSpiritApi.md#getSessionBattleSpirit) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**listSessionBattleSpirits**](SessionBattleSpiritApi.md#listSessionBattleSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 



## getSessionBattleSpirit



Get Spirit

### Example

```bash
 getSessionBattleSpirit sessionName=value battleName=value spiritName=value
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

