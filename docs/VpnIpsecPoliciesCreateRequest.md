# VpnIpsecPoliciesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PfsGroup** | Pointer to [**NullablePatchedWritableIPSecPolicyRequestPfsGroup**](PatchedWritableIPSecPolicyRequestPfsGroup.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnIpsecPoliciesCreateRequest

`func NewVpnIpsecPoliciesCreateRequest(name string, ) *VpnIpsecPoliciesCreateRequest`

NewVpnIpsecPoliciesCreateRequest instantiates a new VpnIpsecPoliciesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnIpsecPoliciesCreateRequestWithDefaults

`func NewVpnIpsecPoliciesCreateRequestWithDefaults() *VpnIpsecPoliciesCreateRequest`

NewVpnIpsecPoliciesCreateRequestWithDefaults instantiates a new VpnIpsecPoliciesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnIpsecPoliciesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnIpsecPoliciesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnIpsecPoliciesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VpnIpsecPoliciesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnIpsecPoliciesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnIpsecPoliciesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnIpsecPoliciesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProposals

`func (o *VpnIpsecPoliciesCreateRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *VpnIpsecPoliciesCreateRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *VpnIpsecPoliciesCreateRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *VpnIpsecPoliciesCreateRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPfsGroup

`func (o *VpnIpsecPoliciesCreateRequest) GetPfsGroup() PatchedWritableIPSecPolicyRequestPfsGroup`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *VpnIpsecPoliciesCreateRequest) GetPfsGroupOk() (*PatchedWritableIPSecPolicyRequestPfsGroup, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *VpnIpsecPoliciesCreateRequest) SetPfsGroup(v PatchedWritableIPSecPolicyRequestPfsGroup)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *VpnIpsecPoliciesCreateRequest) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### SetPfsGroupNil

`func (o *VpnIpsecPoliciesCreateRequest) SetPfsGroupNil(b bool)`

 SetPfsGroupNil sets the value for PfsGroup to be an explicit nil

### UnsetPfsGroup
`func (o *VpnIpsecPoliciesCreateRequest) UnsetPfsGroup()`

UnsetPfsGroup ensures that no value is present for PfsGroup, not even an explicit nil
### GetOwner

`func (o *VpnIpsecPoliciesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VpnIpsecPoliciesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VpnIpsecPoliciesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VpnIpsecPoliciesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VpnIpsecPoliciesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VpnIpsecPoliciesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VpnIpsecPoliciesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VpnIpsecPoliciesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VpnIpsecPoliciesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VpnIpsecPoliciesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VpnIpsecPoliciesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnIpsecPoliciesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnIpsecPoliciesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnIpsecPoliciesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnIpsecPoliciesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnIpsecPoliciesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnIpsecPoliciesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnIpsecPoliciesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


