# PatchedWritablePowerOutletTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **NullableInt32** |  | [optional] 
**ModuleType** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to **string** | * &#x60;iec-60320-c5&#x60; - C5 * &#x60;iec-60320-c7&#x60; - C7 * &#x60;iec-60320-c13&#x60; - C13 * &#x60;iec-60320-c15&#x60; - C15 * &#x60;iec-60320-c19&#x60; - C19 * &#x60;iec-60320-c21&#x60; - C21 * &#x60;iec-60309-p-n-e-4h&#x60; - P+N+E 4H * &#x60;iec-60309-p-n-e-6h&#x60; - P+N+E 6H * &#x60;iec-60309-p-n-e-9h&#x60; - P+N+E 9H * &#x60;iec-60309-2p-e-4h&#x60; - 2P+E 4H * &#x60;iec-60309-2p-e-6h&#x60; - 2P+E 6H * &#x60;iec-60309-2p-e-9h&#x60; - 2P+E 9H * &#x60;iec-60309-3p-e-4h&#x60; - 3P+E 4H * &#x60;iec-60309-3p-e-6h&#x60; - 3P+E 6H * &#x60;iec-60309-3p-e-9h&#x60; - 3P+E 9H * &#x60;iec-60309-3p-n-e-4h&#x60; - 3P+N+E 4H * &#x60;iec-60309-3p-n-e-6h&#x60; - 3P+N+E 6H * &#x60;iec-60309-3p-n-e-9h&#x60; - 3P+N+E 9H * &#x60;iec-60906-1&#x60; - IEC 60906-1 * &#x60;nbr-14136-10a&#x60; - 2P+T 10A (NBR 14136) * &#x60;nbr-14136-20a&#x60; - 2P+T 20A (NBR 14136) * &#x60;nema-1-15r&#x60; - NEMA 1-15R * &#x60;nema-5-15r&#x60; - NEMA 5-15R * &#x60;nema-5-20r&#x60; - NEMA 5-20R * &#x60;nema-5-30r&#x60; - NEMA 5-30R * &#x60;nema-5-50r&#x60; - NEMA 5-50R * &#x60;nema-6-15r&#x60; - NEMA 6-15R * &#x60;nema-6-20r&#x60; - NEMA 6-20R * &#x60;nema-6-30r&#x60; - NEMA 6-30R * &#x60;nema-6-50r&#x60; - NEMA 6-50R * &#x60;nema-10-30r&#x60; - NEMA 10-30R * &#x60;nema-10-50r&#x60; - NEMA 10-50R * &#x60;nema-14-20r&#x60; - NEMA 14-20R * &#x60;nema-14-30r&#x60; - NEMA 14-30R * &#x60;nema-14-50r&#x60; - NEMA 14-50R * &#x60;nema-14-60r&#x60; - NEMA 14-60R * &#x60;nema-15-15r&#x60; - NEMA 15-15R * &#x60;nema-15-20r&#x60; - NEMA 15-20R * &#x60;nema-15-30r&#x60; - NEMA 15-30R * &#x60;nema-15-50r&#x60; - NEMA 15-50R * &#x60;nema-15-60r&#x60; - NEMA 15-60R * &#x60;nema-l1-15r&#x60; - NEMA L1-15R * &#x60;nema-l5-15r&#x60; - NEMA L5-15R * &#x60;nema-l5-20r&#x60; - NEMA L5-20R * &#x60;nema-l5-30r&#x60; - NEMA L5-30R * &#x60;nema-l5-50r&#x60; - NEMA L5-50R * &#x60;nema-l6-15r&#x60; - NEMA L6-15R * &#x60;nema-l6-20r&#x60; - NEMA L6-20R * &#x60;nema-l6-30r&#x60; - NEMA L6-30R * &#x60;nema-l6-50r&#x60; - NEMA L6-50R * &#x60;nema-l10-30r&#x60; - NEMA L10-30R * &#x60;nema-l14-20r&#x60; - NEMA L14-20R * &#x60;nema-l14-30r&#x60; - NEMA L14-30R * &#x60;nema-l14-50r&#x60; - NEMA L14-50R * &#x60;nema-l14-60r&#x60; - NEMA L14-60R * &#x60;nema-l15-20r&#x60; - NEMA L15-20R * &#x60;nema-l15-30r&#x60; - NEMA L15-30R * &#x60;nema-l15-50r&#x60; - NEMA L15-50R * &#x60;nema-l15-60r&#x60; - NEMA L15-60R * &#x60;nema-l21-20r&#x60; - NEMA L21-20R * &#x60;nema-l21-30r&#x60; - NEMA L21-30R * &#x60;nema-l22-30r&#x60; - NEMA L22-30R * &#x60;CS6360C&#x60; - CS6360C * &#x60;CS6364C&#x60; - CS6364C * &#x60;CS8164C&#x60; - CS8164C * &#x60;CS8264C&#x60; - CS8264C * &#x60;CS8364C&#x60; - CS8364C * &#x60;CS8464C&#x60; - CS8464C * &#x60;ita-e&#x60; - ITA Type E (CEE 7/5) * &#x60;ita-f&#x60; - ITA Type F (CEE 7/3) * &#x60;ita-g&#x60; - ITA Type G (BS 1363) * &#x60;ita-h&#x60; - ITA Type H * &#x60;ita-i&#x60; - ITA Type I * &#x60;ita-j&#x60; - ITA Type J * &#x60;ita-k&#x60; - ITA Type K * &#x60;ita-l&#x60; - ITA Type L (CEI 23-50) * &#x60;ita-m&#x60; - ITA Type M (BS 546) * &#x60;ita-n&#x60; - ITA Type N * &#x60;ita-o&#x60; - ITA Type O * &#x60;ita-multistandard&#x60; - ITA Multistandard * &#x60;usb-a&#x60; - USB Type A * &#x60;usb-micro-b&#x60; - USB Micro B * &#x60;usb-c&#x60; - USB Type C * &#x60;dc-terminal&#x60; - DC Terminal * &#x60;hdot-cx&#x60; - HDOT Cx * &#x60;saf-d-grid&#x60; - Saf-D-Grid * &#x60;neutrik-powercon-20a&#x60; - Neutrik powerCON (20A) * &#x60;neutrik-powercon-32a&#x60; - Neutrik powerCON (32A) * &#x60;neutrik-powercon-true1&#x60; - Neutrik powerCON TRUE1 * &#x60;neutrik-powercon-true1-top&#x60; - Neutrik powerCON TRUE1 TOP * &#x60;ubiquiti-smartpower&#x60; - Ubiquiti SmartPower * &#x60;hardwired&#x60; - Hardwired * &#x60;other&#x60; - Other | [optional] 
**PowerPort** | Pointer to **NullableInt32** |  | [optional] 
**FeedLeg** | Pointer to **string** | Phase (for three-phase feeds)  * &#x60;A&#x60; - A * &#x60;B&#x60; - B * &#x60;C&#x60; - C | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritablePowerOutletTemplateRequest

`func NewPatchedWritablePowerOutletTemplateRequest() *PatchedWritablePowerOutletTemplateRequest`

NewPatchedWritablePowerOutletTemplateRequest instantiates a new PatchedWritablePowerOutletTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePowerOutletTemplateRequestWithDefaults

`func NewPatchedWritablePowerOutletTemplateRequestWithDefaults() *PatchedWritablePowerOutletTemplateRequest`

NewPatchedWritablePowerOutletTemplateRequestWithDefaults instantiates a new PatchedWritablePowerOutletTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritablePowerOutletTemplateRequest) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritablePowerOutletTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PatchedWritablePowerOutletTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PatchedWritablePowerOutletTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PatchedWritablePowerOutletTemplateRequest) GetModuleType() int32`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetModuleTypeOk() (*int32, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedWritablePowerOutletTemplateRequest) SetModuleType(v int32)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedWritablePowerOutletTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PatchedWritablePowerOutletTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PatchedWritablePowerOutletTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PatchedWritablePowerOutletTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritablePowerOutletTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritablePowerOutletTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritablePowerOutletTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritablePowerOutletTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritablePowerOutletTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritablePowerOutletTemplateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritablePowerOutletTemplateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritablePowerOutletTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPowerPort

`func (o *PatchedWritablePowerOutletTemplateRequest) GetPowerPort() int32`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetPowerPortOk() (*int32, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *PatchedWritablePowerOutletTemplateRequest) SetPowerPort(v int32)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *PatchedWritablePowerOutletTemplateRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *PatchedWritablePowerOutletTemplateRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *PatchedWritablePowerOutletTemplateRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetFeedLeg

`func (o *PatchedWritablePowerOutletTemplateRequest) GetFeedLeg() string`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetFeedLegOk() (*string, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *PatchedWritablePowerOutletTemplateRequest) SetFeedLeg(v string)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *PatchedWritablePowerOutletTemplateRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritablePowerOutletTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritablePowerOutletTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


