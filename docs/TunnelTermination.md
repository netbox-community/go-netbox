# TunnelTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Tunnel** | [**BriefTunnel**](BriefTunnel.md) |  | 
**Role** | [**TunnelTerminationRole**](TunnelTerminationRole.md) |  | 
**TerminationType** | **string** |  | 
**TerminationId** | **NullableInt64** |  | 
**Termination** | **interface{}** |  | [readonly] 
**OutsideIp** | Pointer to [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewTunnelTermination

`func NewTunnelTermination(id int32, url string, displayUrl string, display string, tunnel BriefTunnel, role TunnelTerminationRole, terminationType string, terminationId NullableInt64, termination interface{}, created NullableTime, lastUpdated NullableTime, ) *TunnelTermination`

NewTunnelTermination instantiates a new TunnelTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelTerminationWithDefaults

`func NewTunnelTerminationWithDefaults() *TunnelTermination`

NewTunnelTerminationWithDefaults instantiates a new TunnelTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TunnelTermination) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TunnelTermination) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TunnelTermination) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *TunnelTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TunnelTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TunnelTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *TunnelTermination) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *TunnelTermination) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *TunnelTermination) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *TunnelTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *TunnelTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *TunnelTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetTunnel

`func (o *TunnelTermination) GetTunnel() BriefTunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *TunnelTermination) GetTunnelOk() (*BriefTunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *TunnelTermination) SetTunnel(v BriefTunnel)`

SetTunnel sets Tunnel field to given value.


### GetRole

`func (o *TunnelTermination) GetRole() TunnelTerminationRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TunnelTermination) GetRoleOk() (*TunnelTerminationRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TunnelTermination) SetRole(v TunnelTerminationRole)`

SetRole sets Role field to given value.


### GetTerminationType

`func (o *TunnelTermination) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *TunnelTermination) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *TunnelTermination) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.


### GetTerminationId

`func (o *TunnelTermination) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *TunnelTermination) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *TunnelTermination) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.


### SetTerminationIdNil

`func (o *TunnelTermination) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *TunnelTermination) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetTermination

`func (o *TunnelTermination) GetTermination() interface{}`

GetTermination returns the Termination field if non-nil, zero value otherwise.

### GetTerminationOk

`func (o *TunnelTermination) GetTerminationOk() (*interface{}, bool)`

GetTerminationOk returns a tuple with the Termination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermination

`func (o *TunnelTermination) SetTermination(v interface{})`

SetTermination sets Termination field to given value.


### SetTerminationNil

`func (o *TunnelTermination) SetTerminationNil(b bool)`

 SetTerminationNil sets the value for Termination to be an explicit nil

### UnsetTermination
`func (o *TunnelTermination) UnsetTermination()`

UnsetTermination ensures that no value is present for Termination, not even an explicit nil
### GetOutsideIp

`func (o *TunnelTermination) GetOutsideIp() BriefIPAddress`

GetOutsideIp returns the OutsideIp field if non-nil, zero value otherwise.

### GetOutsideIpOk

`func (o *TunnelTermination) GetOutsideIpOk() (*BriefIPAddress, bool)`

GetOutsideIpOk returns a tuple with the OutsideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideIp

`func (o *TunnelTermination) SetOutsideIp(v BriefIPAddress)`

SetOutsideIp sets OutsideIp field to given value.

### HasOutsideIp

`func (o *TunnelTermination) HasOutsideIp() bool`

HasOutsideIp returns a boolean if a field has been set.

### SetOutsideIpNil

`func (o *TunnelTermination) SetOutsideIpNil(b bool)`

 SetOutsideIpNil sets the value for OutsideIp to be an explicit nil

### UnsetOutsideIp
`func (o *TunnelTermination) UnsetOutsideIp()`

UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
### GetTags

`func (o *TunnelTermination) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TunnelTermination) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TunnelTermination) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TunnelTermination) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *TunnelTermination) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TunnelTermination) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TunnelTermination) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TunnelTermination) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *TunnelTermination) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TunnelTermination) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TunnelTermination) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *TunnelTermination) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *TunnelTermination) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *TunnelTermination) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TunnelTermination) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TunnelTermination) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *TunnelTermination) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *TunnelTermination) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


