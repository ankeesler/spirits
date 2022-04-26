# ActionsApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionBattleSpiritActions**](ActionsApi.md#createSessionBattleSpiritActions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 



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

