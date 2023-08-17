# ConsolePortSpeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** | * &#x60;1200&#x60; - 1200 bps * &#x60;2400&#x60; - 2400 bps * &#x60;4800&#x60; - 4800 bps * &#x60;9600&#x60; - 9600 bps * &#x60;19200&#x60; - 19.2 kbps * &#x60;38400&#x60; - 38.4 kbps * &#x60;57600&#x60; - 57.6 kbps * &#x60;115200&#x60; - 115.2 kbps | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewConsolePortSpeed

`func NewConsolePortSpeed() *ConsolePortSpeed`

NewConsolePortSpeed instantiates a new ConsolePortSpeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolePortSpeedWithDefaults

`func NewConsolePortSpeedWithDefaults() *ConsolePortSpeed`

NewConsolePortSpeedWithDefaults instantiates a new ConsolePortSpeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ConsolePortSpeed) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConsolePortSpeed) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConsolePortSpeed) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConsolePortSpeed) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ConsolePortSpeed) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsolePortSpeed) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsolePortSpeed) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsolePortSpeed) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


