# ObjectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**AppLabel** | **string** |  | 
**AppName** | **string** |  | [readonly] 
**Model** | **string** |  | 
**ModelName** | **string** |  | [readonly] 
**ModelNamePlural** | **string** |  | [readonly] 
**Public** | **bool** |  | [readonly] 
**Features** | **[]string** |  | [readonly] 
**IsPluginModel** | **bool** |  | [readonly] 
**RestApiEndpoint** | **string** |  | [readonly] 
**Description** | **string** |  | [readonly] 

## Methods

### NewObjectType

`func NewObjectType(id int32, url string, display string, appLabel string, appName string, model string, modelName string, modelNamePlural string, public bool, features []string, isPluginModel bool, restApiEndpoint string, description string, ) *ObjectType`

NewObjectType instantiates a new ObjectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectTypeWithDefaults

`func NewObjectTypeWithDefaults() *ObjectType`

NewObjectTypeWithDefaults instantiates a new ObjectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ObjectType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *ObjectType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ObjectType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ObjectType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetAppLabel

`func (o *ObjectType) GetAppLabel() string`

GetAppLabel returns the AppLabel field if non-nil, zero value otherwise.

### GetAppLabelOk

`func (o *ObjectType) GetAppLabelOk() (*string, bool)`

GetAppLabelOk returns a tuple with the AppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLabel

`func (o *ObjectType) SetAppLabel(v string)`

SetAppLabel sets AppLabel field to given value.


### GetAppName

`func (o *ObjectType) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ObjectType) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ObjectType) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetModel

`func (o *ObjectType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ObjectType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ObjectType) SetModel(v string)`

SetModel sets Model field to given value.


### GetModelName

`func (o *ObjectType) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ObjectType) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ObjectType) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelNamePlural

`func (o *ObjectType) GetModelNamePlural() string`

GetModelNamePlural returns the ModelNamePlural field if non-nil, zero value otherwise.

### GetModelNamePluralOk

`func (o *ObjectType) GetModelNamePluralOk() (*string, bool)`

GetModelNamePluralOk returns a tuple with the ModelNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNamePlural

`func (o *ObjectType) SetModelNamePlural(v string)`

SetModelNamePlural sets ModelNamePlural field to given value.


### GetPublic

`func (o *ObjectType) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ObjectType) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ObjectType) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetFeatures

`func (o *ObjectType) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ObjectType) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ObjectType) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetIsPluginModel

`func (o *ObjectType) GetIsPluginModel() bool`

GetIsPluginModel returns the IsPluginModel field if non-nil, zero value otherwise.

### GetIsPluginModelOk

`func (o *ObjectType) GetIsPluginModelOk() (*bool, bool)`

GetIsPluginModelOk returns a tuple with the IsPluginModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPluginModel

`func (o *ObjectType) SetIsPluginModel(v bool)`

SetIsPluginModel sets IsPluginModel field to given value.


### GetRestApiEndpoint

`func (o *ObjectType) GetRestApiEndpoint() string`

GetRestApiEndpoint returns the RestApiEndpoint field if non-nil, zero value otherwise.

### GetRestApiEndpointOk

`func (o *ObjectType) GetRestApiEndpointOk() (*string, bool)`

GetRestApiEndpointOk returns a tuple with the RestApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestApiEndpoint

`func (o *ObjectType) SetRestApiEndpoint(v string)`

SetRestApiEndpoint sets RestApiEndpoint field to given value.


### GetDescription

`func (o *ObjectType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectType) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


