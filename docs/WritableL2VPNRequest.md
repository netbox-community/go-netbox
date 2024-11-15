# WritableL2VPNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Type** | [**BriefL2VPNTypeValue**](BriefL2VPNTypeValue.md) |  | 
**ImportTargets** | Pointer to **[]int32** |  | [optional] 
**ExportTargets** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableL2VPNRequest

`func NewWritableL2VPNRequest(name string, slug string, type_ BriefL2VPNTypeValue, ) *WritableL2VPNRequest`

NewWritableL2VPNRequest instantiates a new WritableL2VPNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableL2VPNRequestWithDefaults

`func NewWritableL2VPNRequestWithDefaults() *WritableL2VPNRequest`

NewWritableL2VPNRequestWithDefaults instantiates a new WritableL2VPNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *WritableL2VPNRequest) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *WritableL2VPNRequest) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *WritableL2VPNRequest) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *WritableL2VPNRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *WritableL2VPNRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *WritableL2VPNRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *WritableL2VPNRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableL2VPNRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableL2VPNRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WritableL2VPNRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WritableL2VPNRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WritableL2VPNRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *WritableL2VPNRequest) GetType() BriefL2VPNTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableL2VPNRequest) GetTypeOk() (*BriefL2VPNTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableL2VPNRequest) SetType(v BriefL2VPNTypeValue)`

SetType sets Type field to given value.


### GetImportTargets

`func (o *WritableL2VPNRequest) GetImportTargets() []int32`

GetImportTargets returns the ImportTargets field if non-nil, zero value otherwise.

### GetImportTargetsOk

`func (o *WritableL2VPNRequest) GetImportTargetsOk() (*[]int32, bool)`

GetImportTargetsOk returns a tuple with the ImportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportTargets

`func (o *WritableL2VPNRequest) SetImportTargets(v []int32)`

SetImportTargets sets ImportTargets field to given value.

### HasImportTargets

`func (o *WritableL2VPNRequest) HasImportTargets() bool`

HasImportTargets returns a boolean if a field has been set.

### GetExportTargets

`func (o *WritableL2VPNRequest) GetExportTargets() []int32`

GetExportTargets returns the ExportTargets field if non-nil, zero value otherwise.

### GetExportTargetsOk

`func (o *WritableL2VPNRequest) GetExportTargetsOk() (*[]int32, bool)`

GetExportTargetsOk returns a tuple with the ExportTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTargets

`func (o *WritableL2VPNRequest) SetExportTargets(v []int32)`

SetExportTargets sets ExportTargets field to given value.

### HasExportTargets

`func (o *WritableL2VPNRequest) HasExportTargets() bool`

HasExportTargets returns a boolean if a field has been set.

### GetDescription

`func (o *WritableL2VPNRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableL2VPNRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableL2VPNRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableL2VPNRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableL2VPNRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableL2VPNRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableL2VPNRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableL2VPNRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTenant

`func (o *WritableL2VPNRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableL2VPNRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableL2VPNRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableL2VPNRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableL2VPNRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableL2VPNRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTags

`func (o *WritableL2VPNRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableL2VPNRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableL2VPNRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableL2VPNRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableL2VPNRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableL2VPNRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableL2VPNRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableL2VPNRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


