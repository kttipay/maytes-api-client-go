# Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gross** | **int32** |  | 
**Discount** | **int32** |  | 
**Net** | **int32** |  | 
**Breakdown** | [**[]FeeBreakdownLine**](FeeBreakdownLine.md) |  | 

## Methods

### NewFee

`func NewFee(gross int32, discount int32, net int32, breakdown []FeeBreakdownLine, ) *Fee`

NewFee instantiates a new Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeWithDefaults

`func NewFeeWithDefaults() *Fee`

NewFeeWithDefaults instantiates a new Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGross

`func (o *Fee) GetGross() int32`

GetGross returns the Gross field if non-nil, zero value otherwise.

### GetGrossOk

`func (o *Fee) GetGrossOk() (*int32, bool)`

GetGrossOk returns a tuple with the Gross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGross

`func (o *Fee) SetGross(v int32)`

SetGross sets Gross field to given value.


### GetDiscount

`func (o *Fee) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Fee) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Fee) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.


### GetNet

`func (o *Fee) GetNet() int32`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *Fee) GetNetOk() (*int32, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *Fee) SetNet(v int32)`

SetNet sets Net field to given value.


### GetBreakdown

`func (o *Fee) GetBreakdown() []FeeBreakdownLine`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Fee) GetBreakdownOk() (*[]FeeBreakdownLine, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Fee) SetBreakdown(v []FeeBreakdownLine)`

SetBreakdown sets Breakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


