# BriefVirtualCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Cid** | **string** | Unique circuit ID | 
**ProviderNetwork** | [**BriefProviderNetwork**](BriefProviderNetwork.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefVirtualCircuit

`func NewBriefVirtualCircuit(id int32, url string, display string, cid string, providerNetwork BriefProviderNetwork, ) *BriefVirtualCircuit`

NewBriefVirtualCircuit instantiates a new BriefVirtualCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefVirtualCircuitWithDefaults

`func NewBriefVirtualCircuitWithDefaults() *BriefVirtualCircuit`

NewBriefVirtualCircuitWithDefaults instantiates a new BriefVirtualCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefVirtualCircuit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefVirtualCircuit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefVirtualCircuit) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefVirtualCircuit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefVirtualCircuit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefVirtualCircuit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefVirtualCircuit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefVirtualCircuit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefVirtualCircuit) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetCid

`func (o *BriefVirtualCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *BriefVirtualCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *BriefVirtualCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProviderNetwork

`func (o *BriefVirtualCircuit) GetProviderNetwork() BriefProviderNetwork`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *BriefVirtualCircuit) GetProviderNetworkOk() (*BriefProviderNetwork, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *BriefVirtualCircuit) SetProviderNetwork(v BriefProviderNetwork)`

SetProviderNetwork sets ProviderNetwork field to given value.


### GetDescription

`func (o *BriefVirtualCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefVirtualCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefVirtualCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefVirtualCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


