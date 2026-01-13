# VpnIkeProposalsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AuthenticationMethod** | [**IKEProposalAuthenticationMethodValue**](IKEProposalAuthenticationMethodValue.md) |  | 
**EncryptionAlgorithm** | [**IKEProposalEncryptionAlgorithmValue**](IKEProposalEncryptionAlgorithmValue.md) |  | 
**AuthenticationAlgorithm** | Pointer to [**NullablePatchedWritableIKEProposalRequestAuthenticationAlgorithm**](PatchedWritableIKEProposalRequestAuthenticationAlgorithm.md) |  | [optional] 
**Group** | [**PatchedWritableIKEProposalRequestGroup**](PatchedWritableIKEProposalRequestGroup.md) |  | 
**SaLifetime** | Pointer to **NullableInt32** | Security association lifetime (in seconds) | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnIkeProposalsCreateRequest

`func NewVpnIkeProposalsCreateRequest(name string, authenticationMethod IKEProposalAuthenticationMethodValue, encryptionAlgorithm IKEProposalEncryptionAlgorithmValue, group PatchedWritableIKEProposalRequestGroup, ) *VpnIkeProposalsCreateRequest`

NewVpnIkeProposalsCreateRequest instantiates a new VpnIkeProposalsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnIkeProposalsCreateRequestWithDefaults

`func NewVpnIkeProposalsCreateRequestWithDefaults() *VpnIkeProposalsCreateRequest`

NewVpnIkeProposalsCreateRequestWithDefaults instantiates a new VpnIkeProposalsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnIkeProposalsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnIkeProposalsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnIkeProposalsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VpnIkeProposalsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnIkeProposalsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnIkeProposalsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnIkeProposalsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *VpnIkeProposalsCreateRequest) GetAuthenticationMethod() IKEProposalAuthenticationMethodValue`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *VpnIkeProposalsCreateRequest) GetAuthenticationMethodOk() (*IKEProposalAuthenticationMethodValue, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *VpnIkeProposalsCreateRequest) SetAuthenticationMethod(v IKEProposalAuthenticationMethodValue)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetEncryptionAlgorithm

`func (o *VpnIkeProposalsCreateRequest) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithmValue`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *VpnIkeProposalsCreateRequest) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithmValue, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *VpnIkeProposalsCreateRequest) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithmValue)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetAuthenticationAlgorithm

`func (o *VpnIkeProposalsCreateRequest) GetAuthenticationAlgorithm() PatchedWritableIKEProposalRequestAuthenticationAlgorithm`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *VpnIkeProposalsCreateRequest) GetAuthenticationAlgorithmOk() (*PatchedWritableIKEProposalRequestAuthenticationAlgorithm, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *VpnIkeProposalsCreateRequest) SetAuthenticationAlgorithm(v PatchedWritableIKEProposalRequestAuthenticationAlgorithm)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *VpnIkeProposalsCreateRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### SetAuthenticationAlgorithmNil

`func (o *VpnIkeProposalsCreateRequest) SetAuthenticationAlgorithmNil(b bool)`

 SetAuthenticationAlgorithmNil sets the value for AuthenticationAlgorithm to be an explicit nil

### UnsetAuthenticationAlgorithm
`func (o *VpnIkeProposalsCreateRequest) UnsetAuthenticationAlgorithm()`

UnsetAuthenticationAlgorithm ensures that no value is present for AuthenticationAlgorithm, not even an explicit nil
### GetGroup

`func (o *VpnIkeProposalsCreateRequest) GetGroup() PatchedWritableIKEProposalRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VpnIkeProposalsCreateRequest) GetGroupOk() (*PatchedWritableIKEProposalRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VpnIkeProposalsCreateRequest) SetGroup(v PatchedWritableIKEProposalRequestGroup)`

SetGroup sets Group field to given value.


### GetSaLifetime

`func (o *VpnIkeProposalsCreateRequest) GetSaLifetime() int32`

GetSaLifetime returns the SaLifetime field if non-nil, zero value otherwise.

### GetSaLifetimeOk

`func (o *VpnIkeProposalsCreateRequest) GetSaLifetimeOk() (*int32, bool)`

GetSaLifetimeOk returns a tuple with the SaLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime

`func (o *VpnIkeProposalsCreateRequest) SetSaLifetime(v int32)`

SetSaLifetime sets SaLifetime field to given value.

### HasSaLifetime

`func (o *VpnIkeProposalsCreateRequest) HasSaLifetime() bool`

HasSaLifetime returns a boolean if a field has been set.

### SetSaLifetimeNil

`func (o *VpnIkeProposalsCreateRequest) SetSaLifetimeNil(b bool)`

 SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil

### UnsetSaLifetime
`func (o *VpnIkeProposalsCreateRequest) UnsetSaLifetime()`

UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
### GetOwner

`func (o *VpnIkeProposalsCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VpnIkeProposalsCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VpnIkeProposalsCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VpnIkeProposalsCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VpnIkeProposalsCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VpnIkeProposalsCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VpnIkeProposalsCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VpnIkeProposalsCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VpnIkeProposalsCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VpnIkeProposalsCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VpnIkeProposalsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnIkeProposalsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnIkeProposalsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnIkeProposalsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnIkeProposalsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnIkeProposalsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnIkeProposalsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnIkeProposalsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


