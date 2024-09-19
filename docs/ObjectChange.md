# ObjectChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Time** | **time.Time** |  | [readonly] 
**User** | [**BriefUser**](BriefUser.md) |  | [readonly] 
**UserName** | **string** |  | [readonly] 
**RequestId** | **string** |  | [readonly] 
**Action** | [**ObjectChangeAction**](ObjectChangeAction.md) |  | 
**ChangedObjectType** | **string** |  | [readonly] 
**ChangedObjectId** | **int64** |  | 
**ChangedObject** | **interface{}** |  | [readonly] 
**PrechangeData** | **interface{}** |  | [readonly] 
**PostchangeData** | **interface{}** |  | [readonly] 

## Methods

### NewObjectChange

`func NewObjectChange(id int32, url string, displayUrl string, display string, time time.Time, user BriefUser, userName string, requestId string, action ObjectChangeAction, changedObjectType string, changedObjectId int64, changedObject interface{}, prechangeData interface{}, postchangeData interface{}, ) *ObjectChange`

NewObjectChange instantiates a new ObjectChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectChangeWithDefaults

`func NewObjectChangeWithDefaults() *ObjectChange`

NewObjectChangeWithDefaults instantiates a new ObjectChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectChange) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectChange) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectChange) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ObjectChange) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectChange) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectChange) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ObjectChange) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ObjectChange) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ObjectChange) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *ObjectChange) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ObjectChange) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ObjectChange) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetTime

`func (o *ObjectChange) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ObjectChange) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ObjectChange) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetUser

`func (o *ObjectChange) GetUser() BriefUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ObjectChange) GetUserOk() (*BriefUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ObjectChange) SetUser(v BriefUser)`

SetUser sets User field to given value.


### GetUserName

`func (o *ObjectChange) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ObjectChange) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ObjectChange) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetRequestId

`func (o *ObjectChange) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ObjectChange) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ObjectChange) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAction

`func (o *ObjectChange) GetAction() ObjectChangeAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ObjectChange) GetActionOk() (*ObjectChangeAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ObjectChange) SetAction(v ObjectChangeAction)`

SetAction sets Action field to given value.


### GetChangedObjectType

`func (o *ObjectChange) GetChangedObjectType() string`

GetChangedObjectType returns the ChangedObjectType field if non-nil, zero value otherwise.

### GetChangedObjectTypeOk

`func (o *ObjectChange) GetChangedObjectTypeOk() (*string, bool)`

GetChangedObjectTypeOk returns a tuple with the ChangedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectType

`func (o *ObjectChange) SetChangedObjectType(v string)`

SetChangedObjectType sets ChangedObjectType field to given value.


### GetChangedObjectId

`func (o *ObjectChange) GetChangedObjectId() int64`

GetChangedObjectId returns the ChangedObjectId field if non-nil, zero value otherwise.

### GetChangedObjectIdOk

`func (o *ObjectChange) GetChangedObjectIdOk() (*int64, bool)`

GetChangedObjectIdOk returns a tuple with the ChangedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObjectId

`func (o *ObjectChange) SetChangedObjectId(v int64)`

SetChangedObjectId sets ChangedObjectId field to given value.


### GetChangedObject

`func (o *ObjectChange) GetChangedObject() interface{}`

GetChangedObject returns the ChangedObject field if non-nil, zero value otherwise.

### GetChangedObjectOk

`func (o *ObjectChange) GetChangedObjectOk() (*interface{}, bool)`

GetChangedObjectOk returns a tuple with the ChangedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedObject

`func (o *ObjectChange) SetChangedObject(v interface{})`

SetChangedObject sets ChangedObject field to given value.


### SetChangedObjectNil

`func (o *ObjectChange) SetChangedObjectNil(b bool)`

 SetChangedObjectNil sets the value for ChangedObject to be an explicit nil

### UnsetChangedObject
`func (o *ObjectChange) UnsetChangedObject()`

UnsetChangedObject ensures that no value is present for ChangedObject, not even an explicit nil
### GetPrechangeData

`func (o *ObjectChange) GetPrechangeData() interface{}`

GetPrechangeData returns the PrechangeData field if non-nil, zero value otherwise.

### GetPrechangeDataOk

`func (o *ObjectChange) GetPrechangeDataOk() (*interface{}, bool)`

GetPrechangeDataOk returns a tuple with the PrechangeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrechangeData

`func (o *ObjectChange) SetPrechangeData(v interface{})`

SetPrechangeData sets PrechangeData field to given value.


### SetPrechangeDataNil

`func (o *ObjectChange) SetPrechangeDataNil(b bool)`

 SetPrechangeDataNil sets the value for PrechangeData to be an explicit nil

### UnsetPrechangeData
`func (o *ObjectChange) UnsetPrechangeData()`

UnsetPrechangeData ensures that no value is present for PrechangeData, not even an explicit nil
### GetPostchangeData

`func (o *ObjectChange) GetPostchangeData() interface{}`

GetPostchangeData returns the PostchangeData field if non-nil, zero value otherwise.

### GetPostchangeDataOk

`func (o *ObjectChange) GetPostchangeDataOk() (*interface{}, bool)`

GetPostchangeDataOk returns a tuple with the PostchangeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostchangeData

`func (o *ObjectChange) SetPostchangeData(v interface{})`

SetPostchangeData sets PostchangeData field to given value.


### SetPostchangeDataNil

`func (o *ObjectChange) SetPostchangeDataNil(b bool)`

 SetPostchangeDataNil sets the value for PostchangeData to be an explicit nil

### UnsetPostchangeData
`func (o *ObjectChange) UnsetPostchangeData()`

UnsetPostchangeData ensures that no value is present for PostchangeData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


