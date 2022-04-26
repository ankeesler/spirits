# Spirit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of this Spirit | 
**Stats** | [**SpiritStats**](SpiritStats.md) |  | 
**Actions** | Pointer to **[]string** | The Action&#39;s that this Spirit can perform | [optional] [default to ["attack"]]
**Intelligence** | Pointer to **string** | The AI setting for this Spirit | [optional] [default to "roundRobin"]

## Methods

### NewSpirit

`func NewSpirit(name string, stats SpiritStats, ) *Spirit`

NewSpirit instantiates a new Spirit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpiritWithDefaults

`func NewSpiritWithDefaults() *Spirit`

NewSpiritWithDefaults instantiates a new Spirit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Spirit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Spirit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Spirit) SetName(v string)`

SetName sets Name field to given value.


### GetStats

`func (o *Spirit) GetStats() SpiritStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Spirit) GetStatsOk() (*SpiritStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Spirit) SetStats(v SpiritStats)`

SetStats sets Stats field to given value.


### GetActions

`func (o *Spirit) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Spirit) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Spirit) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Spirit) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetIntelligence

`func (o *Spirit) GetIntelligence() string`

GetIntelligence returns the Intelligence field if non-nil, zero value otherwise.

### GetIntelligenceOk

`func (o *Spirit) GetIntelligenceOk() (*string, bool)`

GetIntelligenceOk returns a tuple with the Intelligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntelligence

`func (o *Spirit) SetIntelligence(v string)`

SetIntelligence sets Intelligence field to given value.

### HasIntelligence

`func (o *Spirit) HasIntelligence() bool`

HasIntelligence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


