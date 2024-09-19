# DeviceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Manufacturer** | [**BriefManufacturer**](BriefManufacturer.md) |  | 
**DefaultPlatform** | Pointer to [**NullableBriefPlatform**](BriefPlatform.md) |  | [optional] 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**UHeight** | Pointer to **float64** |  | [optional] [default to 1.0]
**ExcludeFromUtilization** | Pointer to **bool** | Devices of this type are excluded when calculating rack utilization. | [optional] 
**IsFullDepth** | Pointer to **bool** | Device consumes both front and rear rack faces. | [optional] 
**SubdeviceRole** | Pointer to [**NullableDeviceTypeSubdeviceRole**](DeviceTypeSubdeviceRole.md) |  | [optional] 
**Airflow** | Pointer to [**NullableDeviceTypeAirflow**](DeviceTypeAirflow.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**WeightUnit** | Pointer to [**NullableDeviceTypeWeightUnit**](DeviceTypeWeightUnit.md) |  | [optional] 
**FrontImage** | Pointer to **NullableString** |  | [optional] 
**RearImage** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**DeviceCount** | Pointer to **int64** |  | [optional] [readonly] 
**ConsolePortTemplateCount** | **int32** |  | [readonly] 
**ConsoleServerPortTemplateCount** | **int32** |  | [readonly] 
**PowerPortTemplateCount** | **int32** |  | [readonly] 
**PowerOutletTemplateCount** | **int32** |  | [readonly] 
**InterfaceTemplateCount** | **int32** |  | [readonly] 
**FrontPortTemplateCount** | **int32** |  | [readonly] 
**RearPortTemplateCount** | **int32** |  | [readonly] 
**DeviceBayTemplateCount** | **int32** |  | [readonly] 
**ModuleBayTemplateCount** | **int32** |  | [readonly] 
**InventoryItemTemplateCount** | **int32** |  | [readonly] 

## Methods

### NewDeviceType

`func NewDeviceType(id int32, url string, displayUrl string, display string, manufacturer BriefManufacturer, model string, slug string, created NullableTime, lastUpdated NullableTime, consolePortTemplateCount int32, consoleServerPortTemplateCount int32, powerPortTemplateCount int32, powerOutletTemplateCount int32, interfaceTemplateCount int32, frontPortTemplateCount int32, rearPortTemplateCount int32, deviceBayTemplateCount int32, moduleBayTemplateCount int32, inventoryItemTemplateCount int32, ) *DeviceType`

NewDeviceType instantiates a new DeviceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTypeWithDefaults

`func NewDeviceTypeWithDefaults() *DeviceType`

NewDeviceTypeWithDefaults instantiates a new DeviceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *DeviceType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *DeviceType) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *DeviceType) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *DeviceType) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *DeviceType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetManufacturer

`func (o *DeviceType) GetManufacturer() BriefManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DeviceType) GetManufacturerOk() (*BriefManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DeviceType) SetManufacturer(v BriefManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetDefaultPlatform

`func (o *DeviceType) GetDefaultPlatform() BriefPlatform`

GetDefaultPlatform returns the DefaultPlatform field if non-nil, zero value otherwise.

### GetDefaultPlatformOk

`func (o *DeviceType) GetDefaultPlatformOk() (*BriefPlatform, bool)`

GetDefaultPlatformOk returns a tuple with the DefaultPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlatform

`func (o *DeviceType) SetDefaultPlatform(v BriefPlatform)`

SetDefaultPlatform sets DefaultPlatform field to given value.

### HasDefaultPlatform

`func (o *DeviceType) HasDefaultPlatform() bool`

HasDefaultPlatform returns a boolean if a field has been set.

### SetDefaultPlatformNil

`func (o *DeviceType) SetDefaultPlatformNil(b bool)`

 SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil

### UnsetDefaultPlatform
`func (o *DeviceType) UnsetDefaultPlatform()`

UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
### GetModel

`func (o *DeviceType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceType) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *DeviceType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DeviceType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DeviceType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPartNumber

`func (o *DeviceType) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *DeviceType) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *DeviceType) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *DeviceType) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetUHeight

`func (o *DeviceType) GetUHeight() float64`

GetUHeight returns the UHeight field if non-nil, zero value otherwise.

### GetUHeightOk

`func (o *DeviceType) GetUHeightOk() (*float64, bool)`

GetUHeightOk returns a tuple with the UHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUHeight

`func (o *DeviceType) SetUHeight(v float64)`

SetUHeight sets UHeight field to given value.

### HasUHeight

`func (o *DeviceType) HasUHeight() bool`

HasUHeight returns a boolean if a field has been set.

### GetExcludeFromUtilization

`func (o *DeviceType) GetExcludeFromUtilization() bool`

GetExcludeFromUtilization returns the ExcludeFromUtilization field if non-nil, zero value otherwise.

### GetExcludeFromUtilizationOk

`func (o *DeviceType) GetExcludeFromUtilizationOk() (*bool, bool)`

GetExcludeFromUtilizationOk returns a tuple with the ExcludeFromUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromUtilization

`func (o *DeviceType) SetExcludeFromUtilization(v bool)`

SetExcludeFromUtilization sets ExcludeFromUtilization field to given value.

### HasExcludeFromUtilization

`func (o *DeviceType) HasExcludeFromUtilization() bool`

HasExcludeFromUtilization returns a boolean if a field has been set.

### GetIsFullDepth

`func (o *DeviceType) GetIsFullDepth() bool`

GetIsFullDepth returns the IsFullDepth field if non-nil, zero value otherwise.

### GetIsFullDepthOk

`func (o *DeviceType) GetIsFullDepthOk() (*bool, bool)`

GetIsFullDepthOk returns a tuple with the IsFullDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullDepth

`func (o *DeviceType) SetIsFullDepth(v bool)`

SetIsFullDepth sets IsFullDepth field to given value.

### HasIsFullDepth

`func (o *DeviceType) HasIsFullDepth() bool`

HasIsFullDepth returns a boolean if a field has been set.

### GetSubdeviceRole

`func (o *DeviceType) GetSubdeviceRole() DeviceTypeSubdeviceRole`

GetSubdeviceRole returns the SubdeviceRole field if non-nil, zero value otherwise.

### GetSubdeviceRoleOk

`func (o *DeviceType) GetSubdeviceRoleOk() (*DeviceTypeSubdeviceRole, bool)`

GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdeviceRole

`func (o *DeviceType) SetSubdeviceRole(v DeviceTypeSubdeviceRole)`

SetSubdeviceRole sets SubdeviceRole field to given value.

### HasSubdeviceRole

`func (o *DeviceType) HasSubdeviceRole() bool`

HasSubdeviceRole returns a boolean if a field has been set.

### SetSubdeviceRoleNil

`func (o *DeviceType) SetSubdeviceRoleNil(b bool)`

 SetSubdeviceRoleNil sets the value for SubdeviceRole to be an explicit nil

### UnsetSubdeviceRole
`func (o *DeviceType) UnsetSubdeviceRole()`

UnsetSubdeviceRole ensures that no value is present for SubdeviceRole, not even an explicit nil
### GetAirflow

`func (o *DeviceType) GetAirflow() DeviceTypeAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *DeviceType) GetAirflowOk() (*DeviceTypeAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *DeviceType) SetAirflow(v DeviceTypeAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *DeviceType) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *DeviceType) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *DeviceType) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetWeight

`func (o *DeviceType) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DeviceType) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DeviceType) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DeviceType) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *DeviceType) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *DeviceType) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *DeviceType) GetWeightUnit() DeviceTypeWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *DeviceType) GetWeightUnitOk() (*DeviceTypeWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *DeviceType) SetWeightUnit(v DeviceTypeWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *DeviceType) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *DeviceType) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *DeviceType) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetFrontImage

`func (o *DeviceType) GetFrontImage() string`

GetFrontImage returns the FrontImage field if non-nil, zero value otherwise.

### GetFrontImageOk

`func (o *DeviceType) GetFrontImageOk() (*string, bool)`

GetFrontImageOk returns a tuple with the FrontImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImage

`func (o *DeviceType) SetFrontImage(v string)`

SetFrontImage sets FrontImage field to given value.

### HasFrontImage

`func (o *DeviceType) HasFrontImage() bool`

HasFrontImage returns a boolean if a field has been set.

### SetFrontImageNil

`func (o *DeviceType) SetFrontImageNil(b bool)`

 SetFrontImageNil sets the value for FrontImage to be an explicit nil

### UnsetFrontImage
`func (o *DeviceType) UnsetFrontImage()`

UnsetFrontImage ensures that no value is present for FrontImage, not even an explicit nil
### GetRearImage

`func (o *DeviceType) GetRearImage() string`

GetRearImage returns the RearImage field if non-nil, zero value otherwise.

### GetRearImageOk

`func (o *DeviceType) GetRearImageOk() (*string, bool)`

GetRearImageOk returns a tuple with the RearImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearImage

`func (o *DeviceType) SetRearImage(v string)`

SetRearImage sets RearImage field to given value.

### HasRearImage

`func (o *DeviceType) HasRearImage() bool`

HasRearImage returns a boolean if a field has been set.

### SetRearImageNil

`func (o *DeviceType) SetRearImageNil(b bool)`

 SetRearImageNil sets the value for RearImage to be an explicit nil

### UnsetRearImage
`func (o *DeviceType) UnsetRearImage()`

UnsetRearImage ensures that no value is present for RearImage, not even an explicit nil
### GetDescription

`func (o *DeviceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *DeviceType) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DeviceType) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DeviceType) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DeviceType) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *DeviceType) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceType) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceType) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DeviceType) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DeviceType) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DeviceType) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DeviceType) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *DeviceType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceType) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceType) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *DeviceType) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *DeviceType) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *DeviceType) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceType) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceType) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *DeviceType) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *DeviceType) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetDeviceCount

`func (o *DeviceType) GetDeviceCount() int64`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *DeviceType) GetDeviceCountOk() (*int64, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *DeviceType) SetDeviceCount(v int64)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *DeviceType) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetConsolePortTemplateCount

`func (o *DeviceType) GetConsolePortTemplateCount() int32`

GetConsolePortTemplateCount returns the ConsolePortTemplateCount field if non-nil, zero value otherwise.

### GetConsolePortTemplateCountOk

`func (o *DeviceType) GetConsolePortTemplateCountOk() (*int32, bool)`

GetConsolePortTemplateCountOk returns a tuple with the ConsolePortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePortTemplateCount

`func (o *DeviceType) SetConsolePortTemplateCount(v int32)`

SetConsolePortTemplateCount sets ConsolePortTemplateCount field to given value.


### GetConsoleServerPortTemplateCount

`func (o *DeviceType) GetConsoleServerPortTemplateCount() int32`

GetConsoleServerPortTemplateCount returns the ConsoleServerPortTemplateCount field if non-nil, zero value otherwise.

### GetConsoleServerPortTemplateCountOk

`func (o *DeviceType) GetConsoleServerPortTemplateCountOk() (*int32, bool)`

GetConsoleServerPortTemplateCountOk returns a tuple with the ConsoleServerPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPortTemplateCount

`func (o *DeviceType) SetConsoleServerPortTemplateCount(v int32)`

SetConsoleServerPortTemplateCount sets ConsoleServerPortTemplateCount field to given value.


### GetPowerPortTemplateCount

`func (o *DeviceType) GetPowerPortTemplateCount() int32`

GetPowerPortTemplateCount returns the PowerPortTemplateCount field if non-nil, zero value otherwise.

### GetPowerPortTemplateCountOk

`func (o *DeviceType) GetPowerPortTemplateCountOk() (*int32, bool)`

GetPowerPortTemplateCountOk returns a tuple with the PowerPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPortTemplateCount

`func (o *DeviceType) SetPowerPortTemplateCount(v int32)`

SetPowerPortTemplateCount sets PowerPortTemplateCount field to given value.


### GetPowerOutletTemplateCount

`func (o *DeviceType) GetPowerOutletTemplateCount() int32`

GetPowerOutletTemplateCount returns the PowerOutletTemplateCount field if non-nil, zero value otherwise.

### GetPowerOutletTemplateCountOk

`func (o *DeviceType) GetPowerOutletTemplateCountOk() (*int32, bool)`

GetPowerOutletTemplateCountOk returns a tuple with the PowerOutletTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutletTemplateCount

`func (o *DeviceType) SetPowerOutletTemplateCount(v int32)`

SetPowerOutletTemplateCount sets PowerOutletTemplateCount field to given value.


### GetInterfaceTemplateCount

`func (o *DeviceType) GetInterfaceTemplateCount() int32`

GetInterfaceTemplateCount returns the InterfaceTemplateCount field if non-nil, zero value otherwise.

### GetInterfaceTemplateCountOk

`func (o *DeviceType) GetInterfaceTemplateCountOk() (*int32, bool)`

GetInterfaceTemplateCountOk returns a tuple with the InterfaceTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceTemplateCount

`func (o *DeviceType) SetInterfaceTemplateCount(v int32)`

SetInterfaceTemplateCount sets InterfaceTemplateCount field to given value.


### GetFrontPortTemplateCount

`func (o *DeviceType) GetFrontPortTemplateCount() int32`

GetFrontPortTemplateCount returns the FrontPortTemplateCount field if non-nil, zero value otherwise.

### GetFrontPortTemplateCountOk

`func (o *DeviceType) GetFrontPortTemplateCountOk() (*int32, bool)`

GetFrontPortTemplateCountOk returns a tuple with the FrontPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPortTemplateCount

`func (o *DeviceType) SetFrontPortTemplateCount(v int32)`

SetFrontPortTemplateCount sets FrontPortTemplateCount field to given value.


### GetRearPortTemplateCount

`func (o *DeviceType) GetRearPortTemplateCount() int32`

GetRearPortTemplateCount returns the RearPortTemplateCount field if non-nil, zero value otherwise.

### GetRearPortTemplateCountOk

`func (o *DeviceType) GetRearPortTemplateCountOk() (*int32, bool)`

GetRearPortTemplateCountOk returns a tuple with the RearPortTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortTemplateCount

`func (o *DeviceType) SetRearPortTemplateCount(v int32)`

SetRearPortTemplateCount sets RearPortTemplateCount field to given value.


### GetDeviceBayTemplateCount

`func (o *DeviceType) GetDeviceBayTemplateCount() int32`

GetDeviceBayTemplateCount returns the DeviceBayTemplateCount field if non-nil, zero value otherwise.

### GetDeviceBayTemplateCountOk

`func (o *DeviceType) GetDeviceBayTemplateCountOk() (*int32, bool)`

GetDeviceBayTemplateCountOk returns a tuple with the DeviceBayTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceBayTemplateCount

`func (o *DeviceType) SetDeviceBayTemplateCount(v int32)`

SetDeviceBayTemplateCount sets DeviceBayTemplateCount field to given value.


### GetModuleBayTemplateCount

`func (o *DeviceType) GetModuleBayTemplateCount() int32`

GetModuleBayTemplateCount returns the ModuleBayTemplateCount field if non-nil, zero value otherwise.

### GetModuleBayTemplateCountOk

`func (o *DeviceType) GetModuleBayTemplateCountOk() (*int32, bool)`

GetModuleBayTemplateCountOk returns a tuple with the ModuleBayTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayTemplateCount

`func (o *DeviceType) SetModuleBayTemplateCount(v int32)`

SetModuleBayTemplateCount sets ModuleBayTemplateCount field to given value.


### GetInventoryItemTemplateCount

`func (o *DeviceType) GetInventoryItemTemplateCount() int32`

GetInventoryItemTemplateCount returns the InventoryItemTemplateCount field if non-nil, zero value otherwise.

### GetInventoryItemTemplateCountOk

`func (o *DeviceType) GetInventoryItemTemplateCountOk() (*int32, bool)`

GetInventoryItemTemplateCountOk returns a tuple with the InventoryItemTemplateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemTemplateCount

`func (o *DeviceType) SetInventoryItemTemplateCount(v int32)`

SetInventoryItemTemplateCount sets InventoryItemTemplateCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


