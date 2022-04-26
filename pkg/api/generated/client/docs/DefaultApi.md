# \DefaultApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RootGet**](DefaultApi.md#RootGet) | **Get** / | 
[**SessionsGet**](DefaultApi.md#SessionsGet) | **Get** /sessions | 
[**SessionsPost**](DefaultApi.md#SessionsPost) | **Post** /sessions | 
[**SessionsSessionNameBattlesBattleNameDelete**](DefaultApi.md#SessionsSessionNameBattlesBattleNameDelete) | **Delete** /sessions/{sessionName}/battles/{battleName} | 
[**SessionsSessionNameBattlesBattleNameGet**](DefaultApi.md#SessionsSessionNameBattlesBattleNameGet) | **Get** /sessions/{sessionName}/battles/{battleName} | 
[**SessionsSessionNameBattlesBattleNameSpiritsGet**](DefaultApi.md#SessionsSessionNameBattlesBattleNameSpiritsGet) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits | 
[**SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost**](DefaultApi.md#SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost) | **Post** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
[**SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet**](DefaultApi.md#SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet) | **Get** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
[**SessionsSessionNameBattlesGet**](DefaultApi.md#SessionsSessionNameBattlesGet) | **Get** /sessions/{sessionName}/battles | 
[**SessionsSessionNameBattlesPost**](DefaultApi.md#SessionsSessionNameBattlesPost) | **Post** /sessions/{sessionName}/battles | 
[**SessionsSessionNameDelete**](DefaultApi.md#SessionsSessionNameDelete) | **Delete** /sessions/{sessionName} | 
[**SessionsSessionNameGet**](DefaultApi.md#SessionsSessionNameGet) | **Get** /sessions/{sessionName} | 
[**SessionsSessionNamePut**](DefaultApi.md#SessionsSessionNamePut) | **Put** /sessions/{sessionName} | 
[**SessionsSessionNameTeamsGet**](DefaultApi.md#SessionsSessionNameTeamsGet) | **Get** /sessions/{sessionName}/teams | 
[**SessionsSessionNameTeamsPost**](DefaultApi.md#SessionsSessionNameTeamsPost) | **Post** /sessions/{sessionName}/teams | 
[**SessionsSessionNameTeamsTeamNameDelete**](DefaultApi.md#SessionsSessionNameTeamsTeamNameDelete) | **Delete** /sessions/{sessionName}/teams/{teamName} | 
[**SessionsSessionNameTeamsTeamNameGet**](DefaultApi.md#SessionsSessionNameTeamsTeamNameGet) | **Get** /sessions/{sessionName}/teams/{teamName} | 
[**SessionsSessionNameTeamsTeamNamePut**](DefaultApi.md#SessionsSessionNameTeamsTeamNamePut) | **Put** /sessions/{sessionName}/teams/{teamName} | 
[**SessionsSessionNameTeamsTeamNameSpiritsGet**](DefaultApi.md#SessionsSessionNameTeamsTeamNameSpiritsGet) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**SessionsSessionNameTeamsTeamNameSpiritsPost**](DefaultApi.md#SessionsSessionNameTeamsTeamNameSpiritsPost) | **Post** /sessions/{sessionName}/teams/{teamName}/spirits | 
[**SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete**](DefaultApi.md#SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete) | **Delete** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet**](DefaultApi.md#SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet) | **Get** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
[**SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut**](DefaultApi.md#SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut) | **Put** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 



## RootGet

> map[string]interface{} RootGet(ctx).Execute()





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
    resp, r, err := apiClient.DefaultApi.RootGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RootGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RootGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsGet

> Session SessionsGet(ctx).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsGet`: Session
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsGetRequest struct via the builder pattern


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


## SessionsPost

> Session SessionsPost(ctx).Session(session).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsPost(context.Background()).Session(session).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsPost`: Session
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionsPostRequest struct via the builder pattern


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


## SessionsSessionNameBattlesBattleNameDelete

> Battle SessionsSessionNameBattlesBattleNameDelete(ctx, sessionName, battleName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameBattlesBattleNameDelete(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameBattlesBattleNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameBattlesBattleNameDelete`: Battle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameBattlesBattleNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 
**battleName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameBattlesBattleNameDeleteRequest struct via the builder pattern


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


## SessionsSessionNameBattlesBattleNameGet

> Battle SessionsSessionNameBattlesBattleNameGet(ctx, sessionName, battleName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameBattlesBattleNameGet(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameBattlesBattleNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameBattlesBattleNameGet`: Battle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameBattlesBattleNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 
**battleName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameBattlesBattleNameGetRequest struct via the builder pattern


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


## SessionsSessionNameBattlesBattleNameSpiritsGet

> Spirit SessionsSessionNameBattlesBattleNameSpiritsGet(ctx, sessionName, battleName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsGet(context.Background(), sessionName, battleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameBattlesBattleNameSpiritsGet`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**battleName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameBattlesBattleNameSpiritsGetRequest struct via the builder pattern


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


## SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost

> Action SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost(ctx, sessionName, battleName, spiritName).Action(action).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost(context.Background(), sessionName, battleName, spiritName).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost`: Action
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSessionsSessionNameBattlesBattleNameSpiritsSpiritNameActionsPostRequest struct via the builder pattern


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


## SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet

> Spirit SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet(ctx, sessionName, battleName, spiritName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet(context.Background(), sessionName, battleName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameBattlesBattleNameSpiritsSpiritNameGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSessionsSessionNameBattlesBattleNameSpiritsSpiritNameGetRequest struct via the builder pattern


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


## SessionsSessionNameBattlesGet

> Battle SessionsSessionNameBattlesGet(ctx, sessionName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameBattlesGet(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameBattlesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameBattlesGet`: Battle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameBattlesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameBattlesGetRequest struct via the builder pattern


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


## SessionsSessionNameBattlesPost

> Battle SessionsSessionNameBattlesPost(ctx, sessionName).Battle(battle).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameBattlesPost(context.Background(), sessionName).Battle(battle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameBattlesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameBattlesPost`: Battle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameBattlesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Battle name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameBattlesPostRequest struct via the builder pattern


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


## SessionsSessionNameDelete

> Session SessionsSessionNameDelete(ctx, sessionName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameDelete(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameDelete`: Session
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Session name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameDeleteRequest struct via the builder pattern


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


## SessionsSessionNameGet

> Session SessionsSessionNameGet(ctx, sessionName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameGet(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameGet`: Session
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Session name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameGetRequest struct via the builder pattern


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


## SessionsSessionNamePut

> Session SessionsSessionNamePut(ctx, sessionName).Session(session).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNamePut(context.Background(), sessionName).Session(session).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNamePut`: Session
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Session name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNamePutRequest struct via the builder pattern


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


## SessionsSessionNameTeamsGet

> Team SessionsSessionNameTeamsGet(ctx, sessionName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsGet(context.Background(), sessionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsGet`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsGetRequest struct via the builder pattern


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


## SessionsSessionNameTeamsPost

> Team SessionsSessionNameTeamsPost(ctx, sessionName).Team(team).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsPost(context.Background(), sessionName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsPost`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsPostRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNameDelete

> Team SessionsSessionNameTeamsTeamNameDelete(ctx, sessionName, teamName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNameDelete(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNameDelete`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNameDeleteRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNameGet

> Team SessionsSessionNameTeamsTeamNameGet(ctx, sessionName, teamName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNameGet(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNameGet`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNameGetRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNamePut

> Team SessionsSessionNameTeamsTeamNamePut(ctx, sessionName, teamName).Team(team).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNamePut(context.Background(), sessionName, teamName).Team(team).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNamePut`: Team
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Team name | 
**teamName** | **string** | Team name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNamePutRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNameSpiritsGet

> Spirit SessionsSessionNameTeamsTeamNameSpiritsGet(ctx, sessionName, teamName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsGet(context.Background(), sessionName, teamName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNameSpiritsGet`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNameSpiritsGetRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNameSpiritsPost

> Spirit SessionsSessionNameTeamsTeamNameSpiritsPost(ctx, sessionName, teamName).Spirit(spirit).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsPost(context.Background(), sessionName, teamName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNameSpiritsPost`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionName** | **string** | Spirit name | 
**teamName** | **string** | Spirit name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNameSpiritsPostRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete

> Spirit SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete(ctx, sessionName, teamName, spiritName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNameDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNameSpiritsSpiritNameDeleteRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet

> Spirit SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet(ctx, sessionName, teamName, spiritName).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet(context.Background(), sessionName, teamName, spiritName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNameGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNameSpiritsSpiritNameGetRequest struct via the builder pattern


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


## SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut

> Spirit SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut(ctx, sessionName, teamName, spiritName).Spirit(spirit).Execute()





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
    resp, r, err := apiClient.DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut(context.Background(), sessionName, teamName, spiritName).Spirit(spirit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut`: Spirit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionNameTeamsTeamNameSpiritsSpiritNamePut`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiSessionsSessionNameTeamsTeamNameSpiritsSpiritNamePutRequest struct via the builder pattern


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

