# WritableSiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Full name of the site | 
**Slug** | **string** |  | 
**Status** | Pointer to [**LocationStatusValue**](LocationStatusValue.md) |  | [optional] 
**Region** | Pointer to [**NullableBriefRegionRequest**](BriefRegionRequest.md) |  | [optional] 
**Group** | Pointer to [**NullableBriefSiteGroupRequest**](BriefSiteGroupRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Facility** | Pointer to **string** | Local facility ID or description | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PhysicalAddress** | Pointer to **string** | Physical location of the building | [optional] 
**ShippingAddress** | Pointer to **string** | If different from the physical address | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Asns** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableSiteRequest

`func NewWritableSiteRequest(name string, slug string, ) *WritableSiteRequest`

NewWritableSiteRequest instantiates a new WritableSiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableSiteRequestWithDefaults

`func NewWritableSiteRequestWithDefaults() *WritableSiteRequest`

NewWritableSiteRequestWithDefaults instantiates a new WritableSiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableSiteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableSiteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableSiteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritableSiteRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableSiteRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableSiteRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetStatus

`func (o *WritableSiteRequest) GetStatus() LocationStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableSiteRequest) GetStatusOk() (*LocationStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableSiteRequest) SetStatus(v LocationStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableSiteRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegion

`func (o *WritableSiteRequest) GetRegion() BriefRegionRequest`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *WritableSiteRequest) GetRegionOk() (*BriefRegionRequest, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *WritableSiteRequest) SetRegion(v BriefRegionRequest)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *WritableSiteRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *WritableSiteRequest) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *WritableSiteRequest) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetGroup

`func (o *WritableSiteRequest) GetGroup() BriefSiteGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableSiteRequest) GetGroupOk() (*BriefSiteGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableSiteRequest) SetGroup(v BriefSiteGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WritableSiteRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WritableSiteRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WritableSiteRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetTenant

`func (o *WritableSiteRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableSiteRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableSiteRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableSiteRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableSiteRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableSiteRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetFacility

`func (o *WritableSiteRequest) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *WritableSiteRequest) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *WritableSiteRequest) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *WritableSiteRequest) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetTimeZone

`func (o *WritableSiteRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *WritableSiteRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *WritableSiteRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *WritableSiteRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *WritableSiteRequest) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *WritableSiteRequest) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetDescription

`func (o *WritableSiteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableSiteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableSiteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableSiteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhysicalAddress

`func (o *WritableSiteRequest) GetPhysicalAddress() string`

GetPhysicalAddress returns the PhysicalAddress field if non-nil, zero value otherwise.

### GetPhysicalAddressOk

`func (o *WritableSiteRequest) GetPhysicalAddressOk() (*string, bool)`

GetPhysicalAddressOk returns a tuple with the PhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAddress

`func (o *WritableSiteRequest) SetPhysicalAddress(v string)`

SetPhysicalAddress sets PhysicalAddress field to given value.

### HasPhysicalAddress

`func (o *WritableSiteRequest) HasPhysicalAddress() bool`

HasPhysicalAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *WritableSiteRequest) GetShippingAddress() string`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *WritableSiteRequest) GetShippingAddressOk() (*string, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *WritableSiteRequest) SetShippingAddress(v string)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *WritableSiteRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetLatitude

`func (o *WritableSiteRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *WritableSiteRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *WritableSiteRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *WritableSiteRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *WritableSiteRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *WritableSiteRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *WritableSiteRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *WritableSiteRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *WritableSiteRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *WritableSiteRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *WritableSiteRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *WritableSiteRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetComments

`func (o *WritableSiteRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableSiteRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableSiteRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableSiteRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAsns

`func (o *WritableSiteRequest) GetAsns() []int32`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *WritableSiteRequest) GetAsnsOk() (*[]int32, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *WritableSiteRequest) SetAsns(v []int32)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *WritableSiteRequest) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetTags

`func (o *WritableSiteRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableSiteRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableSiteRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableSiteRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableSiteRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableSiteRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableSiteRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableSiteRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


