# \SessionTeamApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionTeam**](SessionTeamApi.md#CreateSessionTeam) | **Post** /sessions/{sessionName}/teams | 
[**DeleteSessionTeam**](SessionTeamApi.md#DeleteSessionTeam) | **Delete** /sessions/{sessionName}/teams/{teamName} | 
[**GetSessionTeam**](SessionTeamApi.md#GetSessionTeam) | **Get** /sessions/{sessionName}/teams/{teamName} | 
[**ListSessionTeams**](SessionTeamApi.md#ListSessionTeams) | **Get** /sessions/{sessionName}/teams | 
[**UpdateSessionTeam**](SessionTeamApi.md#UpdateSessionTeam) | **Put** /sessions/{sessionName}/teams/{teamName} | 



## CreateSessionTeam

> Team CreateSessionTeam(ctx, sessionName).Team(team).Execute()





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
    resp, r, err := apiClient.SessionTeamApi.CreateSessionTeam(context.Background(), sessionName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamApi.CreateSessionTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeam`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamApi.CreateSessionTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionTeamRequest struct via the builder pattern


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


## DeleteSessionTeam

> Team DeleteSessionTeam(ctx, sessionName, teamName).Execute()





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
    resp, r, err := apiClient.SessionTeamApi.DeleteSessionTeam(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamApi.DeleteSessionTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeam`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamApi.DeleteSessionTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionTeamRequest struct via the builder pattern


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


## GetSessionTeam

> Team GetSessionTeam(ctx, sessionName, teamName).Execute()





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
    resp, r, err := apiClient.SessionTeamApi.GetSessionTeam(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamApi.GetSessionTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeam`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamApi.GetSessionTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionTeamRequest struct via the builder pattern


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
    resp, r, err := apiClient.SessionTeamApi.ListSessionTeams(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamApi.ListSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamApi.ListSessionTeams`: %v\n", resp)
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


## UpdateSessionTeam

> Team UpdateSessionTeam(ctx, sessionName, teamName).Team(team).Execute()





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
    resp, r, err := apiClient.SessionTeamApi.UpdateSessionTeam(context.Background(), sessionName, teamName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionTeamApi.UpdateSessionTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeam`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionTeamApi.UpdateSessionTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionTeamRequest struct via the builder pattern


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

