# DeviceTypeAirflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;front-to-rear&#x60; - Front to rear * &#x60;rear-to-front&#x60; - Rear to front * &#x60;left-to-right&#x60; - Left to right * &#x60;right-to-left&#x60; - Right to left * &#x60;side-to-rear&#x60; - Side to rear * &#x60;passive&#x60; - Passive * &#x60;mixed&#x60; - Mixed | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceTypeAirflow

`func NewDeviceTypeAirflow() *DeviceTypeAirflow`

NewDeviceTypeAirflow instantiates a new DeviceTypeAirflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTypeAirflowWithDefaults

`func NewDeviceTypeAirflowWithDefaults() *DeviceTypeAirflow`

NewDeviceTypeAirflowWithDefaults instantiates a new DeviceTypeAirflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DeviceTypeAirflow) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeviceTypeAirflow) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeviceTypeAirflow) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeviceTypeAirflow) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *DeviceTypeAirflow) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeviceTypeAirflow) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeviceTypeAirflow) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeviceTypeAirflow) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


