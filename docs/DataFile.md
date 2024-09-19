# DataFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Source** | [**BriefDataSource**](BriefDataSource.md) |  | [readonly] 
**Path** | **string** | File path relative to the data source&#39;s root | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**Size** | **int32** |  | [readonly] 
**Hash** | **string** | SHA256 hash of the file data | [readonly] 

## Methods

### NewDataFile

`func NewDataFile(id int32, url string, displayUrl string, display string, source BriefDataSource, path string, lastUpdated time.Time, size int32, hash string, ) *DataFile`

NewDataFile instantiates a new DataFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFileWithDefaults

`func NewDataFileWithDefaults() *DataFile`

NewDataFileWithDefaults instantiates a new DataFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataFile) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *DataFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DataFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DataFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *DataFile) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *DataFile) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *DataFile) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *DataFile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DataFile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DataFile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetSource

`func (o *DataFile) GetSource() BriefDataSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataFile) GetSourceOk() (*BriefDataSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataFile) SetSource(v BriefDataSource)`

SetSource sets Source field to given value.


### GetPath

`func (o *DataFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DataFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DataFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetLastUpdated

`func (o *DataFile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DataFile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DataFile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetSize

`func (o *DataFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DataFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DataFile) SetSize(v int32)`

SetSize sets Size field to given value.


### GetHash

`func (o *DataFile) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DataFile) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DataFile) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


