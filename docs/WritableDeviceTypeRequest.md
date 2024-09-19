# WritableDeviceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | [**BriefManufacturerRequest**](BriefManufacturerRequest.md) |  | 
**DefaultPlatform** | Pointer to [**NullableBriefPlatformRequest**](BriefPlatformRequest.md) |  | [optional] 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**UHeight** | Pointer to **float64** |  | [optional] [default to 1.0]
**ExcludeFromUtilization** | Pointer to **bool** | Devices of this type are excluded when calculating rack utilization. | [optional] 
**IsFullDepth** | Pointer to **bool** | Device consumes both front and rear rack faces. | [optional] 
**SubdeviceRole** | Pointer to [**ParentChildStatus1**](ParentChildStatus1.md) |  | [optional] 
**Airflow** | Pointer to [**DeviceAirflowValue**](DeviceAirflowValue.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**WeightUnit** | Pointer to [**DeviceTypeWeightUnitValue**](DeviceTypeWeightUnitValue.md) |  | [optional] 
**FrontImage** | Pointer to ***os.File** |  | [optional] 
**RearImage** | Pointer to ***os.File** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableDeviceTypeRequest

`func NewWritableDeviceTypeRequest(manufacturer BriefManufacturerRequest, model string, slug string, ) *WritableDeviceTypeRequest`

NewWritableDeviceTypeRequest instantiates a new WritableDeviceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableDeviceTypeRequestWithDefaults

`func NewWritableDeviceTypeRequestWithDefaults() *WritableDeviceTypeRequest`

NewWritableDeviceTypeRequestWithDefaults instantiates a new WritableDeviceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *WritableDeviceTypeRequest) GetManufacturer() BriefManufacturerRequest`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *WritableDeviceTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *WritableDeviceTypeRequest) SetManufacturer(v BriefManufacturerRequest)`

SetManufacturer sets Manufacturer field to given value.


### GetDefaultPlatform

`func (o *WritableDeviceTypeRequest) GetDefaultPlatform() BriefPlatformRequest`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *WritableDeviceTypeRequest) GetDefaultPlatformOk() (*BriefPlatformRequest, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *WritableDeviceTypeRequest) SetDefaultPlatform(v BriefPlatformRequest)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *WritableDeviceTypeRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *WritableDeviceTypeRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *WritableDeviceTypeRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetModel

`func (o *WritableDeviceTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WritableDeviceTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WritableDeviceTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *WritableDeviceTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableDeviceTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableDeviceTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPartNumber

`func (o *WritableDeviceTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *WritableDeviceTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *WritableDeviceTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *WritableDeviceTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetUHeight

`func (o *WritableDeviceTypeRequest) GetUHeight() float64`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *WritableDeviceTypeRequest) GetUHeightOk() (*float64, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *WritableDeviceTypeRequest) SetUHeight(v float64)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *WritableDeviceTypeRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetExcludeFromUtilization

`func (o *WritableDeviceTypeRequest) GetExcludeFromUtilization() bool`

GetExcludeFromUtilization returns the ExcludeFromUtilization field if non-nil, zero value otherwise.

### GetExcludeFromUtilizationOk

`func (o *WritableDeviceTypeRequest) GetExcludeFromUtilizationOk() (*bool, bool)`

GetExcludeFromUtilizationOk returns a tuple with the ExcludeFromUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromUtilization

`func (o *WritableDeviceTypeRequest) SetExcludeFromUtilization(v bool)`

SetExcludeFromUtilization sets ExcludeFromUtilization field to given value.

### HasExcludeFromUtilization

`func (o *WritableDeviceTypeRequest) HasExcludeFromUtilization() bool`

HasExcludeFromUtilization returns a boolean if a field has been set.

### GetIsFullDepth

`func (o *WritableDeviceTypeRequest) GetIsFullDepth() bool`

GetIsFullDepth returns the IsFullDepth field if non-nil, zero value otherwise.

### GetIsFullDepthOk

`func (o *WritableDeviceTypeRequest) GetIsFullDepthOk() (*bool, bool)`

GetIsFullDepthOk returns a tuple with the IsFullDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDepth

`func (o *WritableDeviceTypeRequest) SetIsFullDepth(v bool)`

SetIsFullDepth sets IsFullDepth field to given value.

### HasIsFullDepth

`func (o *WritableDeviceTypeRequest) HasIsFullDepth() bool`

HasIsFullDepth returns a boolean if a field has been set.

### GetSubdeviceRole

`func (o *WritableDeviceTypeRequest) GetSubdeviceRole() ParentChildStatus1`

GetSubdeviceRole returns the SubdeviceRole field if non-nil, zero value otherwise.

### GetSubdeviceRoleOk

`func (o *WritableDeviceTypeRequest) GetSubdeviceRoleOk() (*ParentChildStatus1, bool)`

GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdeviceRole

`func (o *WritableDeviceTypeRequest) SetSubdeviceRole(v ParentChildStatus1)`

SetSubdeviceRole sets SubdeviceRole field to given value.

### HasSubdeviceRole

`func (o *WritableDeviceTypeRequest) HasSubdeviceRole() bool`

HasSubdeviceRole returns a boolean if a field has been set.

### GetAirflow

`func (o *WritableDeviceTypeRequest) GetAirflow() DeviceAirflowValue`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *WritableDeviceTypeRequest) GetAirflowOk() (*DeviceAirflowValue, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *WritableDeviceTypeRequest) SetAirflow(v DeviceAirflowValue)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *WritableDeviceTypeRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetWeight

`func (o *WritableDeviceTypeRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WritableDeviceTypeRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WritableDeviceTypeRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WritableDeviceTypeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *WritableDeviceTypeRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *WritableDeviceTypeRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *WritableDeviceTypeRequest) GetWeightUnit() DeviceTypeWeightUnitValue`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *WritableDeviceTypeRequest) GetWeightUnitOk() (*DeviceTypeWeightUnitValue, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *WritableDeviceTypeRequest) SetWeightUnit(v DeviceTypeWeightUnitValue)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *WritableDeviceTypeRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetFrontImage

`func (o *WritableDeviceTypeRequest) GetFrontImage() *os.File`

GetFrontImage returns the FrontImage field if non-nil, zero value otherwise.

### GetFrontImageOk

`func (o *WritableDeviceTypeRequest) GetFrontImageOk() (**os.File, bool)`

GetFrontImageOk returns a tuple with the FrontImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImage

`func (o *WritableDeviceTypeRequest) SetFrontImage(v *os.File)`

SetFrontImage sets FrontImage field to given value.

### HasFrontImage

`func (o *WritableDeviceTypeRequest) HasFrontImage() bool`

HasFrontImage returns a boolean if a field has been set.

### GetRearImage

`func (o *WritableDeviceTypeRequest) GetRearImage() *os.File`

GetRearImage returns the RearImage field if non-nil, zero value otherwise.

### GetRearImageOk

`func (o *WritableDeviceTypeRequest) GetRearImageOk() (**os.File, bool)`

GetRearImageOk returns a tuple with the RearImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearImage

`func (o *WritableDeviceTypeRequest) SetRearImage(v *os.File)`

SetRearImage sets RearImage field to given value.

### HasRearImage

`func (o *WritableDeviceTypeRequest) HasRearImage() bool`

HasRearImage returns a boolean if a field has been set.

### GetDescription

`func (o *WritableDeviceTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableDeviceTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableDeviceTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableDeviceTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableDeviceTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableDeviceTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableDeviceTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableDeviceTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableDeviceTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableDeviceTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableDeviceTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableDeviceTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableDeviceTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableDeviceTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableDeviceTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableDeviceTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


