# CheckoutDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Scope** | **string** |  | 
**Amount** | **int32** |  | 
**Currency** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**AppliesTo** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCheckoutDiscount

`func NewCheckoutDiscount(name string, scope string, amount int32, currency string, ) *CheckoutDiscount`

NewCheckoutDiscount instantiates a new CheckoutDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutDiscountWithDefaults

`func NewCheckoutDiscountWithDefaults() *CheckoutDiscount`

NewCheckoutDiscountWithDefaults instantiates a new CheckoutDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CheckoutDiscount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckoutDiscount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckoutDiscount) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CheckoutDiscount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckoutDiscount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckoutDiscount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckoutDiscount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CheckoutDiscount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CheckoutDiscount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScope

`func (o *CheckoutDiscount) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CheckoutDiscount) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CheckoutDiscount) SetScope(v string)`

SetScope sets Scope field to given value.


### GetAmount

`func (o *CheckoutDiscount) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutDiscount) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutDiscount) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CheckoutDiscount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutDiscount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutDiscount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMetadata

`func (o *CheckoutDiscount) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutDiscount) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutDiscount) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutDiscount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CheckoutDiscount) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CheckoutDiscount) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAppliesTo

`func (o *CheckoutDiscount) GetAppliesTo() []string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *CheckoutDiscount) GetAppliesToOk() (*[]string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *CheckoutDiscount) SetAppliesTo(v []string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *CheckoutDiscount) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesToNil

`func (o *CheckoutDiscount) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *CheckoutDiscount) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


