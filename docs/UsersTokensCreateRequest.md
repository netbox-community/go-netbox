# UsersTokensCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**PatchedTokenRequestVersion**](PatchedTokenRequestVersion.md) |  | [optional] 
**User** | [**BookmarkRequestUser**](BookmarkRequestUser.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**LastUsed** | Pointer to **NullableTime** |  | [optional] 
**Enabled** | Pointer to **bool** | Disable to temporarily revoke this token without deleting it. | [optional] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this key | [optional] 
**PepperId** | Pointer to **NullableInt32** | ID of the cryptographic pepper used to hash the token (v2 only) | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewUsersTokensCreateRequest

`func NewUsersTokensCreateRequest(user BookmarkRequestUser, ) *UsersTokensCreateRequest`

NewUsersTokensCreateRequest instantiates a new UsersTokensCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersTokensCreateRequestWithDefaults

`func NewUsersTokensCreateRequestWithDefaults() *UsersTokensCreateRequest`

NewUsersTokensCreateRequestWithDefaults instantiates a new UsersTokensCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UsersTokensCreateRequest) GetVersion() PatchedTokenRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UsersTokensCreateRequest) GetVersionOk() (*PatchedTokenRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UsersTokensCreateRequest) SetVersion(v PatchedTokenRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UsersTokensCreateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUser

`func (o *UsersTokensCreateRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UsersTokensCreateRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UsersTokensCreateRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.


### GetDescription

`func (o *UsersTokensCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UsersTokensCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UsersTokensCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UsersTokensCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpires

`func (o *UsersTokensCreateRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *UsersTokensCreateRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *UsersTokensCreateRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *UsersTokensCreateRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *UsersTokensCreateRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *UsersTokensCreateRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetLastUsed

`func (o *UsersTokensCreateRequest) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *UsersTokensCreateRequest) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *UsersTokensCreateRequest) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *UsersTokensCreateRequest) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *UsersTokensCreateRequest) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *UsersTokensCreateRequest) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetEnabled

`func (o *UsersTokensCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UsersTokensCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UsersTokensCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UsersTokensCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *UsersTokensCreateRequest) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *UsersTokensCreateRequest) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *UsersTokensCreateRequest) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *UsersTokensCreateRequest) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetPepperId

`func (o *UsersTokensCreateRequest) GetPepperId() int32`

GetPepperId returns the PepperId field if non-nil, zero value otherwise.

### GetPepperIdOk

`func (o *UsersTokensCreateRequest) GetPepperIdOk() (*int32, bool)`

GetPepperIdOk returns a tuple with the PepperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPepperId

`func (o *UsersTokensCreateRequest) SetPepperId(v int32)`

SetPepperId sets PepperId field to given value.

### HasPepperId

`func (o *UsersTokensCreateRequest) HasPepperId() bool`

HasPepperId returns a boolean if a field has been set.

### SetPepperIdNil

`func (o *UsersTokensCreateRequest) SetPepperIdNil(b bool)`

 SetPepperIdNil sets the value for PepperId to be an explicit nil

### UnsetPepperId
`func (o *UsersTokensCreateRequest) UnsetPepperId()`

UnsetPepperId ensures that no value is present for PepperId, not even an explicit nil
### GetToken

`func (o *UsersTokensCreateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UsersTokensCreateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UsersTokensCreateRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UsersTokensCreateRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


