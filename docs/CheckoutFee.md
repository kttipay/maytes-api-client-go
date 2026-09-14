# CheckoutFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**FeeType** | Pointer to **NullableString** |  | [optional] 
**Scope** | **string** |  | 
**Amount** | **int32** |  | 
**Currency** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**AppliesTo** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCheckoutFee

`func NewCheckoutFee(name string, scope string, amount int32, currency string, ) *CheckoutFee`

NewCheckoutFee instantiates a new CheckoutFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutFeeWithDefaults

`func NewCheckoutFeeWithDefaults() *CheckoutFee`

NewCheckoutFeeWithDefaults instantiates a new CheckoutFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CheckoutFee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckoutFee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckoutFee) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CheckoutFee) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckoutFee) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckoutFee) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckoutFee) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CheckoutFee) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CheckoutFee) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFeeType

`func (o *CheckoutFee) GetFeeType() string`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *CheckoutFee) GetFeeTypeOk() (*string, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *CheckoutFee) SetFeeType(v string)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *CheckoutFee) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### SetFeeTypeNil

`func (o *CheckoutFee) SetFeeTypeNil(b bool)`

 SetFeeTypeNil sets the value for FeeType to be an explicit nil

### UnsetFeeType
`func (o *CheckoutFee) UnsetFeeType()`

UnsetFeeType ensures that no value is present for FeeType, not even an explicit nil
### GetScope

`func (o *CheckoutFee) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CheckoutFee) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CheckoutFee) SetScope(v string)`

SetScope sets Scope field to given value.


### GetAmount

`func (o *CheckoutFee) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutFee) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutFee) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CheckoutFee) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutFee) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutFee) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMetadata

`func (o *CheckoutFee) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutFee) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutFee) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutFee) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CheckoutFee) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CheckoutFee) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAppliesTo

`func (o *CheckoutFee) GetAppliesTo() []string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *CheckoutFee) GetAppliesToOk() (*[]string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *CheckoutFee) SetAppliesTo(v []string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *CheckoutFee) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *CheckoutFee) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *CheckoutFee) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


