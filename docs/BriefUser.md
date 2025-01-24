# BriefUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | 

## Methods

### NewBriefUser

`func NewBriefUser(id int32, url string, display string, username string, ) *BriefUser`

NewBriefUser instantiates a new BriefUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefUserWithDefaults

`func NewBriefUserWithDefaults() *BriefUser`

NewBriefUserWithDefaults instantiates a new BriefUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefUser) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefUser) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefUser) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefUser) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefUser) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUsername

`func (o *BriefUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BriefUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BriefUser) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


