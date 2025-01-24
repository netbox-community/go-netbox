# VLANTranslationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Rules** | [**[]VLANTranslationRule**](VLANTranslationRule.md) |  | [readonly] 

## Methods

### NewVLANTranslationPolicy

`func NewVLANTranslationPolicy(id int32, url string, display string, name string, rules []VLANTranslationRule, ) *VLANTranslationPolicy`

NewVLANTranslationPolicy instantiates a new VLANTranslationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVLANTranslationPolicyWithDefaults

`func NewVLANTranslationPolicyWithDefaults() *VLANTranslationPolicy`

NewVLANTranslationPolicyWithDefaults instantiates a new VLANTranslationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VLANTranslationPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VLANTranslationPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VLANTranslationPolicy) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *VLANTranslationPolicy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VLANTranslationPolicy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VLANTranslationPolicy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *VLANTranslationPolicy) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VLANTranslationPolicy) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VLANTranslationPolicy) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *VLANTranslationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VLANTranslationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VLANTranslationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VLANTranslationPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VLANTranslationPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VLANTranslationPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VLANTranslationPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRules

`func (o *VLANTranslationPolicy) GetRules() []VLANTranslationRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *VLANTranslationPolicy) GetRulesOk() (*[]VLANTranslationRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *VLANTranslationPolicy) SetRules(v []VLANTranslationRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


