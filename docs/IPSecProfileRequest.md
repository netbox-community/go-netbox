# IPSecProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | [**IPSecProfileModeValue**](IPSecProfileModeValue.md) |  | 
**IkePolicy** | [**BriefIKEPolicyRequest**](BriefIKEPolicyRequest.md) |  | 
**IpsecPolicy** | [**BriefIPSecPolicyRequest**](BriefIPSecPolicyRequest.md) |  | 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIPSecProfileRequest

`func NewIPSecProfileRequest(name string, mode IPSecProfileModeValue, ikePolicy BriefIKEPolicyRequest, ipsecPolicy BriefIPSecPolicyRequest, ) *IPSecProfileRequest`

NewIPSecProfileRequest instantiates a new IPSecProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecProfileRequestWithDefaults

`func NewIPSecProfileRequestWithDefaults() *IPSecProfileRequest`

NewIPSecProfileRequestWithDefaults instantiates a new IPSecProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPSecProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPSecProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPSecProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IPSecProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPSecProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPSecProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPSecProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *IPSecProfileRequest) GetMode() IPSecProfileModeValue`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IPSecProfileRequest) GetModeOk() (*IPSecProfileModeValue, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IPSecProfileRequest) SetMode(v IPSecProfileModeValue)`

SetMode sets Mode field to given value.


### GetIkePolicy

`func (o *IPSecProfileRequest) GetIkePolicy() BriefIKEPolicyRequest`

GetIkePolicy returns the IkePolicy field if non-nil, zero value otherwise.

### GetIkePolicyOk

`func (o *IPSecProfileRequest) GetIkePolicyOk() (*BriefIKEPolicyRequest, bool)`

GetIkePolicyOk returns a tuple with the IkePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePolicy

`func (o *IPSecProfileRequest) SetIkePolicy(v BriefIKEPolicyRequest)`

SetIkePolicy sets IkePolicy field to given value.


### GetIpsecPolicy

`func (o *IPSecProfileRequest) GetIpsecPolicy() BriefIPSecPolicyRequest`

GetIpsecPolicy returns the IpsecPolicy field if non-nil, zero value otherwise.

### GetIpsecPolicyOk

`func (o *IPSecProfileRequest) GetIpsecPolicyOk() (*BriefIPSecPolicyRequest, bool)`

GetIpsecPolicyOk returns a tuple with the IpsecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicy

`func (o *IPSecProfileRequest) SetIpsecPolicy(v BriefIPSecPolicyRequest)`

SetIpsecPolicy sets IpsecPolicy field to given value.


### GetComments

`func (o *IPSecProfileRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IPSecProfileRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IPSecProfileRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IPSecProfileRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *IPSecProfileRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPSecProfileRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPSecProfileRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPSecProfileRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IPSecProfileRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPSecProfileRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPSecProfileRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPSecProfileRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


