# FHRPGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Protocol** | [**FHRPGroupProtocol**](FHRPGroupProtocol.md) |  | 
**GroupId** | **int32** |  | 
**AuthType** | Pointer to [**AuthenticationType**](AuthenticationType.md) |  | [optional] 
**AuthKey** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**IpAddresses** | [**[]NestedIPAddress**](NestedIPAddress.md) |  | [readonly] 

## Methods

### NewFHRPGroup

`func NewFHRPGroup(id int32, url string, display string, protocol FHRPGroupProtocol, groupId int32, created NullableTime, lastUpdated NullableTime, ipAddresses []NestedIPAddress, ) *FHRPGroup`

NewFHRPGroup instantiates a new FHRPGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFHRPGroupWithDefaults

`func NewFHRPGroupWithDefaults() *FHRPGroup`

NewFHRPGroupWithDefaults instantiates a new FHRPGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FHRPGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FHRPGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FHRPGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FHRPGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FHRPGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FHRPGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FHRPGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *FHRPGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FHRPGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FHRPGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *FHRPGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *FHRPGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *FHRPGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetProtocol

`func (o *FHRPGroup) GetProtocol() FHRPGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FHRPGroup) GetProtocolOk() (*FHRPGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FHRPGroup) SetProtocol(v FHRPGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetGroupId

`func (o *FHRPGroup) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FHRPGroup) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FHRPGroup) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetAuthType

`func (o *FHRPGroup) GetAuthType() AuthenticationType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *FHRPGroup) GetAuthTypeOk() (*AuthenticationType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *FHRPGroup) SetAuthType(v AuthenticationType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *FHRPGroup) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthKey

`func (o *FHRPGroup) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *FHRPGroup) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *FHRPGroup) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *FHRPGroup) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetDescription

`func (o *FHRPGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FHRPGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FHRPGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FHRPGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *FHRPGroup) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *FHRPGroup) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *FHRPGroup) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *FHRPGroup) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *FHRPGroup) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FHRPGroup) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FHRPGroup) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FHRPGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *FHRPGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *FHRPGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *FHRPGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *FHRPGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *FHRPGroup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FHRPGroup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FHRPGroup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *FHRPGroup) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FHRPGroup) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *FHRPGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FHRPGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FHRPGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *FHRPGroup) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *FHRPGroup) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetIpAddresses

`func (o *FHRPGroup) GetIpAddresses() []NestedIPAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *FHRPGroup) GetIpAddressesOk() (*[]NestedIPAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *FHRPGroup) SetIpAddresses(v []NestedIPAddress)`

SetIpAddresses sets IpAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


