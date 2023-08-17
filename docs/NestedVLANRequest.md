# NestedVLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vid** | **int32** | Numeric VLAN ID (1-4094) | 
**Name** | **string** |  | 

## Methods

### NewNestedVLANRequest

`func NewNestedVLANRequest(vid int32, name string, ) *NestedVLANRequest`

NewNestedVLANRequest instantiates a new NestedVLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedVLANRequestWithDefaults

`func NewNestedVLANRequestWithDefaults() *NestedVLANRequest`

NewNestedVLANRequestWithDefaults instantiates a new NestedVLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVid

`func (o *NestedVLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *NestedVLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *NestedVLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *NestedVLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedVLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedVLANRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


