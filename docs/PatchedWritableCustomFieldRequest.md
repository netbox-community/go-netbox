# PatchedWritableCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**PatchedWritableCustomFieldRequestType**](PatchedWritableCustomFieldRequestType.md) |  | [optional] 
**RelatedObjectType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** | Internal field name | [optional] 
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

### NewPatchedWritableCustomFieldRequest

`func NewPatchedWritableCustomFieldRequest() *PatchedWritableCustomFieldRequest`

NewPatchedWritableCustomFieldRequest instantiates a new PatchedWritableCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCustomFieldRequestWithDefaults

`func NewPatchedWritableCustomFieldRequestWithDefaults() *PatchedWritableCustomFieldRequest`

NewPatchedWritableCustomFieldRequestWithDefaults instantiates a new PatchedWritableCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *PatchedWritableCustomFieldRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedWritableCustomFieldRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedWritableCustomFieldRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedWritableCustomFieldRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableCustomFieldRequest) GetType() PatchedWritableCustomFieldRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableCustomFieldRequest) GetTypeOk() (*PatchedWritableCustomFieldRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableCustomFieldRequest) SetType(v PatchedWritableCustomFieldRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableCustomFieldRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelatedObjectType

`func (o *PatchedWritableCustomFieldRequest) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *PatchedWritableCustomFieldRequest) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *PatchedWritableCustomFieldRequest) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *PatchedWritableCustomFieldRequest) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### SetRelatedObjectTypeNil

`func (o *PatchedWritableCustomFieldRequest) SetRelatedObjectTypeNil(b bool)`

 SetRelatedObjectTypeNil sets the value for RelatedObjectType to be an explicit nil

### UnsetRelatedObjectType
`func (o *PatchedWritableCustomFieldRequest) UnsetRelatedObjectType()`

UnsetRelatedObjectType ensures that no value is present for RelatedObjectType, not even an explicit nil
### GetName

`func (o *PatchedWritableCustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableCustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableCustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableCustomFieldRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableCustomFieldRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableCustomFieldRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableCustomFieldRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableCustomFieldRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupName

`func (o *PatchedWritableCustomFieldRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PatchedWritableCustomFieldRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PatchedWritableCustomFieldRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PatchedWritableCustomFieldRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableCustomFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableCustomFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableCustomFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableCustomFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *PatchedWritableCustomFieldRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PatchedWritableCustomFieldRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PatchedWritableCustomFieldRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PatchedWritableCustomFieldRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnique

`func (o *PatchedWritableCustomFieldRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *PatchedWritableCustomFieldRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *PatchedWritableCustomFieldRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *PatchedWritableCustomFieldRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetSearchWeight

`func (o *PatchedWritableCustomFieldRequest) GetSearchWeight() int32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *PatchedWritableCustomFieldRequest) GetSearchWeightOk() (*int32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *PatchedWritableCustomFieldRequest) SetSearchWeight(v int32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *PatchedWritableCustomFieldRequest) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetFilterLogic

`func (o *PatchedWritableCustomFieldRequest) GetFilterLogic() PatchedWritableCustomFieldRequestFilterLogic`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *PatchedWritableCustomFieldRequest) GetFilterLogicOk() (*PatchedWritableCustomFieldRequestFilterLogic, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *PatchedWritableCustomFieldRequest) SetFilterLogic(v PatchedWritableCustomFieldRequestFilterLogic)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *PatchedWritableCustomFieldRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetUiVisible

`func (o *PatchedWritableCustomFieldRequest) GetUiVisible() PatchedWritableCustomFieldRequestUiVisible`

GetUiVisible returns the UiVisible field if non-nil, zero value otherwise.

### GetUiVisibleOk

`func (o *PatchedWritableCustomFieldRequest) GetUiVisibleOk() (*PatchedWritableCustomFieldRequestUiVisible, bool)`

GetUiVisibleOk returns a tuple with the UiVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVisible

`func (o *PatchedWritableCustomFieldRequest) SetUiVisible(v PatchedWritableCustomFieldRequestUiVisible)`

SetUiVisible sets UiVisible field to given value.

### HasUiVisible

`func (o *PatchedWritableCustomFieldRequest) HasUiVisible() bool`

HasUiVisible returns a boolean if a field has been set.

### GetUiEditable

`func (o *PatchedWritableCustomFieldRequest) GetUiEditable() PatchedWritableCustomFieldRequestUiEditable`

GetUiEditable returns the UiEditable field if non-nil, zero value otherwise.

### GetUiEditableOk

`func (o *PatchedWritableCustomFieldRequest) GetUiEditableOk() (*PatchedWritableCustomFieldRequestUiEditable, bool)`

GetUiEditableOk returns a tuple with the UiEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiEditable

`func (o *PatchedWritableCustomFieldRequest) SetUiEditable(v PatchedWritableCustomFieldRequestUiEditable)`

SetUiEditable sets UiEditable field to given value.

### HasUiEditable

`func (o *PatchedWritableCustomFieldRequest) HasUiEditable() bool`

HasUiEditable returns a boolean if a field has been set.

### GetIsCloneable

`func (o *PatchedWritableCustomFieldRequest) GetIsCloneable() bool`

GetIsCloneable returns the IsCloneable field if non-nil, zero value otherwise.

### GetIsCloneableOk

`func (o *PatchedWritableCustomFieldRequest) GetIsCloneableOk() (*bool, bool)`

GetIsCloneableOk returns a tuple with the IsCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloneable

`func (o *PatchedWritableCustomFieldRequest) SetIsCloneable(v bool)`

SetIsCloneable sets IsCloneable field to given value.

### HasIsCloneable

`func (o *PatchedWritableCustomFieldRequest) HasIsCloneable() bool`

HasIsCloneable returns a boolean if a field has been set.

### GetDefault

`func (o *PatchedWritableCustomFieldRequest) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PatchedWritableCustomFieldRequest) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PatchedWritableCustomFieldRequest) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PatchedWritableCustomFieldRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *PatchedWritableCustomFieldRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *PatchedWritableCustomFieldRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetRelatedObjectFilter

`func (o *PatchedWritableCustomFieldRequest) GetRelatedObjectFilter() interface{}`

GetRelatedObjectFilter returns the RelatedObjectFilter field if non-nil, zero value otherwise.

### GetRelatedObjectFilterOk

`func (o *PatchedWritableCustomFieldRequest) GetRelatedObjectFilterOk() (*interface{}, bool)`

GetRelatedObjectFilterOk returns a tuple with the RelatedObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectFilter

`func (o *PatchedWritableCustomFieldRequest) SetRelatedObjectFilter(v interface{})`

SetRelatedObjectFilter sets RelatedObjectFilter field to given value.

### HasRelatedObjectFilter

`func (o *PatchedWritableCustomFieldRequest) HasRelatedObjectFilter() bool`

HasRelatedObjectFilter returns a boolean if a field has been set.

### SetRelatedObjectFilterNil

`func (o *PatchedWritableCustomFieldRequest) SetRelatedObjectFilterNil(b bool)`

 SetRelatedObjectFilterNil sets the value for RelatedObjectFilter to be an explicit nil

### UnsetRelatedObjectFilter
`func (o *PatchedWritableCustomFieldRequest) UnsetRelatedObjectFilter()`

UnsetRelatedObjectFilter ensures that no value is present for RelatedObjectFilter, not even an explicit nil
### GetWeight

`func (o *PatchedWritableCustomFieldRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedWritableCustomFieldRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedWritableCustomFieldRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedWritableCustomFieldRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetValidationMinimum

`func (o *PatchedWritableCustomFieldRequest) GetValidationMinimum() int64`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *PatchedWritableCustomFieldRequest) GetValidationMinimumOk() (*int64, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *PatchedWritableCustomFieldRequest) SetValidationMinimum(v int64)`

SetValidationMinimum sets ValidationMinimum field to given value.

### HasValidationMinimum

`func (o *PatchedWritableCustomFieldRequest) HasValidationMinimum() bool`

HasValidationMinimum returns a boolean if a field has been set.

### SetValidationMinimumNil

`func (o *PatchedWritableCustomFieldRequest) SetValidationMinimumNil(b bool)`

 SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil

### UnsetValidationMinimum
`func (o *PatchedWritableCustomFieldRequest) UnsetValidationMinimum()`

UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
### GetValidationMaximum

`func (o *PatchedWritableCustomFieldRequest) GetValidationMaximum() int64`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *PatchedWritableCustomFieldRequest) GetValidationMaximumOk() (*int64, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *PatchedWritableCustomFieldRequest) SetValidationMaximum(v int64)`

SetValidationMaximum sets ValidationMaximum field to given value.

### HasValidationMaximum

`func (o *PatchedWritableCustomFieldRequest) HasValidationMaximum() bool`

HasValidationMaximum returns a boolean if a field has been set.

### SetValidationMaximumNil

`func (o *PatchedWritableCustomFieldRequest) SetValidationMaximumNil(b bool)`

 SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil

### UnsetValidationMaximum
`func (o *PatchedWritableCustomFieldRequest) UnsetValidationMaximum()`

UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
### GetValidationRegex

`func (o *PatchedWritableCustomFieldRequest) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *PatchedWritableCustomFieldRequest) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *PatchedWritableCustomFieldRequest) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *PatchedWritableCustomFieldRequest) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetChoiceSet

`func (o *PatchedWritableCustomFieldRequest) GetChoiceSet() BriefCustomFieldChoiceSetRequest`

GetChoiceSet returns the ChoiceSet field if non-nil, zero value otherwise.

### GetChoiceSetOk

`func (o *PatchedWritableCustomFieldRequest) GetChoiceSetOk() (*BriefCustomFieldChoiceSetRequest, bool)`

GetChoiceSetOk returns a tuple with the ChoiceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceSet

`func (o *PatchedWritableCustomFieldRequest) SetChoiceSet(v BriefCustomFieldChoiceSetRequest)`

SetChoiceSet sets ChoiceSet field to given value.

### HasChoiceSet

`func (o *PatchedWritableCustomFieldRequest) HasChoiceSet() bool`

HasChoiceSet returns a boolean if a field has been set.

### SetChoiceSetNil

`func (o *PatchedWritableCustomFieldRequest) SetChoiceSetNil(b bool)`

 SetChoiceSetNil sets the value for ChoiceSet to be an explicit nil

### UnsetChoiceSet
`func (o *PatchedWritableCustomFieldRequest) UnsetChoiceSet()`

UnsetChoiceSet ensures that no value is present for ChoiceSet, not even an explicit nil
### GetComments

`func (o *PatchedWritableCustomFieldRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableCustomFieldRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableCustomFieldRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableCustomFieldRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


