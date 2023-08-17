# JournalEntryKind

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;info&#x60; - Info * &#x60;success&#x60; - Success * &#x60;warning&#x60; - Warning * &#x60;danger&#x60; - Danger | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewJournalEntryKind

`func NewJournalEntryKind() *JournalEntryKind`

NewJournalEntryKind instantiates a new JournalEntryKind object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryKindWithDefaults

`func NewJournalEntryKindWithDefaults() *JournalEntryKind`

NewJournalEntryKindWithDefaults instantiates a new JournalEntryKind object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *JournalEntryKind) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JournalEntryKind) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JournalEntryKind) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *JournalEntryKind) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *JournalEntryKind) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *JournalEntryKind) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *JournalEntryKind) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *JournalEntryKind) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


