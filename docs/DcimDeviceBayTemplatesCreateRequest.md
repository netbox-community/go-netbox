# DcimDeviceBayTemplatesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | [**DeviceBayTemplateRequestDeviceType**](DeviceBayTemplateRequestDeviceType.md) |  | 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewDcimDeviceBayTemplatesCreateRequest

`func NewDcimDeviceBayTemplatesCreateRequest(deviceType DeviceBayTemplateRequestDeviceType, name string, ) *DcimDeviceBayTemplatesCreateRequest`

NewDcimDeviceBayTemplatesCreateRequest instantiates a new DcimDeviceBayTemplatesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimDeviceBayTemplatesCreateRequestWithDefaults

`func NewDcimDeviceBayTemplatesCreateRequestWithDefaults() *DcimDeviceBayTemplatesCreateRequest`

NewDcimDeviceBayTemplatesCreateRequestWithDefaults instantiates a new DcimDeviceBayTemplatesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *DcimDeviceBayTemplatesCreateRequest) GetDeviceType() DeviceBayTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DcimDeviceBayTemplatesCreateRequest) GetDeviceTypeOk() (*DeviceBayTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DcimDeviceBayTemplatesCreateRequest) SetDeviceType(v DeviceBayTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetName

`func (o *DcimDeviceBayTemplatesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimDeviceBayTemplatesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimDeviceBayTemplatesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *DcimDeviceBayTemplatesCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DcimDeviceBayTemplatesCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DcimDeviceBayTemplatesCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DcimDeviceBayTemplatesCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *DcimDeviceBayTemplatesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimDeviceBayTemplatesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimDeviceBayTemplatesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimDeviceBayTemplatesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


