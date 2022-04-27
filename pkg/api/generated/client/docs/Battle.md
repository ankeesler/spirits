# Battle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of this Battle | 
**Teams** | Pointer to **[]string** | The Team&#39;s involved in this Battle. | [optional] 

## Methods

### NewBattle

`func NewBattle(name string, ) *Battle`

NewBattle instantiates a new Battle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBattleWithDefaults

`func NewBattleWithDefaults() *Battle`

NewBattleWithDefaults instantiates a new Battle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Battle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Battle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Battle) SetName(v string)`

SetName sets Name field to given value.


### GetTeams

`func (o *Battle) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Battle) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Battle) SetTeams(v []string)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Battle) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


