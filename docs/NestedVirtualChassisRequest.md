# NestedVirtualChassisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Master** | [**NestedDeviceRequest**](NestedDeviceRequest.md) |  | 

## Methods

### NewNestedVirtualChassisRequest

`func NewNestedVirtualChassisRequest(name string, master NestedDeviceRequest, ) *NestedVirtualChassisRequest`

NewNestedVirtualChassisRequest instantiates a new NestedVirtualChassisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedVirtualChassisRequestWithDefaults

`func NewNestedVirtualChassisRequestWithDefaults() *NestedVirtualChassisRequest`

NewNestedVirtualChassisRequestWithDefaults instantiates a new NestedVirtualChassisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NestedVirtualChassisRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedVirtualChassisRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedVirtualChassisRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMaster

`func (o *NestedVirtualChassisRequest) GetMaster() NestedDeviceRequest`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *NestedVirtualChassisRequest) GetMasterOk() (*NestedDeviceRequest, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *NestedVirtualChassisRequest) SetMaster(v NestedDeviceRequest)`

SetMaster sets Master field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


