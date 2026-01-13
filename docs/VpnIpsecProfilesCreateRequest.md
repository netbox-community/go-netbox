# VpnIpsecProfilesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | [**IPSecProfileModeValue**](IPSecProfileModeValue.md) |  | 
**IkePolicy** | [**IPSecProfileRequestIkePolicy**](IPSecProfileRequestIkePolicy.md) |  | 
**IpsecPolicy** | [**IPSecProfileRequestIpsecPolicy**](IPSecProfileRequestIpsecPolicy.md) |  | 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnIpsecProfilesCreateRequest

`func NewVpnIpsecProfilesCreateRequest(name string, mode IPSecProfileModeValue, ikePolicy IPSecProfileRequestIkePolicy, ipsecPolicy IPSecProfileRequestIpsecPolicy, ) *VpnIpsecProfilesCreateRequest`

NewVpnIpsecProfilesCreateRequest instantiates a new VpnIpsecProfilesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnIpsecProfilesCreateRequestWithDefaults

`func NewVpnIpsecProfilesCreateRequestWithDefaults() *VpnIpsecProfilesCreateRequest`

NewVpnIpsecProfilesCreateRequestWithDefaults instantiates a new VpnIpsecProfilesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnIpsecProfilesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnIpsecProfilesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnIpsecProfilesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VpnIpsecProfilesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnIpsecProfilesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnIpsecProfilesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnIpsecProfilesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *VpnIpsecProfilesCreateRequest) GetMode() IPSecProfileModeValue`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VpnIpsecProfilesCreateRequest) GetModeOk() (*IPSecProfileModeValue, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VpnIpsecProfilesCreateRequest) SetMode(v IPSecProfileModeValue)`

SetMode sets Mode field to given value.


### GetIkePolicy

`func (o *VpnIpsecProfilesCreateRequest) GetIkePolicy() IPSecProfileRequestIkePolicy`

GetIkePolicy returns the IkePolicy field if non-nil, zero value otherwise.

### GetIkePolicyOk

`func (o *VpnIpsecProfilesCreateRequest) GetIkePolicyOk() (*IPSecProfileRequestIkePolicy, bool)`

GetIkePolicyOk returns a tuple with the IkePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkePolicy

`func (o *VpnIpsecProfilesCreateRequest) SetIkePolicy(v IPSecProfileRequestIkePolicy)`

SetIkePolicy sets IkePolicy field to given value.


### GetIpsecPolicy

`func (o *VpnIpsecProfilesCreateRequest) GetIpsecPolicy() IPSecProfileRequestIpsecPolicy`

GetIpsecPolicy returns the IpsecPolicy field if non-nil, zero value otherwise.

### GetIpsecPolicyOk

`func (o *VpnIpsecProfilesCreateRequest) GetIpsecPolicyOk() (*IPSecProfileRequestIpsecPolicy, bool)`

GetIpsecPolicyOk returns a tuple with the IpsecPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecPolicy

`func (o *VpnIpsecProfilesCreateRequest) SetIpsecPolicy(v IPSecProfileRequestIpsecPolicy)`

SetIpsecPolicy sets IpsecPolicy field to given value.


### GetOwner

`func (o *VpnIpsecProfilesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VpnIpsecProfilesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VpnIpsecProfilesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VpnIpsecProfilesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VpnIpsecProfilesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VpnIpsecProfilesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VpnIpsecProfilesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VpnIpsecProfilesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VpnIpsecProfilesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VpnIpsecProfilesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VpnIpsecProfilesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnIpsecProfilesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnIpsecProfilesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnIpsecProfilesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnIpsecProfilesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnIpsecProfilesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnIpsecProfilesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnIpsecProfilesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


