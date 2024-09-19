# IPSecProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | [**IPSecProfileMode**](IPSecProfileMode.md) |  | 
**IkePolicy** | [**BriefIKEPolicy**](BriefIKEPolicy.md) |  | 
**IpsecPolicy** | [**BriefIPSecPolicy**](BriefIPSecPolicy.md) |  | 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewIPSecProfile

`func NewIPSecProfile(id int32, url string, displayUrl string, display string, name string, mode IPSecProfileMode, ikePolicy BriefIKEPolicy, ipsecPolicy BriefIPSecPolicy, created NullableTime, lastUpdated NullableTime, ) *IPSecProfile`

NewIPSecProfile instantiates a new IPSecProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecProfileWithDefaults

`func NewIPSecProfileWithDefaults() *IPSecProfile`

NewIPSecProfileWithDefaults instantiates a new IPSecProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPSecProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSecProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSecProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *IPSecProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPSecProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPSecProfile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *IPSecProfile) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *IPSecProfile) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *IPSecProfile) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *IPSecProfile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPSecProfile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPSecProfile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *IPSecProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPSecProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPSecProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IPSecProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPSecProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPSecProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPSecProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *IPSecProfile) GetMode() IPSecProfileMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IPSecProfile) GetModeOk() (*IPSecProfileMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IPSecProfile) SetMode(v IPSecProfileMode)`

SetMode sets Mode field to given value.


### GetIkePolicy

`func (o *IPSecProfile) GetIkePolicy() BriefIKEPolicy`

GetIkePolicy returns the IkePolicy field if non-nil, zero value otherwise.

### GetIkePolicyOk

`func (o *IPSecProfile) GetIkePolicyOk() (*BriefIKEPolicy, bool)`

GetIkePolicyOk returns a tuple with the IkePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePolicy

`func (o *IPSecProfile) SetIkePolicy(v BriefIKEPolicy)`

SetIkePolicy sets IkePolicy field to given value.


### GetIpsecPolicy

`func (o *IPSecProfile) GetIpsecPolicy() BriefIPSecPolicy`

GetIpsecPolicy returns the IpsecPolicy field if non-nil, zero value otherwise.

### GetIpsecPolicyOk

`func (o *IPSecProfile) GetIpsecPolicyOk() (*BriefIPSecPolicy, bool)`

GetIpsecPolicyOk returns a tuple with the IpsecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicy

`func (o *IPSecProfile) SetIpsecPolicy(v BriefIPSecPolicy)`

SetIpsecPolicy sets IpsecPolicy field to given value.


### GetComments

`func (o *IPSecProfile) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IPSecProfile) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IPSecProfile) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IPSecProfile) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *IPSecProfile) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPSecProfile) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPSecProfile) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPSecProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IPSecProfile) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPSecProfile) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPSecProfile) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPSecProfile) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *IPSecProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPSecProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPSecProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *IPSecProfile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *IPSecProfile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *IPSecProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IPSecProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IPSecProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *IPSecProfile) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *IPSecProfile) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


