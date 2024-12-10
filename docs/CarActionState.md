# CarActionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CarId** | **int32** |  | 
**Timestamp** | Pointer to **int64** | A Unix timestamp in milliseconds. The timestamp is used to determine the time of creation of an object. | [optional] 
**ActionStatus** | [**CarActionStatus**](CarActionStatus.md) |  | 

## Methods

### NewCarActionState

`func NewCarActionState(carId int32, actionStatus CarActionStatus, ) *CarActionState`

NewCarActionState instantiates a new CarActionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarActionStateWithDefaults

`func NewCarActionStateWithDefaults() *CarActionState`

NewCarActionStateWithDefaults instantiates a new CarActionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CarActionState) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CarActionState) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CarActionState) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CarActionState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCarId

`func (o *CarActionState) GetCarId() int32`

GetCarId returns the CarId field if non-nil, zero value otherwise.

### GetCarIdOk

`func (o *CarActionState) GetCarIdOk() (*int32, bool)`

GetCarIdOk returns a tuple with the CarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarId

`func (o *CarActionState) SetCarId(v int32)`

SetCarId sets CarId field to given value.


### GetTimestamp

`func (o *CarActionState) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CarActionState) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CarActionState) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CarActionState) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetActionStatus

`func (o *CarActionState) GetActionStatus() CarActionStatus`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *CarActionState) GetActionStatusOk() (*CarActionStatus, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *CarActionState) SetActionStatus(v CarActionStatus)`

SetActionStatus sets ActionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


