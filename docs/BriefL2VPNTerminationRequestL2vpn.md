# BriefL2VPNTerminationRequestL2vpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Type** | Pointer to [**BriefL2VPNTypeValue**](BriefL2VPNTypeValue.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefL2VPNTerminationRequestL2vpn

`func NewBriefL2VPNTerminationRequestL2vpn(name string, slug string, ) *BriefL2VPNTerminationRequestL2vpn`

NewBriefL2VPNTerminationRequestL2vpn instantiates a new BriefL2VPNTerminationRequestL2vpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefL2VPNTerminationRequestL2vpnWithDefaults

`func NewBriefL2VPNTerminationRequestL2vpnWithDefaults() *BriefL2VPNTerminationRequestL2vpn`

NewBriefL2VPNTerminationRequestL2vpnWithDefaults instantiates a new BriefL2VPNTerminationRequestL2vpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *BriefL2VPNTerminationRequestL2vpn) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BriefL2VPNTerminationRequestL2vpn) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BriefL2VPNTerminationRequestL2vpn) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BriefL2VPNTerminationRequestL2vpn) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *BriefL2VPNTerminationRequestL2vpn) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *BriefL2VPNTerminationRequestL2vpn) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *BriefL2VPNTerminationRequestL2vpn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefL2VPNTerminationRequestL2vpn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefL2VPNTerminationRequestL2vpn) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefL2VPNTerminationRequestL2vpn) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefL2VPNTerminationRequestL2vpn) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefL2VPNTerminationRequestL2vpn) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *BriefL2VPNTerminationRequestL2vpn) GetType() BriefL2VPNTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BriefL2VPNTerminationRequestL2vpn) GetTypeOk() (*BriefL2VPNTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BriefL2VPNTerminationRequestL2vpn) SetType(v BriefL2VPNTypeValue)`

SetType sets Type field to given value.

### HasType

`func (o *BriefL2VPNTerminationRequestL2vpn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *BriefL2VPNTerminationRequestL2vpn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefL2VPNTerminationRequestL2vpn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefL2VPNTerminationRequestL2vpn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefL2VPNTerminationRequestL2vpn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


