# OAuthTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuthTokenRequest

`func NewOAuthTokenRequest(grantType string, clientId string, ) *OAuthTokenRequest`

NewOAuthTokenRequest instantiates a new OAuthTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthTokenRequestWithDefaults

`func NewOAuthTokenRequestWithDefaults() *OAuthTokenRequest`

NewOAuthTokenRequestWithDefaults instantiates a new OAuthTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *OAuthTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuthTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuthTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetClientId

`func (o *OAuthTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuthTokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthTokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthTokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthTokenRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetUsername

`func (o *OAuthTokenRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OAuthTokenRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OAuthTokenRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OAuthTokenRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *OAuthTokenRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OAuthTokenRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OAuthTokenRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OAuthTokenRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OAuthTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuthTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuthTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuthTokenRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScope

`func (o *OAuthTokenRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthTokenRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthTokenRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuthTokenRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


