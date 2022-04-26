# SessionAuthOidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** | OIDC issuer to use for authentication | [optional] 

## Methods

### NewSessionAuthOidc

`func NewSessionAuthOidc() *SessionAuthOidc`

NewSessionAuthOidc instantiates a new SessionAuthOidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAuthOidcWithDefaults

`func NewSessionAuthOidcWithDefaults() *SessionAuthOidc`

NewSessionAuthOidcWithDefaults instantiates a new SessionAuthOidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *SessionAuthOidc) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SessionAuthOidc) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SessionAuthOidc) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SessionAuthOidc) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


