# OAuthTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**RefreshToken** | Pointer to **string** |  | [optional] 
**ExpiresIn** | **float32** |  | 
**TokenType** | **string** |  | 
**Scope** | **string** |  | 

## Methods

### NewOAuthTokenResponse

`func NewOAuthTokenResponse(accessToken string, expiresIn float32, tokenType string, scope string, ) *OAuthTokenResponse`

NewOAuthTokenResponse instantiates a new OAuthTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthTokenResponseWithDefaults

`func NewOAuthTokenResponseWithDefaults() *OAuthTokenResponse`

NewOAuthTokenResponseWithDefaults instantiates a new OAuthTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuthTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *OAuthTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuthTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuthTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuthTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *OAuthTokenResponse) GetExpiresIn() float32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *OAuthTokenResponse) GetExpiresInOk() (*float32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *OAuthTokenResponse) SetExpiresIn(v float32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetTokenType

`func (o *OAuthTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *OAuthTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *OAuthTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetScope

`func (o *OAuthTokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthTokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthTokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


