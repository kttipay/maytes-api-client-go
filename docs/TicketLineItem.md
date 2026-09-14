# TicketLineItem

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
**CustomData** | [**TicketCustomData**](TicketCustomData.md) |  | 

## Methods

### NewTicketLineItem

`func NewTicketLineItem(itemRef string, name string, quantity int32, unitPrice int32, currency string, category string, customData TicketCustomData, ) *TicketLineItem`

NewTicketLineItem instantiates a new TicketLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketLineItemWithDefaults

`func NewTicketLineItemWithDefaults() *TicketLineItem`

NewTicketLineItemWithDefaults instantiates a new TicketLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemRef

`func (o *TicketLineItem) GetItemRef() string`

GetItemRef returns the ItemRef field if non-nil, zero value otherwise.

### GetItemRefOk

`func (o *TicketLineItem) GetItemRefOk() (*string, bool)`

GetItemRefOk returns a tuple with the ItemRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRef

`func (o *TicketLineItem) SetItemRef(v string)`

SetItemRef sets ItemRef field to given value.


### GetParentRef

`func (o *TicketLineItem) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *TicketLineItem) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *TicketLineItem) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *TicketLineItem) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### SetParentRefNil

`func (o *TicketLineItem) SetParentRefNil(b bool)`

 SetParentRefNil sets the value for ParentRef to be an explicit nil

### UnsetParentRef
`func (o *TicketLineItem) UnsetParentRef()`

UnsetParentRef ensures that no value is present for ParentRef, not even an explicit nil
### GetSku

`func (o *TicketLineItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *TicketLineItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *TicketLineItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *TicketLineItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *TicketLineItem) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *TicketLineItem) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetName

`func (o *TicketLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TicketLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TicketLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TicketLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TicketLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TicketLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TicketLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TicketLineItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TicketLineItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *TicketLineItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *TicketLineItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *TicketLineItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *TicketLineItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *TicketLineItem) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *TicketLineItem) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetQuantity

`func (o *TicketLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TicketLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TicketLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *TicketLineItem) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *TicketLineItem) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *TicketLineItem) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetCurrency

`func (o *TicketLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TicketLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TicketLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCategory

`func (o *TicketLineItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TicketLineItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TicketLineItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCustomData

`func (o *TicketLineItem) GetCustomData() TicketCustomData`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *TicketLineItem) GetCustomDataOk() (*TicketCustomData, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *TicketLineItem) SetCustomData(v TicketCustomData)`

SetCustomData sets CustomData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


