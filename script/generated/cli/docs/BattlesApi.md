# BattlesApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSessionBattleSpiritActions**](BattlesApi.md#createSessionBattleSpiritActions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**createSessionBattles**](BattlesApi.md#createSessionBattles) | **POST** /sessions/{sessionName}/battles | 
[**deleteSessionBattles**](BattlesApi.md#deleteSessionBattles) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
[**getSessionBattleSpirits**](BattlesApi.md#getSessionBattleSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**getSessionBattles**](BattlesApi.md#getSessionBattles) | **GET** /sessions/{sessionName}/battles/{battleName} | 
[**listSessionsBattles**](BattlesApi.md#listSessionsBattles) | **GET** /sessions/{sessionName}/battles | 
[**listSessionsBattlesSpirits**](BattlesApi.md#listSessionsBattlesSpirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 



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


## listSessionsBattles



List Battles

### Example

```bash
 listSessionsBattles sessionName=value
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

