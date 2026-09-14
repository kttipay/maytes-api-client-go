# DigitalGoodsLineItem

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
**CustomData** | Pointer to [**DigitalGoodsCustomData**](DigitalGoodsCustomData.md) |  | [optional] 

## Methods

### NewDigitalGoodsLineItem

`func NewDigitalGoodsLineItem(itemRef string, name string, quantity int32, unitPrice int32, currency string, category string, ) *DigitalGoodsLineItem`

NewDigitalGoodsLineItem instantiates a new DigitalGoodsLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalGoodsLineItemWithDefaults

`func NewDigitalGoodsLineItemWithDefaults() *DigitalGoodsLineItem`

NewDigitalGoodsLineItemWithDefaults instantiates a new DigitalGoodsLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemRef

`func (o *DigitalGoodsLineItem) GetItemRef() string`

GetItemRef returns the ItemRef field if non-nil, zero value otherwise.

### GetItemRefOk

`func (o *DigitalGoodsLineItem) GetItemRefOk() (*string, bool)`

GetItemRefOk returns a tuple with the ItemRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRef

`func (o *DigitalGoodsLineItem) SetItemRef(v string)`

SetItemRef sets ItemRef field to given value.


### GetParentRef

`func (o *DigitalGoodsLineItem) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *DigitalGoodsLineItem) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *DigitalGoodsLineItem) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *DigitalGoodsLineItem) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### SetParentRefNil

`func (o *DigitalGoodsLineItem) SetParentRefNil(b bool)`

 SetParentRefNil sets the value for ParentRef to be an explicit nil

### UnsetParentRef
`func (o *DigitalGoodsLineItem) UnsetParentRef()`

UnsetParentRef ensures that no value is present for ParentRef, not even an explicit nil
### GetSku

`func (o *DigitalGoodsLineItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *DigitalGoodsLineItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *DigitalGoodsLineItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *DigitalGoodsLineItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *DigitalGoodsLineItem) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *DigitalGoodsLineItem) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetName

`func (o *DigitalGoodsLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DigitalGoodsLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DigitalGoodsLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DigitalGoodsLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DigitalGoodsLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DigitalGoodsLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DigitalGoodsLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DigitalGoodsLineItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DigitalGoodsLineItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *DigitalGoodsLineItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *DigitalGoodsLineItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *DigitalGoodsLineItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *DigitalGoodsLineItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *DigitalGoodsLineItem) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *DigitalGoodsLineItem) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetQuantity

`func (o *DigitalGoodsLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DigitalGoodsLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DigitalGoodsLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *DigitalGoodsLineItem) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *DigitalGoodsLineItem) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *DigitalGoodsLineItem) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetCurrency

`func (o *DigitalGoodsLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DigitalGoodsLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DigitalGoodsLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCategory

`func (o *DigitalGoodsLineItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DigitalGoodsLineItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DigitalGoodsLineItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCustomData

`func (o *DigitalGoodsLineItem) GetCustomData() DigitalGoodsCustomData`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *DigitalGoodsLineItem) GetCustomDataOk() (*DigitalGoodsCustomData, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *DigitalGoodsLineItem) SetCustomData(v DigitalGoodsCustomData)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *DigitalGoodsLineItem) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


