# IPSecPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PfsGroup** | Pointer to [**IKEProposalGroupValue**](IKEProposalGroupValue.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIPSecPolicyRequest

`func NewIPSecPolicyRequest(name string, ) *IPSecPolicyRequest`

NewIPSecPolicyRequest instantiates a new IPSecPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecPolicyRequestWithDefaults

`func NewIPSecPolicyRequestWithDefaults() *IPSecPolicyRequest`

NewIPSecPolicyRequestWithDefaults instantiates a new IPSecPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPSecPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPSecPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPSecPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IPSecPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPSecPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPSecPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPSecPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProposals

`func (o *IPSecPolicyRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *IPSecPolicyRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *IPSecPolicyRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *IPSecPolicyRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPfsGroup

`func (o *IPSecPolicyRequest) GetPfsGroup() IKEProposalGroupValue`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *IPSecPolicyRequest) GetPfsGroupOk() (*IKEProposalGroupValue, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *IPSecPolicyRequest) SetPfsGroup(v IKEProposalGroupValue)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *IPSecPolicyRequest) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### GetComments

`func (o *IPSecPolicyRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IPSecPolicyRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IPSecPolicyRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IPSecPolicyRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *IPSecPolicyRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPSecPolicyRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPSecPolicyRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPSecPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IPSecPolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPSecPolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPSecPolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPSecPolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


