# PatchedProviderNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**BriefProviderRequest**](BriefProviderRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedProviderNetworkRequest

`func NewPatchedProviderNetworkRequest() *PatchedProviderNetworkRequest`

NewPatchedProviderNetworkRequest instantiates a new PatchedProviderNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProviderNetworkRequestWithDefaults

`func NewPatchedProviderNetworkRequestWithDefaults() *PatchedProviderNetworkRequest`

NewPatchedProviderNetworkRequestWithDefaults instantiates a new PatchedProviderNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *PatchedProviderNetworkRequest) GetProvider() BriefProviderRequest`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PatchedProviderNetworkRequest) GetProviderOk() (*BriefProviderRequest, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PatchedProviderNetworkRequest) SetProvider(v BriefProviderRequest)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PatchedProviderNetworkRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetName

`func (o *PatchedProviderNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedProviderNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedProviderNetworkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedProviderNetworkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceId

`func (o *PatchedProviderNetworkRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PatchedProviderNetworkRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PatchedProviderNetworkRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PatchedProviderNetworkRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedProviderNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedProviderNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedProviderNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedProviderNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedProviderNetworkRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedProviderNetworkRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedProviderNetworkRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedProviderNetworkRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedProviderNetworkRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedProviderNetworkRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedProviderNetworkRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedProviderNetworkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedProviderNetworkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedProviderNetworkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedProviderNetworkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedProviderNetworkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


