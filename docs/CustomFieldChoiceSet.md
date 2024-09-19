# CustomFieldChoiceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**BaseChoices** | Pointer to [**CustomFieldChoiceSetBaseChoices**](CustomFieldChoiceSetBaseChoices.md) |  | [optional] 
**ExtraChoices** | **[][]interface{}** |  | 
**OrderAlphabetically** | Pointer to **bool** | Choices are automatically ordered alphabetically | [optional] 
**ChoicesCount** | **string** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewCustomFieldChoiceSet

`func NewCustomFieldChoiceSet(id int32, url string, displayUrl string, display string, name string, extraChoices [][]interface{}, choicesCount string, created NullableTime, lastUpdated NullableTime, ) *CustomFieldChoiceSet`

NewCustomFieldChoiceSet instantiates a new CustomFieldChoiceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldChoiceSetWithDefaults

`func NewCustomFieldChoiceSetWithDefaults() *CustomFieldChoiceSet`

NewCustomFieldChoiceSetWithDefaults instantiates a new CustomFieldChoiceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldChoiceSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldChoiceSet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldChoiceSet) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CustomFieldChoiceSet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomFieldChoiceSet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomFieldChoiceSet) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *CustomFieldChoiceSet) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *CustomFieldChoiceSet) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *CustomFieldChoiceSet) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *CustomFieldChoiceSet) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CustomFieldChoiceSet) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CustomFieldChoiceSet) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *CustomFieldChoiceSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldChoiceSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldChoiceSet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomFieldChoiceSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomFieldChoiceSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomFieldChoiceSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomFieldChoiceSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseChoices

`func (o *CustomFieldChoiceSet) GetBaseChoices() CustomFieldChoiceSetBaseChoices`

GetBaseChoices returns the BaseChoices field if non-nil, zero value otherwise.

### GetBaseChoicesOk

`func (o *CustomFieldChoiceSet) GetBaseChoicesOk() (*CustomFieldChoiceSetBaseChoices, bool)`

GetBaseChoicesOk returns a tuple with the BaseChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseChoices

`func (o *CustomFieldChoiceSet) SetBaseChoices(v CustomFieldChoiceSetBaseChoices)`

SetBaseChoices sets BaseChoices field to given value.

### HasBaseChoices

`func (o *CustomFieldChoiceSet) HasBaseChoices() bool`

HasBaseChoices returns a boolean if a field has been set.

### GetExtraChoices

`func (o *CustomFieldChoiceSet) GetExtraChoices() [][]interface{}`

GetExtraChoices returns the ExtraChoices field if non-nil, zero value otherwise.

### GetExtraChoicesOk

`func (o *CustomFieldChoiceSet) GetExtraChoicesOk() (*[][]interface{}, bool)`

GetExtraChoicesOk returns a tuple with the ExtraChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraChoices

`func (o *CustomFieldChoiceSet) SetExtraChoices(v [][]interface{})`

SetExtraChoices sets ExtraChoices field to given value.


### GetOrderAlphabetically

`func (o *CustomFieldChoiceSet) GetOrderAlphabetically() bool`

GetOrderAlphabetically returns the OrderAlphabetically field if non-nil, zero value otherwise.

### GetOrderAlphabeticallyOk

`func (o *CustomFieldChoiceSet) GetOrderAlphabeticallyOk() (*bool, bool)`

GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAlphabetically

`func (o *CustomFieldChoiceSet) SetOrderAlphabetically(v bool)`

SetOrderAlphabetically sets OrderAlphabetically field to given value.

### HasOrderAlphabetically

`func (o *CustomFieldChoiceSet) HasOrderAlphabetically() bool`

HasOrderAlphabetically returns a boolean if a field has been set.

### GetChoicesCount

`func (o *CustomFieldChoiceSet) GetChoicesCount() string`

GetChoicesCount returns the ChoicesCount field if non-nil, zero value otherwise.

### GetChoicesCountOk

`func (o *CustomFieldChoiceSet) GetChoicesCountOk() (*string, bool)`

GetChoicesCountOk returns a tuple with the ChoicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoicesCount

`func (o *CustomFieldChoiceSet) SetChoicesCount(v string)`

SetChoicesCount sets ChoicesCount field to given value.


### GetCreated

`func (o *CustomFieldChoiceSet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomFieldChoiceSet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomFieldChoiceSet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CustomFieldChoiceSet) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CustomFieldChoiceSet) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CustomFieldChoiceSet) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CustomFieldChoiceSet) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CustomFieldChoiceSet) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CustomFieldChoiceSet) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CustomFieldChoiceSet) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


