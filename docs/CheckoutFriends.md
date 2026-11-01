# CheckoutFriends

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | How many friends will split, not counting the customer. Give this or &#x60;people&#x60;. | [optional] 
**People** | Pointer to [**[]CheckoutFriend**](CheckoutFriend.md) | The friends, optionally with their shares. Give this or &#x60;count&#x60;. | [optional] 

## Methods

### NewCheckoutFriends

`func NewCheckoutFriends() *CheckoutFriends`

NewCheckoutFriends instantiates a new CheckoutFriends object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutFriendsWithDefaults

`func NewCheckoutFriendsWithDefaults() *CheckoutFriends`

NewCheckoutFriendsWithDefaults instantiates a new CheckoutFriends object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CheckoutFriends) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CheckoutFriends) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CheckoutFriends) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CheckoutFriends) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPeople

`func (o *CheckoutFriends) GetPeople() []CheckoutFriend`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *CheckoutFriends) GetPeopleOk() (*[]CheckoutFriend, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *CheckoutFriends) SetPeople(v []CheckoutFriend)`

SetPeople sets People field to given value.

### HasPeople

`func (o *CheckoutFriends) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


