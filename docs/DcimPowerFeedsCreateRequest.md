# DcimPowerFeedsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerPanel** | [**PatchedWritablePowerFeedRequestPowerPanel**](PatchedWritablePowerFeedRequestPowerPanel.md) |  | 
**Rack** | Pointer to [**NullableDeviceWithConfigContextRequestRack**](DeviceWithConfigContextRequestRack.md) |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to [**PatchedWritablePowerFeedRequestStatus**](PatchedWritablePowerFeedRequestStatus.md) |  | [optional] 
**Type** | Pointer to [**PatchedWritablePowerFeedRequestType**](PatchedWritablePowerFeedRequestType.md) |  | [optional] 
**Supply** | Pointer to [**PatchedWritablePowerFeedRequestSupply**](PatchedWritablePowerFeedRequestSupply.md) |  | [optional] 
**Phase** | Pointer to [**PatchedWritablePowerFeedRequestPhase**](PatchedWritablePowerFeedRequestPhase.md) |  | [optional] 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimPowerFeedsCreateRequest

`func NewDcimPowerFeedsCreateRequest(powerPanel PatchedWritablePowerFeedRequestPowerPanel, name string, ) *DcimPowerFeedsCreateRequest`

NewDcimPowerFeedsCreateRequest instantiates a new DcimPowerFeedsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimPowerFeedsCreateRequestWithDefaults

`func NewDcimPowerFeedsCreateRequestWithDefaults() *DcimPowerFeedsCreateRequest`

NewDcimPowerFeedsCreateRequestWithDefaults instantiates a new DcimPowerFeedsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerPanel

`func (o *DcimPowerFeedsCreateRequest) GetPowerPanel() PatchedWritablePowerFeedRequestPowerPanel`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *DcimPowerFeedsCreateRequest) GetPowerPanelOk() (*PatchedWritablePowerFeedRequestPowerPanel, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *DcimPowerFeedsCreateRequest) SetPowerPanel(v PatchedWritablePowerFeedRequestPowerPanel)`

SetPowerPanel sets PowerPanel field to given value.


### GetRack

`func (o *DcimPowerFeedsCreateRequest) GetRack() DeviceWithConfigContextRequestRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *DcimPowerFeedsCreateRequest) GetRackOk() (*DeviceWithConfigContextRequestRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *DcimPowerFeedsCreateRequest) SetRack(v DeviceWithConfigContextRequestRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *DcimPowerFeedsCreateRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *DcimPowerFeedsCreateRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *DcimPowerFeedsCreateRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *DcimPowerFeedsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimPowerFeedsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimPowerFeedsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DcimPowerFeedsCreateRequest) GetStatus() PatchedWritablePowerFeedRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DcimPowerFeedsCreateRequest) GetStatusOk() (*PatchedWritablePowerFeedRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DcimPowerFeedsCreateRequest) SetStatus(v PatchedWritablePowerFeedRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DcimPowerFeedsCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DcimPowerFeedsCreateRequest) GetType() PatchedWritablePowerFeedRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcimPowerFeedsCreateRequest) GetTypeOk() (*PatchedWritablePowerFeedRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcimPowerFeedsCreateRequest) SetType(v PatchedWritablePowerFeedRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *DcimPowerFeedsCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupply

`func (o *DcimPowerFeedsCreateRequest) GetSupply() PatchedWritablePowerFeedRequestSupply`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *DcimPowerFeedsCreateRequest) GetSupplyOk() (*PatchedWritablePowerFeedRequestSupply, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *DcimPowerFeedsCreateRequest) SetSupply(v PatchedWritablePowerFeedRequestSupply)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *DcimPowerFeedsCreateRequest) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *DcimPowerFeedsCreateRequest) GetPhase() PatchedWritablePowerFeedRequestPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *DcimPowerFeedsCreateRequest) GetPhaseOk() (*PatchedWritablePowerFeedRequestPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *DcimPowerFeedsCreateRequest) SetPhase(v PatchedWritablePowerFeedRequestPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *DcimPowerFeedsCreateRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVoltage

`func (o *DcimPowerFeedsCreateRequest) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *DcimPowerFeedsCreateRequest) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *DcimPowerFeedsCreateRequest) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *DcimPowerFeedsCreateRequest) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *DcimPowerFeedsCreateRequest) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *DcimPowerFeedsCreateRequest) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *DcimPowerFeedsCreateRequest) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *DcimPowerFeedsCreateRequest) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *DcimPowerFeedsCreateRequest) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *DcimPowerFeedsCreateRequest) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *DcimPowerFeedsCreateRequest) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *DcimPowerFeedsCreateRequest) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetMarkConnected

`func (o *DcimPowerFeedsCreateRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *DcimPowerFeedsCreateRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *DcimPowerFeedsCreateRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *DcimPowerFeedsCreateRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetDescription

`func (o *DcimPowerFeedsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimPowerFeedsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimPowerFeedsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimPowerFeedsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenant

`func (o *DcimPowerFeedsCreateRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DcimPowerFeedsCreateRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DcimPowerFeedsCreateRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DcimPowerFeedsCreateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *DcimPowerFeedsCreateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *DcimPowerFeedsCreateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetOwner

`func (o *DcimPowerFeedsCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimPowerFeedsCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimPowerFeedsCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimPowerFeedsCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimPowerFeedsCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimPowerFeedsCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *DcimPowerFeedsCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DcimPowerFeedsCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DcimPowerFeedsCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DcimPowerFeedsCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *DcimPowerFeedsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimPowerFeedsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimPowerFeedsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimPowerFeedsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimPowerFeedsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimPowerFeedsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimPowerFeedsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimPowerFeedsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


