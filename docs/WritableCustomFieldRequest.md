# WritableCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | **[]string** |  | 
**Type** | Pointer to [**PatchedWritableCustomFieldRequestType**](PatchedWritableCustomFieldRequestType.md) |  | [optional] 
**RelatedObjectType** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Internal field name | 
**Label** | Pointer to **string** | Name of the field as displayed to users (if not provided, &#39;the field&#39;s name will be used) | [optional] 
**GroupName** | Pointer to **string** | Custom fields within the same group will be displayed together | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** | This field is required when creating new objects or editing an existing object. | [optional] 
**Unique** | Pointer to **bool** | The value of this field must be unique for the assigned object | [optional] 
**SearchWeight** | Pointer to **int32** | Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored. | [optional] 
**FilterLogic** | Pointer to [**PatchedWritableCustomFieldRequestFilterLogic**](PatchedWritableCustomFieldRequestFilterLogic.md) |  | [optional] 
**UiVisible** | Pointer to [**PatchedWritableCustomFieldRequestUiVisible**](PatchedWritableCustomFieldRequestUiVisible.md) |  | [optional] 
**UiEditable** | Pointer to [**PatchedWritableCustomFieldRequestUiEditable**](PatchedWritableCustomFieldRequestUiEditable.md) |  | [optional] 
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

### NewWritableCustomFieldRequest

`func NewWritableCustomFieldRequest(objectTypes []string, name string, ) *WritableCustomFieldRequest`

NewWritableCustomFieldRequest instantiates a new WritableCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCustomFieldRequestWithDefaults

`func NewWritableCustomFieldRequestWithDefaults() *WritableCustomFieldRequest`

NewWritableCustomFieldRequestWithDefaults instantiates a new WritableCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *WritableCustomFieldRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *WritableCustomFieldRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *WritableCustomFieldRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetType

`func (o *WritableCustomFieldRequest) GetType() PatchedWritableCustomFieldRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCustomFieldRequest) GetTypeOk() (*PatchedWritableCustomFieldRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCustomFieldRequest) SetType(v PatchedWritableCustomFieldRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *WritableCustomFieldRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelatedObjectType

`func (o *WritableCustomFieldRequest) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *WritableCustomFieldRequest) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *WritableCustomFieldRequest) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *WritableCustomFieldRequest) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### SetRelatedObjectTypeNil

`func (o *WritableCustomFieldRequest) SetRelatedObjectTypeNil(b bool)`

 SetRelatedObjectTypeNil sets the value for RelatedObjectType to be an explicit nil

### UnsetRelatedObjectType
`func (o *WritableCustomFieldRequest) UnsetRelatedObjectType()`

UnsetRelatedObjectType ensures that no value is present for RelatedObjectType, not even an explicit nil
### GetName

`func (o *WritableCustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableCustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableCustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableCustomFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableCustomFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableCustomFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableCustomFieldRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupName

`func (o *WritableCustomFieldRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *WritableCustomFieldRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *WritableCustomFieldRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *WritableCustomFieldRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDescription

`func (o *WritableCustomFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableCustomFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableCustomFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableCustomFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *WritableCustomFieldRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *WritableCustomFieldRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *WritableCustomFieldRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *WritableCustomFieldRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnique

`func (o *WritableCustomFieldRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *WritableCustomFieldRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *WritableCustomFieldRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *WritableCustomFieldRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetSearchWeight

`func (o *WritableCustomFieldRequest) GetSearchWeight() int32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *WritableCustomFieldRequest) GetSearchWeightOk() (*int32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *WritableCustomFieldRequest) SetSearchWeight(v int32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *WritableCustomFieldRequest) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetFilterLogic

`func (o *WritableCustomFieldRequest) GetFilterLogic() PatchedWritableCustomFieldRequestFilterLogic`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *WritableCustomFieldRequest) GetFilterLogicOk() (*PatchedWritableCustomFieldRequestFilterLogic, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *WritableCustomFieldRequest) SetFilterLogic(v PatchedWritableCustomFieldRequestFilterLogic)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *WritableCustomFieldRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetUiVisible

`func (o *WritableCustomFieldRequest) GetUiVisible() PatchedWritableCustomFieldRequestUiVisible`

GetUiVisible returns the UiVisible field if non-nil, zero value otherwise.

### GetUiVisibleOk

`func (o *WritableCustomFieldRequest) GetUiVisibleOk() (*PatchedWritableCustomFieldRequestUiVisible, bool)`

GetUiVisibleOk returns a tuple with the UiVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVisible

`func (o *WritableCustomFieldRequest) SetUiVisible(v PatchedWritableCustomFieldRequestUiVisible)`

SetUiVisible sets UiVisible field to given value.

### HasUiVisible

`func (o *WritableCustomFieldRequest) HasUiVisible() bool`

HasUiVisible returns a boolean if a field has been set.

### GetUiEditable

`func (o *WritableCustomFieldRequest) GetUiEditable() PatchedWritableCustomFieldRequestUiEditable`

GetUiEditable returns the UiEditable field if non-nil, zero value otherwise.

### GetUiEditableOk

`func (o *WritableCustomFieldRequest) GetUiEditableOk() (*PatchedWritableCustomFieldRequestUiEditable, bool)`

GetUiEditableOk returns a tuple with the UiEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiEditable

`func (o *WritableCustomFieldRequest) SetUiEditable(v PatchedWritableCustomFieldRequestUiEditable)`

SetUiEditable sets UiEditable field to given value.

### HasUiEditable

`func (o *WritableCustomFieldRequest) HasUiEditable() bool`

HasUiEditable returns a boolean if a field has been set.

### GetIsCloneable

`func (o *WritableCustomFieldRequest) GetIsCloneable() bool`

GetIsCloneable returns the IsCloneable field if non-nil, zero value otherwise.

### GetIsCloneableOk

`func (o *WritableCustomFieldRequest) GetIsCloneableOk() (*bool, bool)`

GetIsCloneableOk returns a tuple with the IsCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloneable

`func (o *WritableCustomFieldRequest) SetIsCloneable(v bool)`

SetIsCloneable sets IsCloneable field to given value.

### HasIsCloneable

`func (o *WritableCustomFieldRequest) HasIsCloneable() bool`

HasIsCloneable returns a boolean if a field has been set.

### GetDefault

`func (o *WritableCustomFieldRequest) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WritableCustomFieldRequest) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WritableCustomFieldRequest) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WritableCustomFieldRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *WritableCustomFieldRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *WritableCustomFieldRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetRelatedObjectFilter

`func (o *WritableCustomFieldRequest) GetRelatedObjectFilter() interface{}`

GetRelatedObjectFilter returns the RelatedObjectFilter field if non-nil, zero value otherwise.

### GetRelatedObjectFilterOk

`func (o *WritableCustomFieldRequest) GetRelatedObjectFilterOk() (*interface{}, bool)`

GetRelatedObjectFilterOk returns a tuple with the RelatedObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectFilter

`func (o *WritableCustomFieldRequest) SetRelatedObjectFilter(v interface{})`

SetRelatedObjectFilter sets RelatedObjectFilter field to given value.

### HasRelatedObjectFilter

`func (o *WritableCustomFieldRequest) HasRelatedObjectFilter() bool`

HasRelatedObjectFilter returns a boolean if a field has been set.

### SetRelatedObjectFilterNil

`func (o *WritableCustomFieldRequest) SetRelatedObjectFilterNil(b bool)`

 SetRelatedObjectFilterNil sets the value for RelatedObjectFilter to be an explicit nil

### UnsetRelatedObjectFilter
`func (o *WritableCustomFieldRequest) UnsetRelatedObjectFilter()`

UnsetRelatedObjectFilter ensures that no value is present for RelatedObjectFilter, not even an explicit nil
### GetWeight

`func (o *WritableCustomFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WritableCustomFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WritableCustomFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WritableCustomFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetValidationMinimum

`func (o *WritableCustomFieldRequest) GetValidationMinimum() int64`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *WritableCustomFieldRequest) GetValidationMinimumOk() (*int64, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *WritableCustomFieldRequest) SetValidationMinimum(v int64)`

SetValidationMinimum sets ValidationMinimum field to given value.

### HasValidationMinimum

`func (o *WritableCustomFieldRequest) HasValidationMinimum() bool`

HasValidationMinimum returns a boolean if a field has been set.

### SetValidationMinimumNil

`func (o *WritableCustomFieldRequest) SetValidationMinimumNil(b bool)`

 SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil

### UnsetValidationMinimum
`func (o *WritableCustomFieldRequest) UnsetValidationMinimum()`

UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
### GetValidationMaximum

`func (o *WritableCustomFieldRequest) GetValidationMaximum() int64`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *WritableCustomFieldRequest) GetValidationMaximumOk() (*int64, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *WritableCustomFieldRequest) SetValidationMaximum(v int64)`

SetValidationMaximum sets ValidationMaximum field to given value.

### HasValidationMaximum

`func (o *WritableCustomFieldRequest) HasValidationMaximum() bool`

HasValidationMaximum returns a boolean if a field has been set.

### SetValidationMaximumNil

`func (o *WritableCustomFieldRequest) SetValidationMaximumNil(b bool)`

 SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil

### UnsetValidationMaximum
`func (o *WritableCustomFieldRequest) UnsetValidationMaximum()`

UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
### GetValidationRegex

`func (o *WritableCustomFieldRequest) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *WritableCustomFieldRequest) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *WritableCustomFieldRequest) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *WritableCustomFieldRequest) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetChoiceSet

`func (o *WritableCustomFieldRequest) GetChoiceSet() BriefCustomFieldChoiceSetRequest`

GetChoiceSet returns the ChoiceSet field if non-nil, zero value otherwise.

### GetChoiceSetOk

`func (o *WritableCustomFieldRequest) GetChoiceSetOk() (*BriefCustomFieldChoiceSetRequest, bool)`

GetChoiceSetOk returns a tuple with the ChoiceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceSet

`func (o *WritableCustomFieldRequest) SetChoiceSet(v BriefCustomFieldChoiceSetRequest)`

SetChoiceSet sets ChoiceSet field to given value.

### HasChoiceSet

`func (o *WritableCustomFieldRequest) HasChoiceSet() bool`

HasChoiceSet returns a boolean if a field has been set.

### SetChoiceSetNil

`func (o *WritableCustomFieldRequest) SetChoiceSetNil(b bool)`

 SetChoiceSetNil sets the value for ChoiceSet to be an explicit nil

### UnsetChoiceSet
`func (o *WritableCustomFieldRequest) UnsetChoiceSet()`

UnsetChoiceSet ensures that no value is present for ChoiceSet, not even an explicit nil
### GetComments

`func (o *WritableCustomFieldRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableCustomFieldRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableCustomFieldRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableCustomFieldRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


