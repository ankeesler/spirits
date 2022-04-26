# \SessionTeamSpiritsApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionTeamSpirits**](SessionTeamSpiritsApi.md#CreateSessionTeamSpirits) | **Post** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**DeleteSessionTeamSpirits**](SessionTeamSpiritsApi.md#DeleteSessionTeamSpirits) | **Delete** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**GetSessionTeamSpirits**](SessionTeamSpiritsApi.md#GetSessionTeamSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**ListSessionTeamSpirits**](SessionTeamSpiritsApi.md#ListSessionTeamSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**UpdateSessionTeamSpirits**](SessionTeamSpiritsApi.md#UpdateSessionTeamSpirits) | **Put** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 



## CreateSessionTeamSpirits

> Spirit CreateSessionTeamSpirits(ctx, sessionName, teamName).Spirit(spirit).Execute()





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
    teamName := "teamName_example" // string | Spirit name
    spirit := *openapiclient.NewSpirit("Name_example", *openapiclient.NewSpiritStats(int64(123))) // Spirit | Spirit to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamSpiritsApi.CreateSessionTeamSpirits(context.Background(), sessionName, teamName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamSpiritsApi.CreateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamSpiritsApi.CreateSessionTeamSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionTeamSpiritsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spirit** | [**Spirit**](Spirit.md) | Spirit to create | 

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSessionTeamSpirits

> Spirit DeleteSessionTeamSpirits(ctx, sessionName, teamName, spiritName).Execute()





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
    teamName := "teamName_example" // string | Spirit name
    spiritName := "spiritName_example" // string | Spirit name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamSpiritsApi.DeleteSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamSpiritsApi.DeleteSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamSpiritsApi.DeleteSessionTeamSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 
**spiritName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionTeamSpiritsRequest struct via the builder pattern


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


## GetSessionTeamSpirits

> Spirit GetSessionTeamSpirits(ctx, sessionName, teamName, spiritName).Execute()





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
    teamName := "teamName_example" // string | Spirit name
    spiritName := "spiritName_example" // string | Spirit name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamSpiritsApi.GetSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamSpiritsApi.GetSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamSpiritsApi.GetSessionTeamSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 
**spiritName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionTeamSpiritsRequest struct via the builder pattern


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


## ListSessionTeamSpirits

> Spirit ListSessionTeamSpirits(ctx, sessionName, teamName).Execute()





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
    teamName := "teamName_example" // string | Spirit name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamSpiritsApi.ListSessionTeamSpirits(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamSpiritsApi.ListSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamSpiritsApi.ListSessionTeamSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionTeamSpiritsRequest struct via the builder pattern


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


## UpdateSessionTeamSpirits

> Spirit UpdateSessionTeamSpirits(ctx, sessionName, teamName, spiritName).Spirit(spirit).Execute()





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
    teamName := "teamName_example" // string | Spirit name
    spiritName := "spiritName_example" // string | Spirit name
    spirit := *openapiclient.NewSpirit("Name_example", *openapiclient.NewSpiritStats(int64(123))) // Spirit | Spirit to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamSpiritsApi.UpdateSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamSpiritsApi.UpdateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamSpiritsApi.UpdateSessionTeamSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 
**spiritName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionTeamSpiritsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **spirit** | [**Spirit**](Spirit.md) | Spirit to update | 

### Return type

[**Spirit**](Spirit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

