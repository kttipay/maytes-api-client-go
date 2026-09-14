# CheckoutCaptured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutUuid** | **string** |  | 
**MerchantOrderId** | **NullableString** |  | 
**Status** | **string** |  | 
**CapturedAt** | **NullableTime** |  | 
**Amount** | **int32** |  | 
**Currency** | **string** |  | 

## Methods

### NewCheckoutCaptured

`func NewCheckoutCaptured(checkoutUuid string, merchantOrderId NullableString, status string, capturedAt NullableTime, amount int32, currency string, ) *CheckoutCaptured`

NewCheckoutCaptured instantiates a new CheckoutCaptured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutCapturedWithDefaults

`func NewCheckoutCapturedWithDefaults() *CheckoutCaptured`

NewCheckoutCapturedWithDefaults instantiates a new CheckoutCaptured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutUuid

`func (o *CheckoutCaptured) GetCheckoutUuid() string`

GetCheckoutUuid returns the CheckoutUuid field if non-nil, zero value otherwise.

### GetCheckoutUuidOk

`func (o *CheckoutCaptured) GetCheckoutUuidOk() (*string, bool)`

GetCheckoutUuidOk returns a tuple with the CheckoutUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUuid

`func (o *CheckoutCaptured) SetCheckoutUuid(v string)`

SetCheckoutUuid sets CheckoutUuid field to given value.


### GetMerchantOrderId

`func (o *CheckoutCaptured) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *CheckoutCaptured) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *CheckoutCaptured) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.


### SetMerchantOrderIdNil

`func (o *CheckoutCaptured) SetMerchantOrderIdNil(b bool)`

 SetMerchantOrderIdNil sets the value for MerchantOrderId to be an explicit nil

### UnsetMerchantOrderId
`func (o *CheckoutCaptured) UnsetMerchantOrderId()`

UnsetMerchantOrderId ensures that no value is present for MerchantOrderId, not even an explicit nil
### GetStatus

`func (o *CheckoutCaptured) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutCaptured) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutCaptured) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCapturedAt

`func (o *CheckoutCaptured) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *CheckoutCaptured) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *CheckoutCaptured) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.


### SetCapturedAtNil

`func (o *CheckoutCaptured) SetCapturedAtNil(b bool)`

 SetCapturedAtNil sets the value for CapturedAt to be an explicit nil

### UnsetCapturedAt
`func (o *CheckoutCaptured) UnsetCapturedAt()`

UnsetCapturedAt ensures that no value is present for CapturedAt, not even an explicit nil
### GetAmount

`func (o *CheckoutCaptured) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutCaptured) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutCaptured) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CheckoutCaptured) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutCaptured) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutCaptured) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


