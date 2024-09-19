# BriefL2VPN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Identifier** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Type** | Pointer to [**BriefL2VPNType**](BriefL2VPNType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefL2VPN

`func NewBriefL2VPN(id int32, url string, display string, name string, slug string, ) *BriefL2VPN`

NewBriefL2VPN instantiates a new BriefL2VPN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefL2VPNWithDefaults

`func NewBriefL2VPNWithDefaults() *BriefL2VPN`

NewBriefL2VPNWithDefaults instantiates a new BriefL2VPN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefL2VPN) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefL2VPN) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefL2VPN) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefL2VPN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefL2VPN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefL2VPN) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefL2VPN) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefL2VPN) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefL2VPN) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetIdentifier

`func (o *BriefL2VPN) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BriefL2VPN) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BriefL2VPN) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BriefL2VPN) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *BriefL2VPN) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *BriefL2VPN) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *BriefL2VPN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefL2VPN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefL2VPN) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefL2VPN) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefL2VPN) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefL2VPN) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *BriefL2VPN) GetType() BriefL2VPNType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BriefL2VPN) GetTypeOk() (*BriefL2VPNType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BriefL2VPN) SetType(v BriefL2VPNType)`

SetType sets Type field to given value.

### HasType

`func (o *BriefL2VPN) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *BriefL2VPN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefL2VPN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefL2VPN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefL2VPN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


