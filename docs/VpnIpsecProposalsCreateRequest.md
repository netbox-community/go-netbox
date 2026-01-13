# VpnIpsecProposalsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EncryptionAlgorithm** | Pointer to [**NullableEncryption**](Encryption.md) |  | [optional] 
**AuthenticationAlgorithm** | Pointer to [**NullableAuthentication**](Authentication.md) |  | [optional] 
**SaLifetimeSeconds** | Pointer to **NullableInt32** | Security association lifetime (seconds) | [optional] 
**SaLifetimeData** | Pointer to **NullableInt32** | Security association lifetime (in kilobytes) | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnIpsecProposalsCreateRequest

`func NewVpnIpsecProposalsCreateRequest(name string, ) *VpnIpsecProposalsCreateRequest`

NewVpnIpsecProposalsCreateRequest instantiates a new VpnIpsecProposalsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnIpsecProposalsCreateRequestWithDefaults

`func NewVpnIpsecProposalsCreateRequestWithDefaults() *VpnIpsecProposalsCreateRequest`

NewVpnIpsecProposalsCreateRequestWithDefaults instantiates a new VpnIpsecProposalsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnIpsecProposalsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnIpsecProposalsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnIpsecProposalsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VpnIpsecProposalsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnIpsecProposalsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnIpsecProposalsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnIpsecProposalsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *VpnIpsecProposalsCreateRequest) GetEncryptionAlgorithm() Encryption`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *VpnIpsecProposalsCreateRequest) GetEncryptionAlgorithmOk() (*Encryption, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *VpnIpsecProposalsCreateRequest) SetEncryptionAlgorithm(v Encryption)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *VpnIpsecProposalsCreateRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### SetEncryptionAlgorithmNil

`func (o *VpnIpsecProposalsCreateRequest) SetEncryptionAlgorithmNil(b bool)`

 SetEncryptionAlgorithmNil sets the value for EncryptionAlgorithm to be an explicit nil

### UnsetEncryptionAlgorithm
`func (o *VpnIpsecProposalsCreateRequest) UnsetEncryptionAlgorithm()`

UnsetEncryptionAlgorithm ensures that no value is present for EncryptionAlgorithm, not even an explicit nil
### GetAuthenticationAlgorithm

`func (o *VpnIpsecProposalsCreateRequest) GetAuthenticationAlgorithm() Authentication`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *VpnIpsecProposalsCreateRequest) GetAuthenticationAlgorithmOk() (*Authentication, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *VpnIpsecProposalsCreateRequest) SetAuthenticationAlgorithm(v Authentication)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *VpnIpsecProposalsCreateRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### SetAuthenticationAlgorithmNil

`func (o *VpnIpsecProposalsCreateRequest) SetAuthenticationAlgorithmNil(b bool)`

 SetAuthenticationAlgorithmNil sets the value for AuthenticationAlgorithm to be an explicit nil

### UnsetAuthenticationAlgorithm
`func (o *VpnIpsecProposalsCreateRequest) UnsetAuthenticationAlgorithm()`

UnsetAuthenticationAlgorithm ensures that no value is present for AuthenticationAlgorithm, not even an explicit nil
### GetSaLifetimeSeconds

`func (o *VpnIpsecProposalsCreateRequest) GetSaLifetimeSeconds() int32`

GetSaLifetimeSeconds returns the SaLifetimeSeconds field if non-nil, zero value otherwise.

### GetSaLifetimeSecondsOk

`func (o *VpnIpsecProposalsCreateRequest) GetSaLifetimeSecondsOk() (*int32, bool)`

GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeSeconds

`func (o *VpnIpsecProposalsCreateRequest) SetSaLifetimeSeconds(v int32)`

SetSaLifetimeSeconds sets SaLifetimeSeconds field to given value.

### HasSaLifetimeSeconds

`func (o *VpnIpsecProposalsCreateRequest) HasSaLifetimeSeconds() bool`

HasSaLifetimeSeconds returns a boolean if a field has been set.

### SetSaLifetimeSecondsNil

`func (o *VpnIpsecProposalsCreateRequest) SetSaLifetimeSecondsNil(b bool)`

 SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil

### UnsetSaLifetimeSeconds
`func (o *VpnIpsecProposalsCreateRequest) UnsetSaLifetimeSeconds()`

UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
### GetSaLifetimeData

`func (o *VpnIpsecProposalsCreateRequest) GetSaLifetimeData() int32`

GetSaLifetimeData returns the SaLifetimeData field if non-nil, zero value otherwise.

### GetSaLifetimeDataOk

`func (o *VpnIpsecProposalsCreateRequest) GetSaLifetimeDataOk() (*int32, bool)`

GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeData

`func (o *VpnIpsecProposalsCreateRequest) SetSaLifetimeData(v int32)`

SetSaLifetimeData sets SaLifetimeData field to given value.

### HasSaLifetimeData

`func (o *VpnIpsecProposalsCreateRequest) HasSaLifetimeData() bool`

HasSaLifetimeData returns a boolean if a field has been set.

### SetSaLifetimeDataNil

`func (o *VpnIpsecProposalsCreateRequest) SetSaLifetimeDataNil(b bool)`

 SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil

### UnsetSaLifetimeData
`func (o *VpnIpsecProposalsCreateRequest) UnsetSaLifetimeData()`

UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
### GetOwner

`func (o *VpnIpsecProposalsCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VpnIpsecProposalsCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VpnIpsecProposalsCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VpnIpsecProposalsCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VpnIpsecProposalsCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VpnIpsecProposalsCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VpnIpsecProposalsCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VpnIpsecProposalsCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VpnIpsecProposalsCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VpnIpsecProposalsCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VpnIpsecProposalsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnIpsecProposalsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnIpsecProposalsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnIpsecProposalsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnIpsecProposalsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnIpsecProposalsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnIpsecProposalsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnIpsecProposalsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


