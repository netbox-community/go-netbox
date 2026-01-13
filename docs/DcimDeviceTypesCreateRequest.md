# DcimDeviceTypesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | [**BriefDeviceTypeRequestManufacturer**](BriefDeviceTypeRequestManufacturer.md) |  | 
**DefaultPlatform** | Pointer to [**NullableDeviceTypeRequestDefaultPlatform**](DeviceTypeRequestDefaultPlatform.md) |  | [optional] 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**UHeight** | Pointer to **float64** |  | [optional] [default to 1.0]
**ExcludeFromUtilization** | Pointer to **bool** | Devices of this type are excluded when calculating rack utilization. | [optional] 
**IsFullDepth** | Pointer to **bool** | Device consumes both front and rear rack faces. | [optional] 
**SubdeviceRole** | Pointer to [**NullableParentChildStatus1**](ParentChildStatus1.md) |  | [optional] 
**Airflow** | Pointer to [**NullableDeviceTypeRequestAirflow**](DeviceTypeRequestAirflow.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**WeightUnit** | Pointer to [**NullableDeviceTypeRequestWeightUnit**](DeviceTypeRequestWeightUnit.md) |  | [optional] 
**FrontImage** | Pointer to ***os.File** |  | [optional] 
**RearImage** | Pointer to ***os.File** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimDeviceTypesCreateRequest

`func NewDcimDeviceTypesCreateRequest(manufacturer BriefDeviceTypeRequestManufacturer, model string, slug string, ) *DcimDeviceTypesCreateRequest`

NewDcimDeviceTypesCreateRequest instantiates a new DcimDeviceTypesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimDeviceTypesCreateRequestWithDefaults

`func NewDcimDeviceTypesCreateRequestWithDefaults() *DcimDeviceTypesCreateRequest`

NewDcimDeviceTypesCreateRequestWithDefaults instantiates a new DcimDeviceTypesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *DcimDeviceTypesCreateRequest) GetManufacturer() BriefDeviceTypeRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DcimDeviceTypesCreateRequest) GetManufacturerOk() (*BriefDeviceTypeRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DcimDeviceTypesCreateRequest) SetManufacturer(v BriefDeviceTypeRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetDefaultPlatform

`func (o *DcimDeviceTypesCreateRequest) GetDefaultPlatform() DeviceTypeRequestDefaultPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *DcimDeviceTypesCreateRequest) GetDefaultPlatformOk() (*DeviceTypeRequestDefaultPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *DcimDeviceTypesCreateRequest) SetDefaultPlatform(v DeviceTypeRequestDefaultPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *DcimDeviceTypesCreateRequest) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *DcimDeviceTypesCreateRequest) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *DcimDeviceTypesCreateRequest) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetModel

`func (o *DcimDeviceTypesCreateRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DcimDeviceTypesCreateRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DcimDeviceTypesCreateRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *DcimDeviceTypesCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DcimDeviceTypesCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DcimDeviceTypesCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPartNumber

`func (o *DcimDeviceTypesCreateRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *DcimDeviceTypesCreateRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *DcimDeviceTypesCreateRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *DcimDeviceTypesCreateRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetUHeight

`func (o *DcimDeviceTypesCreateRequest) GetUHeight() float64`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *DcimDeviceTypesCreateRequest) GetUHeightOk() (*float64, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *DcimDeviceTypesCreateRequest) SetUHeight(v float64)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *DcimDeviceTypesCreateRequest) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetExcludeFromUtilization

`func (o *DcimDeviceTypesCreateRequest) GetExcludeFromUtilization() bool`

GetExcludeFromUtilization returns the ExcludeFromUtilization field if non-nil, zero value otherwise.

### GetExcludeFromUtilizationOk

`func (o *DcimDeviceTypesCreateRequest) GetExcludeFromUtilizationOk() (*bool, bool)`

GetExcludeFromUtilizationOk returns a tuple with the ExcludeFromUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromUtilization

`func (o *DcimDeviceTypesCreateRequest) SetExcludeFromUtilization(v bool)`

SetExcludeFromUtilization sets ExcludeFromUtilization field to given value.

### HasExcludeFromUtilization

`func (o *DcimDeviceTypesCreateRequest) HasExcludeFromUtilization() bool`

HasExcludeFromUtilization returns a boolean if a field has been set.

### GetIsFullDepth

`func (o *DcimDeviceTypesCreateRequest) GetIsFullDepth() bool`

GetIsFullDepth returns the IsFullDepth field if non-nil, zero value otherwise.

### GetIsFullDepthOk

`func (o *DcimDeviceTypesCreateRequest) GetIsFullDepthOk() (*bool, bool)`

GetIsFullDepthOk returns a tuple with the IsFullDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDepth

`func (o *DcimDeviceTypesCreateRequest) SetIsFullDepth(v bool)`

SetIsFullDepth sets IsFullDepth field to given value.

### HasIsFullDepth

`func (o *DcimDeviceTypesCreateRequest) HasIsFullDepth() bool`

HasIsFullDepth returns a boolean if a field has been set.

### GetSubdeviceRole

`func (o *DcimDeviceTypesCreateRequest) GetSubdeviceRole() ParentChildStatus1`

GetSubdeviceRole returns the SubdeviceRole field if non-nil, zero value otherwise.

### GetSubdeviceRoleOk

`func (o *DcimDeviceTypesCreateRequest) GetSubdeviceRoleOk() (*ParentChildStatus1, bool)`

GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdeviceRole

`func (o *DcimDeviceTypesCreateRequest) SetSubdeviceRole(v ParentChildStatus1)`

SetSubdeviceRole sets SubdeviceRole field to given value.

### HasSubdeviceRole

`func (o *DcimDeviceTypesCreateRequest) HasSubdeviceRole() bool`

HasSubdeviceRole returns a boolean if a field has been set.

### SetSubdeviceRoleNil

`func (o *DcimDeviceTypesCreateRequest) SetSubdeviceRoleNil(b bool)`

 SetSubdeviceRoleNil sets the value for SubdeviceRole to be an explicit nil

### UnsetSubdeviceRole
`func (o *DcimDeviceTypesCreateRequest) UnsetSubdeviceRole()`

UnsetSubdeviceRole ensures that no value is present for SubdeviceRole, not even an explicit nil
### GetAirflow

`func (o *DcimDeviceTypesCreateRequest) GetAirflow() DeviceTypeRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *DcimDeviceTypesCreateRequest) GetAirflowOk() (*DeviceTypeRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *DcimDeviceTypesCreateRequest) SetAirflow(v DeviceTypeRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *DcimDeviceTypesCreateRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *DcimDeviceTypesCreateRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *DcimDeviceTypesCreateRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetWeight

`func (o *DcimDeviceTypesCreateRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DcimDeviceTypesCreateRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DcimDeviceTypesCreateRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DcimDeviceTypesCreateRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *DcimDeviceTypesCreateRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *DcimDeviceTypesCreateRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *DcimDeviceTypesCreateRequest) GetWeightUnit() DeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *DcimDeviceTypesCreateRequest) GetWeightUnitOk() (*DeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *DcimDeviceTypesCreateRequest) SetWeightUnit(v DeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *DcimDeviceTypesCreateRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *DcimDeviceTypesCreateRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *DcimDeviceTypesCreateRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetFrontImage

`func (o *DcimDeviceTypesCreateRequest) GetFrontImage() *os.File`

GetFrontImage returns the FrontImage field if non-nil, zero value otherwise.

### GetFrontImageOk

`func (o *DcimDeviceTypesCreateRequest) GetFrontImageOk() (**os.File, bool)`

GetFrontImageOk returns a tuple with the FrontImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImage

`func (o *DcimDeviceTypesCreateRequest) SetFrontImage(v *os.File)`

SetFrontImage sets FrontImage field to given value.

### HasFrontImage

`func (o *DcimDeviceTypesCreateRequest) HasFrontImage() bool`

HasFrontImage returns a boolean if a field has been set.

### GetRearImage

`func (o *DcimDeviceTypesCreateRequest) GetRearImage() *os.File`

GetRearImage returns the RearImage field if non-nil, zero value otherwise.

### GetRearImageOk

`func (o *DcimDeviceTypesCreateRequest) GetRearImageOk() (**os.File, bool)`

GetRearImageOk returns a tuple with the RearImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearImage

`func (o *DcimDeviceTypesCreateRequest) SetRearImage(v *os.File)`

SetRearImage sets RearImage field to given value.

### HasRearImage

`func (o *DcimDeviceTypesCreateRequest) HasRearImage() bool`

HasRearImage returns a boolean if a field has been set.

### GetDescription

`func (o *DcimDeviceTypesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimDeviceTypesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimDeviceTypesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimDeviceTypesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *DcimDeviceTypesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimDeviceTypesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimDeviceTypesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimDeviceTypesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimDeviceTypesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimDeviceTypesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *DcimDeviceTypesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DcimDeviceTypesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DcimDeviceTypesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DcimDeviceTypesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *DcimDeviceTypesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimDeviceTypesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimDeviceTypesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimDeviceTypesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimDeviceTypesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimDeviceTypesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimDeviceTypesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimDeviceTypesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


