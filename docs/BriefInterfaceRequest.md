# BriefInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefDeviceRequest**](BriefDeviceRequest.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefInterfaceRequest

`func NewBriefInterfaceRequest(device BriefDeviceRequest, name string, ) *BriefInterfaceRequest`

NewBriefInterfaceRequest instantiates a new BriefInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefInterfaceRequestWithDefaults

`func NewBriefInterfaceRequestWithDefaults() *BriefInterfaceRequest`

NewBriefInterfaceRequestWithDefaults instantiates a new BriefInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *BriefInterfaceRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BriefInterfaceRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BriefInterfaceRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.


### GetName

`func (o *BriefInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


