# \BattlesApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionBattleSpiritActions**](BattlesApi.md#CreateSessionBattleSpiritActions) | **Post** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**CreateSessionBattles**](BattlesApi.md#CreateSessionBattles) | **Post** /sessions/{sessionName}/battles | 
[**DeleteSessionBattles**](BattlesApi.md#DeleteSessionBattles) | **Delete** /sessions/{sessionName}/battles/{battleName} | 
[**GetSessionBattleSpirits**](BattlesApi.md#GetSessionBattleSpirits) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**GetSessionBattles**](BattlesApi.md#GetSessionBattles) | **Get** /sessions/{sessionName}/battles/{battleName} | 
[**ListSessionsBattles**](BattlesApi.md#ListSessionsBattles) | **Get** /sessions/{sessionName}/battles | 
[**ListSessionsBattlesSpirits**](BattlesApi.md#ListSessionsBattlesSpirits) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits | 



## CreateSessionBattleSpiritActions

> Action CreateSessionBattleSpiritActions(ctx, sessionName, battleName, spiritName).Action(action).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionName := "sessionName_example" // string | Action name
    battleName := "battleName_example" // string | Action name
    spiritName := "spiritName_example" // string | Action name
    action := *openapiclient.NewAction("Name_example") // Action | Action to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BattlesApi.CreateSessionBattleSpiritActions(context.Background(), sessionName, battleName, spiritName).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BattlesApi.CreateSessionBattleSpiritActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattleSpiritActions`: Action
    fmt.Fprintf(os.Stdout, "Response from `BattlesApi.CreateSessionBattleSpiritActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Action name | 
**battleName** | **string** | Action name | 
**spiritName** | **string** | Action name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionBattleSpiritActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **action** | [**Action**](Action.md) | Action to create | 

### Return type

[**Action**](Action.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionBattles

> Battle CreateSessionBattles(ctx, sessionName).Battle(battle).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionName := "sessionName_example" // string | Battle name
    battle := *openapiclient.NewBattle("Name_example", []string{"Spirits_example"}) // Battle | Battle to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BattlesApi.CreateSessionBattles(context.Background(), sessionName).Battle(battle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BattlesApi.CreateSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `BattlesApi.CreateSessionBattles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionBattlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **battle** | [**Battle**](Battle.md) | Battle to create | 

### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSessionBattles

> Battle DeleteSessionBattles(ctx, sessionName, battleName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionName := "sessionName_example" // string | Battle name
    battleName := "battleName_example" // string | Battle name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BattlesApi.DeleteSessionBattles(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BattlesApi.DeleteSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `BattlesApi.DeleteSessionBattles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 
**battleName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionBattlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionBattleSpirits

> Spirit GetSessionBattleSpirits(ctx, sessionName, battleName, spiritName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionName := "sessionName_example" // string | Spirit name
    battleName := "battleName_example" // string | Spirit name
    spiritName := "spiritName_example" // string | Spirit name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BattlesApi.GetSessionBattleSpirits(context.Background(), sessionName, battleName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BattlesApi.GetSessionBattleSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionBattleSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `BattlesApi.GetSessionBattleSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**battleName** | **string** | Spirit name | 
**spiritName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionBattleSpiritsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionBattles

> Battle GetSessionBattles(ctx, sessionName, battleName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionName := "sessionName_example" // string | Battle name
    battleName := "battleName_example" // string | Battle name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BattlesApi.GetSessionBattles(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BattlesApi.GetSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `BattlesApi.GetSessionBattles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 
**battleName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionBattlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionsBattles

> Battle ListSessionsBattles(ctx, sessionName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionName := "sessionName_example" // string | Battle name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BattlesApi.ListSessionsBattles(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BattlesApi.ListSessionsBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `BattlesApi.ListSessionsBattles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsBattlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Battle**](Battle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionsBattlesSpirits

> Spirit ListSessionsBattlesSpirits(ctx, sessionName, battleName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionName := "sessionName_example" // string | Spirit name
    battleName := "battleName_example" // string | Spirit name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BattlesApi.ListSessionsBattlesSpirits(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BattlesApi.ListSessionsBattlesSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsBattlesSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `BattlesApi.ListSessionsBattlesSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**battleName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsBattlesSpiritsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

