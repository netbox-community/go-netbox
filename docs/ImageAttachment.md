# ImageAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Parent** | **interface{}** |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Image** | **string** |  | 
**ImageHeight** | **int32** |  | [readonly] 
**ImageWidth** | **int32** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewImageAttachment

`func NewImageAttachment(id int32, url string, display string, objectType string, objectId int64, parent interface{}, image string, imageHeight int32, imageWidth int32, created NullableTime, lastUpdated NullableTime, ) *ImageAttachment`

NewImageAttachment instantiates a new ImageAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAttachmentWithDefaults

`func NewImageAttachmentWithDefaults() *ImageAttachment`

NewImageAttachmentWithDefaults instantiates a new ImageAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageAttachment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageAttachment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageAttachment) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ImageAttachment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageAttachment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageAttachment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *ImageAttachment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ImageAttachment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ImageAttachment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetObjectType

`func (o *ImageAttachment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ImageAttachment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ImageAttachment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ImageAttachment) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ImageAttachment) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ImageAttachment) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetParent

`func (o *ImageAttachment) GetParent() interface{}`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ImageAttachment) GetParentOk() (*interface{}, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ImageAttachment) SetParent(v interface{})`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *ImageAttachment) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ImageAttachment) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *ImageAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *ImageAttachment) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageAttachment) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageAttachment) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageHeight

`func (o *ImageAttachment) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *ImageAttachment) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *ImageAttachment) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.


### GetImageWidth

`func (o *ImageAttachment) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *ImageAttachment) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *ImageAttachment) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.


### GetCreated

`func (o *ImageAttachment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ImageAttachment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ImageAttachment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ImageAttachment) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ImageAttachment) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ImageAttachment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ImageAttachment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ImageAttachment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ImageAttachment) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ImageAttachment) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


