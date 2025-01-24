# ConfigTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentParams** | Pointer to **interface{}** | Any &lt;a href&#x3D;\&quot;https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment\&quot;&gt;additional parameters&lt;/a&gt; to pass when constructing the Jinja2 environment. | [optional] 
**TemplateCode** | **string** | Jinja2 template code. | 
**DataSource** | Pointer to [**BriefDataSourceRequest**](BriefDataSourceRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewConfigTemplateRequest

`func NewConfigTemplateRequest(name string, templateCode string, ) *ConfigTemplateRequest`

NewConfigTemplateRequest instantiates a new ConfigTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigTemplateRequestWithDefaults

`func NewConfigTemplateRequestWithDefaults() *ConfigTemplateRequest`

NewConfigTemplateRequestWithDefaults instantiates a new ConfigTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentParams

`func (o *ConfigTemplateRequest) GetEnvironmentParams() interface{}`

GetEnvironmentParams returns the EnvironmentParams field if non-nil, zero value otherwise.

### GetEnvironmentParamsOk

`func (o *ConfigTemplateRequest) GetEnvironmentParamsOk() (*interface{}, bool)`

GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParams

`func (o *ConfigTemplateRequest) SetEnvironmentParams(v interface{})`

SetEnvironmentParams sets EnvironmentParams field to given value.

### HasEnvironmentParams

`func (o *ConfigTemplateRequest) HasEnvironmentParams() bool`

HasEnvironmentParams returns a boolean if a field has been set.

### SetEnvironmentParamsNil

`func (o *ConfigTemplateRequest) SetEnvironmentParamsNil(b bool)`

 SetEnvironmentParamsNil sets the value for EnvironmentParams to be an explicit nil

### UnsetEnvironmentParams
`func (o *ConfigTemplateRequest) UnsetEnvironmentParams()`

UnsetEnvironmentParams ensures that no value is present for EnvironmentParams, not even an explicit nil
### GetTemplateCode

`func (o *ConfigTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *ConfigTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *ConfigTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetDataSource

`func (o *ConfigTemplateRequest) GetDataSource() BriefDataSourceRequest`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ConfigTemplateRequest) GetDataSourceOk() (*BriefDataSourceRequest, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ConfigTemplateRequest) SetDataSource(v BriefDataSourceRequest)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ConfigTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetTags

`func (o *ConfigTemplateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigTemplateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigTemplateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


