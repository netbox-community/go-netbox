# ExtrasImageAttachmentsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Image** | ***os.File** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewExtrasImageAttachmentsCreateRequest

`func NewExtrasImageAttachmentsCreateRequest(objectType string, objectId int64, image *os.File, ) *ExtrasImageAttachmentsCreateRequest`

NewExtrasImageAttachmentsCreateRequest instantiates a new ExtrasImageAttachmentsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasImageAttachmentsCreateRequestWithDefaults

`func NewExtrasImageAttachmentsCreateRequestWithDefaults() *ExtrasImageAttachmentsCreateRequest`

NewExtrasImageAttachmentsCreateRequestWithDefaults instantiates a new ExtrasImageAttachmentsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *ExtrasImageAttachmentsCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExtrasImageAttachmentsCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExtrasImageAttachmentsCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ExtrasImageAttachmentsCreateRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ExtrasImageAttachmentsCreateRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ExtrasImageAttachmentsCreateRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetName

`func (o *ExtrasImageAttachmentsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtrasImageAttachmentsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtrasImageAttachmentsCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtrasImageAttachmentsCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *ExtrasImageAttachmentsCreateRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ExtrasImageAttachmentsCreateRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ExtrasImageAttachmentsCreateRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.


### GetDescription

`func (o *ExtrasImageAttachmentsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtrasImageAttachmentsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtrasImageAttachmentsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtrasImageAttachmentsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


