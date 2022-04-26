# \ActionsApi

All URIs are relative to *https://oh-great-spirits.herokuapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionBattleSpiritActions**](ActionsApi.md#CreateSessionBattleSpiritActions) | **Post** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 



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
    resp, r, err := apiClient.ActionsApi.CreateSessionBattleSpiritActions(context.Background(), sessionName, battleName, spiritName).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionsApi.CreateSessionBattleSpiritActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionBattleSpiritActions`: Action
    fmt.Fprintf(os.Stdout, "Response from `ActionsApi.CreateSessionBattleSpiritActions`: %v\n", resp)
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

