# DcimCablesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableCableType**](CableType.md) |  | [optional] 
**ATerminations** | Pointer to [**[]GenericObjectRequest**](GenericObjectRequest.md) |  | [optional] 
**BTerminations** | Pointer to [**[]GenericObjectRequest**](GenericObjectRequest.md) |  | [optional] 
**Status** | Pointer to [**CableStatusValue**](CableStatusValue.md) |  | [optional] 
**Profile** | Pointer to [**PatchedWritableCableRequestProfile**](PatchedWritableCableRequestProfile.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **NullableFloat64** |  | [optional] 
**LengthUnit** | Pointer to [**NullableCableRequestLengthUnit**](CableRequestLengthUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimCablesCreateRequest

`func NewDcimCablesCreateRequest() *DcimCablesCreateRequest`

NewDcimCablesCreateRequest instantiates a new DcimCablesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimCablesCreateRequestWithDefaults

`func NewDcimCablesCreateRequestWithDefaults() *DcimCablesCreateRequest`

NewDcimCablesCreateRequestWithDefaults instantiates a new DcimCablesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DcimCablesCreateRequest) GetType() CableType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcimCablesCreateRequest) GetTypeOk() (*CableType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcimCablesCreateRequest) SetType(v CableType)`

SetType sets Type field to given value.

### HasType

`func (o *DcimCablesCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DcimCablesCreateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DcimCablesCreateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetATerminations

`func (o *DcimCablesCreateRequest) GetATerminations() []GenericObjectRequest`

GetATerminations returns the ATerminations field if non-nil, zero value otherwise.

### GetATerminationsOk

`func (o *DcimCablesCreateRequest) GetATerminationsOk() (*[]GenericObjectRequest, bool)`

GetATerminationsOk returns a tuple with the ATerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetATerminations

`func (o *DcimCablesCreateRequest) SetATerminations(v []GenericObjectRequest)`

SetATerminations sets ATerminations field to given value.

### HasATerminations

`func (o *DcimCablesCreateRequest) HasATerminations() bool`

HasATerminations returns a boolean if a field has been set.

### GetBTerminations

`func (o *DcimCablesCreateRequest) GetBTerminations() []GenericObjectRequest`

GetBTerminations returns the BTerminations field if non-nil, zero value otherwise.

### GetBTerminationsOk

`func (o *DcimCablesCreateRequest) GetBTerminationsOk() (*[]GenericObjectRequest, bool)`

GetBTerminationsOk returns a tuple with the BTerminations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBTerminations

`func (o *DcimCablesCreateRequest) SetBTerminations(v []GenericObjectRequest)`

SetBTerminations sets BTerminations field to given value.

### HasBTerminations

`func (o *DcimCablesCreateRequest) HasBTerminations() bool`

HasBTerminations returns a boolean if a field has been set.

### GetStatus

`func (o *DcimCablesCreateRequest) GetStatus() CableStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DcimCablesCreateRequest) GetStatusOk() (*CableStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DcimCablesCreateRequest) SetStatus(v CableStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DcimCablesCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProfile

`func (o *DcimCablesCreateRequest) GetProfile() PatchedWritableCableRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DcimCablesCreateRequest) GetProfileOk() (*PatchedWritableCableRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DcimCablesCreateRequest) SetProfile(v PatchedWritableCableRequestProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *DcimCablesCreateRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTenant

`func (o *DcimCablesCreateRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DcimCablesCreateRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DcimCablesCreateRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DcimCablesCreateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *DcimCablesCreateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *DcimCablesCreateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLabel

`func (o *DcimCablesCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DcimCablesCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DcimCablesCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DcimCablesCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetColor

`func (o *DcimCablesCreateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DcimCablesCreateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DcimCablesCreateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DcimCablesCreateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLength

`func (o *DcimCablesCreateRequest) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *DcimCablesCreateRequest) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *DcimCablesCreateRequest) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *DcimCablesCreateRequest) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *DcimCablesCreateRequest) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *DcimCablesCreateRequest) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetLengthUnit

`func (o *DcimCablesCreateRequest) GetLengthUnit() CableRequestLengthUnit`

GetLengthUnit returns the LengthUnit field if non-nil, zero value otherwise.

### GetLengthUnitOk

`func (o *DcimCablesCreateRequest) GetLengthUnitOk() (*CableRequestLengthUnit, bool)`

GetLengthUnitOk returns a tuple with the LengthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnit

`func (o *DcimCablesCreateRequest) SetLengthUnit(v CableRequestLengthUnit)`

SetLengthUnit sets LengthUnit field to given value.

### HasLengthUnit

`func (o *DcimCablesCreateRequest) HasLengthUnit() bool`

HasLengthUnit returns a boolean if a field has been set.

### SetLengthUnitNil

`func (o *DcimCablesCreateRequest) SetLengthUnitNil(b bool)`

 SetLengthUnitNil sets the value for LengthUnit to be an explicit nil

### UnsetLengthUnit
`func (o *DcimCablesCreateRequest) UnsetLengthUnit()`

UnsetLengthUnit ensures that no value is present for LengthUnit, not even an explicit nil
### GetDescription

`func (o *DcimCablesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimCablesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimCablesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimCablesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *DcimCablesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimCablesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimCablesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimCablesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimCablesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimCablesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *DcimCablesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DcimCablesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DcimCablesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DcimCablesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *DcimCablesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimCablesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimCablesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimCablesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimCablesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimCablesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimCablesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimCablesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


