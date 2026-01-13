# BriefProviderAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] [default to ""]
**Account** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefProviderAccount

`func NewBriefProviderAccount(id int32, url string, display string, account string, ) *BriefProviderAccount`

NewBriefProviderAccount instantiates a new BriefProviderAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefProviderAccountWithDefaults

`func NewBriefProviderAccountWithDefaults() *BriefProviderAccount`

NewBriefProviderAccountWithDefaults instantiates a new BriefProviderAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefProviderAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefProviderAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefProviderAccount) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefProviderAccount) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefProviderAccount) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefProviderAccount) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefProviderAccount) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefProviderAccount) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefProviderAccount) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefProviderAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefProviderAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefProviderAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BriefProviderAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *BriefProviderAccount) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BriefProviderAccount) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BriefProviderAccount) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetDescription

`func (o *BriefProviderAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefProviderAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefProviderAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefProviderAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


