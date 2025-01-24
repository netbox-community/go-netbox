# WritableEventRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**TypeCreate** | Pointer to **bool** | Triggers when a matching object is created. | [optional] 
**TypeUpdate** | Pointer to **bool** | Triggers when a matching object is updated. | [optional] 
**TypeDelete** | Pointer to **bool** | Triggers when a matching object is deleted. | [optional] 
**TypeJobStart** | Pointer to **bool** | Triggers when a job for a matching object is started. | [optional] 
**TypeJobEnd** | Pointer to **bool** | Triggers when a job for a matching object terminates. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Conditions** | Pointer to **interface{}** | A set of conditions which determine whether the event will be generated. | [optional] 
**ActionType** | Pointer to [**EventRuleActionTypeValue**](EventRuleActionTypeValue.md) |  | [optional] 
**ActionObjectType** | **string** |  | 
**ActionObjectId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewWritableEventRuleRequest

`func NewWritableEventRuleRequest(objectTypes []string, name string, actionObjectType string, ) *WritableEventRuleRequest`

NewWritableEventRuleRequest instantiates a new WritableEventRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableEventRuleRequestWithDefaults

`func NewWritableEventRuleRequestWithDefaults() *WritableEventRuleRequest`

NewWritableEventRuleRequestWithDefaults instantiates a new WritableEventRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *WritableEventRuleRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *WritableEventRuleRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *WritableEventRuleRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *WritableEventRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableEventRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableEventRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTypeCreate

`func (o *WritableEventRuleRequest) GetTypeCreate() bool`

GetTypeCreate returns the TypeCreate field if non-nil, zero value otherwise.

### GetTypeCreateOk

`func (o *WritableEventRuleRequest) GetTypeCreateOk() (*bool, bool)`

GetTypeCreateOk returns a tuple with the TypeCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCreate

`func (o *WritableEventRuleRequest) SetTypeCreate(v bool)`

SetTypeCreate sets TypeCreate field to given value.

### HasTypeCreate

`func (o *WritableEventRuleRequest) HasTypeCreate() bool`

HasTypeCreate returns a boolean if a field has been set.

### GetTypeUpdate

`func (o *WritableEventRuleRequest) GetTypeUpdate() bool`

GetTypeUpdate returns the TypeUpdate field if non-nil, zero value otherwise.

### GetTypeUpdateOk

`func (o *WritableEventRuleRequest) GetTypeUpdateOk() (*bool, bool)`

GetTypeUpdateOk returns a tuple with the TypeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUpdate

`func (o *WritableEventRuleRequest) SetTypeUpdate(v bool)`

SetTypeUpdate sets TypeUpdate field to given value.

### HasTypeUpdate

`func (o *WritableEventRuleRequest) HasTypeUpdate() bool`

HasTypeUpdate returns a boolean if a field has been set.

### GetTypeDelete

`func (o *WritableEventRuleRequest) GetTypeDelete() bool`

GetTypeDelete returns the TypeDelete field if non-nil, zero value otherwise.

### GetTypeDeleteOk

`func (o *WritableEventRuleRequest) GetTypeDeleteOk() (*bool, bool)`

GetTypeDeleteOk returns a tuple with the TypeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDelete

`func (o *WritableEventRuleRequest) SetTypeDelete(v bool)`

SetTypeDelete sets TypeDelete field to given value.

### HasTypeDelete

`func (o *WritableEventRuleRequest) HasTypeDelete() bool`

HasTypeDelete returns a boolean if a field has been set.

### GetTypeJobStart

`func (o *WritableEventRuleRequest) GetTypeJobStart() bool`

GetTypeJobStart returns the TypeJobStart field if non-nil, zero value otherwise.

### GetTypeJobStartOk

`func (o *WritableEventRuleRequest) GetTypeJobStartOk() (*bool, bool)`

GetTypeJobStartOk returns a tuple with the TypeJobStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeJobStart

`func (o *WritableEventRuleRequest) SetTypeJobStart(v bool)`

SetTypeJobStart sets TypeJobStart field to given value.

### HasTypeJobStart

`func (o *WritableEventRuleRequest) HasTypeJobStart() bool`

HasTypeJobStart returns a boolean if a field has been set.

### GetTypeJobEnd

`func (o *WritableEventRuleRequest) GetTypeJobEnd() bool`

GetTypeJobEnd returns the TypeJobEnd field if non-nil, zero value otherwise.

### GetTypeJobEndOk

`func (o *WritableEventRuleRequest) GetTypeJobEndOk() (*bool, bool)`

GetTypeJobEndOk returns a tuple with the TypeJobEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeJobEnd

`func (o *WritableEventRuleRequest) SetTypeJobEnd(v bool)`

SetTypeJobEnd sets TypeJobEnd field to given value.

### HasTypeJobEnd

`func (o *WritableEventRuleRequest) HasTypeJobEnd() bool`

HasTypeJobEnd returns a boolean if a field has been set.

### GetEnabled

`func (o *WritableEventRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WritableEventRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WritableEventRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WritableEventRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConditions

`func (o *WritableEventRuleRequest) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WritableEventRuleRequest) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WritableEventRuleRequest) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WritableEventRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *WritableEventRuleRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *WritableEventRuleRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetActionType

`func (o *WritableEventRuleRequest) GetActionType() EventRuleActionTypeValue`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WritableEventRuleRequest) GetActionTypeOk() (*EventRuleActionTypeValue, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WritableEventRuleRequest) SetActionType(v EventRuleActionTypeValue)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *WritableEventRuleRequest) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionObjectType

`func (o *WritableEventRuleRequest) GetActionObjectType() string`

GetActionObjectType returns the ActionObjectType field if non-nil, zero value otherwise.

### GetActionObjectTypeOk

`func (o *WritableEventRuleRequest) GetActionObjectTypeOk() (*string, bool)`

GetActionObjectTypeOk returns a tuple with the ActionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectType

`func (o *WritableEventRuleRequest) SetActionObjectType(v string)`

SetActionObjectType sets ActionObjectType field to given value.


### GetActionObjectId

`func (o *WritableEventRuleRequest) GetActionObjectId() int64`

GetActionObjectId returns the ActionObjectId field if non-nil, zero value otherwise.

### GetActionObjectIdOk

`func (o *WritableEventRuleRequest) GetActionObjectIdOk() (*int64, bool)`

GetActionObjectIdOk returns a tuple with the ActionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionObjectId

`func (o *WritableEventRuleRequest) SetActionObjectId(v int64)`

SetActionObjectId sets ActionObjectId field to given value.

### HasActionObjectId

`func (o *WritableEventRuleRequest) HasActionObjectId() bool`

HasActionObjectId returns a boolean if a field has been set.

### SetActionObjectIdNil

`func (o *WritableEventRuleRequest) SetActionObjectIdNil(b bool)`

 SetActionObjectIdNil sets the value for ActionObjectId to be an explicit nil

### UnsetActionObjectId
`func (o *WritableEventRuleRequest) UnsetActionObjectId()`

UnsetActionObjectId ensures that no value is present for ActionObjectId, not even an explicit nil
### GetDescription

`func (o *WritableEventRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableEventRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableEventRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableEventRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableEventRuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableEventRuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableEventRuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableEventRuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *WritableEventRuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableEventRuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableEventRuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableEventRuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


