# PatchedWritableIKEPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**PatchedWritableIKEPolicyRequestVersion**](PatchedWritableIKEPolicyRequestVersion.md) |  | [optional] 
**Mode** | Pointer to [**PatchedWritableIKEPolicyRequestMode**](PatchedWritableIKEPolicyRequestMode.md) |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableIKEPolicyRequest

`func NewPatchedWritableIKEPolicyRequest() *PatchedWritableIKEPolicyRequest`

NewPatchedWritableIKEPolicyRequest instantiates a new PatchedWritableIKEPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableIKEPolicyRequestWithDefaults

`func NewPatchedWritableIKEPolicyRequestWithDefaults() *PatchedWritableIKEPolicyRequest`

NewPatchedWritableIKEPolicyRequestWithDefaults instantiates a new PatchedWritableIKEPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableIKEPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableIKEPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableIKEPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableIKEPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableIKEPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableIKEPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableIKEPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableIKEPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *PatchedWritableIKEPolicyRequest) GetVersion() PatchedWritableIKEPolicyRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedWritableIKEPolicyRequest) GetVersionOk() (*PatchedWritableIKEPolicyRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedWritableIKEPolicyRequest) SetVersion(v PatchedWritableIKEPolicyRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedWritableIKEPolicyRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMode

`func (o *PatchedWritableIKEPolicyRequest) GetMode() PatchedWritableIKEPolicyRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedWritableIKEPolicyRequest) GetModeOk() (*PatchedWritableIKEPolicyRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedWritableIKEPolicyRequest) SetMode(v PatchedWritableIKEPolicyRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedWritableIKEPolicyRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProposals

`func (o *PatchedWritableIKEPolicyRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *PatchedWritableIKEPolicyRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *PatchedWritableIKEPolicyRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *PatchedWritableIKEPolicyRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPresharedKey

`func (o *PatchedWritableIKEPolicyRequest) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *PatchedWritableIKEPolicyRequest) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *PatchedWritableIKEPolicyRequest) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *PatchedWritableIKEPolicyRequest) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableIKEPolicyRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableIKEPolicyRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableIKEPolicyRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableIKEPolicyRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableIKEPolicyRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableIKEPolicyRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableIKEPolicyRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableIKEPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableIKEPolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableIKEPolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableIKEPolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableIKEPolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


