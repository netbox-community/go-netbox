# EventRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**EventTypes** | [**[]EventRuleEventTypesInner**](EventRuleEventTypesInner.md) | The types of event which will trigger this rule. | 
**Conditions** | Pointer to **interface{}** | A set of conditions which determine whether the event will be generated. | [optional] 
**ActionType** | [**EventRuleActionType**](EventRuleActionType.md) |  | 
**ActionObjectType** | **string** |  | 
**ActionObjectId** | Pointer to **NullableInt64** |  | [optional] 
**ActionObject** | **map[string]interface{}** |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewEventRule

`func NewEventRule(id int32, url string, displayUrl string, display string, objectTypes []string, name string, eventTypes []EventRuleEventTypesInner, actionType EventRuleActionType, actionObjectType string, actionObject map[string]interface{}, created NullableTime, lastUpdated NullableTime, ) *EventRule`

NewEventRule instantiates a new EventRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleWithDefaults

`func NewEventRuleWithDefaults() *EventRule`

NewEventRuleWithDefaults instantiates a new EventRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventRule) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *EventRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventRule) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *EventRule) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *EventRule) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *EventRule) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *EventRule) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *EventRule) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *EventRule) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetObjectTypes

`func (o *EventRule) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *EventRule) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *EventRule) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *EventRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventRule) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *EventRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EventRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EventRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EventRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventRule) GetEventTypes() []EventRuleEventTypesInner`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventRule) GetEventTypesOk() (*[]EventRuleEventTypesInner, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventRule) SetEventTypes(v []EventRuleEventTypesInner)`

SetEventTypes sets EventTypes field to given value.


### GetConditions

`func (o *EventRule) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EventRule) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EventRule) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *EventRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *EventRule) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *EventRule) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetActionType

`func (o *EventRule) GetActionType() EventRuleActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *EventRule) GetActionTypeOk() (*EventRuleActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *EventRule) SetActionType(v EventRuleActionType)`

SetActionType sets ActionType field to given value.


### GetActionObjectType

`func (o *EventRule) GetActionObjectType() string`

GetActionObjectType returns the ActionObjectType field if non-nil, zero value otherwise.

### GetActionObjectTypeOk

`func (o *EventRule) GetActionObjectTypeOk() (*string, bool)`

GetActionObjectTypeOk returns a tuple with the ActionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectType

`func (o *EventRule) SetActionObjectType(v string)`

SetActionObjectType sets ActionObjectType field to given value.


### GetActionObjectId

`func (o *EventRule) GetActionObjectId() int64`

GetActionObjectId returns the ActionObjectId field if non-nil, zero value otherwise.

### GetActionObjectIdOk

`func (o *EventRule) GetActionObjectIdOk() (*int64, bool)`

GetActionObjectIdOk returns a tuple with the ActionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectId

`func (o *EventRule) SetActionObjectId(v int64)`

SetActionObjectId sets ActionObjectId field to given value.

### HasActionObjectId

`func (o *EventRule) HasActionObjectId() bool`

HasActionObjectId returns a boolean if a field has been set.

### SetActionObjectIdNil

`func (o *EventRule) SetActionObjectIdNil(b bool)`

 SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil

### UnsetActionObjectId
`func (o *EventRule) UnsetActionObjectId()`

UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
### GetActionObject

`func (o *EventRule) GetActionObject() map[string]interface{}`

GetActionObject returns the ActionObject field if non-nil, zero value otherwise.

### GetActionObjectOk

`func (o *EventRule) GetActionObjectOk() (*map[string]interface{}, bool)`

GetActionObjectOk returns a tuple with the ActionObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObject

`func (o *EventRule) SetActionObject(v map[string]interface{})`

SetActionObject sets ActionObject field to given value.


### GetDescription

`func (o *EventRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *EventRule) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EventRule) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EventRule) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EventRule) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *EventRule) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventRule) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventRule) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreated

`func (o *EventRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EventRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *EventRule) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *EventRule) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *EventRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EventRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EventRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *EventRule) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *EventRule) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


