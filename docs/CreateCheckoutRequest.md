# CreateCheckoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderId** | Pointer to **NullableString** |  | [optional] 
**MerchantExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**TotalAmount** | **int32** |  | 
**Currency** | **string** |  | 
**ReturnUrl** | **string** |  | 
**CancelUrl** | **string** |  | 
**Items** | [**[]LineItem**](LineItem.md) |  | 
**Fees** | Pointer to [**[]CheckoutFee**](CheckoutFee.md) |  | [optional] 
**Discounts** | Pointer to [**[]CheckoutDiscount**](CheckoutDiscount.md) |  | [optional] 
**Tax** | Pointer to [**NullableTax**](Tax.md) |  | [optional] 
**ShippingAmount** | Pointer to **NullableInt32** |  | [optional] 
**CustomerData** | Pointer to [**NullableCustomerData**](CustomerData.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**AllocationModel** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateCheckoutRequest

`func NewCreateCheckoutRequest(totalAmount int32, currency string, returnUrl string, cancelUrl string, items []LineItem, ) *CreateCheckoutRequest`

NewCreateCheckoutRequest instantiates a new CreateCheckoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutRequestWithDefaults

`func NewCreateCheckoutRequestWithDefaults() *CreateCheckoutRequest`

NewCreateCheckoutRequestWithDefaults instantiates a new CreateCheckoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderId

`func (o *CreateCheckoutRequest) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *CreateCheckoutRequest) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *CreateCheckoutRequest) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.

### HasMerchantOrderId

`func (o *CreateCheckoutRequest) HasMerchantOrderId() bool`

HasMerchantOrderId returns a boolean if a field has been set.

### SetMerchantOrderIdNil

`func (o *CreateCheckoutRequest) SetMerchantOrderIdNil(b bool)`

 SetMerchantOrderIdNil sets the value for MerchantOrderId to be an explicit nil

### UnsetMerchantOrderId
`func (o *CreateCheckoutRequest) UnsetMerchantOrderId()`

UnsetMerchantOrderId ensures that no value is present for MerchantOrderId, not even an explicit nil
### GetMerchantExpiresAt

`func (o *CreateCheckoutRequest) GetMerchantExpiresAt() time.Time`

GetMerchantExpiresAt returns the MerchantExpiresAt field if non-nil, zero value otherwise.

### GetMerchantExpiresAtOk

`func (o *CreateCheckoutRequest) GetMerchantExpiresAtOk() (*time.Time, bool)`

GetMerchantExpiresAtOk returns a tuple with the MerchantExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantExpiresAt

`func (o *CreateCheckoutRequest) SetMerchantExpiresAt(v time.Time)`

SetMerchantExpiresAt sets MerchantExpiresAt field to given value.

### HasMerchantExpiresAt

`func (o *CreateCheckoutRequest) HasMerchantExpiresAt() bool`

HasMerchantExpiresAt returns a boolean if a field has been set.

### SetMerchantExpiresAtNil

`func (o *CreateCheckoutRequest) SetMerchantExpiresAtNil(b bool)`

 SetMerchantExpiresAtNil sets the value for MerchantExpiresAt to be an explicit nil

### UnsetMerchantExpiresAt
`func (o *CreateCheckoutRequest) UnsetMerchantExpiresAt()`

UnsetMerchantExpiresAt ensures that no value is present for MerchantExpiresAt, not even an explicit nil
### GetTotalAmount

`func (o *CreateCheckoutRequest) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CreateCheckoutRequest) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CreateCheckoutRequest) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetCurrency

`func (o *CreateCheckoutRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCheckoutRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCheckoutRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReturnUrl

`func (o *CreateCheckoutRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreateCheckoutRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreateCheckoutRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetCancelUrl

`func (o *CreateCheckoutRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CreateCheckoutRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CreateCheckoutRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetItems

`func (o *CreateCheckoutRequest) GetItems() []LineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateCheckoutRequest) GetItemsOk() (*[]LineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateCheckoutRequest) SetItems(v []LineItem)`

SetItems sets Items field to given value.


### GetFees

`func (o *CreateCheckoutRequest) GetFees() []CheckoutFee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *CreateCheckoutRequest) GetFeesOk() (*[]CheckoutFee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *CreateCheckoutRequest) SetFees(v []CheckoutFee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *CreateCheckoutRequest) HasFees() bool`

HasFees returns a boolean if a field has been set.

### SetFeesNil

`func (o *CreateCheckoutRequest) SetFeesNil(b bool)`

 SetFeesNil sets the value for Fees to be an explicit nil

### UnsetFees
`func (o *CreateCheckoutRequest) UnsetFees()`

UnsetFees ensures that no value is present for Fees, not even an explicit nil
### GetDiscounts

`func (o *CreateCheckoutRequest) GetDiscounts() []CheckoutDiscount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *CreateCheckoutRequest) GetDiscountsOk() (*[]CheckoutDiscount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *CreateCheckoutRequest) SetDiscounts(v []CheckoutDiscount)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *CreateCheckoutRequest) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscountsNil

`func (o *CreateCheckoutRequest) SetDiscountsNil(b bool)`

 SetDiscountsNil sets the value for Discounts to be an explicit nil

### UnsetDiscounts
`func (o *CreateCheckoutRequest) UnsetDiscounts()`

UnsetDiscounts ensures that no value is present for Discounts, not even an explicit nil
### GetTax

`func (o *CreateCheckoutRequest) GetTax() Tax`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *CreateCheckoutRequest) GetTaxOk() (*Tax, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *CreateCheckoutRequest) SetTax(v Tax)`

SetTax sets Tax field to given value.

### HasTax

`func (o *CreateCheckoutRequest) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *CreateCheckoutRequest) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *CreateCheckoutRequest) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetShippingAmount

`func (o *CreateCheckoutRequest) GetShippingAmount() int32`

GetShippingAmount returns the ShippingAmount field if non-nil, zero value otherwise.

### GetShippingAmountOk

`func (o *CreateCheckoutRequest) GetShippingAmountOk() (*int32, bool)`

GetShippingAmountOk returns a tuple with the ShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmount

`func (o *CreateCheckoutRequest) SetShippingAmount(v int32)`

SetShippingAmount sets ShippingAmount field to given value.

### HasShippingAmount

`func (o *CreateCheckoutRequest) HasShippingAmount() bool`

HasShippingAmount returns a boolean if a field has been set.

### SetShippingAmountNil

`func (o *CreateCheckoutRequest) SetShippingAmountNil(b bool)`

 SetShippingAmountNil sets the value for ShippingAmount to be an explicit nil

### UnsetShippingAmount
`func (o *CreateCheckoutRequest) UnsetShippingAmount()`

UnsetShippingAmount ensures that no value is present for ShippingAmount, not even an explicit nil
### GetCustomerData

`func (o *CreateCheckoutRequest) GetCustomerData() CustomerData`

GetCustomerData returns the CustomerData field if non-nil, zero value otherwise.

### GetCustomerDataOk

`func (o *CreateCheckoutRequest) GetCustomerDataOk() (*CustomerData, bool)`

GetCustomerDataOk returns a tuple with the CustomerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerData

`func (o *CreateCheckoutRequest) SetCustomerData(v CustomerData)`

SetCustomerData sets CustomerData field to given value.

### HasCustomerData

`func (o *CreateCheckoutRequest) HasCustomerData() bool`

HasCustomerData returns a boolean if a field has been set.

### SetCustomerDataNil

`func (o *CreateCheckoutRequest) SetCustomerDataNil(b bool)`

 SetCustomerDataNil sets the value for CustomerData to be an explicit nil

### UnsetCustomerData
`func (o *CreateCheckoutRequest) UnsetCustomerData()`

UnsetCustomerData ensures that no value is present for CustomerData, not even an explicit nil
### GetMetadata

`func (o *CreateCheckoutRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCheckoutRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCheckoutRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCheckoutRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateCheckoutRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateCheckoutRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAllocationModel

`func (o *CreateCheckoutRequest) GetAllocationModel() string`

GetAllocationModel returns the AllocationModel field if non-nil, zero value otherwise.

### GetAllocationModelOk

`func (o *CreateCheckoutRequest) GetAllocationModelOk() (*string, bool)`

GetAllocationModelOk returns a tuple with the AllocationModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationModel

`func (o *CreateCheckoutRequest) SetAllocationModel(v string)`

SetAllocationModel sets AllocationModel field to given value.

### HasAllocationModel

`func (o *CreateCheckoutRequest) HasAllocationModel() bool`

HasAllocationModel returns a boolean if a field has been set.

### SetAllocationModelNil

`func (o *CreateCheckoutRequest) SetAllocationModelNil(b bool)`

 SetAllocationModelNil sets the value for AllocationModel to be an explicit nil

### UnsetAllocationModel
`func (o *CreateCheckoutRequest) UnsetAllocationModel()`

UnsetAllocationModel ensures that no value is present for AllocationModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


