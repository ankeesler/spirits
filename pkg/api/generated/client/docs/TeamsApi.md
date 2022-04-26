# \TeamsApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionTeamSpirits**](TeamsApi.md#CreateSessionTeamSpirits) | **Post** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**CreateSessionTeams**](TeamsApi.md#CreateSessionTeams) | **Post** /sessions/{sessionName}/teams | 
[**DeleteSessionTeamSpirits**](TeamsApi.md#DeleteSessionTeamSpirits) | **Delete** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**DeleteSessionTeams**](TeamsApi.md#DeleteSessionTeams) | **Delete** /sessions/{sessionName}/teams/{teamName} | 
[**GetSessionTeamSpirits**](TeamsApi.md#GetSessionTeamSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**GetSessionTeams**](TeamsApi.md#GetSessionTeams) | **Get** /sessions/{sessionName}/teams/{teamName} | 
[**ListSessionsTeams**](TeamsApi.md#ListSessionsTeams) | **Get** /sessions/{sessionName}/teams | 
[**ListSessionsTeamsSpirits**](TeamsApi.md#ListSessionsTeamsSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**UpdateSessionTeamSpirits**](TeamsApi.md#UpdateSessionTeamSpirits) | **Put** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**UpdateSessionTeams**](TeamsApi.md#UpdateSessionTeams) | **Put** /sessions/{sessionName}/teams/{teamName} | 



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
    resp, r, err := apiClient.TeamsApi.CreateSessionTeamSpirits(context.Background(), sessionName, teamName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.TeamsApi.CreateSessionTeams(context.Background(), sessionName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateSessionTeams`: %v\n", resp)
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
    resp, r, err := apiClient.TeamsApi.DeleteSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.DeleteSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.DeleteSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.TeamsApi.DeleteSessionTeams(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.DeleteSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.DeleteSessionTeams`: %v\n", resp)
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
    resp, r, err := apiClient.TeamsApi.GetSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.TeamsApi.GetSessionTeams(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetSessionTeams`: %v\n", resp)
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


## ListSessionsTeams

> Team ListSessionsTeams(ctx, sessionName).Execute()





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
    resp, r, err := apiClient.TeamsApi.ListSessionsTeams(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListSessionsTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListSessionsTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsTeamsRequest struct via the builder pattern


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


## ListSessionsTeamsSpirits

> Spirit ListSessionsTeamsSpirits(ctx, sessionName, teamName).Execute()





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
    resp, r, err := apiClient.TeamsApi.ListSessionsTeamsSpirits(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListSessionsTeamsSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsTeamsSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListSessionsTeamsSpirits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsTeamsSpiritsRequest struct via the builder pattern


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
    resp, r, err := apiClient.TeamsApi.UpdateSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.TeamsApi.UpdateSessionTeams(context.Background(), sessionName, teamName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateSessionTeams`: %v\n", resp)
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

