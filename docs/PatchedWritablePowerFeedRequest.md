# PatchedWritablePowerFeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerPanel** | Pointer to [**BriefPowerPanelRequest**](BriefPowerPanelRequest.md) |  | [optional] 
**Rack** | Pointer to [**NullableBriefRackRequest**](BriefRackRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PatchedWritablePowerFeedRequestStatus**](PatchedWritablePowerFeedRequestStatus.md) |  | [optional] 
**Type** | Pointer to [**PatchedWritablePowerFeedRequestType**](PatchedWritablePowerFeedRequestType.md) |  | [optional] 
**Supply** | Pointer to [**PatchedWritablePowerFeedRequestSupply**](PatchedWritablePowerFeedRequestSupply.md) |  | [optional] 
**Phase** | Pointer to [**PatchedWritablePowerFeedRequestPhase**](PatchedWritablePowerFeedRequestPhase.md) |  | [optional] 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritablePowerFeedRequest

`func NewPatchedWritablePowerFeedRequest() *PatchedWritablePowerFeedRequest`

NewPatchedWritablePowerFeedRequest instantiates a new PatchedWritablePowerFeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePowerFeedRequestWithDefaults

`func NewPatchedWritablePowerFeedRequestWithDefaults() *PatchedWritablePowerFeedRequest`

NewPatchedWritablePowerFeedRequestWithDefaults instantiates a new PatchedWritablePowerFeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerPanel

`func (o *PatchedWritablePowerFeedRequest) GetPowerPanel() BriefPowerPanelRequest`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *PatchedWritablePowerFeedRequest) GetPowerPanelOk() (*BriefPowerPanelRequest, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *PatchedWritablePowerFeedRequest) SetPowerPanel(v BriefPowerPanelRequest)`

SetPowerPanel sets PowerPanel field to given value.

### HasPowerPanel

`func (o *PatchedWritablePowerFeedRequest) HasPowerPanel() bool`

HasPowerPanel returns a boolean if a field has been set.

### GetRack

`func (o *PatchedWritablePowerFeedRequest) GetRack() BriefRackRequest`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritablePowerFeedRequest) GetRackOk() (*BriefRackRequest, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritablePowerFeedRequest) SetRack(v BriefRackRequest)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritablePowerFeedRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedWritablePowerFeedRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedWritablePowerFeedRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *PatchedWritablePowerFeedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritablePowerFeedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritablePowerFeedRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritablePowerFeedRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritablePowerFeedRequest) GetStatus() PatchedWritablePowerFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritablePowerFeedRequest) GetStatusOk() (*PatchedWritablePowerFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritablePowerFeedRequest) SetStatus(v PatchedWritablePowerFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritablePowerFeedRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritablePowerFeedRequest) GetType() PatchedWritablePowerFeedRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritablePowerFeedRequest) GetTypeOk() (*PatchedWritablePowerFeedRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritablePowerFeedRequest) SetType(v PatchedWritablePowerFeedRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritablePowerFeedRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupply

`func (o *PatchedWritablePowerFeedRequest) GetSupply() PatchedWritablePowerFeedRequestSupply`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *PatchedWritablePowerFeedRequest) GetSupplyOk() (*PatchedWritablePowerFeedRequestSupply, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *PatchedWritablePowerFeedRequest) SetSupply(v PatchedWritablePowerFeedRequestSupply)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *PatchedWritablePowerFeedRequest) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *PatchedWritablePowerFeedRequest) GetPhase() PatchedWritablePowerFeedRequestPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PatchedWritablePowerFeedRequest) GetPhaseOk() (*PatchedWritablePowerFeedRequestPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PatchedWritablePowerFeedRequest) SetPhase(v PatchedWritablePowerFeedRequestPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PatchedWritablePowerFeedRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVoltage

`func (o *PatchedWritablePowerFeedRequest) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *PatchedWritablePowerFeedRequest) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *PatchedWritablePowerFeedRequest) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *PatchedWritablePowerFeedRequest) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *PatchedWritablePowerFeedRequest) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *PatchedWritablePowerFeedRequest) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *PatchedWritablePowerFeedRequest) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *PatchedWritablePowerFeedRequest) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *PatchedWritablePowerFeedRequest) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *PatchedWritablePowerFeedRequest) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *PatchedWritablePowerFeedRequest) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *PatchedWritablePowerFeedRequest) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PatchedWritablePowerFeedRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PatchedWritablePowerFeedRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PatchedWritablePowerFeedRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PatchedWritablePowerFeedRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritablePowerFeedRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritablePowerFeedRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritablePowerFeedRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritablePowerFeedRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritablePowerFeedRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritablePowerFeedRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritablePowerFeedRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritablePowerFeedRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritablePowerFeedRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritablePowerFeedRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetComments

`func (o *PatchedWritablePowerFeedRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritablePowerFeedRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritablePowerFeedRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritablePowerFeedRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritablePowerFeedRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePowerFeedRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePowerFeedRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePowerFeedRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritablePowerFeedRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePowerFeedRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePowerFeedRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePowerFeedRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


