# TicketCustomData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**EventStartAt** | **time.Time** |  | 
**Timezone** | **string** |  | 
**TicketClass** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EventEndAt** | Pointer to **time.Time** |  | [optional] 
**DeliveryMode** | Pointer to **string** |  | [optional] 
**VenueName** | Pointer to **string** |  | [optional] 
**VenueAddress** | Pointer to **map[string]interface{}** |  | [optional] 
**Seating** | Pointer to [**TicketSeating**](TicketSeating.md) |  | [optional] 
**TermsUrl** | Pointer to **string** |  | [optional] 
**AttendeeRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewTicketCustomData

`func NewTicketCustomData(eventId string, eventStartAt time.Time, timezone string, ) *TicketCustomData`

NewTicketCustomData instantiates a new TicketCustomData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCustomDataWithDefaults

`func NewTicketCustomDataWithDefaults() *TicketCustomData`

NewTicketCustomDataWithDefaults instantiates a new TicketCustomData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *TicketCustomData) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TicketCustomData) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TicketCustomData) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventStartAt

`func (o *TicketCustomData) GetEventStartAt() time.Time`

GetEventStartAt returns the EventStartAt field if non-nil, zero value otherwise.

### GetEventStartAtOk

`func (o *TicketCustomData) GetEventStartAtOk() (*time.Time, bool)`

GetEventStartAtOk returns a tuple with the EventStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStartAt

`func (o *TicketCustomData) SetEventStartAt(v time.Time)`

SetEventStartAt sets EventStartAt field to given value.


### GetTimezone

`func (o *TicketCustomData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *TicketCustomData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *TicketCustomData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetTicketClass

`func (o *TicketCustomData) GetTicketClass() string`

GetTicketClass returns the TicketClass field if non-nil, zero value otherwise.

### GetTicketClassOk

`func (o *TicketCustomData) GetTicketClassOk() (*string, bool)`

GetTicketClassOk returns a tuple with the TicketClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketClass

`func (o *TicketCustomData) SetTicketClass(v string)`

SetTicketClass sets TicketClass field to given value.

### HasTicketClass

`func (o *TicketCustomData) HasTicketClass() bool`

HasTicketClass returns a boolean if a field has been set.

### GetName

`func (o *TicketCustomData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TicketCustomData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TicketCustomData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TicketCustomData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventEndAt

`func (o *TicketCustomData) GetEventEndAt() time.Time`

GetEventEndAt returns the EventEndAt field if non-nil, zero value otherwise.

### GetEventEndAtOk

`func (o *TicketCustomData) GetEventEndAtOk() (*time.Time, bool)`

GetEventEndAtOk returns a tuple with the EventEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEndAt

`func (o *TicketCustomData) SetEventEndAt(v time.Time)`

SetEventEndAt sets EventEndAt field to given value.

### HasEventEndAt

`func (o *TicketCustomData) HasEventEndAt() bool`

HasEventEndAt returns a boolean if a field has been set.

### GetDeliveryMode

`func (o *TicketCustomData) GetDeliveryMode() string`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *TicketCustomData) GetDeliveryModeOk() (*string, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *TicketCustomData) SetDeliveryMode(v string)`

SetDeliveryMode sets DeliveryMode field to given value.

### HasDeliveryMode

`func (o *TicketCustomData) HasDeliveryMode() bool`

HasDeliveryMode returns a boolean if a field has been set.

### GetVenueName

`func (o *TicketCustomData) GetVenueName() string`

GetVenueName returns the VenueName field if non-nil, zero value otherwise.

### GetVenueNameOk

`func (o *TicketCustomData) GetVenueNameOk() (*string, bool)`

GetVenueNameOk returns a tuple with the VenueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueName

`func (o *TicketCustomData) SetVenueName(v string)`

SetVenueName sets VenueName field to given value.

### HasVenueName

`func (o *TicketCustomData) HasVenueName() bool`

HasVenueName returns a boolean if a field has been set.

### GetVenueAddress

`func (o *TicketCustomData) GetVenueAddress() map[string]interface{}`

GetVenueAddress returns the VenueAddress field if non-nil, zero value otherwise.

### GetVenueAddressOk

`func (o *TicketCustomData) GetVenueAddressOk() (*map[string]interface{}, bool)`

GetVenueAddressOk returns a tuple with the VenueAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueAddress

`func (o *TicketCustomData) SetVenueAddress(v map[string]interface{})`

SetVenueAddress sets VenueAddress field to given value.

### HasVenueAddress

`func (o *TicketCustomData) HasVenueAddress() bool`

HasVenueAddress returns a boolean if a field has been set.

### GetSeating

`func (o *TicketCustomData) GetSeating() TicketSeating`

GetSeating returns the Seating field if non-nil, zero value otherwise.

### GetSeatingOk

`func (o *TicketCustomData) GetSeatingOk() (*TicketSeating, bool)`

GetSeatingOk returns a tuple with the Seating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeating

`func (o *TicketCustomData) SetSeating(v TicketSeating)`

SetSeating sets Seating field to given value.

### HasSeating

`func (o *TicketCustomData) HasSeating() bool`

HasSeating returns a boolean if a field has been set.

### GetTermsUrl

`func (o *TicketCustomData) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *TicketCustomData) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *TicketCustomData) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *TicketCustomData) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### GetAttendeeRequired

`func (o *TicketCustomData) GetAttendeeRequired() bool`

GetAttendeeRequired returns the AttendeeRequired field if non-nil, zero value otherwise.

### GetAttendeeRequiredOk

`func (o *TicketCustomData) GetAttendeeRequiredOk() (*bool, bool)`

GetAttendeeRequiredOk returns a tuple with the AttendeeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeeRequired

`func (o *TicketCustomData) SetAttendeeRequired(v bool)`

SetAttendeeRequired sets AttendeeRequired field to given value.

### HasAttendeeRequired

`func (o *TicketCustomData) HasAttendeeRequired() bool`

HasAttendeeRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


