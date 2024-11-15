# ProviderNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | [**BriefProviderRequest**](BriefProviderRequest.md) |  | 
**Name** | **string** |  | 
**ServiceId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProviderNetworkRequest

`func NewProviderNetworkRequest(provider BriefProviderRequest, name string, ) *ProviderNetworkRequest`

NewProviderNetworkRequest instantiates a new ProviderNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderNetworkRequestWithDefaults

`func NewProviderNetworkRequestWithDefaults() *ProviderNetworkRequest`

NewProviderNetworkRequestWithDefaults instantiates a new ProviderNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ProviderNetworkRequest) GetProvider() BriefProviderRequest`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderNetworkRequest) GetProviderOk() (*BriefProviderRequest, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderNetworkRequest) SetProvider(v BriefProviderRequest)`

SetProvider sets Provider field to given value.


### GetName

`func (o *ProviderNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderNetworkRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServiceId

`func (o *ProviderNetworkRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ProviderNetworkRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ProviderNetworkRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ProviderNetworkRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDescription

`func (o *ProviderNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *ProviderNetworkRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProviderNetworkRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProviderNetworkRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProviderNetworkRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *ProviderNetworkRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProviderNetworkRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProviderNetworkRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProviderNetworkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProviderNetworkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProviderNetworkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProviderNetworkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProviderNetworkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


