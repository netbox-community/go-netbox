# VirtualCircuitType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | Pointer to **string** |  | [optional] [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 
**VirtualCircuitCount** | **int64** |  | [readonly] 

## Methods

### NewVirtualCircuitType

`func NewVirtualCircuitType(id int32, url string, display string, name string, slug string, virtualCircuitCount int64, ) *VirtualCircuitType`

NewVirtualCircuitType instantiates a new VirtualCircuitType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitTypeWithDefaults

`func NewVirtualCircuitTypeWithDefaults() *VirtualCircuitType`

NewVirtualCircuitTypeWithDefaults instantiates a new VirtualCircuitType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualCircuitType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCircuitType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCircuitType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *VirtualCircuitType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualCircuitType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualCircuitType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *VirtualCircuitType) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *VirtualCircuitType) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *VirtualCircuitType) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *VirtualCircuitType) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *VirtualCircuitType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualCircuitType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualCircuitType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *VirtualCircuitType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCircuitType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCircuitType) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *VirtualCircuitType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VirtualCircuitType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VirtualCircuitType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *VirtualCircuitType) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *VirtualCircuitType) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *VirtualCircuitType) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *VirtualCircuitType) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualCircuitType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuitType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuitType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCircuitType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *VirtualCircuitType) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuitType) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuitType) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualCircuitType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualCircuitType) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualCircuitType) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualCircuitType) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualCircuitType) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *VirtualCircuitType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VirtualCircuitType) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VirtualCircuitType) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *VirtualCircuitType) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *VirtualCircuitType) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VirtualCircuitType) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VirtualCircuitType) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualCircuitType) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualCircuitType) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VirtualCircuitType) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *VirtualCircuitType) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VirtualCircuitType) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetVirtualCircuitCount

`func (o *VirtualCircuitType) GetVirtualCircuitCount() int64`

GetVirtualCircuitCount returns the VirtualCircuitCount field if non-nil, zero value otherwise.

### GetVirtualCircuitCountOk

`func (o *VirtualCircuitType) GetVirtualCircuitCountOk() (*int64, bool)`

GetVirtualCircuitCountOk returns a tuple with the VirtualCircuitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuitCount

`func (o *VirtualCircuitType) SetVirtualCircuitCount(v int64)`

SetVirtualCircuitCount sets VirtualCircuitCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


