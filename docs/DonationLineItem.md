# DonationLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemRef** | **string** |  | 
**ParentRef** | Pointer to **NullableString** |  | [optional] 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Quantity** | **int32** |  | 
**UnitPrice** | **int32** |  | 
**Currency** | **string** |  | 
**Category** | **string** |  | 

## Methods

### NewDonationLineItem

`func NewDonationLineItem(itemRef string, name string, quantity int32, unitPrice int32, currency string, category string, ) *DonationLineItem`

NewDonationLineItem instantiates a new DonationLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDonationLineItemWithDefaults

`func NewDonationLineItemWithDefaults() *DonationLineItem`

NewDonationLineItemWithDefaults instantiates a new DonationLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemRef

`func (o *DonationLineItem) GetItemRef() string`

GetItemRef returns the ItemRef field if non-nil, zero value otherwise.

### GetItemRefOk

`func (o *DonationLineItem) GetItemRefOk() (*string, bool)`

GetItemRefOk returns a tuple with the ItemRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRef

`func (o *DonationLineItem) SetItemRef(v string)`

SetItemRef sets ItemRef field to given value.


### GetParentRef

`func (o *DonationLineItem) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *DonationLineItem) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *DonationLineItem) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *DonationLineItem) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### SetParentRefNil

`func (o *DonationLineItem) SetParentRefNil(b bool)`

 SetParentRefNil sets the value for ParentRef to be an explicit nil

### UnsetParentRef
`func (o *DonationLineItem) UnsetParentRef()`

UnsetParentRef ensures that no value is present for ParentRef, not even an explicit nil
### GetSku

`func (o *DonationLineItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *DonationLineItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *DonationLineItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *DonationLineItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *DonationLineItem) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *DonationLineItem) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetName

`func (o *DonationLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DonationLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DonationLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DonationLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DonationLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DonationLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DonationLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DonationLineItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DonationLineItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *DonationLineItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *DonationLineItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *DonationLineItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *DonationLineItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *DonationLineItem) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *DonationLineItem) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetQuantity

`func (o *DonationLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DonationLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DonationLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *DonationLineItem) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *DonationLineItem) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *DonationLineItem) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetCurrency

`func (o *DonationLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DonationLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DonationLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCategory

`func (o *DonationLineItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DonationLineItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DonationLineItem) SetCategory(v string)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


