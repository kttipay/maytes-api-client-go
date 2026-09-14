# CaptureCheckoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOrderId** | Pointer to **NullableString** |  | [optional] 
**ExpectedTotalAmount** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCaptureCheckoutRequest

`func NewCaptureCheckoutRequest() *CaptureCheckoutRequest`

NewCaptureCheckoutRequest instantiates a new CaptureCheckoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureCheckoutRequestWithDefaults

`func NewCaptureCheckoutRequestWithDefaults() *CaptureCheckoutRequest`

NewCaptureCheckoutRequestWithDefaults instantiates a new CaptureCheckoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOrderId

`func (o *CaptureCheckoutRequest) GetMerchantOrderId() string`

GetMerchantOrderId returns the MerchantOrderId field if non-nil, zero value otherwise.

### GetMerchantOrderIdOk

`func (o *CaptureCheckoutRequest) GetMerchantOrderIdOk() (*string, bool)`

GetMerchantOrderIdOk returns a tuple with the MerchantOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderId

`func (o *CaptureCheckoutRequest) SetMerchantOrderId(v string)`

SetMerchantOrderId sets MerchantOrderId field to given value.

### HasMerchantOrderId

`func (o *CaptureCheckoutRequest) HasMerchantOrderId() bool`

HasMerchantOrderId returns a boolean if a field has been set.

### SetMerchantOrderIdNil

`func (o *CaptureCheckoutRequest) SetMerchantOrderIdNil(b bool)`

 SetMerchantOrderIdNil sets the value for MerchantOrderId to be an explicit nil

### UnsetMerchantOrderId
`func (o *CaptureCheckoutRequest) UnsetMerchantOrderId()`

UnsetMerchantOrderId ensures that no value is present for MerchantOrderId, not even an explicit nil
### GetExpectedTotalAmount

`func (o *CaptureCheckoutRequest) GetExpectedTotalAmount() int32`

GetExpectedTotalAmount returns the ExpectedTotalAmount field if non-nil, zero value otherwise.

### GetExpectedTotalAmountOk

`func (o *CaptureCheckoutRequest) GetExpectedTotalAmountOk() (*int32, bool)`

GetExpectedTotalAmountOk returns a tuple with the ExpectedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTotalAmount

`func (o *CaptureCheckoutRequest) SetExpectedTotalAmount(v int32)`

SetExpectedTotalAmount sets ExpectedTotalAmount field to given value.

### HasExpectedTotalAmount

`func (o *CaptureCheckoutRequest) HasExpectedTotalAmount() bool`

HasExpectedTotalAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CaptureCheckoutRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CaptureCheckoutRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CaptureCheckoutRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CaptureCheckoutRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMetadata

`func (o *CaptureCheckoutRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CaptureCheckoutRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CaptureCheckoutRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CaptureCheckoutRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CaptureCheckoutRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CaptureCheckoutRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


