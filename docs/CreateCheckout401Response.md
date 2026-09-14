# CreateCheckout401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Error** | **string** |  | 
**Code** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateCheckout401Response

`func NewCreateCheckout401Response(message string, error_ string, code string, ) *CreateCheckout401Response`

NewCreateCheckout401Response instantiates a new CreateCheckout401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckout401ResponseWithDefaults

`func NewCreateCheckout401ResponseWithDefaults() *CreateCheckout401Response`

NewCreateCheckout401ResponseWithDefaults instantiates a new CreateCheckout401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateCheckout401Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateCheckout401Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateCheckout401Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *CreateCheckout401Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateCheckout401Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateCheckout401Response) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *CreateCheckout401Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateCheckout401Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateCheckout401Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetMetadata

`func (o *CreateCheckout401Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCheckout401Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCheckout401Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCheckout401Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


