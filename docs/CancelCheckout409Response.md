# CancelCheckout409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Error** | **string** |  | 
**Code** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCancelCheckout409Response

`func NewCancelCheckout409Response(message string, error_ string, code string, ) *CancelCheckout409Response`

NewCancelCheckout409Response instantiates a new CancelCheckout409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelCheckout409ResponseWithDefaults

`func NewCancelCheckout409ResponseWithDefaults() *CancelCheckout409Response`

NewCancelCheckout409ResponseWithDefaults instantiates a new CancelCheckout409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CancelCheckout409Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CancelCheckout409Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CancelCheckout409Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *CancelCheckout409Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CancelCheckout409Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CancelCheckout409Response) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *CancelCheckout409Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CancelCheckout409Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CancelCheckout409Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetMetadata

`func (o *CancelCheckout409Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CancelCheckout409Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CancelCheckout409Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CancelCheckout409Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


