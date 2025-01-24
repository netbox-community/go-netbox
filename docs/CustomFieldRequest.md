# CustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | **[]string** |  | 
**Type** | [**CustomFieldTypeValue**](CustomFieldTypeValue.md) |  | 
**RelatedObjectType** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Internal field name | 
**Label** | Pointer to **string** | Name of the field as displayed to users (if not provided, &#39;the field&#39;s name will be used) | [optional] 
**GroupName** | Pointer to **string** | Custom fields within the same group will be displayed together | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** | This field is required when creating new objects or editing an existing object. | [optional] 
**Unique** | Pointer to **bool** | The value of this field must be unique for the assigned object | [optional] 
**SearchWeight** | Pointer to **int32** | Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored. | [optional] 
**FilterLogic** | Pointer to [**CustomFieldFilterLogicValue**](CustomFieldFilterLogicValue.md) |  | [optional] 
**UiVisible** | Pointer to [**CustomFieldUiVisibleValue**](CustomFieldUiVisibleValue.md) |  | [optional] 
**UiEditable** | Pointer to [**CustomFieldUiEditableValue**](CustomFieldUiEditableValue.md) |  | [optional] 
**IsCloneable** | Pointer to **bool** | Replicate this value when cloning objects | [optional] 
**Default** | Pointer to **interface{}** | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**RelatedObjectFilter** | Pointer to **interface{}** | Filter the object selection choices using a query_params dict (must be a JSON value).Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**Weight** | Pointer to **int32** | Fields with higher weights appear lower in a form. | [optional] 
**ValidationMinimum** | Pointer to **NullableInt64** | Minimum allowed value (for numeric fields) | [optional] 
**ValidationMaximum** | Pointer to **NullableInt64** | Maximum allowed value (for numeric fields) | [optional] 
**ValidationRegex** | Pointer to **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. | [optional] 
**ChoiceSet** | Pointer to [**NullableBriefCustomFieldChoiceSetRequest**](BriefCustomFieldChoiceSetRequest.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomFieldRequest

`func NewCustomFieldRequest(objectTypes []string, type_ CustomFieldTypeValue, name string, ) *CustomFieldRequest`

NewCustomFieldRequest instantiates a new CustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldRequestWithDefaults

`func NewCustomFieldRequestWithDefaults() *CustomFieldRequest`

NewCustomFieldRequestWithDefaults instantiates a new CustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *CustomFieldRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *CustomFieldRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *CustomFieldRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetType

`func (o *CustomFieldRequest) GetType() CustomFieldTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldRequest) GetTypeOk() (*CustomFieldTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldRequest) SetType(v CustomFieldTypeValue)`

SetType sets Type field to given value.


### GetRelatedObjectType

`func (o *CustomFieldRequest) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *CustomFieldRequest) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *CustomFieldRequest) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *CustomFieldRequest) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### SetRelatedObjectTypeNil

`func (o *CustomFieldRequest) SetRelatedObjectTypeNil(b bool)`

 SetRelatedObjectTypeNil sets the value for RelatedObjectType to be an explicit nil

### UnsetRelatedObjectType
`func (o *CustomFieldRequest) UnsetRelatedObjectType()`

UnsetRelatedObjectType ensures that no value is present for RelatedObjectType, not even an explicit nil
### GetName

`func (o *CustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CustomFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomFieldRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupName

`func (o *CustomFieldRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CustomFieldRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CustomFieldRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CustomFieldRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDescription

`func (o *CustomFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *CustomFieldRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CustomFieldRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CustomFieldRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CustomFieldRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnique

`func (o *CustomFieldRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *CustomFieldRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *CustomFieldRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *CustomFieldRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetSearchWeight

`func (o *CustomFieldRequest) GetSearchWeight() int32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *CustomFieldRequest) GetSearchWeightOk() (*int32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *CustomFieldRequest) SetSearchWeight(v int32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *CustomFieldRequest) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetFilterLogic

`func (o *CustomFieldRequest) GetFilterLogic() CustomFieldFilterLogicValue`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *CustomFieldRequest) GetFilterLogicOk() (*CustomFieldFilterLogicValue, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *CustomFieldRequest) SetFilterLogic(v CustomFieldFilterLogicValue)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *CustomFieldRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetUiVisible

`func (o *CustomFieldRequest) GetUiVisible() CustomFieldUiVisibleValue`

GetUiVisible returns the UiVisible field if non-nil, zero value otherwise.

### GetUiVisibleOk

`func (o *CustomFieldRequest) GetUiVisibleOk() (*CustomFieldUiVisibleValue, bool)`

GetUiVisibleOk returns a tuple with the UiVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVisible

`func (o *CustomFieldRequest) SetUiVisible(v CustomFieldUiVisibleValue)`

SetUiVisible sets UiVisible field to given value.

### HasUiVisible

`func (o *CustomFieldRequest) HasUiVisible() bool`

HasUiVisible returns a boolean if a field has been set.

### GetUiEditable

`func (o *CustomFieldRequest) GetUiEditable() CustomFieldUiEditableValue`

GetUiEditable returns the UiEditable field if non-nil, zero value otherwise.

### GetUiEditableOk

`func (o *CustomFieldRequest) GetUiEditableOk() (*CustomFieldUiEditableValue, bool)`

GetUiEditableOk returns a tuple with the UiEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiEditable

`func (o *CustomFieldRequest) SetUiEditable(v CustomFieldUiEditableValue)`

SetUiEditable sets UiEditable field to given value.

### HasUiEditable

`func (o *CustomFieldRequest) HasUiEditable() bool`

HasUiEditable returns a boolean if a field has been set.

### GetIsCloneable

`func (o *CustomFieldRequest) GetIsCloneable() bool`

GetIsCloneable returns the IsCloneable field if non-nil, zero value otherwise.

### GetIsCloneableOk

`func (o *CustomFieldRequest) GetIsCloneableOk() (*bool, bool)`

GetIsCloneableOk returns a tuple with the IsCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloneable

`func (o *CustomFieldRequest) SetIsCloneable(v bool)`

SetIsCloneable sets IsCloneable field to given value.

### HasIsCloneable

`func (o *CustomFieldRequest) HasIsCloneable() bool`

HasIsCloneable returns a boolean if a field has been set.

### GetDefault

`func (o *CustomFieldRequest) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomFieldRequest) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomFieldRequest) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomFieldRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *CustomFieldRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *CustomFieldRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetRelatedObjectFilter

`func (o *CustomFieldRequest) GetRelatedObjectFilter() interface{}`

GetRelatedObjectFilter returns the RelatedObjectFilter field if non-nil, zero value otherwise.

### GetRelatedObjectFilterOk

`func (o *CustomFieldRequest) GetRelatedObjectFilterOk() (*interface{}, bool)`

GetRelatedObjectFilterOk returns a tuple with the RelatedObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectFilter

`func (o *CustomFieldRequest) SetRelatedObjectFilter(v interface{})`

SetRelatedObjectFilter sets RelatedObjectFilter field to given value.

### HasRelatedObjectFilter

`func (o *CustomFieldRequest) HasRelatedObjectFilter() bool`

HasRelatedObjectFilter returns a boolean if a field has been set.

### SetRelatedObjectFilterNil

`func (o *CustomFieldRequest) SetRelatedObjectFilterNil(b bool)`

 SetRelatedObjectFilterNil sets the value for RelatedObjectFilter to be an explicit nil

### UnsetRelatedObjectFilter
`func (o *CustomFieldRequest) UnsetRelatedObjectFilter()`

UnsetRelatedObjectFilter ensures that no value is present for RelatedObjectFilter, not even an explicit nil
### GetWeight

`func (o *CustomFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CustomFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CustomFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CustomFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetValidationMinimum

`func (o *CustomFieldRequest) GetValidationMinimum() int64`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *CustomFieldRequest) GetValidationMinimumOk() (*int64, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *CustomFieldRequest) SetValidationMinimum(v int64)`

SetValidationMinimum sets ValidationMinimum field to given value.

### HasValidationMinimum

`func (o *CustomFieldRequest) HasValidationMinimum() bool`

HasValidationMinimum returns a boolean if a field has been set.

### SetValidationMinimumNil

`func (o *CustomFieldRequest) SetValidationMinimumNil(b bool)`

 SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil

### UnsetValidationMinimum
`func (o *CustomFieldRequest) UnsetValidationMinimum()`

UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
### GetValidationMaximum

`func (o *CustomFieldRequest) GetValidationMaximum() int64`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *CustomFieldRequest) GetValidationMaximumOk() (*int64, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *CustomFieldRequest) SetValidationMaximum(v int64)`

SetValidationMaximum sets ValidationMaximum field to given value.

### HasValidationMaximum

`func (o *CustomFieldRequest) HasValidationMaximum() bool`

HasValidationMaximum returns a boolean if a field has been set.

### SetValidationMaximumNil

`func (o *CustomFieldRequest) SetValidationMaximumNil(b bool)`

 SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil

### UnsetValidationMaximum
`func (o *CustomFieldRequest) UnsetValidationMaximum()`

UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
### GetValidationRegex

`func (o *CustomFieldRequest) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *CustomFieldRequest) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *CustomFieldRequest) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *CustomFieldRequest) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetChoiceSet

`func (o *CustomFieldRequest) GetChoiceSet() BriefCustomFieldChoiceSetRequest`

GetChoiceSet returns the ChoiceSet field if non-nil, zero value otherwise.

### GetChoiceSetOk

`func (o *CustomFieldRequest) GetChoiceSetOk() (*BriefCustomFieldChoiceSetRequest, bool)`

GetChoiceSetOk returns a tuple with the ChoiceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceSet

`func (o *CustomFieldRequest) SetChoiceSet(v BriefCustomFieldChoiceSetRequest)`

SetChoiceSet sets ChoiceSet field to given value.

### HasChoiceSet

`func (o *CustomFieldRequest) HasChoiceSet() bool`

HasChoiceSet returns a boolean if a field has been set.

### SetChoiceSetNil

`func (o *CustomFieldRequest) SetChoiceSetNil(b bool)`

 SetChoiceSetNil sets the value for ChoiceSet to be an explicit nil

### UnsetChoiceSet
`func (o *CustomFieldRequest) UnsetChoiceSet()`

UnsetChoiceSet ensures that no value is present for ChoiceSet, not even an explicit nil
### GetComments

`func (o *CustomFieldRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CustomFieldRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CustomFieldRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CustomFieldRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


