# NestedFHRPGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Protocol** | [**FHRPGroupProtocol**](FHRPGroupProtocol.md) |  | 
**GroupId** | **int32** |  | 

## Methods

### NewNestedFHRPGroup

`func NewNestedFHRPGroup(id int32, url string, display string, protocol FHRPGroupProtocol, groupId int32, ) *NestedFHRPGroup`

NewNestedFHRPGroup instantiates a new NestedFHRPGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedFHRPGroupWithDefaults

`func NewNestedFHRPGroupWithDefaults() *NestedFHRPGroup`

NewNestedFHRPGroupWithDefaults instantiates a new NestedFHRPGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedFHRPGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedFHRPGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedFHRPGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedFHRPGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedFHRPGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedFHRPGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedFHRPGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedFHRPGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedFHRPGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetProtocol

`func (o *NestedFHRPGroup) GetProtocol() FHRPGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NestedFHRPGroup) GetProtocolOk() (*FHRPGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NestedFHRPGroup) SetProtocol(v FHRPGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetGroupId

`func (o *NestedFHRPGroup) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NestedFHRPGroup) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NestedFHRPGroup) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


