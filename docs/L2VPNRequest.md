# L2VPNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Type** | Pointer to [**L2VPNTypeValue**](L2VPNTypeValue.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewL2VPNRequest

`func NewL2VPNRequest(name string, slug string, ) *L2VPNRequest`

NewL2VPNRequest instantiates a new L2VPNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL2VPNRequestWithDefaults

`func NewL2VPNRequestWithDefaults() *L2VPNRequest`

NewL2VPNRequestWithDefaults instantiates a new L2VPNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *L2VPNRequest) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *L2VPNRequest) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *L2VPNRequest) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *L2VPNRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *L2VPNRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *L2VPNRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *L2VPNRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *L2VPNRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *L2VPNRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *L2VPNRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *L2VPNRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *L2VPNRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *L2VPNRequest) GetType() L2VPNTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *L2VPNRequest) GetTypeOk() (*L2VPNTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *L2VPNRequest) SetType(v L2VPNTypeValue)`

SetType sets Type field to given value.

### HasType

`func (o *L2VPNRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *L2VPNRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *L2VPNRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *L2VPNRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *L2VPNRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


