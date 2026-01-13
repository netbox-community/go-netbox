# BriefModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**ModuleBay** | [**NestedModuleBay**](NestedModuleBay.md) |  | 

## Methods

### NewBriefModule

`func NewBriefModule(id int32, url string, display string, device BriefDevice, moduleBay NestedModuleBay, ) *BriefModule`

NewBriefModule instantiates a new BriefModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefModuleWithDefaults

`func NewBriefModuleWithDefaults() *BriefModule`

NewBriefModuleWithDefaults instantiates a new BriefModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefModule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefModule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefModule) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefModule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefModule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefModule) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefModule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefModule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefModule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *BriefModule) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BriefModule) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BriefModule) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetModuleBay

`func (o *BriefModule) GetModuleBay() NestedModuleBay`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *BriefModule) GetModuleBayOk() (*NestedModuleBay, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *BriefModule) SetModuleBay(v NestedModuleBay)`

SetModuleBay sets ModuleBay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


