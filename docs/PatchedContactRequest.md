# PatchedContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**NullableBriefContactGroupRequest**](BriefContactGroupRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedContactRequest

`func NewPatchedContactRequest() *PatchedContactRequest`

NewPatchedContactRequest instantiates a new PatchedContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedContactRequestWithDefaults

`func NewPatchedContactRequestWithDefaults() *PatchedContactRequest`

NewPatchedContactRequestWithDefaults instantiates a new PatchedContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *PatchedContactRequest) GetGroup() BriefContactGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedContactRequest) GetGroupOk() (*BriefContactGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedContactRequest) SetGroup(v BriefContactGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedContactRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedContactRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedContactRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetName

`func (o *PatchedContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedContactRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedContactRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *PatchedContactRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchedContactRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchedContactRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchedContactRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPhone

`func (o *PatchedContactRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PatchedContactRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PatchedContactRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PatchedContactRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedContactRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedContactRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedContactRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedContactRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAddress

`func (o *PatchedContactRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchedContactRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchedContactRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchedContactRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLink

`func (o *PatchedContactRequest) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PatchedContactRequest) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PatchedContactRequest) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PatchedContactRequest) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedContactRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedContactRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedContactRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedContactRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedContactRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedContactRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedContactRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedContactRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedContactRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedContactRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedContactRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedContactRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedContactRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedContactRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedContactRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedContactRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


