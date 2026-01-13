# TokenRequest

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

### NewTokenRequest

`func NewTokenRequest(user BookmarkRequestUser, ) *TokenRequest`

NewTokenRequest instantiates a new TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestWithDefaults

`func NewTokenRequestWithDefaults() *TokenRequest`

NewTokenRequestWithDefaults instantiates a new TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *TokenRequest) GetVersion() PatchedTokenRequestVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TokenRequest) GetVersionOk() (*PatchedTokenRequestVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TokenRequest) SetVersion(v PatchedTokenRequestVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TokenRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUser

`func (o *TokenRequest) GetUser() BookmarkRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TokenRequest) GetUserOk() (*BookmarkRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TokenRequest) SetUser(v BookmarkRequestUser)`

SetUser sets User field to given value.


### GetDescription

`func (o *TokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpires

`func (o *TokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *TokenRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *TokenRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetLastUsed

`func (o *TokenRequest) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *TokenRequest) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *TokenRequest) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *TokenRequest) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *TokenRequest) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *TokenRequest) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetEnabled

`func (o *TokenRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TokenRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TokenRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TokenRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWriteEnabled

`func (o *TokenRequest) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *TokenRequest) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *TokenRequest) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *TokenRequest) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetPepperId

`func (o *TokenRequest) GetPepperId() int32`

GetPepperId returns the PepperId field if non-nil, zero value otherwise.

### GetPepperIdOk

`func (o *TokenRequest) GetPepperIdOk() (*int32, bool)`

GetPepperIdOk returns a tuple with the PepperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPepperId

`func (o *TokenRequest) SetPepperId(v int32)`

SetPepperId sets PepperId field to given value.

### HasPepperId

`func (o *TokenRequest) HasPepperId() bool`

HasPepperId returns a boolean if a field has been set.

### SetPepperIdNil

`func (o *TokenRequest) SetPepperIdNil(b bool)`

 SetPepperIdNil sets the value for PepperId to be an explicit nil

### UnsetPepperId
`func (o *TokenRequest) UnsetPepperId()`

UnsetPepperId ensures that no value is present for PepperId, not even an explicit nil
### GetToken

`func (o *TokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


