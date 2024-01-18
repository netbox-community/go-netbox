# NestedModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**NestedDevice**](NestedDevice.md) |  | [readonly] 
**ModuleBay** | [**ModuleNestedModuleBay**](ModuleNestedModuleBay.md) |  | [readonly] 
**ModuleType** | [**NestedModuleType**](NestedModuleType.md) |  | [readonly] 

## Methods

### NewNestedModule

`func NewNestedModule(id int32, url string, display string, device NestedDevice, moduleBay ModuleNestedModuleBay, moduleType NestedModuleType, ) *NestedModule`

NewNestedModule instantiates a new NestedModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedModuleWithDefaults

`func NewNestedModuleWithDefaults() *NestedModule`

NewNestedModuleWithDefaults instantiates a new NestedModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedModule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedModule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedModule) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedModule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedModule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedModule) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedModule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedModule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedModule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *NestedModule) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NestedModule) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NestedModule) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.


### GetModuleBay

`func (o *NestedModule) GetModuleBay() ModuleNestedModuleBay`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *NestedModule) GetModuleBayOk() (*ModuleNestedModuleBay, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *NestedModule) SetModuleBay(v ModuleNestedModuleBay)`

SetModuleBay sets ModuleBay field to given value.


### GetModuleType

`func (o *NestedModule) GetModuleType() NestedModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *NestedModule) GetModuleTypeOk() (*NestedModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *NestedModule) SetModuleType(v NestedModuleType)`

SetModuleType sets ModuleType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


