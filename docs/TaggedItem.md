# TaggedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**ObjectId** | **int32** |  | 
**Object** | **interface{}** |  | [readonly] 
**Tag** | [**BriefTag**](BriefTag.md) |  | [readonly] 

## Methods

### NewTaggedItem

`func NewTaggedItem(id int32, url string, display string, objectType string, objectId int32, object interface{}, tag BriefTag, ) *TaggedItem`

NewTaggedItem instantiates a new TaggedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedItemWithDefaults

`func NewTaggedItemWithDefaults() *TaggedItem`

NewTaggedItemWithDefaults instantiates a new TaggedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaggedItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaggedItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaggedItem) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *TaggedItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaggedItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaggedItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *TaggedItem) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *TaggedItem) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *TaggedItem) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetObjectType

`func (o *TaggedItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TaggedItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TaggedItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *TaggedItem) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *TaggedItem) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *TaggedItem) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.


### GetObject

`func (o *TaggedItem) GetObject() interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TaggedItem) GetObjectOk() (*interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TaggedItem) SetObject(v interface{})`

SetObject sets Object field to given value.


### SetObjectNil

`func (o *TaggedItem) SetObjectNil(b bool)`

 SetObjectNil sets the value for Object to be an explicit nil

### UnsetObject
`func (o *TaggedItem) UnsetObject()`

UnsetObject ensures that no value is present for Object, not even an explicit nil
### GetTag

`func (o *TaggedItem) GetTag() BriefTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TaggedItem) GetTagOk() (*BriefTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TaggedItem) SetTag(v BriefTag)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


