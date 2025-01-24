# VirtualCircuitTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | Pointer to **string** |  | [optional] [readonly] 
**Display** | **string** |  | [readonly] 
**VirtualCircuit** | [**BriefVirtualCircuit**](BriefVirtualCircuit.md) |  | 
**Role** | Pointer to [**TunnelTerminationRole**](TunnelTerminationRole.md) |  | [optional] 
**Interface** | [**BriefInterface**](BriefInterface.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewVirtualCircuitTermination

`func NewVirtualCircuitTermination(id int32, url string, display string, virtualCircuit BriefVirtualCircuit, interface_ BriefInterface, ) *VirtualCircuitTermination`

NewVirtualCircuitTermination instantiates a new VirtualCircuitTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitTerminationWithDefaults

`func NewVirtualCircuitTerminationWithDefaults() *VirtualCircuitTermination`

NewVirtualCircuitTerminationWithDefaults instantiates a new VirtualCircuitTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualCircuitTermination) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCircuitTermination) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCircuitTermination) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *VirtualCircuitTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualCircuitTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualCircuitTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *VirtualCircuitTermination) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *VirtualCircuitTermination) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *VirtualCircuitTermination) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *VirtualCircuitTermination) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *VirtualCircuitTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualCircuitTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualCircuitTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetVirtualCircuit

`func (o *VirtualCircuitTermination) GetVirtualCircuit() BriefVirtualCircuit`

GetVirtualCircuit returns the VirtualCircuit field if non-nil, zero value otherwise.

### GetVirtualCircuitOk

`func (o *VirtualCircuitTermination) GetVirtualCircuitOk() (*BriefVirtualCircuit, bool)`

GetVirtualCircuitOk returns a tuple with the VirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuit

`func (o *VirtualCircuitTermination) SetVirtualCircuit(v BriefVirtualCircuit)`

SetVirtualCircuit sets VirtualCircuit field to given value.


### GetRole

`func (o *VirtualCircuitTermination) GetRole() TunnelTerminationRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VirtualCircuitTermination) GetRoleOk() (*TunnelTerminationRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VirtualCircuitTermination) SetRole(v TunnelTerminationRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *VirtualCircuitTermination) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetInterface

`func (o *VirtualCircuitTermination) GetInterface() BriefInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *VirtualCircuitTermination) GetInterfaceOk() (*BriefInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *VirtualCircuitTermination) SetInterface(v BriefInterface)`

SetInterface sets Interface field to given value.


### GetDescription

`func (o *VirtualCircuitTermination) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuitTermination) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuitTermination) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCircuitTermination) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *VirtualCircuitTermination) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuitTermination) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuitTermination) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualCircuitTermination) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualCircuitTermination) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualCircuitTermination) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualCircuitTermination) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualCircuitTermination) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *VirtualCircuitTermination) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VirtualCircuitTermination) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VirtualCircuitTermination) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *VirtualCircuitTermination) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *VirtualCircuitTermination) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VirtualCircuitTermination) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VirtualCircuitTermination) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualCircuitTermination) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualCircuitTermination) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VirtualCircuitTermination) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *VirtualCircuitTermination) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VirtualCircuitTermination) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


