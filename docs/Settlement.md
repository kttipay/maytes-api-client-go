# Settlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**MerchantUuid** | **string** |  | 
**PeriodFrom** | **time.Time** |  | 
**PeriodTo** | **time.Time** |  | 
**IdempotencyKey** | **string** |  | 
**MerchantPayableAmount** | **int32** |  | 
**SettledAmount** | **int32** |  | 
**Currency** | **string** |  | 
**Status** | **string** |  | 
**StripeTransferId** | **NullableString** |  | 
**Notes** | **NullableString** |  | 
**SettledAt** | **NullableTime** |  | 
**ExpiresAt** | **NullableTime** |  | 
**ReviewedByUserUuid** | **NullableString** |  | 
**ReviewedAt** | **NullableTime** |  | 
**RejectionReason** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewSettlement

`func NewSettlement(uuid string, merchantUuid string, periodFrom time.Time, periodTo time.Time, idempotencyKey string, merchantPayableAmount int32, settledAmount int32, currency string, status string, stripeTransferId NullableString, notes NullableString, settledAt NullableTime, expiresAt NullableTime, reviewedByUserUuid NullableString, reviewedAt NullableTime, rejectionReason NullableString, createdAt time.Time, updatedAt time.Time, ) *Settlement`

NewSettlement instantiates a new Settlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementWithDefaults

`func NewSettlementWithDefaults() *Settlement`

NewSettlementWithDefaults instantiates a new Settlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Settlement) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Settlement) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Settlement) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetMerchantUuid

`func (o *Settlement) GetMerchantUuid() string`

GetMerchantUuid returns the MerchantUuid field if non-nil, zero value otherwise.

### GetMerchantUuidOk

`func (o *Settlement) GetMerchantUuidOk() (*string, bool)`

GetMerchantUuidOk returns a tuple with the MerchantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUuid

`func (o *Settlement) SetMerchantUuid(v string)`

SetMerchantUuid sets MerchantUuid field to given value.


### GetPeriodFrom

`func (o *Settlement) GetPeriodFrom() time.Time`

GetPeriodFrom returns the PeriodFrom field if non-nil, zero value otherwise.

### GetPeriodFromOk

`func (o *Settlement) GetPeriodFromOk() (*time.Time, bool)`

GetPeriodFromOk returns a tuple with the PeriodFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodFrom

`func (o *Settlement) SetPeriodFrom(v time.Time)`

SetPeriodFrom sets PeriodFrom field to given value.


### GetPeriodTo

`func (o *Settlement) GetPeriodTo() time.Time`

GetPeriodTo returns the PeriodTo field if non-nil, zero value otherwise.

### GetPeriodToOk

`func (o *Settlement) GetPeriodToOk() (*time.Time, bool)`

GetPeriodToOk returns a tuple with the PeriodTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodTo

`func (o *Settlement) SetPeriodTo(v time.Time)`

SetPeriodTo sets PeriodTo field to given value.


### GetIdempotencyKey

`func (o *Settlement) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *Settlement) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *Settlement) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetMerchantPayableAmount

`func (o *Settlement) GetMerchantPayableAmount() int32`

GetMerchantPayableAmount returns the MerchantPayableAmount field if non-nil, zero value otherwise.

### GetMerchantPayableAmountOk

`func (o *Settlement) GetMerchantPayableAmountOk() (*int32, bool)`

GetMerchantPayableAmountOk returns a tuple with the MerchantPayableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantPayableAmount

`func (o *Settlement) SetMerchantPayableAmount(v int32)`

SetMerchantPayableAmount sets MerchantPayableAmount field to given value.


### GetSettledAmount

`func (o *Settlement) GetSettledAmount() int32`

GetSettledAmount returns the SettledAmount field if non-nil, zero value otherwise.

### GetSettledAmountOk

`func (o *Settlement) GetSettledAmountOk() (*int32, bool)`

GetSettledAmountOk returns a tuple with the SettledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAmount

`func (o *Settlement) SetSettledAmount(v int32)`

SetSettledAmount sets SettledAmount field to given value.


### GetCurrency

`func (o *Settlement) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Settlement) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Settlement) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *Settlement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Settlement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Settlement) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStripeTransferId

`func (o *Settlement) GetStripeTransferId() string`

GetStripeTransferId returns the StripeTransferId field if non-nil, zero value otherwise.

### GetStripeTransferIdOk

`func (o *Settlement) GetStripeTransferIdOk() (*string, bool)`

GetStripeTransferIdOk returns a tuple with the StripeTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeTransferId

`func (o *Settlement) SetStripeTransferId(v string)`

SetStripeTransferId sets StripeTransferId field to given value.


### SetStripeTransferIdNil

`func (o *Settlement) SetStripeTransferIdNil(b bool)`

 SetStripeTransferIdNil sets the value for StripeTransferId to be an explicit nil

### UnsetStripeTransferId
`func (o *Settlement) UnsetStripeTransferId()`

UnsetStripeTransferId ensures that no value is present for StripeTransferId, not even an explicit nil
### GetNotes

`func (o *Settlement) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Settlement) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Settlement) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *Settlement) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Settlement) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetSettledAt

`func (o *Settlement) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *Settlement) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *Settlement) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.


### SetSettledAtNil

`func (o *Settlement) SetSettledAtNil(b bool)`

 SetSettledAtNil sets the value for SettledAt to be an explicit nil

### UnsetSettledAt
`func (o *Settlement) UnsetSettledAt()`

UnsetSettledAt ensures that no value is present for SettledAt, not even an explicit nil
### GetExpiresAt

`func (o *Settlement) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Settlement) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Settlement) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *Settlement) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Settlement) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetReviewedByUserUuid

`func (o *Settlement) GetReviewedByUserUuid() string`

GetReviewedByUserUuid returns the ReviewedByUserUuid field if non-nil, zero value otherwise.

### GetReviewedByUserUuidOk

`func (o *Settlement) GetReviewedByUserUuidOk() (*string, bool)`

GetReviewedByUserUuidOk returns a tuple with the ReviewedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedByUserUuid

`func (o *Settlement) SetReviewedByUserUuid(v string)`

SetReviewedByUserUuid sets ReviewedByUserUuid field to given value.


### SetReviewedByUserUuidNil

`func (o *Settlement) SetReviewedByUserUuidNil(b bool)`

 SetReviewedByUserUuidNil sets the value for ReviewedByUserUuid to be an explicit nil

### UnsetReviewedByUserUuid
`func (o *Settlement) UnsetReviewedByUserUuid()`

UnsetReviewedByUserUuid ensures that no value is present for ReviewedByUserUuid, not even an explicit nil
### GetReviewedAt

`func (o *Settlement) GetReviewedAt() time.Time`

GetReviewedAt returns the ReviewedAt field if non-nil, zero value otherwise.

### GetReviewedAtOk

`func (o *Settlement) GetReviewedAtOk() (*time.Time, bool)`

GetReviewedAtOk returns a tuple with the ReviewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAt

`func (o *Settlement) SetReviewedAt(v time.Time)`

SetReviewedAt sets ReviewedAt field to given value.


### SetReviewedAtNil

`func (o *Settlement) SetReviewedAtNil(b bool)`

 SetReviewedAtNil sets the value for ReviewedAt to be an explicit nil

### UnsetReviewedAt
`func (o *Settlement) UnsetReviewedAt()`

UnsetReviewedAt ensures that no value is present for ReviewedAt, not even an explicit nil
### GetRejectionReason

`func (o *Settlement) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *Settlement) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *Settlement) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.


### SetRejectionReasonNil

`func (o *Settlement) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *Settlement) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetCreatedAt

`func (o *Settlement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Settlement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Settlement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Settlement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Settlement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Settlement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


