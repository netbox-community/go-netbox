# PatchedWritableIPSecProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**IPSecProfileModeValue**](IPSecProfileModeValue.md) |  | [optional] 
**IkePolicy** | Pointer to [**BriefIKEPolicyRequest**](BriefIKEPolicyRequest.md) |  | [optional] 
**IpsecPolicy** | Pointer to [**BriefIPSecPolicyRequest**](BriefIPSecPolicyRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableIPSecProfileRequest

`func NewPatchedWritableIPSecProfileRequest() *PatchedWritableIPSecProfileRequest`

NewPatchedWritableIPSecProfileRequest instantiates a new PatchedWritableIPSecProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableIPSecProfileRequestWithDefaults

`func NewPatchedWritableIPSecProfileRequestWithDefaults() *PatchedWritableIPSecProfileRequest`

NewPatchedWritableIPSecProfileRequestWithDefaults instantiates a new PatchedWritableIPSecProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableIPSecProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableIPSecProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableIPSecProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableIPSecProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableIPSecProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableIPSecProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableIPSecProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableIPSecProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *PatchedWritableIPSecProfileRequest) GetMode() IPSecProfileModeValue`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedWritableIPSecProfileRequest) GetModeOk() (*IPSecProfileModeValue, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedWritableIPSecProfileRequest) SetMode(v IPSecProfileModeValue)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedWritableIPSecProfileRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIkePolicy

`func (o *PatchedWritableIPSecProfileRequest) GetIkePolicy() BriefIKEPolicyRequest`

GetIkePolicy returns the IkePolicy field if non-nil, zero value otherwise.

### GetIkePolicyOk

`func (o *PatchedWritableIPSecProfileRequest) GetIkePolicyOk() (*BriefIKEPolicyRequest, bool)`

GetIkePolicyOk returns a tuple with the IkePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePolicy

`func (o *PatchedWritableIPSecProfileRequest) SetIkePolicy(v BriefIKEPolicyRequest)`

SetIkePolicy sets IkePolicy field to given value.

### HasIkePolicy

`func (o *PatchedWritableIPSecProfileRequest) HasIkePolicy() bool`

HasIkePolicy returns a boolean if a field has been set.

### GetIpsecPolicy

`func (o *PatchedWritableIPSecProfileRequest) GetIpsecPolicy() BriefIPSecPolicyRequest`

GetIpsecPolicy returns the IpsecPolicy field if non-nil, zero value otherwise.

### GetIpsecPolicyOk

`func (o *PatchedWritableIPSecProfileRequest) GetIpsecPolicyOk() (*BriefIPSecPolicyRequest, bool)`

GetIpsecPolicyOk returns a tuple with the IpsecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicy

`func (o *PatchedWritableIPSecProfileRequest) SetIpsecPolicy(v BriefIPSecPolicyRequest)`

SetIpsecPolicy sets IpsecPolicy field to given value.

### HasIpsecPolicy

`func (o *PatchedWritableIPSecProfileRequest) HasIpsecPolicy() bool`

HasIpsecPolicy returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableIPSecProfileRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableIPSecProfileRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableIPSecProfileRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableIPSecProfileRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableIPSecProfileRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableIPSecProfileRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableIPSecProfileRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableIPSecProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableIPSecProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableIPSecProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableIPSecProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableIPSecProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


