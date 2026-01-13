# VpnL2vpnTerminationsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L2vpn** | [**BriefL2VPNTerminationRequestL2vpn**](BriefL2VPNTerminationRequestL2vpn.md) |  | 
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int64** |  | 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnL2vpnTerminationsCreateRequest

`func NewVpnL2vpnTerminationsCreateRequest(l2vpn BriefL2VPNTerminationRequestL2vpn, assignedObjectType string, assignedObjectId int64, ) *VpnL2vpnTerminationsCreateRequest`

NewVpnL2vpnTerminationsCreateRequest instantiates a new VpnL2vpnTerminationsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnL2vpnTerminationsCreateRequestWithDefaults

`func NewVpnL2vpnTerminationsCreateRequestWithDefaults() *VpnL2vpnTerminationsCreateRequest`

NewVpnL2vpnTerminationsCreateRequestWithDefaults instantiates a new VpnL2vpnTerminationsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL2vpn

`func (o *VpnL2vpnTerminationsCreateRequest) GetL2vpn() BriefL2VPNTerminationRequestL2vpn`

GetL2vpn returns the L2vpn field if non-nil, zero value otherwise.

### GetL2vpnOk

`func (o *VpnL2vpnTerminationsCreateRequest) GetL2vpnOk() (*BriefL2VPNTerminationRequestL2vpn, bool)`

GetL2vpnOk returns a tuple with the L2vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2vpn

`func (o *VpnL2vpnTerminationsCreateRequest) SetL2vpn(v BriefL2VPNTerminationRequestL2vpn)`

SetL2vpn sets L2vpn field to given value.


### GetAssignedObjectType

`func (o *VpnL2vpnTerminationsCreateRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *VpnL2vpnTerminationsCreateRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *VpnL2vpnTerminationsCreateRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *VpnL2vpnTerminationsCreateRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *VpnL2vpnTerminationsCreateRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *VpnL2vpnTerminationsCreateRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetTags

`func (o *VpnL2vpnTerminationsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnL2vpnTerminationsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnL2vpnTerminationsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnL2vpnTerminationsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnL2vpnTerminationsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnL2vpnTerminationsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnL2vpnTerminationsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnL2vpnTerminationsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


