# CheckoutFriendItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemRef** | **string** |  | 
**Quantity** | Pointer to **int32** | How many units of this line this friend pays for. Omit for the whole line. | [optional] 

## Methods

### NewCheckoutFriendItem

`func NewCheckoutFriendItem(itemRef string, ) *CheckoutFriendItem`

NewCheckoutFriendItem instantiates a new CheckoutFriendItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutFriendItemWithDefaults

`func NewCheckoutFriendItemWithDefaults() *CheckoutFriendItem`

NewCheckoutFriendItemWithDefaults instantiates a new CheckoutFriendItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemRef

`func (o *CheckoutFriendItem) GetItemRef() string`

GetItemRef returns the ItemRef field if non-nil, zero value otherwise.

### GetItemRefOk

`func (o *CheckoutFriendItem) GetItemRefOk() (*string, bool)`

GetItemRefOk returns a tuple with the ItemRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRef

`func (o *CheckoutFriendItem) SetItemRef(v string)`

SetItemRef sets ItemRef field to given value.


### GetQuantity

`func (o *CheckoutFriendItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CheckoutFriendItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CheckoutFriendItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CheckoutFriendItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


