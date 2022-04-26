# \SessionsApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionBattleSpiritActions**](SessionsApi.md#CreateSessionBattleSpiritActions) | **Post** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**CreateSessionBattles**](SessionsApi.md#CreateSessionBattles) | **Post** /sessions/{sessionName}/battles | 
[**CreateSessionTeamSpirits**](SessionsApi.md#CreateSessionTeamSpirits) | **Post** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**CreateSessionTeams**](SessionsApi.md#CreateSessionTeams) | **Post** /sessions/{sessionName}/teams | 
[**CreateSessions**](SessionsApi.md#CreateSessions) | **Post** /sessions | 
[**DeleteSessionBattles**](SessionsApi.md#DeleteSessionBattles) | **Delete** /sessions/{sessionName}/battles/{battleName} | 
[**DeleteSessionTeamSpirits**](SessionsApi.md#DeleteSessionTeamSpirits) | **Delete** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**DeleteSessionTeams**](SessionsApi.md#DeleteSessionTeams) | **Delete** /sessions/{sessionName}/teams/{teamName} | 
[**DeleteSessions**](SessionsApi.md#DeleteSessions) | **Delete** /sessions/{sessionName} | 
[**GetSessionBattleSpirits**](SessionsApi.md#GetSessionBattleSpirits) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**GetSessionBattles**](SessionsApi.md#GetSessionBattles) | **Get** /sessions/{sessionName}/battles/{battleName} | 
[**GetSessionTeamSpirits**](SessionsApi.md#GetSessionTeamSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**GetSessionTeams**](SessionsApi.md#GetSessionTeams) | **Get** /sessions/{sessionName}/teams/{teamName} | 
[**GetSessions**](SessionsApi.md#GetSessions) | **Get** /sessions/{sessionName} | 
[**ListSessions**](SessionsApi.md#ListSessions) | **Get** /sessions | 
[**ListSessionsBattles**](SessionsApi.md#ListSessionsBattles) | **Get** /sessions/{sessionName}/battles | 
[**ListSessionsBattlesSpirits**](SessionsApi.md#ListSessionsBattlesSpirits) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits | 
[**ListSessionsTeams**](SessionsApi.md#ListSessionsTeams) | **Get** /sessions/{sessionName}/teams | 
[**ListSessionsTeamsSpirits**](SessionsApi.md#ListSessionsTeamsSpirits) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**UpdateSessionTeamSpirits**](SessionsApi.md#UpdateSessionTeamSpirits) | **Put** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**UpdateSessionTeams**](SessionsApi.md#UpdateSessionTeams) | **Put** /sessions/{sessionName}/teams/{teamName} | 
[**UpdateSessions**](SessionsApi.md#UpdateSessions) | **Put** /sessions/{sessionName} | 



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
    resp, r, err := apiClient.SessionsApi.CreateSessionBattleSpiritActions(context.Background(), sessionName, battleName, spiritName).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CreateSessionBattleSpiritActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattleSpiritActions`: Action
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CreateSessionBattleSpiritActions`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.CreateSessionBattles(context.Background(), sessionName).Battle(battle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CreateSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CreateSessionBattles`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.CreateSessionTeamSpirits(context.Background(), sessionName, teamName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CreateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CreateSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.CreateSessionTeams(context.Background(), sessionName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CreateSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CreateSessionTeams`: %v\n", resp)
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


## CreateSessions

> Session CreateSessions(ctx).Session(session).Execute()





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
    session := *openapiclient.NewSession("Name_example") // Session | Session to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.CreateSessions(context.Background()).Session(session).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.CreateSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessions`: Session
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.CreateSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **session** | [**Session**](Session.md) | Session to create | 

### Return type

[**Session**](Session.md)

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
    resp, r, err := apiClient.SessionsApi.DeleteSessionBattles(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.DeleteSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.DeleteSessionBattles`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.DeleteSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.DeleteSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.DeleteSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.DeleteSessionTeams(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.DeleteSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.DeleteSessionTeams`: %v\n", resp)
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


## DeleteSessions

> Session DeleteSessions(ctx, sessionName).Execute()





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
    sessionName := "sessionName_example" // string | Session name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.DeleteSessions(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.DeleteSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSessions`: Session
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.DeleteSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Session name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Session**](Session.md)

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
    resp, r, err := apiClient.SessionsApi.GetSessionBattleSpirits(context.Background(), sessionName, battleName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessionBattleSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionBattleSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessionBattleSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.GetSessionBattles(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessionBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessionBattles`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.GetSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.GetSessionTeams(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessionTeams`: %v\n", resp)
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


## GetSessions

> Session GetSessions(ctx, sessionName).Execute()





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
    sessionName := "sessionName_example" // string | Session name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.GetSessions(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.GetSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessions`: Session
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.GetSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Session name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessions

> Session ListSessions(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.ListSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.ListSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessions`: Session
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.ListSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsRequest struct via the builder pattern


### Return type

[**Session**](Session.md)

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
    resp, r, err := apiClient.SessionsApi.ListSessionsBattles(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.ListSessionsBattles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsBattles`: Battle
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.ListSessionsBattles`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.ListSessionsBattlesSpirits(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.ListSessionsBattlesSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsBattlesSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.ListSessionsBattlesSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.ListSessionsTeams(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.ListSessionsTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.ListSessionsTeams`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.ListSessionsTeamsSpirits(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.ListSessionsTeamsSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionsTeamsSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.ListSessionsTeamsSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.UpdateSessionTeamSpirits(context.Background(), sessionName, teamName, spiritName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.UpdateSessionTeamSpirits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeamSpirits`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.UpdateSessionTeamSpirits`: %v\n", resp)
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
    resp, r, err := apiClient.SessionsApi.UpdateSessionTeams(context.Background(), sessionName, teamName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.UpdateSessionTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionTeams`: Team
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.UpdateSessionTeams`: %v\n", resp)
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


## UpdateSessions

> Session UpdateSessions(ctx, sessionName).Session(session).Execute()





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
    sessionName := "sessionName_example" // string | Session name
    session := *openapiclient.NewSession("Name_example") // Session | Session to update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.UpdateSessions(context.Background(), sessionName).Session(session).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.UpdateSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessions`: Session
    fmt.Fprintf(os.Stdout, "Response from `SessionsApi.UpdateSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Session name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **session** | [**Session**](Session.md) | Session to update | 

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

