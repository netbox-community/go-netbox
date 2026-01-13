# VpnIkePoliciesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**PatchedWritableIKEPolicyRequestVersion**](PatchedWritableIKEPolicyRequestVersion.md) |  | [optional] 
**Mode** | Pointer to [**NullablePatchedWritableIKEPolicyRequestMode**](PatchedWritableIKEPolicyRequestMode.md) |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnIkePoliciesCreateRequest

`func NewVpnIkePoliciesCreateRequest(name string, ) *VpnIkePoliciesCreateRequest`

NewVpnIkePoliciesCreateRequest instantiates a new VpnIkePoliciesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnIkePoliciesCreateRequestWithDefaults

`func NewVpnIkePoliciesCreateRequestWithDefaults() *VpnIkePoliciesCreateRequest`

NewVpnIkePoliciesCreateRequestWithDefaults instantiates a new VpnIkePoliciesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnIkePoliciesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnIkePoliciesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnIkePoliciesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VpnIkePoliciesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnIkePoliciesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnIkePoliciesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnIkePoliciesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *VpnIkePoliciesCreateRequest) GetVersion() PatchedWritableIKEPolicyRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VpnIkePoliciesCreateRequest) GetVersionOk() (*PatchedWritableIKEPolicyRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VpnIkePoliciesCreateRequest) SetVersion(v PatchedWritableIKEPolicyRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VpnIkePoliciesCreateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMode

`func (o *VpnIkePoliciesCreateRequest) GetMode() PatchedWritableIKEPolicyRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VpnIkePoliciesCreateRequest) GetModeOk() (*PatchedWritableIKEPolicyRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VpnIkePoliciesCreateRequest) SetMode(v PatchedWritableIKEPolicyRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VpnIkePoliciesCreateRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *VpnIkePoliciesCreateRequest) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *VpnIkePoliciesCreateRequest) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetProposals

`func (o *VpnIkePoliciesCreateRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *VpnIkePoliciesCreateRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *VpnIkePoliciesCreateRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *VpnIkePoliciesCreateRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPresharedKey

`func (o *VpnIkePoliciesCreateRequest) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *VpnIkePoliciesCreateRequest) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *VpnIkePoliciesCreateRequest) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *VpnIkePoliciesCreateRequest) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetOwner

`func (o *VpnIkePoliciesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VpnIkePoliciesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VpnIkePoliciesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VpnIkePoliciesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VpnIkePoliciesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VpnIkePoliciesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VpnIkePoliciesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VpnIkePoliciesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VpnIkePoliciesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VpnIkePoliciesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VpnIkePoliciesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnIkePoliciesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnIkePoliciesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnIkePoliciesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnIkePoliciesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnIkePoliciesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnIkePoliciesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnIkePoliciesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


