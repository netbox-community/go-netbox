# BriefModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefDeviceRequest**](BriefDeviceRequest.md) |  | 
**ModuleBay** | [**NestedModuleBayRequest**](NestedModuleBayRequest.md) |  | 

## Methods

### NewBriefModuleRequest

`func NewBriefModuleRequest(device BriefDeviceRequest, moduleBay NestedModuleBayRequest, ) *BriefModuleRequest`

NewBriefModuleRequest instantiates a new BriefModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefModuleRequestWithDefaults

`func NewBriefModuleRequestWithDefaults() *BriefModuleRequest`

NewBriefModuleRequestWithDefaults instantiates a new BriefModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *BriefModuleRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BriefModuleRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BriefModuleRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.


### GetModuleBay

`func (o *BriefModuleRequest) GetModuleBay() NestedModuleBayRequest`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *BriefModuleRequest) GetModuleBayOk() (*NestedModuleBayRequest, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *BriefModuleRequest) SetModuleBay(v NestedModuleBayRequest)`

SetModuleBay sets ModuleBay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


