# WritableConfigTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentParams** | Pointer to **map[string]interface{}** | Any &lt;a href&#x3D;\&quot;https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment\&quot;&gt;additional parameters&lt;/a&gt; to pass when constructing the Jinja2 environment. | [optional] 
**TemplateCode** | **string** | Jinja2 template code. | 
**DataSource** | Pointer to **NullableInt32** | Remote data source | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewWritableConfigTemplateRequest

`func NewWritableConfigTemplateRequest(name string, templateCode string, ) *WritableConfigTemplateRequest`

NewWritableConfigTemplateRequest instantiates a new WritableConfigTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableConfigTemplateRequestWithDefaults

`func NewWritableConfigTemplateRequestWithDefaults() *WritableConfigTemplateRequest`

NewWritableConfigTemplateRequestWithDefaults instantiates a new WritableConfigTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableConfigTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableConfigTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableConfigTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableConfigTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableConfigTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableConfigTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableConfigTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentParams

`func (o *WritableConfigTemplateRequest) GetEnvironmentParams() map[string]interface{}`

GetEnvironmentParams returns the EnvironmentParams field if non-nil, zero value otherwise.

### GetEnvironmentParamsOk

`func (o *WritableConfigTemplateRequest) GetEnvironmentParamsOk() (*map[string]interface{}, bool)`

GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParams

`func (o *WritableConfigTemplateRequest) SetEnvironmentParams(v map[string]interface{})`

SetEnvironmentParams sets EnvironmentParams field to given value.

### HasEnvironmentParams

`func (o *WritableConfigTemplateRequest) HasEnvironmentParams() bool`

HasEnvironmentParams returns a boolean if a field has been set.

### SetEnvironmentParamsNil

`func (o *WritableConfigTemplateRequest) SetEnvironmentParamsNil(b bool)`

 SetEnvironmentParamsNil sets the value for EnvironmentParams to be an explicit nil

### UnsetEnvironmentParams
`func (o *WritableConfigTemplateRequest) UnsetEnvironmentParams()`

UnsetEnvironmentParams ensures that no value is present for EnvironmentParams, not even an explicit nil
### GetTemplateCode

`func (o *WritableConfigTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *WritableConfigTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *WritableConfigTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetDataSource

`func (o *WritableConfigTemplateRequest) GetDataSource() int32`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *WritableConfigTemplateRequest) GetDataSourceOk() (*int32, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *WritableConfigTemplateRequest) SetDataSource(v int32)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *WritableConfigTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *WritableConfigTemplateRequest) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *WritableConfigTemplateRequest) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
### GetTags

`func (o *WritableConfigTemplateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableConfigTemplateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableConfigTemplateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableConfigTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


