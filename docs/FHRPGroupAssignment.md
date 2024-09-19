# FHRPGroupAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Group** | [**BriefFHRPGroup**](BriefFHRPGroup.md) |  | 
**InterfaceType** | **string** |  | 
**InterfaceId** | **int64** |  | 
**Interface** | **interface{}** |  | [readonly] 
**Priority** | **int32** |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewFHRPGroupAssignment

`func NewFHRPGroupAssignment(id int32, url string, display string, group BriefFHRPGroup, interfaceType string, interfaceId int64, interface_ interface{}, priority int32, created NullableTime, lastUpdated NullableTime, ) *FHRPGroupAssignment`

NewFHRPGroupAssignment instantiates a new FHRPGroupAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFHRPGroupAssignmentWithDefaults

`func NewFHRPGroupAssignmentWithDefaults() *FHRPGroupAssignment`

NewFHRPGroupAssignmentWithDefaults instantiates a new FHRPGroupAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FHRPGroupAssignment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FHRPGroupAssignment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FHRPGroupAssignment) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *FHRPGroupAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FHRPGroupAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FHRPGroupAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *FHRPGroupAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *FHRPGroupAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *FHRPGroupAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetGroup

`func (o *FHRPGroupAssignment) GetGroup() BriefFHRPGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FHRPGroupAssignment) GetGroupOk() (*BriefFHRPGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FHRPGroupAssignment) SetGroup(v BriefFHRPGroup)`

SetGroup sets Group field to given value.


### GetInterfaceType

`func (o *FHRPGroupAssignment) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *FHRPGroupAssignment) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *FHRPGroupAssignment) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.


### GetInterfaceId

`func (o *FHRPGroupAssignment) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *FHRPGroupAssignment) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *FHRPGroupAssignment) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.


### GetInterface

`func (o *FHRPGroupAssignment) GetInterface() interface{}`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *FHRPGroupAssignment) GetInterfaceOk() (*interface{}, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *FHRPGroupAssignment) SetInterface(v interface{})`

SetInterface sets Interface field to given value.


### SetInterfaceNil

`func (o *FHRPGroupAssignment) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *FHRPGroupAssignment) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetPriority

`func (o *FHRPGroupAssignment) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *FHRPGroupAssignment) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *FHRPGroupAssignment) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetCreated

`func (o *FHRPGroupAssignment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FHRPGroupAssignment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FHRPGroupAssignment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *FHRPGroupAssignment) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FHRPGroupAssignment) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *FHRPGroupAssignment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FHRPGroupAssignment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FHRPGroupAssignment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *FHRPGroupAssignment) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *FHRPGroupAssignment) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


