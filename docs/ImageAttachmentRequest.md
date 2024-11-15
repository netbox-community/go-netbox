# ImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Image** | ***os.File** |  | 

## Methods

### NewImageAttachmentRequest

`func NewImageAttachmentRequest(objectType string, objectId int64, image *os.File, ) *ImageAttachmentRequest`

NewImageAttachmentRequest instantiates a new ImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAttachmentRequestWithDefaults

`func NewImageAttachmentRequestWithDefaults() *ImageAttachmentRequest`

NewImageAttachmentRequestWithDefaults instantiates a new ImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *ImageAttachmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ImageAttachmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ImageAttachmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ImageAttachmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ImageAttachmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ImageAttachmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetName

`func (o *ImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *ImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


