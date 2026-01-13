# FrontPortMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **int32** |  | 
**RearPort** | **int32** |  | 
**RearPortPosition** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewFrontPortMappingRequest

`func NewFrontPortMappingRequest(position int32, rearPort int32, ) *FrontPortMappingRequest`

NewFrontPortMappingRequest instantiates a new FrontPortMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrontPortMappingRequestWithDefaults

`func NewFrontPortMappingRequestWithDefaults() *FrontPortMappingRequest`

NewFrontPortMappingRequestWithDefaults instantiates a new FrontPortMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *FrontPortMappingRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FrontPortMappingRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FrontPortMappingRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetRearPort

`func (o *FrontPortMappingRequest) GetRearPort() int32`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *FrontPortMappingRequest) GetRearPortOk() (*int32, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *FrontPortMappingRequest) SetRearPort(v int32)`

SetRearPort sets RearPort field to given value.


### GetRearPortPosition

`func (o *FrontPortMappingRequest) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *FrontPortMappingRequest) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *FrontPortMappingRequest) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *FrontPortMappingRequest) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


