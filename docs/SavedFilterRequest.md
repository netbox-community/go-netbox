# SavedFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Parameters** | **map[string]interface{}** |  | 

## Methods

### NewSavedFilterRequest

`func NewSavedFilterRequest(contentTypes []string, name string, slug string, parameters map[string]interface{}, ) *SavedFilterRequest`

NewSavedFilterRequest instantiates a new SavedFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedFilterRequestWithDefaults

`func NewSavedFilterRequestWithDefaults() *SavedFilterRequest`

NewSavedFilterRequestWithDefaults instantiates a new SavedFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *SavedFilterRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *SavedFilterRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *SavedFilterRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetName

`func (o *SavedFilterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedFilterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedFilterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *SavedFilterRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SavedFilterRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SavedFilterRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *SavedFilterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedFilterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedFilterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SavedFilterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUser

`func (o *SavedFilterRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SavedFilterRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SavedFilterRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *SavedFilterRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SavedFilterRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SavedFilterRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetWeight

`func (o *SavedFilterRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SavedFilterRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SavedFilterRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SavedFilterRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetEnabled

`func (o *SavedFilterRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SavedFilterRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SavedFilterRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SavedFilterRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShared

`func (o *SavedFilterRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *SavedFilterRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *SavedFilterRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *SavedFilterRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetParameters

`func (o *SavedFilterRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SavedFilterRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SavedFilterRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


