# PhysicalGoodsLineItem

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
**CustomData** | [**PhysicalGoodsCustomData**](PhysicalGoodsCustomData.md) |  | 

## Methods

### NewPhysicalGoodsLineItem

`func NewPhysicalGoodsLineItem(itemRef string, name string, quantity int32, unitPrice int32, currency string, category string, customData PhysicalGoodsCustomData, ) *PhysicalGoodsLineItem`

NewPhysicalGoodsLineItem instantiates a new PhysicalGoodsLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalGoodsLineItemWithDefaults

`func NewPhysicalGoodsLineItemWithDefaults() *PhysicalGoodsLineItem`

NewPhysicalGoodsLineItemWithDefaults instantiates a new PhysicalGoodsLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemRef

`func (o *PhysicalGoodsLineItem) GetItemRef() string`

GetItemRef returns the ItemRef field if non-nil, zero value otherwise.

### GetItemRefOk

`func (o *PhysicalGoodsLineItem) GetItemRefOk() (*string, bool)`

GetItemRefOk returns a tuple with the ItemRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRef

`func (o *PhysicalGoodsLineItem) SetItemRef(v string)`

SetItemRef sets ItemRef field to given value.


### GetParentRef

`func (o *PhysicalGoodsLineItem) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *PhysicalGoodsLineItem) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *PhysicalGoodsLineItem) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *PhysicalGoodsLineItem) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### SetParentRefNil

`func (o *PhysicalGoodsLineItem) SetParentRefNil(b bool)`

 SetParentRefNil sets the value for ParentRef to be an explicit nil

### UnsetParentRef
`func (o *PhysicalGoodsLineItem) UnsetParentRef()`

UnsetParentRef ensures that no value is present for ParentRef, not even an explicit nil
### GetSku

`func (o *PhysicalGoodsLineItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PhysicalGoodsLineItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PhysicalGoodsLineItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *PhysicalGoodsLineItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *PhysicalGoodsLineItem) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *PhysicalGoodsLineItem) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetName

`func (o *PhysicalGoodsLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhysicalGoodsLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhysicalGoodsLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PhysicalGoodsLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PhysicalGoodsLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PhysicalGoodsLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PhysicalGoodsLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PhysicalGoodsLineItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PhysicalGoodsLineItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *PhysicalGoodsLineItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PhysicalGoodsLineItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PhysicalGoodsLineItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PhysicalGoodsLineItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PhysicalGoodsLineItem) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PhysicalGoodsLineItem) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetQuantity

`func (o *PhysicalGoodsLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PhysicalGoodsLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PhysicalGoodsLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *PhysicalGoodsLineItem) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PhysicalGoodsLineItem) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PhysicalGoodsLineItem) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetCurrency

`func (o *PhysicalGoodsLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PhysicalGoodsLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PhysicalGoodsLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCategory

`func (o *PhysicalGoodsLineItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PhysicalGoodsLineItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PhysicalGoodsLineItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCustomData

`func (o *PhysicalGoodsLineItem) GetCustomData() PhysicalGoodsCustomData`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *PhysicalGoodsLineItem) GetCustomDataOk() (*PhysicalGoodsCustomData, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *PhysicalGoodsLineItem) SetCustomData(v PhysicalGoodsCustomData)`

SetCustomData sets CustomData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


