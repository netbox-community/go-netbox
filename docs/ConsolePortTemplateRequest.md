# ConsolePortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**ConsolePortTypeValue**](ConsolePortTypeValue.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewConsolePortTemplateRequest

`func NewConsolePortTemplateRequest(name string, ) *ConsolePortTemplateRequest`

NewConsolePortTemplateRequest instantiates a new ConsolePortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolePortTemplateRequestWithDefaults

`func NewConsolePortTemplateRequestWithDefaults() *ConsolePortTemplateRequest`

NewConsolePortTemplateRequestWithDefaults instantiates a new ConsolePortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *ConsolePortTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ConsolePortTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ConsolePortTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ConsolePortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *ConsolePortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *ConsolePortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *ConsolePortTemplateRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *ConsolePortTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *ConsolePortTemplateRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *ConsolePortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *ConsolePortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *ConsolePortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *ConsolePortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsolePortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsolePortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConsolePortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsolePortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsolePortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsolePortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *ConsolePortTemplateRequest) GetType() ConsolePortTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsolePortTemplateRequest) GetTypeOk() (*ConsolePortTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsolePortTemplateRequest) SetType(v ConsolePortTypeValue)`

SetType sets Type field to given value.

### HasType

`func (o *ConsolePortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ConsolePortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsolePortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsolePortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsolePortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


