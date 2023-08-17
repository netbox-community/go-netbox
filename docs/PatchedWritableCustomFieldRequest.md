# PatchedWritableCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | The type of data this custom field holds  * &#x60;text&#x60; - Text * &#x60;longtext&#x60; - Text (long) * &#x60;integer&#x60; - Integer * &#x60;decimal&#x60; - Decimal * &#x60;boolean&#x60; - Boolean (true/false) * &#x60;date&#x60; - Date * &#x60;datetime&#x60; - Date &amp; time * &#x60;url&#x60; - URL * &#x60;json&#x60; - JSON * &#x60;select&#x60; - Selection * &#x60;multiselect&#x60; - Multiple selection * &#x60;object&#x60; - Object * &#x60;multiobject&#x60; - Multiple objects | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Internal field name | [optional] 
**Label** | Pointer to **string** | Name of the field as displayed to users (if not provided, the field&#39;s name will be used) | [optional] 
**GroupName** | Pointer to **string** | Custom fields within the same group will be displayed together | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** | If true, this field is required when creating new objects or editing an existing object. | [optional] 
**SearchWeight** | Pointer to **int32** | Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored. | [optional] 
**FilterLogic** | Pointer to **string** | Loose matches any instance of a given string; exact matches the entire field.  * &#x60;disabled&#x60; - Disabled * &#x60;loose&#x60; - Loose * &#x60;exact&#x60; - Exact | [optional] 
**UiVisibility** | Pointer to **string** | Specifies the visibility of custom field in the UI  * &#x60;read-write&#x60; - Read/Write * &#x60;read-only&#x60; - Read-only * &#x60;hidden&#x60; - Hidden * &#x60;hidden-ifunset&#x60; - Hidden (if unset) | [optional] 
**IsCloneable** | Pointer to **bool** | Replicate this value when cloning objects | [optional] 
**Default** | Pointer to **map[string]interface{}** | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**Weight** | Pointer to **int32** | Fields with higher weights appear lower in a form. | [optional] 
**ValidationMinimum** | Pointer to **NullableInt32** | Minimum allowed value (for numeric fields) | [optional] 
**ValidationMaximum** | Pointer to **NullableInt32** | Maximum allowed value (for numeric fields) | [optional] 
**ValidationRegex** | Pointer to **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. | [optional] 
**Choices** | Pointer to **[]string** | Comma-separated list of available choices (for selection fields) | [optional] 

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

### GetContentTypes

`func (o *PatchedWritableCustomFieldRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedWritableCustomFieldRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedWritableCustomFieldRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedWritableCustomFieldRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableCustomFieldRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableCustomFieldRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableCustomFieldRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableCustomFieldRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObjectType

`func (o *PatchedWritableCustomFieldRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedWritableCustomFieldRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedWritableCustomFieldRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedWritableCustomFieldRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

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

`func (o *PatchedWritableCustomFieldRequest) GetFilterLogic() string`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *PatchedWritableCustomFieldRequest) GetFilterLogicOk() (*string, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *PatchedWritableCustomFieldRequest) SetFilterLogic(v string)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *PatchedWritableCustomFieldRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetUiVisibility

`func (o *PatchedWritableCustomFieldRequest) GetUiVisibility() string`

GetUiVisibility returns the UiVisibility field if non-nil, zero value otherwise.

### GetUiVisibilityOk

`func (o *PatchedWritableCustomFieldRequest) GetUiVisibilityOk() (*string, bool)`

GetUiVisibilityOk returns a tuple with the UiVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVisibility

`func (o *PatchedWritableCustomFieldRequest) SetUiVisibility(v string)`

SetUiVisibility sets UiVisibility field to given value.

### HasUiVisibility

`func (o *PatchedWritableCustomFieldRequest) HasUiVisibility() bool`

HasUiVisibility returns a boolean if a field has been set.

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

`func (o *PatchedWritableCustomFieldRequest) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PatchedWritableCustomFieldRequest) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PatchedWritableCustomFieldRequest) SetDefault(v map[string]interface{})`

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

`func (o *PatchedWritableCustomFieldRequest) GetValidationMinimum() int32`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *PatchedWritableCustomFieldRequest) GetValidationMinimumOk() (*int32, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *PatchedWritableCustomFieldRequest) SetValidationMinimum(v int32)`

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

`func (o *PatchedWritableCustomFieldRequest) GetValidationMaximum() int32`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *PatchedWritableCustomFieldRequest) GetValidationMaximumOk() (*int32, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *PatchedWritableCustomFieldRequest) SetValidationMaximum(v int32)`

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

### GetChoices

`func (o *PatchedWritableCustomFieldRequest) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *PatchedWritableCustomFieldRequest) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *PatchedWritableCustomFieldRequest) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *PatchedWritableCustomFieldRequest) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *PatchedWritableCustomFieldRequest) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *PatchedWritableCustomFieldRequest) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


