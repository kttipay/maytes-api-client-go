# CheckoutCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutUuid** | **string** |  | 
**Status** | **string** |  | 
**CheckoutUrl** | **string** |  | 
**ExpiresAt** | **NullableTime** |  | 

## Methods

### NewCheckoutCreated

`func NewCheckoutCreated(checkoutUuid string, status string, checkoutUrl string, expiresAt NullableTime, ) *CheckoutCreated`

NewCheckoutCreated instantiates a new CheckoutCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutCreatedWithDefaults

`func NewCheckoutCreatedWithDefaults() *CheckoutCreated`

NewCheckoutCreatedWithDefaults instantiates a new CheckoutCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutUuid

`func (o *CheckoutCreated) GetCheckoutUuid() string`

GetCheckoutUuid returns the CheckoutUuid field if non-nil, zero value otherwise.

### GetCheckoutUuidOk

`func (o *CheckoutCreated) GetCheckoutUuidOk() (*string, bool)`

GetCheckoutUuidOk returns a tuple with the CheckoutUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUuid

`func (o *CheckoutCreated) SetCheckoutUuid(v string)`

SetCheckoutUuid sets CheckoutUuid field to given value.


### GetStatus

`func (o *CheckoutCreated) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutCreated) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutCreated) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCheckoutUrl

`func (o *CheckoutCreated) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *CheckoutCreated) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *CheckoutCreated) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.


### GetExpiresAt

`func (o *CheckoutCreated) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutCreated) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutCreated) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *CheckoutCreated) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CheckoutCreated) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


