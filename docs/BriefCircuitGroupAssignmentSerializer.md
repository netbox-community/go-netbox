# BriefCircuitGroupAssignmentSerializer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Group** | [**BriefCircuitGroup**](BriefCircuitGroup.md) |  | 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriority**](BriefCircuitGroupAssignmentSerializerPriority.md) |  | [optional] 

## Methods

### NewBriefCircuitGroupAssignmentSerializer

`func NewBriefCircuitGroupAssignmentSerializer(id int32, url string, display string, group BriefCircuitGroup, ) *BriefCircuitGroupAssignmentSerializer`

NewBriefCircuitGroupAssignmentSerializer instantiates a new BriefCircuitGroupAssignmentSerializer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefCircuitGroupAssignmentSerializerWithDefaults

`func NewBriefCircuitGroupAssignmentSerializerWithDefaults() *BriefCircuitGroupAssignmentSerializer`

NewBriefCircuitGroupAssignmentSerializerWithDefaults instantiates a new BriefCircuitGroupAssignmentSerializer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefCircuitGroupAssignmentSerializer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefCircuitGroupAssignmentSerializer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefCircuitGroupAssignmentSerializer) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefCircuitGroupAssignmentSerializer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefCircuitGroupAssignmentSerializer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefCircuitGroupAssignmentSerializer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefCircuitGroupAssignmentSerializer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefCircuitGroupAssignmentSerializer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefCircuitGroupAssignmentSerializer) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetGroup

`func (o *BriefCircuitGroupAssignmentSerializer) GetGroup() BriefCircuitGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BriefCircuitGroupAssignmentSerializer) GetGroupOk() (*BriefCircuitGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BriefCircuitGroupAssignmentSerializer) SetGroup(v BriefCircuitGroup)`

SetGroup sets Group field to given value.


### GetPriority

`func (o *BriefCircuitGroupAssignmentSerializer) GetPriority() BriefCircuitGroupAssignmentSerializerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BriefCircuitGroupAssignmentSerializer) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BriefCircuitGroupAssignmentSerializer) SetPriority(v BriefCircuitGroupAssignmentSerializerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BriefCircuitGroupAssignmentSerializer) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


