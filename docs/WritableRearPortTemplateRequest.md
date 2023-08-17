# WritableRearPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **NullableInt32** |  | [optional] 
**ModuleType** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |          {module} is accepted as a substitution for the module bay position when attached to a module type.          | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | **string** | * &#x60;8p8c&#x60; - 8P8C * &#x60;8p6c&#x60; - 8P6C * &#x60;8p4c&#x60; - 8P4C * &#x60;8p2c&#x60; - 8P2C * &#x60;6p6c&#x60; - 6P6C * &#x60;6p4c&#x60; - 6P4C * &#x60;6p2c&#x60; - 6P2C * &#x60;4p4c&#x60; - 4P4C * &#x60;4p2c&#x60; - 4P2C * &#x60;gg45&#x60; - GG45 * &#x60;tera-4p&#x60; - TERA 4P * &#x60;tera-2p&#x60; - TERA 2P * &#x60;tera-1p&#x60; - TERA 1P * &#x60;110-punch&#x60; - 110 Punch * &#x60;bnc&#x60; - BNC * &#x60;f&#x60; - F Connector * &#x60;n&#x60; - N Connector * &#x60;mrj21&#x60; - MRJ21 * &#x60;fc&#x60; - FC * &#x60;lc&#x60; - LC * &#x60;lc-pc&#x60; - LC/PC * &#x60;lc-upc&#x60; - LC/UPC * &#x60;lc-apc&#x60; - LC/APC * &#x60;lsh&#x60; - LSH * &#x60;lsh-pc&#x60; - LSH/PC * &#x60;lsh-upc&#x60; - LSH/UPC * &#x60;lsh-apc&#x60; - LSH/APC * &#x60;lx5&#x60; - LX.5 * &#x60;lx5-pc&#x60; - LX.5/PC * &#x60;lx5-upc&#x60; - LX.5/UPC * &#x60;lx5-apc&#x60; - LX.5/APC * &#x60;mpo&#x60; - MPO * &#x60;mtrj&#x60; - MTRJ * &#x60;sc&#x60; - SC * &#x60;sc-pc&#x60; - SC/PC * &#x60;sc-upc&#x60; - SC/UPC * &#x60;sc-apc&#x60; - SC/APC * &#x60;st&#x60; - ST * &#x60;cs&#x60; - CS * &#x60;sn&#x60; - SN * &#x60;sma-905&#x60; - SMA 905 * &#x60;sma-906&#x60; - SMA 906 * &#x60;urm-p2&#x60; - URM-P2 * &#x60;urm-p4&#x60; - URM-P4 * &#x60;urm-p8&#x60; - URM-P8 * &#x60;splice&#x60; - Splice * &#x60;other&#x60; - Other | 
**Color** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewWritableRearPortTemplateRequest

`func NewWritableRearPortTemplateRequest(name string, type_ string, ) *WritableRearPortTemplateRequest`

NewWritableRearPortTemplateRequest instantiates a new WritableRearPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableRearPortTemplateRequestWithDefaults

`func NewWritableRearPortTemplateRequestWithDefaults() *WritableRearPortTemplateRequest`

NewWritableRearPortTemplateRequestWithDefaults instantiates a new WritableRearPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *WritableRearPortTemplateRequest) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritableRearPortTemplateRequest) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritableRearPortTemplateRequest) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WritableRearPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *WritableRearPortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *WritableRearPortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *WritableRearPortTemplateRequest) GetModuleType() int32`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *WritableRearPortTemplateRequest) GetModuleTypeOk() (*int32, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *WritableRearPortTemplateRequest) SetModuleType(v int32)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *WritableRearPortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *WritableRearPortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *WritableRearPortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *WritableRearPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableRearPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableRearPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableRearPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableRearPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableRearPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableRearPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritableRearPortTemplateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableRearPortTemplateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableRearPortTemplateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *WritableRearPortTemplateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WritableRearPortTemplateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WritableRearPortTemplateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WritableRearPortTemplateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPositions

`func (o *WritableRearPortTemplateRequest) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *WritableRearPortTemplateRequest) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *WritableRearPortTemplateRequest) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *WritableRearPortTemplateRequest) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetDescription

`func (o *WritableRearPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableRearPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableRearPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableRearPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


