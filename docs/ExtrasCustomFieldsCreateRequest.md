# ExtrasCustomFieldsCreateRequest

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
**ValidationMinimum** | Pointer to **NullableFloat64** | Minimum allowed value (for numeric fields) | [optional] 
**ValidationMaximum** | Pointer to **NullableFloat64** | Maximum allowed value (for numeric fields) | [optional] 
**ValidationRegex** | Pointer to **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. | [optional] 
**ChoiceSet** | Pointer to [**NullableCustomFieldRequestChoiceSet**](CustomFieldRequestChoiceSet.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewExtrasCustomFieldsCreateRequest

`func NewExtrasCustomFieldsCreateRequest(objectTypes []string, name string, ) *ExtrasCustomFieldsCreateRequest`

NewExtrasCustomFieldsCreateRequest instantiates a new ExtrasCustomFieldsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasCustomFieldsCreateRequestWithDefaults

`func NewExtrasCustomFieldsCreateRequestWithDefaults() *ExtrasCustomFieldsCreateRequest`

NewExtrasCustomFieldsCreateRequestWithDefaults instantiates a new ExtrasCustomFieldsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *ExtrasCustomFieldsCreateRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *ExtrasCustomFieldsCreateRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *ExtrasCustomFieldsCreateRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetType

`func (o *ExtrasCustomFieldsCreateRequest) GetType() PatchedWritableCustomFieldRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtrasCustomFieldsCreateRequest) GetTypeOk() (*PatchedWritableCustomFieldRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtrasCustomFieldsCreateRequest) SetType(v PatchedWritableCustomFieldRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *ExtrasCustomFieldsCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelatedObjectType

`func (o *ExtrasCustomFieldsCreateRequest) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *ExtrasCustomFieldsCreateRequest) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *ExtrasCustomFieldsCreateRequest) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *ExtrasCustomFieldsCreateRequest) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### SetRelatedObjectTypeNil

`func (o *ExtrasCustomFieldsCreateRequest) SetRelatedObjectTypeNil(b bool)`

 SetRelatedObjectTypeNil sets the value for RelatedObjectType to be an explicit nil

### UnsetRelatedObjectType
`func (o *ExtrasCustomFieldsCreateRequest) UnsetRelatedObjectType()`

UnsetRelatedObjectType ensures that no value is present for RelatedObjectType, not even an explicit nil
### GetName

`func (o *ExtrasCustomFieldsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtrasCustomFieldsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtrasCustomFieldsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ExtrasCustomFieldsCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtrasCustomFieldsCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtrasCustomFieldsCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ExtrasCustomFieldsCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGroupName

`func (o *ExtrasCustomFieldsCreateRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ExtrasCustomFieldsCreateRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ExtrasCustomFieldsCreateRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ExtrasCustomFieldsCreateRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDescription

`func (o *ExtrasCustomFieldsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtrasCustomFieldsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtrasCustomFieldsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtrasCustomFieldsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *ExtrasCustomFieldsCreateRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ExtrasCustomFieldsCreateRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ExtrasCustomFieldsCreateRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ExtrasCustomFieldsCreateRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnique

`func (o *ExtrasCustomFieldsCreateRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *ExtrasCustomFieldsCreateRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *ExtrasCustomFieldsCreateRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *ExtrasCustomFieldsCreateRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetSearchWeight

`func (o *ExtrasCustomFieldsCreateRequest) GetSearchWeight() int32`

GetSearchWeight returns the SearchWeight field if non-nil, zero value otherwise.

### GetSearchWeightOk

`func (o *ExtrasCustomFieldsCreateRequest) GetSearchWeightOk() (*int32, bool)`

GetSearchWeightOk returns a tuple with the SearchWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWeight

`func (o *ExtrasCustomFieldsCreateRequest) SetSearchWeight(v int32)`

SetSearchWeight sets SearchWeight field to given value.

### HasSearchWeight

`func (o *ExtrasCustomFieldsCreateRequest) HasSearchWeight() bool`

HasSearchWeight returns a boolean if a field has been set.

### GetFilterLogic

`func (o *ExtrasCustomFieldsCreateRequest) GetFilterLogic() PatchedWritableCustomFieldRequestFilterLogic`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *ExtrasCustomFieldsCreateRequest) GetFilterLogicOk() (*PatchedWritableCustomFieldRequestFilterLogic, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *ExtrasCustomFieldsCreateRequest) SetFilterLogic(v PatchedWritableCustomFieldRequestFilterLogic)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *ExtrasCustomFieldsCreateRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetUiVisible

`func (o *ExtrasCustomFieldsCreateRequest) GetUiVisible() PatchedWritableCustomFieldRequestUiVisible`

GetUiVisible returns the UiVisible field if non-nil, zero value otherwise.

### GetUiVisibleOk

`func (o *ExtrasCustomFieldsCreateRequest) GetUiVisibleOk() (*PatchedWritableCustomFieldRequestUiVisible, bool)`

GetUiVisibleOk returns a tuple with the UiVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVisible

`func (o *ExtrasCustomFieldsCreateRequest) SetUiVisible(v PatchedWritableCustomFieldRequestUiVisible)`

SetUiVisible sets UiVisible field to given value.

### HasUiVisible

`func (o *ExtrasCustomFieldsCreateRequest) HasUiVisible() bool`

HasUiVisible returns a boolean if a field has been set.

### GetUiEditable

`func (o *ExtrasCustomFieldsCreateRequest) GetUiEditable() PatchedWritableCustomFieldRequestUiEditable`

GetUiEditable returns the UiEditable field if non-nil, zero value otherwise.

### GetUiEditableOk

`func (o *ExtrasCustomFieldsCreateRequest) GetUiEditableOk() (*PatchedWritableCustomFieldRequestUiEditable, bool)`

GetUiEditableOk returns a tuple with the UiEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiEditable

`func (o *ExtrasCustomFieldsCreateRequest) SetUiEditable(v PatchedWritableCustomFieldRequestUiEditable)`

SetUiEditable sets UiEditable field to given value.

### HasUiEditable

`func (o *ExtrasCustomFieldsCreateRequest) HasUiEditable() bool`

HasUiEditable returns a boolean if a field has been set.

### GetIsCloneable

`func (o *ExtrasCustomFieldsCreateRequest) GetIsCloneable() bool`

GetIsCloneable returns the IsCloneable field if non-nil, zero value otherwise.

### GetIsCloneableOk

`func (o *ExtrasCustomFieldsCreateRequest) GetIsCloneableOk() (*bool, bool)`

GetIsCloneableOk returns a tuple with the IsCloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloneable

`func (o *ExtrasCustomFieldsCreateRequest) SetIsCloneable(v bool)`

SetIsCloneable sets IsCloneable field to given value.

### HasIsCloneable

`func (o *ExtrasCustomFieldsCreateRequest) HasIsCloneable() bool`

HasIsCloneable returns a boolean if a field has been set.

### GetDefault

`func (o *ExtrasCustomFieldsCreateRequest) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ExtrasCustomFieldsCreateRequest) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ExtrasCustomFieldsCreateRequest) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ExtrasCustomFieldsCreateRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *ExtrasCustomFieldsCreateRequest) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *ExtrasCustomFieldsCreateRequest) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetRelatedObjectFilter

`func (o *ExtrasCustomFieldsCreateRequest) GetRelatedObjectFilter() interface{}`

GetRelatedObjectFilter returns the RelatedObjectFilter field if non-nil, zero value otherwise.

### GetRelatedObjectFilterOk

`func (o *ExtrasCustomFieldsCreateRequest) GetRelatedObjectFilterOk() (*interface{}, bool)`

GetRelatedObjectFilterOk returns a tuple with the RelatedObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectFilter

`func (o *ExtrasCustomFieldsCreateRequest) SetRelatedObjectFilter(v interface{})`

SetRelatedObjectFilter sets RelatedObjectFilter field to given value.

### HasRelatedObjectFilter

`func (o *ExtrasCustomFieldsCreateRequest) HasRelatedObjectFilter() bool`

HasRelatedObjectFilter returns a boolean if a field has been set.

### SetRelatedObjectFilterNil

`func (o *ExtrasCustomFieldsCreateRequest) SetRelatedObjectFilterNil(b bool)`

 SetRelatedObjectFilterNil sets the value for RelatedObjectFilter to be an explicit nil

### UnsetRelatedObjectFilter
`func (o *ExtrasCustomFieldsCreateRequest) UnsetRelatedObjectFilter()`

UnsetRelatedObjectFilter ensures that no value is present for RelatedObjectFilter, not even an explicit nil
### GetWeight

`func (o *ExtrasCustomFieldsCreateRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ExtrasCustomFieldsCreateRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ExtrasCustomFieldsCreateRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ExtrasCustomFieldsCreateRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetValidationMinimum

`func (o *ExtrasCustomFieldsCreateRequest) GetValidationMinimum() float64`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *ExtrasCustomFieldsCreateRequest) GetValidationMinimumOk() (*float64, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *ExtrasCustomFieldsCreateRequest) SetValidationMinimum(v float64)`

SetValidationMinimum sets ValidationMinimum field to given value.

### HasValidationMinimum

`func (o *ExtrasCustomFieldsCreateRequest) HasValidationMinimum() bool`

HasValidationMinimum returns a boolean if a field has been set.

### SetValidationMinimumNil

`func (o *ExtrasCustomFieldsCreateRequest) SetValidationMinimumNil(b bool)`

 SetValidationMinimumNil sets the value for ValidationMinimum to be an explicit nil

### UnsetValidationMinimum
`func (o *ExtrasCustomFieldsCreateRequest) UnsetValidationMinimum()`

UnsetValidationMinimum ensures that no value is present for ValidationMinimum, not even an explicit nil
### GetValidationMaximum

`func (o *ExtrasCustomFieldsCreateRequest) GetValidationMaximum() float64`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *ExtrasCustomFieldsCreateRequest) GetValidationMaximumOk() (*float64, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *ExtrasCustomFieldsCreateRequest) SetValidationMaximum(v float64)`

SetValidationMaximum sets ValidationMaximum field to given value.

### HasValidationMaximum

`func (o *ExtrasCustomFieldsCreateRequest) HasValidationMaximum() bool`

HasValidationMaximum returns a boolean if a field has been set.

### SetValidationMaximumNil

`func (o *ExtrasCustomFieldsCreateRequest) SetValidationMaximumNil(b bool)`

 SetValidationMaximumNil sets the value for ValidationMaximum to be an explicit nil

### UnsetValidationMaximum
`func (o *ExtrasCustomFieldsCreateRequest) UnsetValidationMaximum()`

UnsetValidationMaximum ensures that no value is present for ValidationMaximum, not even an explicit nil
### GetValidationRegex

`func (o *ExtrasCustomFieldsCreateRequest) GetValidationRegex() string`

GetValidationRegex returns the ValidationRegex field if non-nil, zero value otherwise.

### GetValidationRegexOk

`func (o *ExtrasCustomFieldsCreateRequest) GetValidationRegexOk() (*string, bool)`

GetValidationRegexOk returns a tuple with the ValidationRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRegex

`func (o *ExtrasCustomFieldsCreateRequest) SetValidationRegex(v string)`

SetValidationRegex sets ValidationRegex field to given value.

### HasValidationRegex

`func (o *ExtrasCustomFieldsCreateRequest) HasValidationRegex() bool`

HasValidationRegex returns a boolean if a field has been set.

### GetChoiceSet

`func (o *ExtrasCustomFieldsCreateRequest) GetChoiceSet() CustomFieldRequestChoiceSet`

GetChoiceSet returns the ChoiceSet field if non-nil, zero value otherwise.

### GetChoiceSetOk

`func (o *ExtrasCustomFieldsCreateRequest) GetChoiceSetOk() (*CustomFieldRequestChoiceSet, bool)`

GetChoiceSetOk returns a tuple with the ChoiceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceSet

`func (o *ExtrasCustomFieldsCreateRequest) SetChoiceSet(v CustomFieldRequestChoiceSet)`

SetChoiceSet sets ChoiceSet field to given value.

### HasChoiceSet

`func (o *ExtrasCustomFieldsCreateRequest) HasChoiceSet() bool`

HasChoiceSet returns a boolean if a field has been set.

### SetChoiceSetNil

`func (o *ExtrasCustomFieldsCreateRequest) SetChoiceSetNil(b bool)`

 SetChoiceSetNil sets the value for ChoiceSet to be an explicit nil

### UnsetChoiceSet
`func (o *ExtrasCustomFieldsCreateRequest) UnsetChoiceSet()`

UnsetChoiceSet ensures that no value is present for ChoiceSet, not even an explicit nil
### GetOwner

`func (o *ExtrasCustomFieldsCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExtrasCustomFieldsCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExtrasCustomFieldsCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExtrasCustomFieldsCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ExtrasCustomFieldsCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ExtrasCustomFieldsCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *ExtrasCustomFieldsCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ExtrasCustomFieldsCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ExtrasCustomFieldsCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ExtrasCustomFieldsCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


