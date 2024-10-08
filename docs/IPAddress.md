# IPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Family** | [**AggregateFamily**](AggregateFamily.md) |  | 
**Address** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewIPAddress

`func NewIPAddress(id int32, url string, display string, family AggregateFamily, address string, ) *IPAddress`

NewIPAddress instantiates a new IPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressWithDefaults

`func NewIPAddressWithDefaults() *IPAddress`

NewIPAddressWithDefaults instantiates a new IPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *IPAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddress) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *IPAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetFamily

`func (o *IPAddress) GetFamily() AggregateFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *IPAddress) GetFamilyOk() (*AggregateFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *IPAddress) SetFamily(v AggregateFamily)`

SetFamily sets Family field to given value.


### GetAddress

`func (o *IPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDescription

`func (o *IPAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


