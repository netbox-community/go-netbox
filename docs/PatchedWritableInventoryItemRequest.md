# PatchedWritableInventoryItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **int32** |  | [optional] 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Role** | Pointer to **NullableInt32** |  | [optional] 
**Manufacturer** | Pointer to **NullableInt32** |  | [optional] 
**PartId** | Pointer to **string** | Manufacturer-assigned part identifier | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this item | [optional] 
**Discovered** | Pointer to **bool** | This item was automatically discovered | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **NullableString** |  | [optional] 
**ComponentId** | Pointer to **NullableInt64** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableInventoryItemRequest

`func NewPatchedWritableInventoryItemRequest() *PatchedWritableInventoryItemRequest`

NewPatchedWritableInventoryItemRequest instantiates a new PatchedWritableInventoryItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableInventoryItemRequestWithDefaults

`func NewPatchedWritableInventoryItemRequestWithDefaults() *PatchedWritableInventoryItemRequest`

NewPatchedWritableInventoryItemRequestWithDefaults instantiates a new PatchedWritableInventoryItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PatchedWritableInventoryItemRequest) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableInventoryItemRequest) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableInventoryItemRequest) SetDevice(v int32)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableInventoryItemRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetParent

`func (o *PatchedWritableInventoryItemRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedWritableInventoryItemRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedWritableInventoryItemRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedWritableInventoryItemRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedWritableInventoryItemRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedWritableInventoryItemRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *PatchedWritableInventoryItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableInventoryItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableInventoryItemRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableInventoryItemRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableInventoryItemRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableInventoryItemRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableInventoryItemRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableInventoryItemRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableInventoryItemRequest) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableInventoryItemRequest) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableInventoryItemRequest) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableInventoryItemRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritableInventoryItemRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritableInventoryItemRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetManufacturer

`func (o *PatchedWritableInventoryItemRequest) GetManufacturer() int32`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedWritableInventoryItemRequest) GetManufacturerOk() (*int32, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedWritableInventoryItemRequest) SetManufacturer(v int32)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedWritableInventoryItemRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *PatchedWritableInventoryItemRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *PatchedWritableInventoryItemRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetPartId

`func (o *PatchedWritableInventoryItemRequest) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *PatchedWritableInventoryItemRequest) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *PatchedWritableInventoryItemRequest) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *PatchedWritableInventoryItemRequest) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetSerial

`func (o *PatchedWritableInventoryItemRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableInventoryItemRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableInventoryItemRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableInventoryItemRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableInventoryItemRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableInventoryItemRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableInventoryItemRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableInventoryItemRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableInventoryItemRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableInventoryItemRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDiscovered

`func (o *PatchedWritableInventoryItemRequest) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *PatchedWritableInventoryItemRequest) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *PatchedWritableInventoryItemRequest) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *PatchedWritableInventoryItemRequest) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableInventoryItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableInventoryItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableInventoryItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableInventoryItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComponentType

`func (o *PatchedWritableInventoryItemRequest) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *PatchedWritableInventoryItemRequest) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *PatchedWritableInventoryItemRequest) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *PatchedWritableInventoryItemRequest) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### SetComponentTypeNil

`func (o *PatchedWritableInventoryItemRequest) SetComponentTypeNil(b bool)`

 SetComponentTypeNil sets the value for ComponentType to be an explicit nil

### UnsetComponentType
`func (o *PatchedWritableInventoryItemRequest) UnsetComponentType()`

UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
### GetComponentId

`func (o *PatchedWritableInventoryItemRequest) GetComponentId() int64`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *PatchedWritableInventoryItemRequest) GetComponentIdOk() (*int64, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *PatchedWritableInventoryItemRequest) SetComponentId(v int64)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *PatchedWritableInventoryItemRequest) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### SetComponentIdNil

`func (o *PatchedWritableInventoryItemRequest) SetComponentIdNil(b bool)`

 SetComponentIdNil sets the value for ComponentId to be an explicit nil

### UnsetComponentId
`func (o *PatchedWritableInventoryItemRequest) UnsetComponentId()`

UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
### GetTags

`func (o *PatchedWritableInventoryItemRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableInventoryItemRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableInventoryItemRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableInventoryItemRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableInventoryItemRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableInventoryItemRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableInventoryItemRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableInventoryItemRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


