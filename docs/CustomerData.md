# CustomerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Mobile** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCustomerData

`func NewCustomerData() *CustomerData`

NewCustomerData instantiates a new CustomerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDataWithDefaults

`func NewCustomerDataWithDefaults() *CustomerData`

NewCustomerDataWithDefaults instantiates a new CustomerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CustomerData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomerData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMobile

`func (o *CustomerData) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *CustomerData) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *CustomerData) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *CustomerData) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### SetMobileNil

`func (o *CustomerData) SetMobileNil(b bool)`

 SetMobileNil sets the value for Mobile to be an explicit nil

### UnsetMobile
`func (o *CustomerData) UnsetMobile()`

UnsetMobile ensures that no value is present for Mobile, not even an explicit nil
### GetEmail

`func (o *CustomerData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CustomerData) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CustomerData) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


