# BriefPowerPortTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefPowerPortTemplate

`func NewBriefPowerPortTemplate(id int32, url string, display string, name string, ) *BriefPowerPortTemplate`

NewBriefPowerPortTemplate instantiates a new BriefPowerPortTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefPowerPortTemplateWithDefaults

`func NewBriefPowerPortTemplateWithDefaults() *BriefPowerPortTemplate`

NewBriefPowerPortTemplateWithDefaults instantiates a new BriefPowerPortTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefPowerPortTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefPowerPortTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefPowerPortTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefPowerPortTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefPowerPortTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefPowerPortTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefPowerPortTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefPowerPortTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefPowerPortTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefPowerPortTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefPowerPortTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefPowerPortTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefPowerPortTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefPowerPortTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefPowerPortTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefPowerPortTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


