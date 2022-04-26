# \SessionTeamsApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionTeams**](SessionTeamsApi.md#CreateSessionTeams) | **Post** /sessions/{sessionName}/teams | 
[**DeleteSessionTeams**](SessionTeamsApi.md#DeleteSessionTeams) | **Delete** /sessions/{sessionName}/teams/{teamName} | 
[**GetSessionTeams**](SessionTeamsApi.md#GetSessionTeams) | **Get** /sessions/{sessionName}/teams/{teamName} | 
[**ListSessionTeams**](SessionTeamsApi.md#ListSessionTeams) | **Get** /sessions/{sessionName}/teams | 
[**UpdateSessionTeams**](SessionTeamsApi.md#UpdateSessionTeams) | **Put** /sessions/{sessionName}/teams/{teamName} | 



## CreateSessionTeams

> Team CreateSessionTeams(ctx, sessionName).Team(team).Execute()





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
    sessionName := "sessionName_example" // string | Team name
    team := *openapiclient.NewTeam("Name_example") // Team | Team to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamsApi.CreateSessionTeams(context.Background(), sessionName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamsApi.CreateSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamsApi.CreateSessionTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **team** | [**Team**](Team.md) | Team to create | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSessionTeams

> Team DeleteSessionTeams(ctx, sessionName, teamName).Execute()





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
    sessionName := "sessionName_example" // string | Team name
    teamName := "teamName_example" // string | Team name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamsApi.DeleteSessionTeams(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamsApi.DeleteSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamsApi.DeleteSessionTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionTeams

> Team GetSessionTeams(ctx, sessionName, teamName).Execute()





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
    sessionName := "sessionName_example" // string | Team name
    teamName := "teamName_example" // string | Team name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamsApi.GetSessionTeams(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamsApi.GetSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamsApi.GetSessionTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionTeams

> Team ListSessionTeams(ctx, sessionName).Execute()





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
    sessionName := "sessionName_example" // string | Team name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamsApi.ListSessionTeams(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamsApi.ListSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamsApi.ListSessionTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSessionTeams

> Team UpdateSessionTeams(ctx, sessionName, teamName).Team(team).Execute()





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
    sessionName := "sessionName_example" // string | Team name
    teamName := "teamName_example" // string | Team name
    team := *openapiclient.NewTeam("Name_example") // Team | Team to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionTeamsApi.UpdateSessionTeams(context.Background(), sessionName, teamName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamsApi.UpdateSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamsApi.UpdateSessionTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **team** | [**Team**](Team.md) | Team to update | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

