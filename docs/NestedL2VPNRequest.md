# NestedL2VPNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Type** | **string** | * &#x60;vpws&#x60; - VPWS * &#x60;vpls&#x60; - VPLS * &#x60;vxlan&#x60; - VXLAN * &#x60;vxlan-evpn&#x60; - VXLAN-EVPN * &#x60;mpls-evpn&#x60; - MPLS EVPN * &#x60;pbb-evpn&#x60; - PBB EVPN * &#x60;epl&#x60; - EPL * &#x60;evpl&#x60; - EVPL * &#x60;ep-lan&#x60; - Ethernet Private LAN * &#x60;evp-lan&#x60; - Ethernet Virtual Private LAN * &#x60;ep-tree&#x60; - Ethernet Private Tree * &#x60;evp-tree&#x60; - Ethernet Virtual Private Tree | 

## Methods

### NewNestedL2VPNRequest

`func NewNestedL2VPNRequest(name string, slug string, type_ string, ) *NestedL2VPNRequest`

NewNestedL2VPNRequest instantiates a new NestedL2VPNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedL2VPNRequestWithDefaults

`func NewNestedL2VPNRequestWithDefaults() *NestedL2VPNRequest`

NewNestedL2VPNRequestWithDefaults instantiates a new NestedL2VPNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *NestedL2VPNRequest) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NestedL2VPNRequest) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NestedL2VPNRequest) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *NestedL2VPNRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *NestedL2VPNRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *NestedL2VPNRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *NestedL2VPNRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedL2VPNRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedL2VPNRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedL2VPNRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedL2VPNRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedL2VPNRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *NestedL2VPNRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NestedL2VPNRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NestedL2VPNRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


