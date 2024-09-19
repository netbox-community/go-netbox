# IKEPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Version** | [**IKEPolicyVersion**](IKEPolicyVersion.md) |  | 
**Mode** | Pointer to [**IKEPolicyMode**](IKEPolicyMode.md) |  | [optional] 
**Proposals** | Pointer to [**[]IKEProposal**](IKEProposal.md) |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewIKEPolicy

`func NewIKEPolicy(id int32, url string, displayUrl string, display string, name string, version IKEPolicyVersion, created NullableTime, lastUpdated NullableTime, ) *IKEPolicy`

NewIKEPolicy instantiates a new IKEPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIKEPolicyWithDefaults

`func NewIKEPolicyWithDefaults() *IKEPolicy`

NewIKEPolicyWithDefaults instantiates a new IKEPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IKEPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IKEPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IKEPolicy) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *IKEPolicy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IKEPolicy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IKEPolicy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *IKEPolicy) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *IKEPolicy) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *IKEPolicy) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *IKEPolicy) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IKEPolicy) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IKEPolicy) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *IKEPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IKEPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IKEPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IKEPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IKEPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IKEPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IKEPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *IKEPolicy) GetVersion() IKEPolicyVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IKEPolicy) GetVersionOk() (*IKEPolicyVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IKEPolicy) SetVersion(v IKEPolicyVersion)`

SetVersion sets Version field to given value.


### GetMode

`func (o *IKEPolicy) GetMode() IKEPolicyMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IKEPolicy) GetModeOk() (*IKEPolicyMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IKEPolicy) SetMode(v IKEPolicyMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *IKEPolicy) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProposals

`func (o *IKEPolicy) GetProposals() []IKEProposal`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *IKEPolicy) GetProposalsOk() (*[]IKEProposal, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *IKEPolicy) SetProposals(v []IKEProposal)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *IKEPolicy) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPresharedKey

`func (o *IKEPolicy) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *IKEPolicy) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *IKEPolicy) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *IKEPolicy) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetComments

`func (o *IKEPolicy) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IKEPolicy) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IKEPolicy) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IKEPolicy) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *IKEPolicy) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IKEPolicy) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IKEPolicy) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IKEPolicy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IKEPolicy) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IKEPolicy) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IKEPolicy) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IKEPolicy) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *IKEPolicy) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IKEPolicy) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IKEPolicy) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *IKEPolicy) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *IKEPolicy) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *IKEPolicy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IKEPolicy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IKEPolicy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *IKEPolicy) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *IKEPolicy) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


