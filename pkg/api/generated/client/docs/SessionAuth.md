# SessionAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oidc** | Pointer to [**SessionAuthOidc**](SessionAuthOidc.md) |  | [optional] 

## Methods

### NewSessionAuth

`func NewSessionAuth() *SessionAuth`

NewSessionAuth instantiates a new SessionAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAuthWithDefaults

`func NewSessionAuthWithDefaults() *SessionAuth`

NewSessionAuthWithDefaults instantiates a new SessionAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidc

`func (o *SessionAuth) GetOidc() SessionAuthOidc`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *SessionAuth) GetOidcOk() (*SessionAuthOidc, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *SessionAuth) SetOidc(v SessionAuthOidc)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *SessionAuth) HasOidc() bool`

HasOidc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


