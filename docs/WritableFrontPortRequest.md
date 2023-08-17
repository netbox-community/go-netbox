# WritableFrontPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **int32** |  | 
**Module** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | **string** | * &#x60;8p8c&#x60; - 8P8C * &#x60;8p6c&#x60; - 8P6C * &#x60;8p4c&#x60; - 8P4C * &#x60;8p2c&#x60; - 8P2C * &#x60;6p6c&#x60; - 6P6C * &#x60;6p4c&#x60; - 6P4C * &#x60;6p2c&#x60; - 6P2C * &#x60;4p4c&#x60; - 4P4C * &#x60;4p2c&#x60; - 4P2C * &#x60;gg45&#x60; - GG45 * &#x60;tera-4p&#x60; - TERA 4P * &#x60;tera-2p&#x60; - TERA 2P * &#x60;tera-1p&#x60; - TERA 1P * &#x60;110-punch&#x60; - 110 Punch * &#x60;bnc&#x60; - BNC * &#x60;f&#x60; - F Connector * &#x60;n&#x60; - N Connector * &#x60;mrj21&#x60; - MRJ21 * &#x60;fc&#x60; - FC * &#x60;lc&#x60; - LC * &#x60;lc-pc&#x60; - LC/PC * &#x60;lc-upc&#x60; - LC/UPC * &#x60;lc-apc&#x60; - LC/APC * &#x60;lsh&#x60; - LSH * &#x60;lsh-pc&#x60; - LSH/PC * &#x60;lsh-upc&#x60; - LSH/UPC * &#x60;lsh-apc&#x60; - LSH/APC * &#x60;lx5&#x60; - LX.5 * &#x60;lx5-pc&#x60; - LX.5/PC * &#x60;lx5-upc&#x60; - LX.5/UPC * &#x60;lx5-apc&#x60; - LX.5/APC * &#x60;mpo&#x60; - MPO * &#x60;mtrj&#x60; - MTRJ * &#x60;sc&#x60; - SC * &#x60;sc-pc&#x60; - SC/PC * &#x60;sc-upc&#x60; - SC/UPC * &#x60;sc-apc&#x60; - SC/APC * &#x60;st&#x60; - ST * &#x60;cs&#x60; - CS * &#x60;sn&#x60; - SN * &#x60;sma-905&#x60; - SMA 905 * &#x60;sma-906&#x60; - SMA 906 * &#x60;urm-p2&#x60; - URM-P2 * &#x60;urm-p4&#x60; - URM-P4 * &#x60;urm-p8&#x60; - URM-P8 * &#x60;splice&#x60; - Splice * &#x60;other&#x60; - Other | 
**Color** | Pointer to **string** |  | [optional] 
**RearPort** | **int32** |  | 
**RearPortPosition** | Pointer to **int32** | Mapped position on corresponding rear port | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableFrontPortRequest

`func NewWritableFrontPortRequest(device int32, name string, type_ string, rearPort int32, ) *WritableFrontPortRequest`

NewWritableFrontPortRequest instantiates a new WritableFrontPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableFrontPortRequestWithDefaults

`func NewWritableFrontPortRequestWithDefaults() *WritableFrontPortRequest`

NewWritableFrontPortRequestWithDefaults instantiates a new WritableFrontPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *WritableFrontPortRequest) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableFrontPortRequest) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableFrontPortRequest) SetDevice(v int32)`

SetDevice sets Device field to given value.


### GetModule

`func (o *WritableFrontPortRequest) GetModule() int32`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WritableFrontPortRequest) GetModuleOk() (*int32, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WritableFrontPortRequest) SetModule(v int32)`

SetModule sets Module field to given value.

### HasModule

`func (o *WritableFrontPortRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *WritableFrontPortRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *WritableFrontPortRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *WritableFrontPortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableFrontPortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableFrontPortRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableFrontPortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableFrontPortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableFrontPortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableFrontPortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritableFrontPortRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableFrontPortRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableFrontPortRequest) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *WritableFrontPortRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WritableFrontPortRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WritableFrontPortRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WritableFrontPortRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetRearPort

`func (o *WritableFrontPortRequest) GetRearPort() int32`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *WritableFrontPortRequest) GetRearPortOk() (*int32, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *WritableFrontPortRequest) SetRearPort(v int32)`

SetRearPort sets RearPort field to given value.


### GetRearPortPosition

`func (o *WritableFrontPortRequest) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *WritableFrontPortRequest) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *WritableFrontPortRequest) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *WritableFrontPortRequest) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.

### GetDescription

`func (o *WritableFrontPortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableFrontPortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableFrontPortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableFrontPortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *WritableFrontPortRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *WritableFrontPortRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *WritableFrontPortRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *WritableFrontPortRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetTags

`func (o *WritableFrontPortRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableFrontPortRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableFrontPortRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableFrontPortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableFrontPortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableFrontPortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableFrontPortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableFrontPortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


