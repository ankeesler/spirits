# SpiritStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | **int64** | A quantitative representation of the energy of the Spirit; when this drops to 0, the Spirit is no longer to participate in a Battle | 
**Power** | Pointer to **int64** | A quantitative representation of the might of the Spirit | [optional] [default to 0]
**Armor** | Pointer to **int64** | A quantitative representation of the defense of the Spirit | [optional] [default to 0]
**Agility** | Pointer to **int64** | A quantitative representation of the speed of the Spirit | [optional] [default to 0]

## Methods

### NewSpiritStats

`func NewSpiritStats(health int64, ) *SpiritStats`

NewSpiritStats instantiates a new SpiritStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpiritStatsWithDefaults

`func NewSpiritStatsWithDefaults() *SpiritStats`

NewSpiritStatsWithDefaults instantiates a new SpiritStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *SpiritStats) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *SpiritStats) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *SpiritStats) SetHealth(v int64)`

SetHealth sets Health field to given value.


### GetPower

`func (o *SpiritStats) GetPower() int64`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *SpiritStats) GetPowerOk() (*int64, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *SpiritStats) SetPower(v int64)`

SetPower sets Power field to given value.

### HasPower

`func (o *SpiritStats) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetArmor

`func (o *SpiritStats) GetArmor() int64`

GetArmor returns the Armor field if non-nil, zero value otherwise.

### GetArmorOk

`func (o *SpiritStats) GetArmorOk() (*int64, bool)`

GetArmorOk returns a tuple with the Armor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmor

`func (o *SpiritStats) SetArmor(v int64)`

SetArmor sets Armor field to given value.

### HasArmor

`func (o *SpiritStats) HasArmor() bool`

HasArmor returns a boolean if a field has been set.

### GetAgility

`func (o *SpiritStats) GetAgility() int64`

GetAgility returns the Agility field if non-nil, zero value otherwise.

### GetAgilityOk

`func (o *SpiritStats) GetAgilityOk() (*int64, bool)`

GetAgilityOk returns a tuple with the Agility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgility

`func (o *SpiritStats) SetAgility(v int64)`

SetAgility sets Agility field to given value.

### HasAgility

`func (o *SpiritStats) HasAgility() bool`

HasAgility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


