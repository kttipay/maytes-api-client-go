# CheckoutCancelled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutUuid** | **string** |  | 
**MerchantOrderId** | **NullableString** |  | 
**Status** | **string** |  | 

## Methods

### NewCheckoutCancelled

`func NewCheckoutCancelled(checkoutUuid string, merchantOrderId NullableString, status string, ) *CheckoutCancelled`

NewCheckoutCancelled instantiates a new CheckoutCancelled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutCancelledWithDefaults

`func NewCheckoutCancelledWithDefaults() *CheckoutCancelled`

NewCheckoutCancelledWithDefaults instantiates a new CheckoutCancelled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutUuid

`func (o *CheckoutCancelled) GetCheckoutUuid() string`

GetCheckoutUuid returns the CheckoutUuid field if non-nil, zero value otherwise.

### GetCheckoutUuidOk

`func (o *CheckoutCancelled) GetCheckoutUuidOk() (*string, bool)`

GetCheckoutUuidOk returns a tuple with the CheckoutUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUuid

`func (o *CheckoutCancelled) SetCheckoutUuid(v string)`

SetCheckoutUuid sets CheckoutUuid field to given value.


### GetMerchantOrderId

`func (o *CheckoutCancelled) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *CheckoutCancelled) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *CheckoutCancelled) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.


### SetMerchantOrderIdNil

`func (o *CheckoutCancelled) SetMerchantOrderIdNil(b bool)`

 SetMerchantOrderIdNil sets the value for MerchantOrderId to be an explicit nil

### UnsetMerchantOrderId
`func (o *CheckoutCancelled) UnsetMerchantOrderId()`

UnsetMerchantOrderId ensures that no value is present for MerchantOrderId, not even an explicit nil
### GetStatus

`func (o *CheckoutCancelled) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutCancelled) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutCancelled) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


