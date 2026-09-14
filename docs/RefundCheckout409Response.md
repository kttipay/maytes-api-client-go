# RefundCheckout409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Error** | **string** |  | 
**Code** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRefundCheckout409Response

`func NewRefundCheckout409Response(message string, error_ string, code string, ) *RefundCheckout409Response`

NewRefundCheckout409Response instantiates a new RefundCheckout409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundCheckout409ResponseWithDefaults

`func NewRefundCheckout409ResponseWithDefaults() *RefundCheckout409Response`

NewRefundCheckout409ResponseWithDefaults instantiates a new RefundCheckout409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RefundCheckout409Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RefundCheckout409Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RefundCheckout409Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *RefundCheckout409Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RefundCheckout409Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RefundCheckout409Response) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *RefundCheckout409Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RefundCheckout409Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RefundCheckout409Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetMetadata

`func (o *RefundCheckout409Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RefundCheckout409Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RefundCheckout409Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RefundCheckout409Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


