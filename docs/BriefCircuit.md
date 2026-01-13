# BriefCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Cid** | **string** | Unique circuit ID | 
**Provider** | [**BriefProvider**](BriefProvider.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefCircuit

`func NewBriefCircuit(id int32, url string, display string, cid string, provider BriefProvider, ) *BriefCircuit`

NewBriefCircuit instantiates a new BriefCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefCircuitWithDefaults

`func NewBriefCircuitWithDefaults() *BriefCircuit`

NewBriefCircuitWithDefaults instantiates a new BriefCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefCircuit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefCircuit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefCircuit) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefCircuit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefCircuit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefCircuit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefCircuit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefCircuit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefCircuit) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetCid

`func (o *BriefCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *BriefCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *BriefCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *BriefCircuit) GetProvider() BriefProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BriefCircuit) GetProviderOk() (*BriefProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BriefCircuit) SetProvider(v BriefProvider)`

SetProvider sets Provider field to given value.


### GetDescription

`func (o *BriefCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


