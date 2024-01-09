# PatchedWritableIKEProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AuthenticationMethod** | Pointer to [**IKEProposalAuthenticationMethodValue**](IKEProposalAuthenticationMethodValue.md) |  | [optional] 
**EncryptionAlgorithm** | Pointer to [**IKEProposalEncryptionAlgorithmValue**](IKEProposalEncryptionAlgorithmValue.md) |  | [optional] 
**AuthenticationAlgorithm** | Pointer to [**PatchedWritableIKEProposalRequestAuthenticationAlgorithm**](PatchedWritableIKEProposalRequestAuthenticationAlgorithm.md) |  | [optional] 
**Group** | Pointer to [**PatchedWritableIKEProposalRequestGroup**](PatchedWritableIKEProposalRequestGroup.md) |  | [optional] 
**SaLifetime** | Pointer to **NullableInt32** | Security association lifetime (in seconds) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableIKEProposalRequest

`func NewPatchedWritableIKEProposalRequest() *PatchedWritableIKEProposalRequest`

NewPatchedWritableIKEProposalRequest instantiates a new PatchedWritableIKEProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableIKEProposalRequestWithDefaults

`func NewPatchedWritableIKEProposalRequestWithDefaults() *PatchedWritableIKEProposalRequest`

NewPatchedWritableIKEProposalRequestWithDefaults instantiates a new PatchedWritableIKEProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableIKEProposalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableIKEProposalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableIKEProposalRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableIKEProposalRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableIKEProposalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableIKEProposalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableIKEProposalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableIKEProposalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *PatchedWritableIKEProposalRequest) GetAuthenticationMethod() IKEProposalAuthenticationMethodValue`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *PatchedWritableIKEProposalRequest) GetAuthenticationMethodOk() (*IKEProposalAuthenticationMethodValue, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *PatchedWritableIKEProposalRequest) SetAuthenticationMethod(v IKEProposalAuthenticationMethodValue)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *PatchedWritableIKEProposalRequest) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *PatchedWritableIKEProposalRequest) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithmValue`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *PatchedWritableIKEProposalRequest) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithmValue, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *PatchedWritableIKEProposalRequest) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithmValue)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *PatchedWritableIKEProposalRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetAuthenticationAlgorithm

`func (o *PatchedWritableIKEProposalRequest) GetAuthenticationAlgorithm() PatchedWritableIKEProposalRequestAuthenticationAlgorithm`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *PatchedWritableIKEProposalRequest) GetAuthenticationAlgorithmOk() (*PatchedWritableIKEProposalRequestAuthenticationAlgorithm, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *PatchedWritableIKEProposalRequest) SetAuthenticationAlgorithm(v PatchedWritableIKEProposalRequestAuthenticationAlgorithm)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *PatchedWritableIKEProposalRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedWritableIKEProposalRequest) GetGroup() PatchedWritableIKEProposalRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableIKEProposalRequest) GetGroupOk() (*PatchedWritableIKEProposalRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableIKEProposalRequest) SetGroup(v PatchedWritableIKEProposalRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableIKEProposalRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSaLifetime

`func (o *PatchedWritableIKEProposalRequest) GetSaLifetime() int32`

GetSaLifetime returns the SaLifetime field if non-nil, zero value otherwise.

### GetSaLifetimeOk

`func (o *PatchedWritableIKEProposalRequest) GetSaLifetimeOk() (*int32, bool)`

GetSaLifetimeOk returns a tuple with the SaLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime

`func (o *PatchedWritableIKEProposalRequest) SetSaLifetime(v int32)`

SetSaLifetime sets SaLifetime field to given value.

### HasSaLifetime

`func (o *PatchedWritableIKEProposalRequest) HasSaLifetime() bool`

HasSaLifetime returns a boolean if a field has been set.

### SetSaLifetimeNil

`func (o *PatchedWritableIKEProposalRequest) SetSaLifetimeNil(b bool)`

 SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil

### UnsetSaLifetime
`func (o *PatchedWritableIKEProposalRequest) UnsetSaLifetime()`

UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
### GetComments

`func (o *PatchedWritableIKEProposalRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableIKEProposalRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableIKEProposalRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableIKEProposalRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableIKEProposalRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableIKEProposalRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableIKEProposalRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableIKEProposalRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableIKEProposalRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableIKEProposalRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableIKEProposalRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableIKEProposalRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


