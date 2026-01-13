# RearPortMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **int32** |  | 
**FrontPort** | **int32** |  | 
**FrontPortPosition** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewRearPortMappingRequest

`func NewRearPortMappingRequest(position int32, frontPort int32, ) *RearPortMappingRequest`

NewRearPortMappingRequest instantiates a new RearPortMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRearPortMappingRequestWithDefaults

`func NewRearPortMappingRequestWithDefaults() *RearPortMappingRequest`

NewRearPortMappingRequestWithDefaults instantiates a new RearPortMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *RearPortMappingRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *RearPortMappingRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *RearPortMappingRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetFrontPort

`func (o *RearPortMappingRequest) GetFrontPort() int32`

GetFrontPort returns the FrontPort field if non-nil, zero value otherwise.

### GetFrontPortOk

`func (o *RearPortMappingRequest) GetFrontPortOk() (*int32, bool)`

GetFrontPortOk returns a tuple with the FrontPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPort

`func (o *RearPortMappingRequest) SetFrontPort(v int32)`

SetFrontPort sets FrontPort field to given value.


### GetFrontPortPosition

`func (o *RearPortMappingRequest) GetFrontPortPosition() int32`

GetFrontPortPosition returns the FrontPortPosition field if non-nil, zero value otherwise.

### GetFrontPortPositionOk

`func (o *RearPortMappingRequest) GetFrontPortPositionOk() (*int32, bool)`

GetFrontPortPositionOk returns a tuple with the FrontPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPortPosition

`func (o *RearPortMappingRequest) SetFrontPortPosition(v int32)`

SetFrontPortPosition sets FrontPortPosition field to given value.

### HasFrontPortPosition

`func (o *RearPortMappingRequest) HasFrontPortPosition() bool`

HasFrontPortPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


