# ConsolePortType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;de-9&#x60; - DE-9 * &#x60;db-25&#x60; - DB-25 * &#x60;rj-11&#x60; - RJ-11 * &#x60;rj-12&#x60; - RJ-12 * &#x60;rj-45&#x60; - RJ-45 * &#x60;mini-din-8&#x60; - Mini-DIN 8 * &#x60;usb-a&#x60; - USB Type A * &#x60;usb-b&#x60; - USB Type B * &#x60;usb-c&#x60; - USB Type C * &#x60;usb-mini-a&#x60; - USB Mini A * &#x60;usb-mini-b&#x60; - USB Mini B * &#x60;usb-micro-a&#x60; - USB Micro A * &#x60;usb-micro-b&#x60; - USB Micro B * &#x60;usb-micro-ab&#x60; - USB Micro AB * &#x60;other&#x60; - Other | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewConsolePortType

`func NewConsolePortType() *ConsolePortType`

NewConsolePortType instantiates a new ConsolePortType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolePortTypeWithDefaults

`func NewConsolePortTypeWithDefaults() *ConsolePortType`

NewConsolePortTypeWithDefaults instantiates a new ConsolePortType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ConsolePortType) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConsolePortType) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConsolePortType) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConsolePortType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ConsolePortType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsolePortType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsolePortType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsolePortType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


