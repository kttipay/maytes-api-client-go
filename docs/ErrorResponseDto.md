# ErrorResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Error** | **string** |  | 
**Code** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewErrorResponseDto

`func NewErrorResponseDto(message string, error_ string, code string, ) *ErrorResponseDto`

NewErrorResponseDto instantiates a new ErrorResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseDtoWithDefaults

`func NewErrorResponseDtoWithDefaults() *ErrorResponseDto`

NewErrorResponseDtoWithDefaults instantiates a new ErrorResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorResponseDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *ErrorResponseDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponseDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponseDto) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *ErrorResponseDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseDto) SetCode(v string)`

SetCode sets Code field to given value.


### GetMetadata

`func (o *ErrorResponseDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ErrorResponseDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ErrorResponseDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ErrorResponseDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


