# DefaultApi

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**rootGet**](DefaultApi.md#rootGet) | **GET** / | 
[**sessionsGet**](DefaultApi.md#sessionsGet) | **GET** /sessions | 
[**sessionsPost**](DefaultApi.md#sessionsPost) | **POST** /sessions | 
[**sessionsSessionNameBattlesBattleNameDelete**](DefaultApi.md#sessionsSessionNameBattlesBattleNameDelete) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
[**sessionsSessionNameBattlesBattleNameGet**](DefaultApi.md#sessionsSessionNameBattlesBattleNameGet) | **GET** /sessions/{sessionName}/battles/{battleName} | 
[**sessionsSessionNameBattlesBattleNameSpiritsGet**](DefaultApi.md#sessionsSessionNameBattlesBattleNameSpiritsGet) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 
[**sessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost**](DefaultApi.md#sessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**sessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet**](DefaultApi.md#sessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**sessionsSessionNameBattlesGet**](DefaultApi.md#sessionsSessionNameBattlesGet) | **GET** /sessions/{sessionName}/battles | 
[**sessionsSessionNameBattlesPost**](DefaultApi.md#sessionsSessionNameBattlesPost) | **POST** /sessions/{sessionName}/battles | 
[**sessionsSessionNameDelete**](DefaultApi.md#sessionsSessionNameDelete) | **DELETE** /sessions/{sessionName} | 
[**sessionsSessionNameGet**](DefaultApi.md#sessionsSessionNameGet) | **GET** /sessions/{sessionName} | 
[**sessionsSessionNamePut**](DefaultApi.md#sessionsSessionNamePut) | **PUT** /sessions/{sessionName} | 
[**sessionsSessionNameTeamsGet**](DefaultApi.md#sessionsSessionNameTeamsGet) | **GET** /sessions/{sessionName}/teams | 
[**sessionsSessionNameTeamsPost**](DefaultApi.md#sessionsSessionNameTeamsPost) | **POST** /sessions/{sessionName}/teams | 
[**sessionsSessionNameTeamsTeamNameDelete**](DefaultApi.md#sessionsSessionNameTeamsTeamNameDelete) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
[**sessionsSessionNameTeamsTeamNameGet**](DefaultApi.md#sessionsSessionNameTeamsTeamNameGet) | **GET** /sessions/{sessionName}/teams/{teamName} | 
[**sessionsSessionNameTeamsTeamNamePut**](DefaultApi.md#sessionsSessionNameTeamsTeamNamePut) | **PUT** /sessions/{sessionName}/teams/{teamName} | 
[**sessionsSessionNameTeamsTeamNameSpiritsGet**](DefaultApi.md#sessionsSessionNameTeamsTeamNameSpiritsGet) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**sessionsSessionNameTeamsTeamNameSpiritsPost**](DefaultApi.md#sessionsSessionNameTeamsTeamNameSpiritsPost) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**sessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete**](DefaultApi.md#sessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**sessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet**](DefaultApi.md#sessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**sessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut**](DefaultApi.md#sessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 



## rootGet



Retrieve the OpenAPI specification currently served

### Example

```bash
 rootGet
```

### Parameters

This endpoint does not need any parameter.

### Return type

**map**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## sessionsGet



Watch Session

### Example

```bash
 sessionsGet
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## sessionsPost



Create a Session

### Example

```bash
 sessionsPost
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session** | [**Session**](Session.md) | Session to create | [optional]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## sessionsSessionNameBattlesBattleNameDelete



Watch Battle

### Example

```bash
 sessionsSessionNameBattlesBattleNameDelete sessionName=value battleName=value
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


## sessionsSessionNameBattlesBattleNameGet



Watch Battle

### Example

```bash
 sessionsSessionNameBattlesBattleNameGet sessionName=value battleName=value
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


## sessionsSessionNameBattlesBattleNameSpiritsGet



Watch Spirit

### Example

```bash
 sessionsSessionNameBattlesBattleNameSpiritsGet sessionName=value battleName=value
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


## sessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost



Create a Action

### Example

```bash
 sessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost sessionName=value battleName=value spiritName=value
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


## sessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet



Watch Spirit

### Example

```bash
 sessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet sessionName=value battleName=value spiritName=value
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


## sessionsSessionNameBattlesGet



Watch Battle

### Example

```bash
 sessionsSessionNameBattlesGet sessionName=value
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


## sessionsSessionNameBattlesPost



Create a Battle

### Example

```bash
 sessionsSessionNameBattlesPost sessionName=value
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


## sessionsSessionNameDelete



Watch Session

### Example

```bash
 sessionsSessionNameDelete sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Session name | [default to null]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## sessionsSessionNameGet



Watch Session

### Example

```bash
 sessionsSessionNameGet sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Session name | [default to null]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not Applicable
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## sessionsSessionNamePut



Update Session

### Example

```bash
 sessionsSessionNamePut sessionName=value
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionName** | **string** | Session name | [default to null]
 **session** | [**Session**](Session.md) | Session to update | [optional]

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## sessionsSessionNameTeamsGet



Watch Team

### Example

```bash
 sessionsSessionNameTeamsGet sessionName=value
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


## sessionsSessionNameTeamsPost



Create a Team

### Example

```bash
 sessionsSessionNameTeamsPost sessionName=value
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


## sessionsSessionNameTeamsTeamNameDelete



Watch Team

### Example

```bash
 sessionsSessionNameTeamsTeamNameDelete sessionName=value teamName=value
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


## sessionsSessionNameTeamsTeamNameGet



Watch Team

### Example

```bash
 sessionsSessionNameTeamsTeamNameGet sessionName=value teamName=value
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


## sessionsSessionNameTeamsTeamNamePut



Update Team

### Example

```bash
 sessionsSessionNameTeamsTeamNamePut sessionName=value teamName=value
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


## sessionsSessionNameTeamsTeamNameSpiritsGet



Watch Spirit

### Example

```bash
 sessionsSessionNameTeamsTeamNameSpiritsGet sessionName=value teamName=value
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


## sessionsSessionNameTeamsTeamNameSpiritsPost



Create a Spirit

### Example

```bash
 sessionsSessionNameTeamsTeamNameSpiritsPost sessionName=value teamName=value
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


## sessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete



Watch Spirit

### Example

```bash
 sessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete sessionName=value teamName=value spiritName=value
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


## sessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet



Watch Spirit

### Example

```bash
 sessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet sessionName=value teamName=value spiritName=value
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


## sessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut



Update Spirit

### Example

```bash
 sessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut sessionName=value teamName=value spiritName=value
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

