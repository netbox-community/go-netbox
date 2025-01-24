# UserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 
**Password** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsStaff** | Pointer to **bool** | Designates whether the user can log into this admin site. | [optional] 
**IsActive** | Pointer to **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] 
**DateJoined** | Pointer to **time.Time** |  | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Permissions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewUserRequest

`func NewUserRequest(username string, password string, ) *UserRequest`

NewUserRequest instantiates a new UserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestWithDefaults

`func NewUserRequestWithDefaults() *UserRequest`

NewUserRequestWithDefaults instantiates a new UserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetFirstName

`func (o *UserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsStaff

`func (o *UserRequest) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *UserRequest) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *UserRequest) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *UserRequest) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsActive

`func (o *UserRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDateJoined

`func (o *UserRequest) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *UserRequest) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *UserRequest) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *UserRequest) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserRequest) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserRequest) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserRequest) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserRequest) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *UserRequest) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *UserRequest) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetGroups

`func (o *UserRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPermissions

`func (o *UserRequest) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserRequest) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserRequest) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


