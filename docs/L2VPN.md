# L2VPN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Identifier** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Type** | Pointer to [**L2VPNType**](L2VPNType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewL2VPN

`func NewL2VPN(id int32, url string, display string, name string, slug string, ) *L2VPN`

NewL2VPN instantiates a new L2VPN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL2VPNWithDefaults

`func NewL2VPNWithDefaults() *L2VPN`

NewL2VPNWithDefaults instantiates a new L2VPN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *L2VPN) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *L2VPN) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *L2VPN) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *L2VPN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *L2VPN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *L2VPN) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *L2VPN) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *L2VPN) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *L2VPN) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetIdentifier

`func (o *L2VPN) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *L2VPN) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *L2VPN) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *L2VPN) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *L2VPN) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *L2VPN) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *L2VPN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *L2VPN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *L2VPN) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *L2VPN) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *L2VPN) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *L2VPN) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *L2VPN) GetType() L2VPNType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *L2VPN) GetTypeOk() (*L2VPNType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *L2VPN) SetType(v L2VPNType)`

SetType sets Type field to given value.

### HasType

`func (o *L2VPN) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *L2VPN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *L2VPN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *L2VPN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *L2VPN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


