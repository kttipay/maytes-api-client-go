# SettlementUpsell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SalesAmount** | **int32** | Upsell order totals paid out in this settlement, gross, minor units | 
**FeeAmount** | **int32** | Upsell fees kept by Maytes in this window, minor units | 
**NetAmount** | **int32** | sales_amount minus fee_amount: the upsell part of the payout | 
**OrderCount** | **int32** |  | 

## Methods

### NewSettlementUpsell

`func NewSettlementUpsell(salesAmount int32, feeAmount int32, netAmount int32, orderCount int32, ) *SettlementUpsell`

NewSettlementUpsell instantiates a new SettlementUpsell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementUpsellWithDefaults

`func NewSettlementUpsellWithDefaults() *SettlementUpsell`

NewSettlementUpsellWithDefaults instantiates a new SettlementUpsell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalesAmount

`func (o *SettlementUpsell) GetSalesAmount() int32`

GetSalesAmount returns the SalesAmount field if non-nil, zero value otherwise.

### GetSalesAmountOk

`func (o *SettlementUpsell) GetSalesAmountOk() (*int32, bool)`

GetSalesAmountOk returns a tuple with the SalesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAmount

`func (o *SettlementUpsell) SetSalesAmount(v int32)`

SetSalesAmount sets SalesAmount field to given value.


### GetFeeAmount

`func (o *SettlementUpsell) GetFeeAmount() int32`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *SettlementUpsell) GetFeeAmountOk() (*int32, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *SettlementUpsell) SetFeeAmount(v int32)`

SetFeeAmount sets FeeAmount field to given value.


### GetNetAmount

`func (o *SettlementUpsell) GetNetAmount() int32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *SettlementUpsell) GetNetAmountOk() (*int32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *SettlementUpsell) SetNetAmount(v int32)`

SetNetAmount sets NetAmount field to given value.


### GetOrderCount

`func (o *SettlementUpsell) GetOrderCount() int32`

GetOrderCount returns the OrderCount field if non-nil, zero value otherwise.

### GetOrderCountOk

`func (o *SettlementUpsell) GetOrderCountOk() (*int32, bool)`

GetOrderCountOk returns a tuple with the OrderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCount

`func (o *SettlementUpsell) SetOrderCount(v int32)`

SetOrderCount sets OrderCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


