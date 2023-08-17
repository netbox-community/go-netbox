# CustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | **[]string** |  | 
**Type** | **string** | * &#x60;text&#x60; - Text * &#x60;longtext&#x60; - Text (long) * &#x60;integer&#x60; - Integer * &#x60;decimal&#x60; - Decimal * &#x60;boolean&#x60; - Boolean (true/false) * &#x60;date&#x60; - Date * &#x60;datetime&#x60; - Date &amp; time * &#x60;url&#x60; - URL * &#x60;json&#x60; - JSON * &#x60;select&#x60; - Selection * &#x60;multiselect&#x60; - Multiple selection * &#x60;object&#x60; - Object * &#x60;multiobject&#x60; - Multiple objects | 
**ObjectType** | Pointer to **string** |  | [optional] 
**Name** | **string** | Internal field name | 
**Label** | Pointer to **string** | Name of the field as displayed to users (if not provided, the field&#39;s name will be used) | [optional] 
**GroupName** | Pointer to **string** | Custom fields within the same group will be displayed together | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** | If true, this field is required when creating new objects or editing an existing object. | [optional] 
**SearchWeight** | Pointer to **int32** | Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored. | [optional] 
**FilterLogic** | Pointer to **string** | * &#x60;disabled&#x60; - Disabled * &#x60;loose&#x60; - Loose * &#x60;exact&#x60; - Exact | [optional] 
**UiVisibility** | Pointer to **string** | * &#x60;read-write&#x60; - Read/Write * &#x60;read-only&#x60; - Read-only * &#x60;hidden&#x60; - Hidden * &#x60;hidden-ifunset&#x60; - Hidden (if unset) | [optional] 
**IsCloneable** | Pointer to **bool** | Replicate this value when cloning objects | [optional] 
**Default** | Pointer to **map[string]interface{}** | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] 
**Weight** | Pointer to **int32** | Fields with higher weights appear lower in a form. | [optional] 
**ValidationMinimum** | Pointer to **NullableInt32** | Minimum allowed value (for numeric fields) | [optional] 
**ValidationMaximum** | Pointer to **NullableInt32** | Maximum allowed value (for numeric fields) | [optional] 
**ValidationRegex** | Pointer to **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. | [optional] 
**Choices** | Pointer to **[]string** | Comma-separated list of available choices (for selection fields) | [optional] 

## Methods

### NewCustomFieldRequest

`func NewCustomFieldRequest(contentTypes []string, type_ string, name string, ) *CustomFieldRequest`

NewCustomFieldRequest instantiates a new CustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldRequestWithDefaults

`func NewCustomFieldRequestWithDefaults() *CustomFieldRequest`

NewCustomFieldRequestWithDefaults instantiates a new CustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *CustomFieldRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *CustomFieldRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *CustomFieldRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetType

`func (o *CustomFieldRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldRequest) SetType(v string)`

SetType sets Type field to given value.


### GetObjectType

`func (o *CustomFieldRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CustomFieldRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CustomFieldRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CustomFieldRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

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

`func (o *CustomFieldRequest) GetFilterLogic() string`

GetFilterLogic returns the FilterLogic field if non-nil, zero value otherwise.

### GetFilterLogicOk

`func (o *CustomFieldRequest) GetFilterLogicOk() (*string, bool)`

GetFilterLogicOk returns a tuple with the FilterLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLogic

`func (o *CustomFieldRequest) SetFilterLogic(v string)`

SetFilterLogic sets FilterLogic field to given value.

### HasFilterLogic

`func (o *CustomFieldRequest) HasFilterLogic() bool`

HasFilterLogic returns a boolean if a field has been set.

### GetUiVisibility

`func (o *CustomFieldRequest) GetUiVisibility() string`

GetUiVisibility returns the UiVisibility field if non-nil, zero value otherwise.

### GetUiVisibilityOk

`func (o *CustomFieldRequest) GetUiVisibilityOk() (*string, bool)`

GetUiVisibilityOk returns a tuple with the UiVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVisibility

`func (o *CustomFieldRequest) SetUiVisibility(v string)`

SetUiVisibility sets UiVisibility field to given value.

### HasUiVisibility

`func (o *CustomFieldRequest) HasUiVisibility() bool`

HasUiVisibility returns a boolean if a field has been set.

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

`func (o *CustomFieldRequest) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomFieldRequest) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomFieldRequest) SetDefault(v map[string]interface{})`

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

`func (o *CustomFieldRequest) GetValidationMinimum() int32`

GetValidationMinimum returns the ValidationMinimum field if non-nil, zero value otherwise.

### GetValidationMinimumOk

`func (o *CustomFieldRequest) GetValidationMinimumOk() (*int32, bool)`

GetValidationMinimumOk returns a tuple with the ValidationMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMinimum

`func (o *CustomFieldRequest) SetValidationMinimum(v int32)`

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

`func (o *CustomFieldRequest) GetValidationMaximum() int32`

GetValidationMaximum returns the ValidationMaximum field if non-nil, zero value otherwise.

### GetValidationMaximumOk

`func (o *CustomFieldRequest) GetValidationMaximumOk() (*int32, bool)`

GetValidationMaximumOk returns a tuple with the ValidationMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMaximum

`func (o *CustomFieldRequest) SetValidationMaximum(v int32)`

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

### GetChoices

`func (o *CustomFieldRequest) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *CustomFieldRequest) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *CustomFieldRequest) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *CustomFieldRequest) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *CustomFieldRequest) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *CustomFieldRequest) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


