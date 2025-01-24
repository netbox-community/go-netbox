# RouteTargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Route target value (formatted in accordance with RFC 4360) | 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRouteTargetRequest

`func NewRouteTargetRequest(name string, ) *RouteTargetRequest`

NewRouteTargetRequest instantiates a new RouteTargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTargetRequestWithDefaults

`func NewRouteTargetRequestWithDefaults() *RouteTargetRequest`

NewRouteTargetRequestWithDefaults instantiates a new RouteTargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouteTargetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteTargetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteTargetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTenant

`func (o *RouteTargetRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RouteTargetRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RouteTargetRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *RouteTargetRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *RouteTargetRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *RouteTargetRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *RouteTargetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteTargetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteTargetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteTargetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *RouteTargetRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RouteTargetRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RouteTargetRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RouteTargetRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *RouteTargetRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RouteTargetRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RouteTargetRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RouteTargetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *RouteTargetRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RouteTargetRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RouteTargetRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RouteTargetRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


