# \SessionBattlesApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionBattles**](SessionBattlesApi.md#CreateSessionBattles) | **Post** /sessions/{sessionName}/battles | 
[**DeleteSessionBattles**](SessionBattlesApi.md#DeleteSessionBattles) | **Delete** /sessions/{sessionName}/battles/{battleName} | 
[**GetSessionBattles**](SessionBattlesApi.md#GetSessionBattles) | **Get** /sessions/{sessionName}/battles/{battleName} | 
[**ListSessionBattles**](SessionBattlesApi.md#ListSessionBattles) | **Get** /sessions/{sessionName}/battles | 



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
    resp, r, err := apiClient.SessionBattlesApi.CreateSessionBattles(context.Background(), sessionName).Battle(battle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattlesApi.CreateSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattlesApi.CreateSessionBattles`: %v\n", resp)
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
    resp, r, err := apiClient.SessionBattlesApi.DeleteSessionBattles(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattlesApi.DeleteSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattlesApi.DeleteSessionBattles`: %v\n", resp)
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
    resp, r, err := apiClient.SessionBattlesApi.GetSessionBattles(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattlesApi.GetSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattlesApi.GetSessionBattles`: %v\n", resp)
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
    resp, r, err := apiClient.SessionBattlesApi.ListSessionBattles(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionBattlesApi.ListSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionBattlesApi.ListSessionBattles`: %v\n", resp)
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

