# ExtrasCustomLinksCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**LinkText** | **string** | Jinja2 template code for link text | 
**LinkUrl** | **string** | Jinja2 template code for link URL | 
**Weight** | Pointer to **int32** |  | [optional] 
**GroupName** | Pointer to **string** | Links with the same group will appear as a dropdown menu | [optional] 
**ButtonClass** | Pointer to [**CustomLinkButtonClass**](CustomLinkButtonClass.md) |  | [optional] 
**NewWindow** | Pointer to **bool** | Force link to open in a new window | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewExtrasCustomLinksCreateRequest

`func NewExtrasCustomLinksCreateRequest(objectTypes []string, name string, linkText string, linkUrl string, ) *ExtrasCustomLinksCreateRequest`

NewExtrasCustomLinksCreateRequest instantiates a new ExtrasCustomLinksCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasCustomLinksCreateRequestWithDefaults

`func NewExtrasCustomLinksCreateRequestWithDefaults() *ExtrasCustomLinksCreateRequest`

NewExtrasCustomLinksCreateRequestWithDefaults instantiates a new ExtrasCustomLinksCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *ExtrasCustomLinksCreateRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *ExtrasCustomLinksCreateRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *ExtrasCustomLinksCreateRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *ExtrasCustomLinksCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtrasCustomLinksCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtrasCustomLinksCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ExtrasCustomLinksCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtrasCustomLinksCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtrasCustomLinksCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExtrasCustomLinksCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLinkText

`func (o *ExtrasCustomLinksCreateRequest) GetLinkText() string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *ExtrasCustomLinksCreateRequest) GetLinkTextOk() (*string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *ExtrasCustomLinksCreateRequest) SetLinkText(v string)`

SetLinkText sets LinkText field to given value.


### GetLinkUrl

`func (o *ExtrasCustomLinksCreateRequest) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *ExtrasCustomLinksCreateRequest) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *ExtrasCustomLinksCreateRequest) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.


### GetWeight

`func (o *ExtrasCustomLinksCreateRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ExtrasCustomLinksCreateRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ExtrasCustomLinksCreateRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ExtrasCustomLinksCreateRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetGroupName

`func (o *ExtrasCustomLinksCreateRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ExtrasCustomLinksCreateRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ExtrasCustomLinksCreateRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ExtrasCustomLinksCreateRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetButtonClass

`func (o *ExtrasCustomLinksCreateRequest) GetButtonClass() CustomLinkButtonClass`

GetButtonClass returns the ButtonClass field if non-nil, zero value otherwise.

### GetButtonClassOk

`func (o *ExtrasCustomLinksCreateRequest) GetButtonClassOk() (*CustomLinkButtonClass, bool)`

GetButtonClassOk returns a tuple with the ButtonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClass

`func (o *ExtrasCustomLinksCreateRequest) SetButtonClass(v CustomLinkButtonClass)`

SetButtonClass sets ButtonClass field to given value.

### HasButtonClass

`func (o *ExtrasCustomLinksCreateRequest) HasButtonClass() bool`

HasButtonClass returns a boolean if a field has been set.

### GetNewWindow

`func (o *ExtrasCustomLinksCreateRequest) GetNewWindow() bool`

GetNewWindow returns the NewWindow field if non-nil, zero value otherwise.

### GetNewWindowOk

`func (o *ExtrasCustomLinksCreateRequest) GetNewWindowOk() (*bool, bool)`

GetNewWindowOk returns a tuple with the NewWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindow

`func (o *ExtrasCustomLinksCreateRequest) SetNewWindow(v bool)`

SetNewWindow sets NewWindow field to given value.

### HasNewWindow

`func (o *ExtrasCustomLinksCreateRequest) HasNewWindow() bool`

HasNewWindow returns a boolean if a field has been set.

### GetOwner

`func (o *ExtrasCustomLinksCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExtrasCustomLinksCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExtrasCustomLinksCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExtrasCustomLinksCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ExtrasCustomLinksCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ExtrasCustomLinksCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


