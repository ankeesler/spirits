# \SpiritsApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionBattleSpiritActions**](SpiritsApi.md#CreateSessionBattleSpiritActions) | **Post** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**CreateSessionTeamSpirits**](SpiritsApi.md#CreateSessionTeamSpirits) | **Post** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**DeleteSessionTeamSpirits**](SpiritsApi.md#DeleteSessionTeamSpirits) | **Delete** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**GetSessionBattleSpirits**](SpiritsApi.md#GetSessionBattleSpirits) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**GetSessionTeamSpirits**](SpiritsApi.md#GetSessionTeamSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**ListSessionsBattlesSpirits**](SpiritsApi.md#ListSessionsBattlesSpirits) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits | 
[**ListSessionsTeamsSpirits**](SpiritsApi.md#ListSessionsTeamsSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**UpdateSessionTeamSpirits**](SpiritsApi.md#UpdateSessionTeamSpirits) | **Put** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 



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
    resp, r, err := apiClient.SpiritsApi.CreateSessionBattleSpiritActions(context.Background(), sessionName, battleName, spiritName).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.CreateSessionBattleSpiritActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattleSpiritActions`: Action
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.CreateSessionBattleSpiritActions`: %v\n", resp)
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
    resp, r, err := apiClient.SpiritsApi.CreateSessionTeamSpirits(context.Background(), sessionName, teamName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.CreateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.CreateSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SpiritsApi.DeleteSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.DeleteSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.DeleteSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SpiritsApi.GetSessionBattleSpirits(context.Background(), sessionName, battleName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.GetSessionBattleSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionBattleSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.GetSessionBattleSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SpiritsApi.GetSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.GetSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.GetSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SpiritsApi.ListSessionsBattlesSpirits(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.ListSessionsBattlesSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsBattlesSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.ListSessionsBattlesSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SpiritsApi.ListSessionsTeamsSpirits(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.ListSessionsTeamsSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsTeamsSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.ListSessionsTeamsSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SpiritsApi.UpdateSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpiritsApi.UpdateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SpiritsApi.UpdateSessionTeamSpirits`: %v\n", resp)
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

