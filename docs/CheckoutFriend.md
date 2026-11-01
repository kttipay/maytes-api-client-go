# CheckoutFriend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Mobile** | **string** | E.164, e.g. +61412345678 | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **int32** | This friend’s share in minor units. Give it for every friend or for none. | [optional] 
**Items** | Pointer to [**[]CheckoutFriendItem**](CheckoutFriendItem.md) | The cart lines this friend pays for (item_assignment only). Give it for every friend or for none. | [optional] 

## Methods

### NewCheckoutFriend

`func NewCheckoutFriend(name string, mobile string, ) *CheckoutFriend`

NewCheckoutFriend instantiates a new CheckoutFriend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutFriendWithDefaults

`func NewCheckoutFriendWithDefaults() *CheckoutFriend`

NewCheckoutFriendWithDefaults instantiates a new CheckoutFriend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CheckoutFriend) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckoutFriend) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckoutFriend) SetName(v string)`

SetName sets Name field to given value.


### GetMobile

`func (o *CheckoutFriend) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *CheckoutFriend) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *CheckoutFriend) SetMobile(v string)`

SetMobile sets Mobile field to given value.


### GetEmail

`func (o *CheckoutFriend) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CheckoutFriend) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CheckoutFriend) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CheckoutFriend) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CheckoutFriend) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CheckoutFriend) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAmount

`func (o *CheckoutFriend) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckoutFriend) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckoutFriend) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CheckoutFriend) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetItems

`func (o *CheckoutFriend) GetItems() []CheckoutFriendItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CheckoutFriend) GetItemsOk() (*[]CheckoutFriendItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CheckoutFriend) SetItems(v []CheckoutFriendItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CheckoutFriend) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


