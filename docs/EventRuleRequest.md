# EventRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**EventTypes** | [**[]EventRuleEventTypesInner**](EventRuleEventTypesInner.md) | The types of event which will trigger this rule. | 
**Conditions** | Pointer to **interface{}** | A set of conditions which determine whether the event will be generated. | [optional] 
**ActionType** | [**EventRuleActionTypeValue**](EventRuleActionTypeValue.md) |  | 
**ActionObjectType** | **string** |  | 
**ActionObjectId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewEventRuleRequest

`func NewEventRuleRequest(objectTypes []string, name string, eventTypes []EventRuleEventTypesInner, actionType EventRuleActionTypeValue, actionObjectType string, ) *EventRuleRequest`

NewEventRuleRequest instantiates a new EventRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleRequestWithDefaults

`func NewEventRuleRequestWithDefaults() *EventRuleRequest`

NewEventRuleRequestWithDefaults instantiates a new EventRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *EventRuleRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *EventRuleRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *EventRuleRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *EventRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *EventRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EventRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EventRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EventRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *EventRuleRequest) GetEventTypes() []EventRuleEventTypesInner`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *EventRuleRequest) GetEventTypesOk() (*[]EventRuleEventTypesInner, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *EventRuleRequest) SetEventTypes(v []EventRuleEventTypesInner)`

SetEventTypes sets EventTypes field to given value.


### GetConditions

`func (o *EventRuleRequest) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EventRuleRequest) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EventRuleRequest) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *EventRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *EventRuleRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *EventRuleRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetActionType

`func (o *EventRuleRequest) GetActionType() EventRuleActionTypeValue`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *EventRuleRequest) GetActionTypeOk() (*EventRuleActionTypeValue, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *EventRuleRequest) SetActionType(v EventRuleActionTypeValue)`

SetActionType sets ActionType field to given value.


### GetActionObjectType

`func (o *EventRuleRequest) GetActionObjectType() string`

GetActionObjectType returns the ActionObjectType field if non-nil, zero value otherwise.

### GetActionObjectTypeOk

`func (o *EventRuleRequest) GetActionObjectTypeOk() (*string, bool)`

GetActionObjectTypeOk returns a tuple with the ActionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectType

`func (o *EventRuleRequest) SetActionObjectType(v string)`

SetActionObjectType sets ActionObjectType field to given value.


### GetActionObjectId

`func (o *EventRuleRequest) GetActionObjectId() int64`

GetActionObjectId returns the ActionObjectId field if non-nil, zero value otherwise.

### GetActionObjectIdOk

`func (o *EventRuleRequest) GetActionObjectIdOk() (*int64, bool)`

GetActionObjectIdOk returns a tuple with the ActionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectId

`func (o *EventRuleRequest) SetActionObjectId(v int64)`

SetActionObjectId sets ActionObjectId field to given value.

### HasActionObjectId

`func (o *EventRuleRequest) HasActionObjectId() bool`

HasActionObjectId returns a boolean if a field has been set.

### SetActionObjectIdNil

`func (o *EventRuleRequest) SetActionObjectIdNil(b bool)`

 SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil

### UnsetActionObjectId
`func (o *EventRuleRequest) UnsetActionObjectId()`

UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
### GetDescription

`func (o *EventRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *EventRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EventRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EventRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EventRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *EventRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


