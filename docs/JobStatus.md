# JobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**JobStatusValue**](JobStatusValue.md) |  | [optional] 
**Label** | Pointer to [**JobStatusLabel**](JobStatusLabel.md) |  | [optional] 

## Methods

### NewJobStatus

`func NewJobStatus() *JobStatus`

NewJobStatus instantiates a new JobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStatusWithDefaults

`func NewJobStatusWithDefaults() *JobStatus`

NewJobStatusWithDefaults instantiates a new JobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *JobStatus) GetValue() JobStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JobStatus) GetValueOk() (*JobStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JobStatus) SetValue(v JobStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *JobStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *JobStatus) GetLabel() JobStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *JobStatus) GetLabelOk() (*JobStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *JobStatus) SetLabel(v JobStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *JobStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


