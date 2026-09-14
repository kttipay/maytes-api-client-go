# Tax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**Percentage** | Pointer to **NullableString** |  | [optional] 
**Inclusive** | Pointer to **bool** |  | [optional] 

## Methods

### NewTax

`func NewTax() *Tax`

NewTax instantiates a new Tax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxWithDefaults

`func NewTaxWithDefaults() *Tax`

NewTaxWithDefaults instantiates a new Tax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Tax) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Tax) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Tax) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Tax) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPercentage

`func (o *Tax) GetPercentage() string`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *Tax) GetPercentageOk() (*string, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *Tax) SetPercentage(v string)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *Tax) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *Tax) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *Tax) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetInclusive

`func (o *Tax) GetInclusive() bool`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *Tax) GetInclusiveOk() (*bool, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *Tax) SetInclusive(v bool)`

SetInclusive sets Inclusive field to given value.

### HasInclusive

`func (o *Tax) HasInclusive() bool`

HasInclusive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


