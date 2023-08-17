# FrontPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableNestedDeviceTypeRequest**](NestedDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableNestedModuleTypeRequest**](NestedModuleTypeRequest.md) |  | [optional] 
**Name** | **string** |          {module} is accepted as a substitution for the module bay position when attached to a module type.          | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | **string** | * &#x60;8p8c&#x60; - 8P8C * &#x60;8p6c&#x60; - 8P6C * &#x60;8p4c&#x60; - 8P4C * &#x60;8p2c&#x60; - 8P2C * &#x60;6p6c&#x60; - 6P6C * &#x60;6p4c&#x60; - 6P4C * &#x60;6p2c&#x60; - 6P2C * &#x60;4p4c&#x60; - 4P4C * &#x60;4p2c&#x60; - 4P2C * &#x60;gg45&#x60; - GG45 * &#x60;tera-4p&#x60; - TERA 4P * &#x60;tera-2p&#x60; - TERA 2P * &#x60;tera-1p&#x60; - TERA 1P * &#x60;110-punch&#x60; - 110 Punch * &#x60;bnc&#x60; - BNC * &#x60;f&#x60; - F Connector * &#x60;n&#x60; - N Connector * &#x60;mrj21&#x60; - MRJ21 * &#x60;fc&#x60; - FC * &#x60;lc&#x60; - LC * &#x60;lc-pc&#x60; - LC/PC * &#x60;lc-upc&#x60; - LC/UPC * &#x60;lc-apc&#x60; - LC/APC * &#x60;lsh&#x60; - LSH * &#x60;lsh-pc&#x60; - LSH/PC * &#x60;lsh-upc&#x60; - LSH/UPC * &#x60;lsh-apc&#x60; - LSH/APC * &#x60;lx5&#x60; - LX.5 * &#x60;lx5-pc&#x60; - LX.5/PC * &#x60;lx5-upc&#x60; - LX.5/UPC * &#x60;lx5-apc&#x60; - LX.5/APC * &#x60;mpo&#x60; - MPO * &#x60;mtrj&#x60; - MTRJ * &#x60;sc&#x60; - SC * &#x60;sc-pc&#x60; - SC/PC * &#x60;sc-upc&#x60; - SC/UPC * &#x60;sc-apc&#x60; - SC/APC * &#x60;st&#x60; - ST * &#x60;cs&#x60; - CS * &#x60;sn&#x60; - SN * &#x60;sma-905&#x60; - SMA 905 * &#x60;sma-906&#x60; - SMA 906 * &#x60;urm-p2&#x60; - URM-P2 * &#x60;urm-p4&#x60; - URM-P4 * &#x60;urm-p8&#x60; - URM-P8 * &#x60;splice&#x60; - Splice * &#x60;other&#x60; - Other | 
**Color** | Pointer to **string** |  | [optional] 
**RearPort** | [**NestedRearPortTemplateRequest**](NestedRearPortTemplateRequest.md) |  | 
**RearPortPosition** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewFrontPortTemplateRequest

`func NewFrontPortTemplateRequest(name string, type_ string, rearPort NestedRearPortTemplateRequest, ) *FrontPortTemplateRequest`

NewFrontPortTemplateRequest instantiates a new FrontPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrontPortTemplateRequestWithDefaults

`func NewFrontPortTemplateRequestWithDefaults() *FrontPortTemplateRequest`

NewFrontPortTemplateRequestWithDefaults instantiates a new FrontPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *FrontPortTemplateRequest) GetDeviceType() NestedDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *FrontPortTemplateRequest) GetDeviceTypeOk() (*NestedDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *FrontPortTemplateRequest) SetDeviceType(v NestedDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *FrontPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *FrontPortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *FrontPortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *FrontPortTemplateRequest) GetModuleType() NestedModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *FrontPortTemplateRequest) GetModuleTypeOk() (*NestedModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *FrontPortTemplateRequest) SetModuleType(v NestedModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *FrontPortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *FrontPortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *FrontPortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *FrontPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrontPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrontPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *FrontPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FrontPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FrontPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FrontPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *FrontPortTemplateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FrontPortTemplateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FrontPortTemplateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *FrontPortTemplateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FrontPortTemplateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FrontPortTemplateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *FrontPortTemplateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetRearPort

`func (o *FrontPortTemplateRequest) GetRearPort() NestedRearPortTemplateRequest`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *FrontPortTemplateRequest) GetRearPortOk() (*NestedRearPortTemplateRequest, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *FrontPortTemplateRequest) SetRearPort(v NestedRearPortTemplateRequest)`

SetRearPort sets RearPort field to given value.


### GetRearPortPosition

`func (o *FrontPortTemplateRequest) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *FrontPortTemplateRequest) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *FrontPortTemplateRequest) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *FrontPortTemplateRequest) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.

### GetDescription

`func (o *FrontPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FrontPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FrontPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FrontPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


