# TicketSeating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Section** | Pointer to **string** |  | [optional] 
**Row** | Pointer to **string** |  | [optional] 
**Seat** | Pointer to **string** |  | [optional] 
**Allocation** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTicketSeating

`func NewTicketSeating() *TicketSeating`

NewTicketSeating instantiates a new TicketSeating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketSeatingWithDefaults

`func NewTicketSeatingWithDefaults() *TicketSeating`

NewTicketSeatingWithDefaults instantiates a new TicketSeating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSection

`func (o *TicketSeating) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *TicketSeating) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *TicketSeating) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *TicketSeating) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetRow

`func (o *TicketSeating) GetRow() string`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *TicketSeating) GetRowOk() (*string, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *TicketSeating) SetRow(v string)`

SetRow sets Row field to given value.

### HasRow

`func (o *TicketSeating) HasRow() bool`

HasRow returns a boolean if a field has been set.

### GetSeat

`func (o *TicketSeating) GetSeat() string`

GetSeat returns the Seat field if non-nil, zero value otherwise.

### GetSeatOk

`func (o *TicketSeating) GetSeatOk() (*string, bool)`

GetSeatOk returns a tuple with the Seat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeat

`func (o *TicketSeating) SetSeat(v string)`

SetSeat sets Seat field to given value.

### HasSeat

`func (o *TicketSeating) HasSeat() bool`

HasSeat returns a boolean if a field has been set.

### GetAllocation

`func (o *TicketSeating) GetAllocation() map[string]interface{}`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *TicketSeating) GetAllocationOk() (*map[string]interface{}, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *TicketSeating) SetAllocation(v map[string]interface{})`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *TicketSeating) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


