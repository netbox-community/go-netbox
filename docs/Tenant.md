# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Group** | Pointer to [**NullableNestedTenantGroup**](NestedTenantGroup.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**CircuitCount** | **int32** |  | [readonly] 
**DeviceCount** | **int32** |  | [readonly] 
**IpaddressCount** | **int32** |  | [readonly] 
**PrefixCount** | **int32** |  | [readonly] 
**RackCount** | **int32** |  | [readonly] 
**SiteCount** | **int32** |  | [readonly] 
**VirtualmachineCount** | **int32** |  | [readonly] 
**VlanCount** | **int32** |  | [readonly] 
**VrfCount** | **int32** |  | [readonly] 
**ClusterCount** | **int32** |  | [readonly] 

## Methods

### NewTenant

`func NewTenant(id int32, url string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, circuitCount int32, deviceCount int32, ipaddressCount int32, prefixCount int32, rackCount int32, siteCount int32, virtualmachineCount int32, vlanCount int32, vrfCount int32, clusterCount int32, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tenant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Tenant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Tenant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Tenant) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *Tenant) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Tenant) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Tenant) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Tenant) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Tenant) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Tenant) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetGroup

`func (o *Tenant) GetGroup() NestedTenantGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Tenant) GetGroupOk() (*NestedTenantGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Tenant) SetGroup(v NestedTenantGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Tenant) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *Tenant) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *Tenant) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetDescription

`func (o *Tenant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tenant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tenant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tenant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *Tenant) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Tenant) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Tenant) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Tenant) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Tenant) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Tenant) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Tenant) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Tenant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Tenant) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Tenant) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Tenant) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Tenant) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Tenant) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Tenant) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Tenant) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Tenant) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Tenant) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Tenant) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Tenant) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Tenant) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Tenant) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Tenant) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetCircuitCount

`func (o *Tenant) GetCircuitCount() int32`

GetCircuitCount returns the CircuitCount field if non-nil, zero value otherwise.

### GetCircuitCountOk

`func (o *Tenant) GetCircuitCountOk() (*int32, bool)`

GetCircuitCountOk returns a tuple with the CircuitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitCount

`func (o *Tenant) SetCircuitCount(v int32)`

SetCircuitCount sets CircuitCount field to given value.


### GetDeviceCount

`func (o *Tenant) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *Tenant) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *Tenant) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.


### GetIpaddressCount

`func (o *Tenant) GetIpaddressCount() int32`

GetIpaddressCount returns the IpaddressCount field if non-nil, zero value otherwise.

### GetIpaddressCountOk

`func (o *Tenant) GetIpaddressCountOk() (*int32, bool)`

GetIpaddressCountOk returns a tuple with the IpaddressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddressCount

`func (o *Tenant) SetIpaddressCount(v int32)`

SetIpaddressCount sets IpaddressCount field to given value.


### GetPrefixCount

`func (o *Tenant) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *Tenant) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *Tenant) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.


### GetRackCount

`func (o *Tenant) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *Tenant) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *Tenant) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.


### GetSiteCount

`func (o *Tenant) GetSiteCount() int32`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *Tenant) GetSiteCountOk() (*int32, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *Tenant) SetSiteCount(v int32)`

SetSiteCount sets SiteCount field to given value.


### GetVirtualmachineCount

`func (o *Tenant) GetVirtualmachineCount() int32`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *Tenant) GetVirtualmachineCountOk() (*int32, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *Tenant) SetVirtualmachineCount(v int32)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.


### GetVlanCount

`func (o *Tenant) GetVlanCount() int32`

GetVlanCount returns the VlanCount field if non-nil, zero value otherwise.

### GetVlanCountOk

`func (o *Tenant) GetVlanCountOk() (*int32, bool)`

GetVlanCountOk returns a tuple with the VlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCount

`func (o *Tenant) SetVlanCount(v int32)`

SetVlanCount sets VlanCount field to given value.


### GetVrfCount

`func (o *Tenant) GetVrfCount() int32`

GetVrfCount returns the VrfCount field if non-nil, zero value otherwise.

### GetVrfCountOk

`func (o *Tenant) GetVrfCountOk() (*int32, bool)`

GetVrfCountOk returns a tuple with the VrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfCount

`func (o *Tenant) SetVrfCount(v int32)`

SetVrfCount sets VrfCount field to given value.


### GetClusterCount

`func (o *Tenant) GetClusterCount() int32`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *Tenant) GetClusterCountOk() (*int32, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *Tenant) SetClusterCount(v int32)`

SetClusterCount sets ClusterCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


