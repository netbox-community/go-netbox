# NestedModuleBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstalledModule** | Pointer to [**NullableModuleBayNestedModuleRequest**](ModuleBayNestedModuleRequest.md) |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewNestedModuleBayRequest

`func NewNestedModuleBayRequest(name string, ) *NestedModuleBayRequest`

NewNestedModuleBayRequest instantiates a new NestedModuleBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedModuleBayRequestWithDefaults

`func NewNestedModuleBayRequestWithDefaults() *NestedModuleBayRequest`

NewNestedModuleBayRequestWithDefaults instantiates a new NestedModuleBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalledModule

`func (o *NestedModuleBayRequest) GetInstalledModule() ModuleBayNestedModuleRequest`

GetInstalledModule returns the InstalledModule field if non-nil, zero value otherwise.

### GetInstalledModuleOk

`func (o *NestedModuleBayRequest) GetInstalledModuleOk() (*ModuleBayNestedModuleRequest, bool)`

GetInstalledModuleOk returns a tuple with the InstalledModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledModule

`func (o *NestedModuleBayRequest) SetInstalledModule(v ModuleBayNestedModuleRequest)`

SetInstalledModule sets InstalledModule field to given value.

### HasInstalledModule

`func (o *NestedModuleBayRequest) HasInstalledModule() bool`

HasInstalledModule returns a boolean if a field has been set.

### SetInstalledModuleNil

`func (o *NestedModuleBayRequest) SetInstalledModuleNil(b bool)`

 SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil

### UnsetInstalledModule
`func (o *NestedModuleBayRequest) UnsetInstalledModule()`

UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
### GetName

`func (o *NestedModuleBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedModuleBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedModuleBayRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


