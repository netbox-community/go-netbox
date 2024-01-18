# PatchedImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Image** | Pointer to ***os.File** |  | [optional] 
**ImageHeight** | Pointer to **int32** |  | [optional] 
**ImageWidth** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedImageAttachmentRequest

`func NewPatchedImageAttachmentRequest() *PatchedImageAttachmentRequest`

NewPatchedImageAttachmentRequest instantiates a new PatchedImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedImageAttachmentRequestWithDefaults

`func NewPatchedImageAttachmentRequestWithDefaults() *PatchedImageAttachmentRequest`

NewPatchedImageAttachmentRequestWithDefaults instantiates a new PatchedImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *PatchedImageAttachmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedImageAttachmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedImageAttachmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedImageAttachmentRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedImageAttachmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedImageAttachmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedImageAttachmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedImageAttachmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetName

`func (o *PatchedImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *PatchedImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PatchedImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PatchedImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.

### HasImage

`func (o *PatchedImageAttachmentRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageHeight

`func (o *PatchedImageAttachmentRequest) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *PatchedImageAttachmentRequest) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *PatchedImageAttachmentRequest) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.

### HasImageHeight

`func (o *PatchedImageAttachmentRequest) HasImageHeight() bool`

HasImageHeight returns a boolean if a field has been set.

### GetImageWidth

`func (o *PatchedImageAttachmentRequest) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *PatchedImageAttachmentRequest) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *PatchedImageAttachmentRequest) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.

### HasImageWidth

`func (o *PatchedImageAttachmentRequest) HasImageWidth() bool`

HasImageWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


