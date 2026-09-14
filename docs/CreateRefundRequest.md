# CreateRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Free-text reason; forwarded to Stripe as refund metadata. | [optional] 
**MerchantRefundRef** | **string** | Idempotency key scoped to the merchant. Repeating the same value returns the existing refund. | 

## Methods

### NewCreateRefundRequest

`func NewCreateRefundRequest(merchantRefundRef string, ) *CreateRefundRequest`

NewCreateRefundRequest instantiates a new CreateRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefundRequestWithDefaults

`func NewCreateRefundRequestWithDefaults() *CreateRefundRequest`

NewCreateRefundRequestWithDefaults instantiates a new CreateRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CreateRefundRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateRefundRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateRefundRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateRefundRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMerchantRefundRef

`func (o *CreateRefundRequest) GetMerchantRefundRef() string`

GetMerchantRefundRef returns the MerchantRefundRef field if non-nil, zero value otherwise.

### GetMerchantRefundRefOk

`func (o *CreateRefundRequest) GetMerchantRefundRefOk() (*string, bool)`

GetMerchantRefundRefOk returns a tuple with the MerchantRefundRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefundRef

`func (o *CreateRefundRequest) SetMerchantRefundRef(v string)`

SetMerchantRefundRef sets MerchantRefundRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


