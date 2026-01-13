# PatchedTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**PatchedTokenRequestVersion**](PatchedTokenRequestVersion.md) |  | [optional] 
**User** | Pointer to [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**LastUsed** | Pointer to **NullableTime** |  | [optional] 
**Enabled** | Pointer to **bool** | Disable to temporarily revoke this token without deleting it. | [optional] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this key | [optional] 
**PepperId** | Pointer to **NullableInt32** | ID of the cryptographic pepper used to hash the token (v2 only) | [optional] 
**Token** | Pointer to **string** |  | [optional] 

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

### GetVersion

`func (o *PatchedTokenRequest) GetVersion() PatchedTokenRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedTokenRequest) GetVersionOk() (*PatchedTokenRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedTokenRequest) SetVersion(v PatchedTokenRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedTokenRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUser

`func (o *PatchedTokenRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedTokenRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedTokenRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedTokenRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

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
### GetEnabled

`func (o *PatchedTokenRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedTokenRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedTokenRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedTokenRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

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

### GetPepperId

`func (o *PatchedTokenRequest) GetPepperId() int32`

GetPepperId returns the PepperId field if non-nil, zero value otherwise.

### GetPepperIdOk

`func (o *PatchedTokenRequest) GetPepperIdOk() (*int32, bool)`

GetPepperIdOk returns a tuple with the PepperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPepperId

`func (o *PatchedTokenRequest) SetPepperId(v int32)`

SetPepperId sets PepperId field to given value.

### HasPepperId

`func (o *PatchedTokenRequest) HasPepperId() bool`

HasPepperId returns a boolean if a field has been set.

### SetPepperIdNil

`func (o *PatchedTokenRequest) SetPepperIdNil(b bool)`

 SetPepperIdNil sets the value for PepperId to be an explicit nil

### UnsetPepperId
`func (o *PatchedTokenRequest) UnsetPepperId()`

UnsetPepperId ensures that no value is present for PepperId, not even an explicit nil
### GetToken

`func (o *PatchedTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PatchedTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PatchedTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PatchedTokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


