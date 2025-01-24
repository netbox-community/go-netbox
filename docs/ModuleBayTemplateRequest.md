# ModuleBayTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Position** | Pointer to **string** | Identifier to reference when renaming installed components | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewModuleBayTemplateRequest

`func NewModuleBayTemplateRequest(name string, ) *ModuleBayTemplateRequest`

NewModuleBayTemplateRequest instantiates a new ModuleBayTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleBayTemplateRequestWithDefaults

`func NewModuleBayTemplateRequestWithDefaults() *ModuleBayTemplateRequest`

NewModuleBayTemplateRequestWithDefaults instantiates a new ModuleBayTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *ModuleBayTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ModuleBayTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ModuleBayTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ModuleBayTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *ModuleBayTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *ModuleBayTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *ModuleBayTemplateRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *ModuleBayTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *ModuleBayTemplateRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *ModuleBayTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *ModuleBayTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *ModuleBayTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *ModuleBayTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModuleBayTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModuleBayTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ModuleBayTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModuleBayTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModuleBayTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ModuleBayTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPosition

`func (o *ModuleBayTemplateRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModuleBayTemplateRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModuleBayTemplateRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModuleBayTemplateRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetDescription

`func (o *ModuleBayTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleBayTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleBayTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleBayTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


