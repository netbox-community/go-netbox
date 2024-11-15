# BriefCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** | Unique circuit ID | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefCircuitRequest

`func NewBriefCircuitRequest(cid string, ) *BriefCircuitRequest`

NewBriefCircuitRequest instantiates a new BriefCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefCircuitRequestWithDefaults

`func NewBriefCircuitRequestWithDefaults() *BriefCircuitRequest`

NewBriefCircuitRequestWithDefaults instantiates a new BriefCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *BriefCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *BriefCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *BriefCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.


### GetDescription

`func (o *BriefCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


