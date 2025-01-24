# PatchedTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**BriefUserRequest**](BriefUserRequest.md) |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**LastUsed** | Pointer to **NullableTime** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this key | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedTokenRequest

`func NewPatchedTokenRequest() *PatchedTokenRequest`

NewPatchedTokenRequest instantiates a new PatchedTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTokenRequestWithDefaults

`func NewPatchedTokenRequestWithDefaults() *PatchedTokenRequest`

NewPatchedTokenRequestWithDefaults instantiates a new PatchedTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PatchedTokenRequest) GetUser() BriefUserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedTokenRequest) GetUserOk() (*BriefUserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedTokenRequest) SetUser(v BriefUserRequest)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedTokenRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetExpires

`func (o *PatchedTokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedTokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedTokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedTokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PatchedTokenRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PatchedTokenRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetLastUsed

`func (o *PatchedTokenRequest) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *PatchedTokenRequest) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *PatchedTokenRequest) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *PatchedTokenRequest) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *PatchedTokenRequest) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *PatchedTokenRequest) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetKey

`func (o *PatchedTokenRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedTokenRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedTokenRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedTokenRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *PatchedTokenRequest) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *PatchedTokenRequest) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *PatchedTokenRequest) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *PatchedTokenRequest) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


