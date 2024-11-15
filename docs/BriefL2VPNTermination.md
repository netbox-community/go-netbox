# BriefL2VPNTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**L2vpn** | [**BriefL2VPN**](BriefL2VPN.md) |  | 

## Methods

### NewBriefL2VPNTermination

`func NewBriefL2VPNTermination(id int32, url string, display string, l2vpn BriefL2VPN, ) *BriefL2VPNTermination`

NewBriefL2VPNTermination instantiates a new BriefL2VPNTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefL2VPNTerminationWithDefaults

`func NewBriefL2VPNTerminationWithDefaults() *BriefL2VPNTermination`

NewBriefL2VPNTerminationWithDefaults instantiates a new BriefL2VPNTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefL2VPNTermination) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefL2VPNTermination) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefL2VPNTermination) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefL2VPNTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefL2VPNTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefL2VPNTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefL2VPNTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefL2VPNTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefL2VPNTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetL2vpn

`func (o *BriefL2VPNTermination) GetL2vpn() BriefL2VPN`

GetL2vpn returns the L2vpn field if non-nil, zero value otherwise.

### GetL2vpnOk

`func (o *BriefL2VPNTermination) GetL2vpnOk() (*BriefL2VPN, bool)`

GetL2vpnOk returns a tuple with the L2vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2vpn

`func (o *BriefL2VPNTermination) SetL2vpn(v BriefL2VPN)`

SetL2vpn sets L2vpn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


