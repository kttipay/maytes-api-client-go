# Checkout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutUuid** | **string** |  | 
**MerchantOrderId** | **NullableString** |  | 
**Status** | **string** |  | 
**TransactionRef** | **NullableString** |  | 
**CheckoutUrl** | **string** |  | 
**MerchantExpiresAt** | **NullableTime** |  | 
**CapturedAt** | **NullableTime** |  | 
**ReturnUrl** | **NullableString** |  | 
**CancelUrl** | **NullableString** |  | 
**Amount** | **int32** |  | 
**Currency** | **string** |  | 
**ServiceFee** | [**Fee**](Fee.md) |  | 
**RefundedAmount** | **int32** |  | 
**RemainingAmount** | **int32** |  | 
**MerchantProcessorFee** | **int32** |  | 
**MerchantServiceFee** | **int32** |  | 
**LastRefundAt** | **NullableTime** |  | 

## Methods

### NewCheckout

`func NewCheckout(checkoutUuid string, merchantOrderId NullableString, status string, transactionRef NullableString, checkoutUrl string, merchantExpiresAt NullableTime, capturedAt NullableTime, returnUrl NullableString, cancelUrl NullableString, amount int32, currency string, serviceFee Fee, refundedAmount int32, remainingAmount int32, merchantProcessorFee int32, merchantServiceFee int32, lastRefundAt NullableTime, ) *Checkout`

NewCheckout instantiates a new Checkout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutWithDefaults

`func NewCheckoutWithDefaults() *Checkout`

NewCheckoutWithDefaults instantiates a new Checkout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutUuid

`func (o *Checkout) GetCheckoutUuid() string`

GetCheckoutUuid returns the CheckoutUuid field if non-nil, zero value otherwise.

### GetCheckoutUuidOk

`func (o *Checkout) GetCheckoutUuidOk() (*string, bool)`

GetCheckoutUuidOk returns a tuple with the CheckoutUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUuid

`func (o *Checkout) SetCheckoutUuid(v string)`

SetCheckoutUuid sets CheckoutUuid field to given value.


### GetMerchantOrderId

`func (o *Checkout) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *Checkout) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *Checkout) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.


### SetMerchantOrderIdNil

`func (o *Checkout) SetMerchantOrderIdNil(b bool)`

 SetMerchantOrderIdNil sets the value for MerchantOrderId to be an explicit nil

### UnsetMerchantOrderId
`func (o *Checkout) UnsetMerchantOrderId()`

UnsetMerchantOrderId ensures that no value is present for MerchantOrderId, not even an explicit nil
### GetStatus

`func (o *Checkout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Checkout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Checkout) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransactionRef

`func (o *Checkout) GetTransactionRef() string`

GetTransactionRef returns the TransactionRef field if non-nil, zero value otherwise.

### GetTransactionRefOk

`func (o *Checkout) GetTransactionRefOk() (*string, bool)`

GetTransactionRefOk returns a tuple with the TransactionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRef

`func (o *Checkout) SetTransactionRef(v string)`

SetTransactionRef sets TransactionRef field to given value.


### SetTransactionRefNil

`func (o *Checkout) SetTransactionRefNil(b bool)`

 SetTransactionRefNil sets the value for TransactionRef to be an explicit nil

### UnsetTransactionRef
`func (o *Checkout) UnsetTransactionRef()`

UnsetTransactionRef ensures that no value is present for TransactionRef, not even an explicit nil
### GetCheckoutUrl

`func (o *Checkout) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *Checkout) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *Checkout) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.


### GetMerchantExpiresAt

`func (o *Checkout) GetMerchantExpiresAt() time.Time`

GetMerchantExpiresAt returns the MerchantExpiresAt field if non-nil, zero value otherwise.

### GetMerchantExpiresAtOk

`func (o *Checkout) GetMerchantExpiresAtOk() (*time.Time, bool)`

GetMerchantExpiresAtOk returns a tuple with the MerchantExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantExpiresAt

`func (o *Checkout) SetMerchantExpiresAt(v time.Time)`

SetMerchantExpiresAt sets MerchantExpiresAt field to given value.


### SetMerchantExpiresAtNil

`func (o *Checkout) SetMerchantExpiresAtNil(b bool)`

 SetMerchantExpiresAtNil sets the value for MerchantExpiresAt to be an explicit nil

### UnsetMerchantExpiresAt
`func (o *Checkout) UnsetMerchantExpiresAt()`

UnsetMerchantExpiresAt ensures that no value is present for MerchantExpiresAt, not even an explicit nil
### GetCapturedAt

`func (o *Checkout) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *Checkout) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *Checkout) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.


### SetCapturedAtNil

`func (o *Checkout) SetCapturedAtNil(b bool)`

 SetCapturedAtNil sets the value for CapturedAt to be an explicit nil

### UnsetCapturedAt
`func (o *Checkout) UnsetCapturedAt()`

UnsetCapturedAt ensures that no value is present for CapturedAt, not even an explicit nil
### GetReturnUrl

`func (o *Checkout) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *Checkout) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *Checkout) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### SetReturnUrlNil

`func (o *Checkout) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *Checkout) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetCancelUrl

`func (o *Checkout) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *Checkout) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *Checkout) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### SetCancelUrlNil

`func (o *Checkout) SetCancelUrlNil(b bool)`

 SetCancelUrlNil sets the value for CancelUrl to be an explicit nil

### UnsetCancelUrl
`func (o *Checkout) UnsetCancelUrl()`

UnsetCancelUrl ensures that no value is present for CancelUrl, not even an explicit nil
### GetAmount

`func (o *Checkout) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Checkout) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Checkout) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Checkout) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Checkout) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Checkout) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetServiceFee

`func (o *Checkout) GetServiceFee() Fee`

GetServiceFee returns the ServiceFee field if non-nil, zero value otherwise.

### GetServiceFeeOk

`func (o *Checkout) GetServiceFeeOk() (*Fee, bool)`

GetServiceFeeOk returns a tuple with the ServiceFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFee

`func (o *Checkout) SetServiceFee(v Fee)`

SetServiceFee sets ServiceFee field to given value.


### GetRefundedAmount

`func (o *Checkout) GetRefundedAmount() int32`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *Checkout) GetRefundedAmountOk() (*int32, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *Checkout) SetRefundedAmount(v int32)`

SetRefundedAmount sets RefundedAmount field to given value.


### GetRemainingAmount

`func (o *Checkout) GetRemainingAmount() int32`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *Checkout) GetRemainingAmountOk() (*int32, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *Checkout) SetRemainingAmount(v int32)`

SetRemainingAmount sets RemainingAmount field to given value.


### GetMerchantProcessorFee

`func (o *Checkout) GetMerchantProcessorFee() int32`

GetMerchantProcessorFee returns the MerchantProcessorFee field if non-nil, zero value otherwise.

### GetMerchantProcessorFeeOk

`func (o *Checkout) GetMerchantProcessorFeeOk() (*int32, bool)`

GetMerchantProcessorFeeOk returns a tuple with the MerchantProcessorFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProcessorFee

`func (o *Checkout) SetMerchantProcessorFee(v int32)`

SetMerchantProcessorFee sets MerchantProcessorFee field to given value.


### GetMerchantServiceFee

`func (o *Checkout) GetMerchantServiceFee() int32`

GetMerchantServiceFee returns the MerchantServiceFee field if non-nil, zero value otherwise.

### GetMerchantServiceFeeOk

`func (o *Checkout) GetMerchantServiceFeeOk() (*int32, bool)`

GetMerchantServiceFeeOk returns a tuple with the MerchantServiceFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantServiceFee

`func (o *Checkout) SetMerchantServiceFee(v int32)`

SetMerchantServiceFee sets MerchantServiceFee field to given value.


### GetLastRefundAt

`func (o *Checkout) GetLastRefundAt() time.Time`

GetLastRefundAt returns the LastRefundAt field if non-nil, zero value otherwise.

### GetLastRefundAtOk

`func (o *Checkout) GetLastRefundAtOk() (*time.Time, bool)`

GetLastRefundAtOk returns a tuple with the LastRefundAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefundAt

`func (o *Checkout) SetLastRefundAt(v time.Time)`

SetLastRefundAt sets LastRefundAt field to given value.


### SetLastRefundAtNil

`func (o *Checkout) SetLastRefundAtNil(b bool)`

 SetLastRefundAtNil sets the value for LastRefundAt to be an explicit nil

### UnsetLastRefundAt
`func (o *Checkout) UnsetLastRefundAt()`

UnsetLastRefundAt ensures that no value is present for LastRefundAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


