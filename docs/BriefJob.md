# BriefJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Status** | [**BriefJobStatus**](BriefJobStatus.md) |  | 
**Created** | **time.Time** |  | [readonly] 
**Completed** | Pointer to **NullableTime** |  | [optional] 
**User** | [**BriefUser**](BriefUser.md) |  | [readonly] 

## Methods

### NewBriefJob

`func NewBriefJob(url string, status BriefJobStatus, created time.Time, user BriefUser, ) *BriefJob`

NewBriefJob instantiates a new BriefJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefJobWithDefaults

`func NewBriefJobWithDefaults() *BriefJob`

NewBriefJobWithDefaults instantiates a new BriefJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *BriefJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefJob) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *BriefJob) GetStatus() BriefJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BriefJob) GetStatusOk() (*BriefJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BriefJob) SetStatus(v BriefJobStatus)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *BriefJob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BriefJob) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BriefJob) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCompleted

`func (o *BriefJob) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *BriefJob) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *BriefJob) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *BriefJob) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *BriefJob) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *BriefJob) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetUser

`func (o *BriefJob) GetUser() BriefUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BriefJob) GetUserOk() (*BriefUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BriefJob) SetUser(v BriefUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


