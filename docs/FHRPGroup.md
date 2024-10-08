# FHRPGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Protocol** | [**FHRPGroupProtocol**](FHRPGroupProtocol.md) |  | 
**GroupId** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewFHRPGroup

`func NewFHRPGroup(id int32, url string, display string, protocol FHRPGroupProtocol, groupId int32, ) *FHRPGroup`

NewFHRPGroup instantiates a new FHRPGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFHRPGroupWithDefaults

`func NewFHRPGroupWithDefaults() *FHRPGroup`

NewFHRPGroupWithDefaults instantiates a new FHRPGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FHRPGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FHRPGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FHRPGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *FHRPGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FHRPGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FHRPGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *FHRPGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *FHRPGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *FHRPGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetProtocol

`func (o *FHRPGroup) GetProtocol() FHRPGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FHRPGroup) GetProtocolOk() (*FHRPGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FHRPGroup) SetProtocol(v FHRPGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetGroupId

`func (o *FHRPGroup) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FHRPGroup) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FHRPGroup) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetDescription

`func (o *FHRPGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FHRPGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FHRPGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FHRPGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


