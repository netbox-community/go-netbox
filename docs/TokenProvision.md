# TokenProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**User** | [**BriefUser**](BriefUser.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**LastUsed** | **time.Time** |  | [readonly] 
**Key** | **string** |  | [readonly] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this key | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenProvision

`func NewTokenProvision(id int32, url string, displayUrl string, display string, user BriefUser, created time.Time, lastUsed time.Time, key string, ) *TokenProvision`

NewTokenProvision instantiates a new TokenProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenProvisionWithDefaults

`func NewTokenProvisionWithDefaults() *TokenProvision`

NewTokenProvisionWithDefaults instantiates a new TokenProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenProvision) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenProvision) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenProvision) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *TokenProvision) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TokenProvision) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TokenProvision) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *TokenProvision) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *TokenProvision) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *TokenProvision) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *TokenProvision) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *TokenProvision) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *TokenProvision) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUser

`func (o *TokenProvision) GetUser() BriefUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TokenProvision) GetUserOk() (*BriefUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TokenProvision) SetUser(v BriefUser)`

SetUser sets User field to given value.


### GetCreated

`func (o *TokenProvision) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TokenProvision) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TokenProvision) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *TokenProvision) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenProvision) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenProvision) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenProvision) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *TokenProvision) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *TokenProvision) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetLastUsed

`func (o *TokenProvision) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *TokenProvision) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *TokenProvision) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### GetKey

`func (o *TokenProvision) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TokenProvision) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TokenProvision) SetKey(v string)`

SetKey sets Key field to given value.


### GetWriteEnabled

`func (o *TokenProvision) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *TokenProvision) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *TokenProvision) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *TokenProvision) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *TokenProvision) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenProvision) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenProvision) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenProvision) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


