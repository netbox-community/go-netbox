# WritableConsoleServerPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **int32** |  | 
**Module** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to **string** | Physical port type  * &#x60;de-9&#x60; - DE-9 * &#x60;db-25&#x60; - DB-25 * &#x60;rj-11&#x60; - RJ-11 * &#x60;rj-12&#x60; - RJ-12 * &#x60;rj-45&#x60; - RJ-45 * &#x60;mini-din-8&#x60; - Mini-DIN 8 * &#x60;usb-a&#x60; - USB Type A * &#x60;usb-b&#x60; - USB Type B * &#x60;usb-c&#x60; - USB Type C * &#x60;usb-mini-a&#x60; - USB Mini A * &#x60;usb-mini-b&#x60; - USB Mini B * &#x60;usb-micro-a&#x60; - USB Micro A * &#x60;usb-micro-b&#x60; - USB Micro B * &#x60;usb-micro-ab&#x60; - USB Micro AB * &#x60;other&#x60; - Other | [optional] 
**Speed** | Pointer to **NullableInt32** | Port speed in bits per second  * &#x60;1200&#x60; - 1200 bps * &#x60;2400&#x60; - 2400 bps * &#x60;4800&#x60; - 4800 bps * &#x60;9600&#x60; - 9600 bps * &#x60;19200&#x60; - 19.2 kbps * &#x60;38400&#x60; - 38.4 kbps * &#x60;57600&#x60; - 57.6 kbps * &#x60;115200&#x60; - 115.2 kbps | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableConsoleServerPortRequest

`func NewWritableConsoleServerPortRequest(device int32, name string, ) *WritableConsoleServerPortRequest`

NewWritableConsoleServerPortRequest instantiates a new WritableConsoleServerPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableConsoleServerPortRequestWithDefaults

`func NewWritableConsoleServerPortRequestWithDefaults() *WritableConsoleServerPortRequest`

NewWritableConsoleServerPortRequestWithDefaults instantiates a new WritableConsoleServerPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *WritableConsoleServerPortRequest) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableConsoleServerPortRequest) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableConsoleServerPortRequest) SetDevice(v int32)`

SetDevice sets Device field to given value.


### GetModule

`func (o *WritableConsoleServerPortRequest) GetModule() int32`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WritableConsoleServerPortRequest) GetModuleOk() (*int32, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WritableConsoleServerPortRequest) SetModule(v int32)`

SetModule sets Module field to given value.

### HasModule

`func (o *WritableConsoleServerPortRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *WritableConsoleServerPortRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *WritableConsoleServerPortRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *WritableConsoleServerPortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableConsoleServerPortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableConsoleServerPortRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableConsoleServerPortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableConsoleServerPortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableConsoleServerPortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableConsoleServerPortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritableConsoleServerPortRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableConsoleServerPortRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableConsoleServerPortRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WritableConsoleServerPortRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSpeed

`func (o *WritableConsoleServerPortRequest) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *WritableConsoleServerPortRequest) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *WritableConsoleServerPortRequest) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *WritableConsoleServerPortRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *WritableConsoleServerPortRequest) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *WritableConsoleServerPortRequest) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDescription

`func (o *WritableConsoleServerPortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableConsoleServerPortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableConsoleServerPortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableConsoleServerPortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *WritableConsoleServerPortRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *WritableConsoleServerPortRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *WritableConsoleServerPortRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *WritableConsoleServerPortRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetTags

`func (o *WritableConsoleServerPortRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableConsoleServerPortRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableConsoleServerPortRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableConsoleServerPortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableConsoleServerPortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableConsoleServerPortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableConsoleServerPortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableConsoleServerPortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


