# IKEPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Version** | [**IKEPolicyVersionValue**](IKEPolicyVersionValue.md) |  | 
**Mode** | Pointer to [**IKEPolicyModeValue**](IKEPolicyModeValue.md) |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PresharedKey** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIKEPolicyRequest

`func NewIKEPolicyRequest(name string, version IKEPolicyVersionValue, ) *IKEPolicyRequest`

NewIKEPolicyRequest instantiates a new IKEPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIKEPolicyRequestWithDefaults

`func NewIKEPolicyRequestWithDefaults() *IKEPolicyRequest`

NewIKEPolicyRequestWithDefaults instantiates a new IKEPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IKEPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IKEPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IKEPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IKEPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IKEPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IKEPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IKEPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *IKEPolicyRequest) GetVersion() IKEPolicyVersionValue`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IKEPolicyRequest) GetVersionOk() (*IKEPolicyVersionValue, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IKEPolicyRequest) SetVersion(v IKEPolicyVersionValue)`

SetVersion sets Version field to given value.


### GetMode

`func (o *IKEPolicyRequest) GetMode() IKEPolicyModeValue`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IKEPolicyRequest) GetModeOk() (*IKEPolicyModeValue, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IKEPolicyRequest) SetMode(v IKEPolicyModeValue)`

SetMode sets Mode field to given value.

### HasMode

`func (o *IKEPolicyRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProposals

`func (o *IKEPolicyRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *IKEPolicyRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *IKEPolicyRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *IKEPolicyRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPresharedKey

`func (o *IKEPolicyRequest) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *IKEPolicyRequest) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *IKEPolicyRequest) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *IKEPolicyRequest) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetComments

`func (o *IKEPolicyRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IKEPolicyRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IKEPolicyRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IKEPolicyRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *IKEPolicyRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IKEPolicyRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IKEPolicyRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IKEPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IKEPolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IKEPolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IKEPolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IKEPolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


