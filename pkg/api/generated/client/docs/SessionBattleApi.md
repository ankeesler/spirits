# \SessionBattleApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionBattle**](SessionBattleApi.md#CreateSessionBattle) | **Post** /sessions/{sessionName}/battles | 
[**DeleteSessionBattle**](SessionBattleApi.md#DeleteSessionBattle) | **Delete** /sessions/{sessionName}/battles/{battleName} | 
[**GetSessionBattle**](SessionBattleApi.md#GetSessionBattle) | **Get** /sessions/{sessionName}/battles/{battleName} | 
[**ListSessionBattles**](SessionBattleApi.md#ListSessionBattles) | **Get** /sessions/{sessionName}/battles | 



## CreateSessionBattle

> Battle CreateSessionBattle(ctx, sessionName).Battle(battle).Execute()





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
    resp, r, err := apiClient.SessionBattleApi.CreateSessionBattle(context.Background(), sessionName).Battle(battle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattleApi.CreateSessionBattle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattle`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattleApi.CreateSessionBattle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionBattleRequest struct via the builder pattern


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


## DeleteSessionBattle

> Battle DeleteSessionBattle(ctx, sessionName, battleName).Execute()





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
    resp, r, err := apiClient.SessionBattleApi.DeleteSessionBattle(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattleApi.DeleteSessionBattle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionBattle`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattleApi.DeleteSessionBattle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 
**battleName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionBattleRequest struct via the builder pattern


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


## GetSessionBattle

> Battle GetSessionBattle(ctx, sessionName, battleName).Execute()





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
    resp, r, err := apiClient.SessionBattleApi.GetSessionBattle(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattleApi.GetSessionBattle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionBattle`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattleApi.GetSessionBattle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 
**battleName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionBattleRequest struct via the builder pattern


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


## ListSessionBattles

> Battle ListSessionBattles(ctx, sessionName).Execute()





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
    resp, r, err := apiClient.SessionBattleApi.ListSessionBattles(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattleApi.ListSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattleApi.ListSessionBattles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionBattlesRequest struct via the builder pattern


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

