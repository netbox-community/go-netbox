# DeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;offline&#x60; - Offline * &#x60;active&#x60; - Active * &#x60;planned&#x60; - Planned * &#x60;staged&#x60; - Staged * &#x60;failed&#x60; - Failed * &#x60;inventory&#x60; - Inventory * &#x60;decommissioning&#x60; - Decommissioning | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceStatus

`func NewDeviceStatus() *DeviceStatus`

NewDeviceStatus instantiates a new DeviceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStatusWithDefaults

`func NewDeviceStatusWithDefaults() *DeviceStatus`

NewDeviceStatusWithDefaults instantiates a new DeviceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DeviceStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeviceStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeviceStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeviceStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *DeviceStatus) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeviceStatus) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeviceStatus) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeviceStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


