# NestedFHRPGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | * &#x60;vrrp2&#x60; - VRRPv2 * &#x60;vrrp3&#x60; - VRRPv3 * &#x60;carp&#x60; - CARP * &#x60;clusterxl&#x60; - ClusterXL * &#x60;hsrp&#x60; - HSRP * &#x60;glbp&#x60; - GLBP * &#x60;other&#x60; - Other | 
**GroupId** | **int32** |  | 

## Methods

### NewNestedFHRPGroupRequest

`func NewNestedFHRPGroupRequest(protocol string, groupId int32, ) *NestedFHRPGroupRequest`

NewNestedFHRPGroupRequest instantiates a new NestedFHRPGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedFHRPGroupRequestWithDefaults

`func NewNestedFHRPGroupRequestWithDefaults() *NestedFHRPGroupRequest`

NewNestedFHRPGroupRequestWithDefaults instantiates a new NestedFHRPGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *NestedFHRPGroupRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NestedFHRPGroupRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NestedFHRPGroupRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetGroupId

`func (o *NestedFHRPGroupRequest) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NestedFHRPGroupRequest) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NestedFHRPGroupRequest) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


