# InventoryItemTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | [**BriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Role** | Pointer to [**NullableBriefInventoryItemRoleRequest**](BriefInventoryItemRoleRequest.md) |  | [optional] 
**Manufacturer** | Pointer to [**NullableBriefManufacturerRequest**](BriefManufacturerRequest.md) |  | [optional] 
**PartId** | Pointer to **string** | Manufacturer-assigned part identifier | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **NullableString** |  | [optional] 
**ComponentId** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewInventoryItemTemplateRequest

`func NewInventoryItemTemplateRequest(deviceType BriefDeviceTypeRequest, name string, ) *InventoryItemTemplateRequest`

NewInventoryItemTemplateRequest instantiates a new InventoryItemTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemTemplateRequestWithDefaults

`func NewInventoryItemTemplateRequestWithDefaults() *InventoryItemTemplateRequest`

NewInventoryItemTemplateRequestWithDefaults instantiates a new InventoryItemTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *InventoryItemTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InventoryItemTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InventoryItemTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.


### GetParent

`func (o *InventoryItemTemplateRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *InventoryItemTemplateRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *InventoryItemTemplateRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *InventoryItemTemplateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *InventoryItemTemplateRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *InventoryItemTemplateRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *InventoryItemTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryItemTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryItemTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *InventoryItemTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InventoryItemTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InventoryItemTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InventoryItemTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRole

`func (o *InventoryItemTemplateRequest) GetRole() BriefInventoryItemRoleRequest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InventoryItemTemplateRequest) GetRoleOk() (*BriefInventoryItemRoleRequest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InventoryItemTemplateRequest) SetRole(v BriefInventoryItemRoleRequest)`

SetRole sets Role field to given value.

### HasRole

`func (o *InventoryItemTemplateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *InventoryItemTemplateRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *InventoryItemTemplateRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetManufacturer

`func (o *InventoryItemTemplateRequest) GetManufacturer() BriefManufacturerRequest`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InventoryItemTemplateRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InventoryItemTemplateRequest) SetManufacturer(v BriefManufacturerRequest)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InventoryItemTemplateRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *InventoryItemTemplateRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *InventoryItemTemplateRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetPartId

`func (o *InventoryItemTemplateRequest) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *InventoryItemTemplateRequest) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *InventoryItemTemplateRequest) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *InventoryItemTemplateRequest) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetDescription

`func (o *InventoryItemTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InventoryItemTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InventoryItemTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InventoryItemTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComponentType

`func (o *InventoryItemTemplateRequest) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *InventoryItemTemplateRequest) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *InventoryItemTemplateRequest) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *InventoryItemTemplateRequest) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### SetComponentTypeNil

`func (o *InventoryItemTemplateRequest) SetComponentTypeNil(b bool)`

 SetComponentTypeNil sets the value for ComponentType to be an explicit nil

### UnsetComponentType
`func (o *InventoryItemTemplateRequest) UnsetComponentType()`

UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
### GetComponentId

`func (o *InventoryItemTemplateRequest) GetComponentId() int64`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *InventoryItemTemplateRequest) GetComponentIdOk() (*int64, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *InventoryItemTemplateRequest) SetComponentId(v int64)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *InventoryItemTemplateRequest) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### SetComponentIdNil

`func (o *InventoryItemTemplateRequest) SetComponentIdNil(b bool)`

 SetComponentIdNil sets the value for ComponentId to be an explicit nil

### UnsetComponentId
`func (o *InventoryItemTemplateRequest) UnsetComponentId()`

UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


