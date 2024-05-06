# NestedModuleBay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**InstalledModule** | Pointer to [**NullableModuleBayNestedModule**](ModuleBayNestedModule.md) |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewNestedModuleBay

`func NewNestedModuleBay(id int32, url string, display string, name string, ) *NestedModuleBay`

NewNestedModuleBay instantiates a new NestedModuleBay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedModuleBayWithDefaults

`func NewNestedModuleBayWithDefaults() *NestedModuleBay`

NewNestedModuleBayWithDefaults instantiates a new NestedModuleBay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedModuleBay) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedModuleBay) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedModuleBay) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedModuleBay) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedModuleBay) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedModuleBay) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedModuleBay) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedModuleBay) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedModuleBay) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetInstalledModule

`func (o *NestedModuleBay) GetInstalledModule() ModuleBayNestedModule`

GetInstalledModule returns the InstalledModule field if non-nil, zero value otherwise.

### GetInstalledModuleOk

`func (o *NestedModuleBay) GetInstalledModuleOk() (*ModuleBayNestedModule, bool)`

GetInstalledModuleOk returns a tuple with the InstalledModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledModule

`func (o *NestedModuleBay) SetInstalledModule(v ModuleBayNestedModule)`

SetInstalledModule sets InstalledModule field to given value.

### HasInstalledModule

`func (o *NestedModuleBay) HasInstalledModule() bool`

HasInstalledModule returns a boolean if a field has been set.

### SetInstalledModuleNil

`func (o *NestedModuleBay) SetInstalledModuleNil(b bool)`

 SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil

### UnsetInstalledModule
`func (o *NestedModuleBay) UnsetInstalledModule()`

UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
### GetName

`func (o *NestedModuleBay) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedModuleBay) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedModuleBay) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


