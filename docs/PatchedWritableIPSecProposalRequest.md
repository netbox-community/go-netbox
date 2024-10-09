# PatchedWritableIPSecProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EncryptionAlgorithm** | Pointer to [**Encryption1**](Encryption1.md) |  | [optional] 
**AuthenticationAlgorithm** | Pointer to [**Authentication1**](Authentication1.md) |  | [optional] 
**SaLifetimeSeconds** | Pointer to **NullableInt32** | Security association lifetime (seconds) | [optional] 
**SaLifetimeData** | Pointer to **NullableInt32** | Security association lifetime (in kilobytes) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableIPSecProposalRequest

`func NewPatchedWritableIPSecProposalRequest() *PatchedWritableIPSecProposalRequest`

NewPatchedWritableIPSecProposalRequest instantiates a new PatchedWritableIPSecProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableIPSecProposalRequestWithDefaults

`func NewPatchedWritableIPSecProposalRequestWithDefaults() *PatchedWritableIPSecProposalRequest`

NewPatchedWritableIPSecProposalRequestWithDefaults instantiates a new PatchedWritableIPSecProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableIPSecProposalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableIPSecProposalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableIPSecProposalRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableIPSecProposalRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableIPSecProposalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableIPSecProposalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableIPSecProposalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableIPSecProposalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *PatchedWritableIPSecProposalRequest) GetEncryptionAlgorithm() Encryption1`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *PatchedWritableIPSecProposalRequest) GetEncryptionAlgorithmOk() (*Encryption1, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *PatchedWritableIPSecProposalRequest) SetEncryptionAlgorithm(v Encryption1)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *PatchedWritableIPSecProposalRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetAuthenticationAlgorithm

`func (o *PatchedWritableIPSecProposalRequest) GetAuthenticationAlgorithm() Authentication1`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *PatchedWritableIPSecProposalRequest) GetAuthenticationAlgorithmOk() (*Authentication1, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *PatchedWritableIPSecProposalRequest) SetAuthenticationAlgorithm(v Authentication1)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *PatchedWritableIPSecProposalRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### GetSaLifetimeSeconds

`func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeSeconds() int32`

GetSaLifetimeSeconds returns the SaLifetimeSeconds field if non-nil, zero value otherwise.

### GetSaLifetimeSecondsOk

`func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeSecondsOk() (*int32, bool)`

GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeSeconds

`func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeSeconds(v int32)`

SetSaLifetimeSeconds sets SaLifetimeSeconds field to given value.

### HasSaLifetimeSeconds

`func (o *PatchedWritableIPSecProposalRequest) HasSaLifetimeSeconds() bool`

HasSaLifetimeSeconds returns a boolean if a field has been set.

### SetSaLifetimeSecondsNil

`func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeSecondsNil(b bool)`

 SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil

### UnsetSaLifetimeSeconds
`func (o *PatchedWritableIPSecProposalRequest) UnsetSaLifetimeSeconds()`

UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
### GetSaLifetimeData

`func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeData() int32`

GetSaLifetimeData returns the SaLifetimeData field if non-nil, zero value otherwise.

### GetSaLifetimeDataOk

`func (o *PatchedWritableIPSecProposalRequest) GetSaLifetimeDataOk() (*int32, bool)`

GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeData

`func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeData(v int32)`

SetSaLifetimeData sets SaLifetimeData field to given value.

### HasSaLifetimeData

`func (o *PatchedWritableIPSecProposalRequest) HasSaLifetimeData() bool`

HasSaLifetimeData returns a boolean if a field has been set.

### SetSaLifetimeDataNil

`func (o *PatchedWritableIPSecProposalRequest) SetSaLifetimeDataNil(b bool)`

 SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil

### UnsetSaLifetimeData
`func (o *PatchedWritableIPSecProposalRequest) UnsetSaLifetimeData()`

UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
### GetComments

`func (o *PatchedWritableIPSecProposalRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableIPSecProposalRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableIPSecProposalRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableIPSecProposalRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableIPSecProposalRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableIPSecProposalRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableIPSecProposalRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableIPSecProposalRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableIPSecProposalRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableIPSecProposalRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableIPSecProposalRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableIPSecProposalRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


