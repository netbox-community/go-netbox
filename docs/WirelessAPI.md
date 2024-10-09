# \WirelessAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WirelessWirelessLanGroupsBulkDestroy**](WirelessAPI.md#WirelessWirelessLanGroupsBulkDestroy) | **Delete** /api/wireless/wireless-lan-groups/ | 
[**WirelessWirelessLanGroupsBulkPartialUpdate**](WirelessAPI.md#WirelessWirelessLanGroupsBulkPartialUpdate) | **Patch** /api/wireless/wireless-lan-groups/ | 
[**WirelessWirelessLanGroupsBulkUpdate**](WirelessAPI.md#WirelessWirelessLanGroupsBulkUpdate) | **Put** /api/wireless/wireless-lan-groups/ | 
[**WirelessWirelessLanGroupsCreate**](WirelessAPI.md#WirelessWirelessLanGroupsCreate) | **Post** /api/wireless/wireless-lan-groups/ | 
[**WirelessWirelessLanGroupsDestroy**](WirelessAPI.md#WirelessWirelessLanGroupsDestroy) | **Delete** /api/wireless/wireless-lan-groups/{id}/ | 
[**WirelessWirelessLanGroupsList**](WirelessAPI.md#WirelessWirelessLanGroupsList) | **Get** /api/wireless/wireless-lan-groups/ | 
[**WirelessWirelessLanGroupsPartialUpdate**](WirelessAPI.md#WirelessWirelessLanGroupsPartialUpdate) | **Patch** /api/wireless/wireless-lan-groups/{id}/ | 
[**WirelessWirelessLanGroupsRetrieve**](WirelessAPI.md#WirelessWirelessLanGroupsRetrieve) | **Get** /api/wireless/wireless-lan-groups/{id}/ | 
[**WirelessWirelessLanGroupsUpdate**](WirelessAPI.md#WirelessWirelessLanGroupsUpdate) | **Put** /api/wireless/wireless-lan-groups/{id}/ | 
[**WirelessWirelessLansBulkDestroy**](WirelessAPI.md#WirelessWirelessLansBulkDestroy) | **Delete** /api/wireless/wireless-lans/ | 
[**WirelessWirelessLansBulkPartialUpdate**](WirelessAPI.md#WirelessWirelessLansBulkPartialUpdate) | **Patch** /api/wireless/wireless-lans/ | 
[**WirelessWirelessLansBulkUpdate**](WirelessAPI.md#WirelessWirelessLansBulkUpdate) | **Put** /api/wireless/wireless-lans/ | 
[**WirelessWirelessLansCreate**](WirelessAPI.md#WirelessWirelessLansCreate) | **Post** /api/wireless/wireless-lans/ | 
[**WirelessWirelessLansDestroy**](WirelessAPI.md#WirelessWirelessLansDestroy) | **Delete** /api/wireless/wireless-lans/{id}/ | 
[**WirelessWirelessLansList**](WirelessAPI.md#WirelessWirelessLansList) | **Get** /api/wireless/wireless-lans/ | 
[**WirelessWirelessLansPartialUpdate**](WirelessAPI.md#WirelessWirelessLansPartialUpdate) | **Patch** /api/wireless/wireless-lans/{id}/ | 
[**WirelessWirelessLansRetrieve**](WirelessAPI.md#WirelessWirelessLansRetrieve) | **Get** /api/wireless/wireless-lans/{id}/ | 
[**WirelessWirelessLansUpdate**](WirelessAPI.md#WirelessWirelessLansUpdate) | **Put** /api/wireless/wireless-lans/{id}/ | 
[**WirelessWirelessLinksBulkDestroy**](WirelessAPI.md#WirelessWirelessLinksBulkDestroy) | **Delete** /api/wireless/wireless-links/ | 
[**WirelessWirelessLinksBulkPartialUpdate**](WirelessAPI.md#WirelessWirelessLinksBulkPartialUpdate) | **Patch** /api/wireless/wireless-links/ | 
[**WirelessWirelessLinksBulkUpdate**](WirelessAPI.md#WirelessWirelessLinksBulkUpdate) | **Put** /api/wireless/wireless-links/ | 
[**WirelessWirelessLinksCreate**](WirelessAPI.md#WirelessWirelessLinksCreate) | **Post** /api/wireless/wireless-links/ | 
[**WirelessWirelessLinksDestroy**](WirelessAPI.md#WirelessWirelessLinksDestroy) | **Delete** /api/wireless/wireless-links/{id}/ | 
[**WirelessWirelessLinksList**](WirelessAPI.md#WirelessWirelessLinksList) | **Get** /api/wireless/wireless-links/ | 
[**WirelessWirelessLinksPartialUpdate**](WirelessAPI.md#WirelessWirelessLinksPartialUpdate) | **Patch** /api/wireless/wireless-links/{id}/ | 
[**WirelessWirelessLinksRetrieve**](WirelessAPI.md#WirelessWirelessLinksRetrieve) | **Get** /api/wireless/wireless-links/{id}/ | 
[**WirelessWirelessLinksUpdate**](WirelessAPI.md#WirelessWirelessLinksUpdate) | **Put** /api/wireless/wireless-links/{id}/ | 



## WirelessWirelessLanGroupsBulkDestroy

> WirelessWirelessLanGroupsBulkDestroy(ctx).WirelessLANGroupRequest(wirelessLANGroupRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLANGroupRequest := []openapiclient.WirelessLANGroupRequest{*openapiclient.NewWirelessLANGroupRequest("Name_example", "Slug_example")} // []WirelessLANGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsBulkDestroy(context.Background()).WirelessLANGroupRequest(wirelessLANGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLANGroupRequest** | [**[]WirelessLANGroupRequest**](WirelessLANGroupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsBulkPartialUpdate

> []WirelessLANGroup WirelessWirelessLanGroupsBulkPartialUpdate(ctx).WirelessLANGroupRequest(wirelessLANGroupRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLANGroupRequest := []openapiclient.WirelessLANGroupRequest{*openapiclient.NewWirelessLANGroupRequest("Name_example", "Slug_example")} // []WirelessLANGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsBulkPartialUpdate(context.Background()).WirelessLANGroupRequest(wirelessLANGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLanGroupsBulkPartialUpdate`: []WirelessLANGroup
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLanGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLANGroupRequest** | [**[]WirelessLANGroupRequest**](WirelessLANGroupRequest.md) |  | 

### Return type

[**[]WirelessLANGroup**](WirelessLANGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsBulkUpdate

> []WirelessLANGroup WirelessWirelessLanGroupsBulkUpdate(ctx).WirelessLANGroupRequest(wirelessLANGroupRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLANGroupRequest := []openapiclient.WirelessLANGroupRequest{*openapiclient.NewWirelessLANGroupRequest("Name_example", "Slug_example")} // []WirelessLANGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsBulkUpdate(context.Background()).WirelessLANGroupRequest(wirelessLANGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLanGroupsBulkUpdate`: []WirelessLANGroup
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLanGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLANGroupRequest** | [**[]WirelessLANGroupRequest**](WirelessLANGroupRequest.md) |  | 

### Return type

[**[]WirelessLANGroup**](WirelessLANGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsCreate

> WirelessLANGroup WirelessWirelessLanGroupsCreate(ctx).WritableWirelessLANGroupRequest(writableWirelessLANGroupRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	writableWirelessLANGroupRequest := *openapiclient.NewWritableWirelessLANGroupRequest("Name_example", "Slug_example", NullableInt32(123)) // WritableWirelessLANGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsCreate(context.Background()).WritableWirelessLANGroupRequest(writableWirelessLANGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLanGroupsCreate`: WirelessLANGroup
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLanGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableWirelessLANGroupRequest** | [**WritableWirelessLANGroupRequest**](WritableWirelessLANGroupRequest.md) |  | 

### Return type

[**WirelessLANGroup**](WirelessLANGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsDestroy

> WirelessWirelessLanGroupsDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsList

> PaginatedWirelessLANGroupList WirelessWirelessLanGroupsList(ctx).Ancestor(ancestor).AncestorN(ancestorN).AncestorId(ancestorId).AncestorIdN(ancestorIdN).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).Offset(offset).Ordering(ordering).Parent(parent).ParentN(parentN).ParentId(parentId).ParentIdN(parentIdN).Q(q).Slug(slug).SlugEmpty(slugEmpty).SlugIc(slugIc).SlugIe(slugIe).SlugIew(slugIew).SlugIsw(slugIsw).SlugN(slugN).SlugNic(slugNic).SlugNie(slugNie).SlugNiew(slugNiew).SlugNisw(slugNisw).Tag(tag).TagN(tagN).UpdatedByRequest(updatedByRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	ancestor := []string{"Inner_example"} // []string |  (optional)
	ancestorN := []string{"Inner_example"} // []string |  (optional)
	ancestorId := []string{"Inner_example"} // []string |  (optional)
	ancestorIdN := []string{"Inner_example"} // []string |  (optional)
	created := []time.Time{time.Now()} // []time.Time |  (optional)
	createdEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdN := []time.Time{time.Now()} // []time.Time |  (optional)
	createdByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	description := []string{"Inner_example"} // []string |  (optional)
	descriptionEmpty := true // bool |  (optional)
	descriptionIc := []string{"Inner_example"} // []string |  (optional)
	descriptionIe := []string{"Inner_example"} // []string |  (optional)
	descriptionIew := []string{"Inner_example"} // []string |  (optional)
	descriptionIsw := []string{"Inner_example"} // []string |  (optional)
	descriptionN := []string{"Inner_example"} // []string |  (optional)
	descriptionNic := []string{"Inner_example"} // []string |  (optional)
	descriptionNie := []string{"Inner_example"} // []string |  (optional)
	descriptionNiew := []string{"Inner_example"} // []string |  (optional)
	descriptionNisw := []string{"Inner_example"} // []string |  (optional)
	id := []int32{int32(123)} // []int32 |  (optional)
	idEmpty := true // bool |  (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	lastUpdated := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedN := []time.Time{time.Now()} // []time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	modifiedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	name := []string{"Inner_example"} // []string |  (optional)
	nameEmpty := true // bool |  (optional)
	nameIc := []string{"Inner_example"} // []string |  (optional)
	nameIe := []string{"Inner_example"} // []string |  (optional)
	nameIew := []string{"Inner_example"} // []string |  (optional)
	nameIsw := []string{"Inner_example"} // []string |  (optional)
	nameN := []string{"Inner_example"} // []string |  (optional)
	nameNic := []string{"Inner_example"} // []string |  (optional)
	nameNie := []string{"Inner_example"} // []string |  (optional)
	nameNiew := []string{"Inner_example"} // []string |  (optional)
	nameNisw := []string{"Inner_example"} // []string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	parent := []string{"Inner_example"} // []string |  (optional)
	parentN := []string{"Inner_example"} // []string |  (optional)
	parentId := []*int32{int32(123)} // []*int32 |  (optional)
	parentIdN := []*int32{int32(123)} // []*int32 |  (optional)
	q := "q_example" // string | Search (optional)
	slug := []string{"Inner_example"} // []string |  (optional)
	slugEmpty := true // bool |  (optional)
	slugIc := []string{"Inner_example"} // []string |  (optional)
	slugIe := []string{"Inner_example"} // []string |  (optional)
	slugIew := []string{"Inner_example"} // []string |  (optional)
	slugIsw := []string{"Inner_example"} // []string |  (optional)
	slugN := []string{"Inner_example"} // []string |  (optional)
	slugNic := []string{"Inner_example"} // []string |  (optional)
	slugNie := []string{"Inner_example"} // []string |  (optional)
	slugNiew := []string{"Inner_example"} // []string |  (optional)
	slugNisw := []string{"Inner_example"} // []string |  (optional)
	tag := []string{"Inner_example"} // []string |  (optional)
	tagN := []string{"Inner_example"} // []string |  (optional)
	updatedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsList(context.Background()).Ancestor(ancestor).AncestorN(ancestorN).AncestorId(ancestorId).AncestorIdN(ancestorIdN).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).Offset(offset).Ordering(ordering).Parent(parent).ParentN(parentN).ParentId(parentId).ParentIdN(parentIdN).Q(q).Slug(slug).SlugEmpty(slugEmpty).SlugIc(slugIc).SlugIe(slugIe).SlugIew(slugIew).SlugIsw(slugIsw).SlugN(slugN).SlugNic(slugNic).SlugNie(slugNie).SlugNiew(slugNiew).SlugNisw(slugNisw).Tag(tag).TagN(tagN).UpdatedByRequest(updatedByRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLanGroupsList`: PaginatedWirelessLANGroupList
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLanGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ancestor** | **[]string** |  | 
 **ancestorN** | **[]string** |  | 
 **ancestorId** | **[]string** |  | 
 **ancestorIdN** | **[]string** |  | 
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdLt** | [**[]time.Time**](time.Time.md) |  | 
 **createdLte** | [**[]time.Time**](time.Time.md) |  | 
 **createdN** | [**[]time.Time**](time.Time.md) |  | 
 **createdByRequest** | **string** |  | 
 **description** | **[]string** |  | 
 **descriptionEmpty** | **bool** |  | 
 **descriptionIc** | **[]string** |  | 
 **descriptionIe** | **[]string** |  | 
 **descriptionIew** | **[]string** |  | 
 **descriptionIsw** | **[]string** |  | 
 **descriptionN** | **[]string** |  | 
 **descriptionNic** | **[]string** |  | 
 **descriptionNie** | **[]string** |  | 
 **descriptionNiew** | **[]string** |  | 
 **descriptionNisw** | **[]string** |  | 
 **id** | **[]int32** |  | 
 **idEmpty** | **bool** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **lastUpdated** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedN** | [**[]time.Time**](time.Time.md) |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedByRequest** | **string** |  | 
 **name** | **[]string** |  | 
 **nameEmpty** | **bool** |  | 
 **nameIc** | **[]string** |  | 
 **nameIe** | **[]string** |  | 
 **nameIew** | **[]string** |  | 
 **nameIsw** | **[]string** |  | 
 **nameN** | **[]string** |  | 
 **nameNic** | **[]string** |  | 
 **nameNie** | **[]string** |  | 
 **nameNiew** | **[]string** |  | 
 **nameNisw** | **[]string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **parent** | **[]string** |  | 
 **parentN** | **[]string** |  | 
 **parentId** | **[]int32** |  | 
 **parentIdN** | **[]int32** |  | 
 **q** | **string** | Search | 
 **slug** | **[]string** |  | 
 **slugEmpty** | **bool** |  | 
 **slugIc** | **[]string** |  | 
 **slugIe** | **[]string** |  | 
 **slugIew** | **[]string** |  | 
 **slugIsw** | **[]string** |  | 
 **slugN** | **[]string** |  | 
 **slugNic** | **[]string** |  | 
 **slugNie** | **[]string** |  | 
 **slugNiew** | **[]string** |  | 
 **slugNisw** | **[]string** |  | 
 **tag** | **[]string** |  | 
 **tagN** | **[]string** |  | 
 **updatedByRequest** | **string** |  | 

### Return type

[**PaginatedWirelessLANGroupList**](PaginatedWirelessLANGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsPartialUpdate

> WirelessLANGroup WirelessWirelessLanGroupsPartialUpdate(ctx, id).PatchedWritableWirelessLANGroupRequest(patchedWritableWirelessLANGroupRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN group.
	patchedWritableWirelessLANGroupRequest := *openapiclient.NewPatchedWritableWirelessLANGroupRequest() // PatchedWritableWirelessLANGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsPartialUpdate(context.Background(), id).PatchedWritableWirelessLANGroupRequest(patchedWritableWirelessLANGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLanGroupsPartialUpdate`: WirelessLANGroup
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLanGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableWirelessLANGroupRequest** | [**PatchedWritableWirelessLANGroupRequest**](PatchedWritableWirelessLANGroupRequest.md) |  | 

### Return type

[**WirelessLANGroup**](WirelessLANGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsRetrieve

> WirelessLANGroup WirelessWirelessLanGroupsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLanGroupsRetrieve`: WirelessLANGroup
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLanGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WirelessLANGroup**](WirelessLANGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLanGroupsUpdate

> WirelessLANGroup WirelessWirelessLanGroupsUpdate(ctx, id).WritableWirelessLANGroupRequest(writableWirelessLANGroupRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN group.
	writableWirelessLANGroupRequest := *openapiclient.NewWritableWirelessLANGroupRequest("Name_example", "Slug_example", NullableInt32(123)) // WritableWirelessLANGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLanGroupsUpdate(context.Background(), id).WritableWirelessLANGroupRequest(writableWirelessLANGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLanGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLanGroupsUpdate`: WirelessLANGroup
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLanGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLanGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableWirelessLANGroupRequest** | [**WritableWirelessLANGroupRequest**](WritableWirelessLANGroupRequest.md) |  | 

### Return type

[**WirelessLANGroup**](WirelessLANGroup.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansBulkDestroy

> WirelessWirelessLansBulkDestroy(ctx).WirelessLANRequest(wirelessLANRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLANRequest := []openapiclient.WirelessLANRequest{*openapiclient.NewWirelessLANRequest("Ssid_example")} // []WirelessLANRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WirelessAPI.WirelessWirelessLansBulkDestroy(context.Background()).WirelessLANRequest(wirelessLANRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLANRequest** | [**[]WirelessLANRequest**](WirelessLANRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansBulkPartialUpdate

> []WirelessLAN WirelessWirelessLansBulkPartialUpdate(ctx).WirelessLANRequest(wirelessLANRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLANRequest := []openapiclient.WirelessLANRequest{*openapiclient.NewWirelessLANRequest("Ssid_example")} // []WirelessLANRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLansBulkPartialUpdate(context.Background()).WirelessLANRequest(wirelessLANRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLansBulkPartialUpdate`: []WirelessLAN
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLansBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLANRequest** | [**[]WirelessLANRequest**](WirelessLANRequest.md) |  | 

### Return type

[**[]WirelessLAN**](WirelessLAN.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansBulkUpdate

> []WirelessLAN WirelessWirelessLansBulkUpdate(ctx).WirelessLANRequest(wirelessLANRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLANRequest := []openapiclient.WirelessLANRequest{*openapiclient.NewWirelessLANRequest("Ssid_example")} // []WirelessLANRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLansBulkUpdate(context.Background()).WirelessLANRequest(wirelessLANRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLansBulkUpdate`: []WirelessLAN
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLansBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLANRequest** | [**[]WirelessLANRequest**](WirelessLANRequest.md) |  | 

### Return type

[**[]WirelessLAN**](WirelessLAN.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansCreate

> WirelessLAN WirelessWirelessLansCreate(ctx).WritableWirelessLANRequest(writableWirelessLANRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	writableWirelessLANRequest := *openapiclient.NewWritableWirelessLANRequest("Ssid_example") // WritableWirelessLANRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLansCreate(context.Background()).WritableWirelessLANRequest(writableWirelessLANRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLansCreate`: WirelessLAN
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLansCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableWirelessLANRequest** | [**WritableWirelessLANRequest**](WritableWirelessLANRequest.md) |  | 

### Return type

[**WirelessLAN**](WirelessLAN.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansDestroy

> WirelessWirelessLansDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WirelessAPI.WirelessWirelessLansDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansList

> PaginatedWirelessLANList WirelessWirelessLansList(ctx).AuthCipher(authCipher).AuthCipherEmpty(authCipherEmpty).AuthCipherIc(authCipherIc).AuthCipherIe(authCipherIe).AuthCipherIew(authCipherIew).AuthCipherIsw(authCipherIsw).AuthCipherN(authCipherN).AuthCipherNic(authCipherNic).AuthCipherNie(authCipherNie).AuthCipherNiew(authCipherNiew).AuthCipherNisw(authCipherNisw).AuthPsk(authPsk).AuthPskEmpty(authPskEmpty).AuthPskIc(authPskIc).AuthPskIe(authPskIe).AuthPskIew(authPskIew).AuthPskIsw(authPskIsw).AuthPskN(authPskN).AuthPskNic(authPskNic).AuthPskNie(authPskNie).AuthPskNiew(authPskNiew).AuthPskNisw(authPskNisw).AuthType(authType).AuthTypeEmpty(authTypeEmpty).AuthTypeIc(authTypeIc).AuthTypeIe(authTypeIe).AuthTypeIew(authTypeIew).AuthTypeIsw(authTypeIsw).AuthTypeN(authTypeN).AuthTypeNic(authTypeNic).AuthTypeNie(authTypeNie).AuthTypeNiew(authTypeNiew).AuthTypeNisw(authTypeNisw).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).InterfaceId(interfaceId).InterfaceIdN(interfaceIdN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Offset(offset).Ordering(ordering).Q(q).Ssid(ssid).SsidEmpty(ssidEmpty).SsidIc(ssidIc).SsidIe(ssidIe).SsidIew(ssidIew).SsidIsw(ssidIsw).SsidN(ssidN).SsidNic(ssidNic).SsidNie(ssidNie).SsidNiew(ssidNiew).SsidNisw(ssidNisw).Status(status).StatusEmpty(statusEmpty).StatusIc(statusIc).StatusIe(statusIe).StatusIew(statusIew).StatusIsw(statusIsw).StatusN(statusN).StatusNic(statusNic).StatusNie(statusNie).StatusNiew(statusNiew).StatusNisw(statusNisw).Tag(tag).TagN(tagN).Tenant(tenant).TenantN(tenantN).TenantGroup(tenantGroup).TenantGroupN(tenantGroupN).TenantGroupId(tenantGroupId).TenantGroupIdN(tenantGroupIdN).TenantId(tenantId).TenantIdN(tenantIdN).UpdatedByRequest(updatedByRequest).VlanId(vlanId).VlanIdN(vlanIdN).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	authCipher := []string{"Inner_example"} // []string |  (optional)
	authCipherEmpty := true // bool |  (optional)
	authCipherIc := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherIe := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherIew := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherIsw := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherN := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNic := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNie := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNiew := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNisw := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authPsk := []string{"Inner_example"} // []string |  (optional)
	authPskEmpty := true // bool |  (optional)
	authPskIc := []string{"Inner_example"} // []string |  (optional)
	authPskIe := []string{"Inner_example"} // []string |  (optional)
	authPskIew := []string{"Inner_example"} // []string |  (optional)
	authPskIsw := []string{"Inner_example"} // []string |  (optional)
	authPskN := []string{"Inner_example"} // []string |  (optional)
	authPskNic := []string{"Inner_example"} // []string |  (optional)
	authPskNie := []string{"Inner_example"} // []string |  (optional)
	authPskNiew := []string{"Inner_example"} // []string |  (optional)
	authPskNisw := []string{"Inner_example"} // []string |  (optional)
	authType := []string{"Inner_example"} // []string |  (optional)
	authTypeEmpty := true // bool |  (optional)
	authTypeIc := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeIe := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeIew := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeIsw := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeN := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNic := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNie := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNiew := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNisw := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	created := []time.Time{time.Now()} // []time.Time |  (optional)
	createdEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdN := []time.Time{time.Now()} // []time.Time |  (optional)
	createdByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	description := []string{"Inner_example"} // []string |  (optional)
	descriptionEmpty := true // bool |  (optional)
	descriptionIc := []string{"Inner_example"} // []string |  (optional)
	descriptionIe := []string{"Inner_example"} // []string |  (optional)
	descriptionIew := []string{"Inner_example"} // []string |  (optional)
	descriptionIsw := []string{"Inner_example"} // []string |  (optional)
	descriptionN := []string{"Inner_example"} // []string |  (optional)
	descriptionNic := []string{"Inner_example"} // []string |  (optional)
	descriptionNie := []string{"Inner_example"} // []string |  (optional)
	descriptionNiew := []string{"Inner_example"} // []string |  (optional)
	descriptionNisw := []string{"Inner_example"} // []string |  (optional)
	group := []string{"Inner_example"} // []string |  (optional)
	groupN := []string{"Inner_example"} // []string |  (optional)
	groupId := []string{"Inner_example"} // []string |  (optional)
	groupIdN := []string{"Inner_example"} // []string |  (optional)
	id := []int32{int32(123)} // []int32 |  (optional)
	idEmpty := true // bool |  (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	interfaceId := []int32{int32(123)} // []int32 |  (optional)
	interfaceIdN := []int32{int32(123)} // []int32 |  (optional)
	lastUpdated := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedN := []time.Time{time.Now()} // []time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	modifiedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	q := "q_example" // string | Search (optional)
	ssid := []string{"Inner_example"} // []string |  (optional)
	ssidEmpty := true // bool |  (optional)
	ssidIc := []string{"Inner_example"} // []string |  (optional)
	ssidIe := []string{"Inner_example"} // []string |  (optional)
	ssidIew := []string{"Inner_example"} // []string |  (optional)
	ssidIsw := []string{"Inner_example"} // []string |  (optional)
	ssidN := []string{"Inner_example"} // []string |  (optional)
	ssidNic := []string{"Inner_example"} // []string |  (optional)
	ssidNie := []string{"Inner_example"} // []string |  (optional)
	ssidNiew := []string{"Inner_example"} // []string |  (optional)
	ssidNisw := []string{"Inner_example"} // []string |  (optional)
	status := []string{"Inner_example"} // []string |  (optional)
	statusEmpty := true // bool |  (optional)
	statusIc := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusIe := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusIew := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusIsw := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusN := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusNic := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusNie := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusNiew := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	statusNisw := []openapiclient.WirelessWirelessLansListStatusIcParameterInner{openapiclient.wireless_wireless_lans_list_status__ic_parameter_inner("active")} // []WirelessWirelessLansListStatusIcParameterInner | * `active` - Active * `reserved` - Reserved * `disabled` - Disabled * `deprecated` - Deprecated (optional)
	tag := []string{"Inner_example"} // []string |  (optional)
	tagN := []string{"Inner_example"} // []string |  (optional)
	tenant := []string{"Inner_example"} // []string | Tenant (slug) (optional)
	tenantN := []string{"Inner_example"} // []string | Tenant (slug) (optional)
	tenantGroup := []string{"Inner_example"} // []string |  (optional)
	tenantGroupN := []string{"Inner_example"} // []string |  (optional)
	tenantGroupId := []string{"Inner_example"} // []string |  (optional)
	tenantGroupIdN := []string{"Inner_example"} // []string |  (optional)
	tenantId := []*int32{int32(123)} // []*int32 | Tenant (ID) (optional)
	tenantIdN := []*int32{int32(123)} // []*int32 | Tenant (ID) (optional)
	updatedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	vlanId := []*int32{int32(123)} // []*int32 |  (optional)
	vlanIdN := []*int32{int32(123)} // []*int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLansList(context.Background()).AuthCipher(authCipher).AuthCipherEmpty(authCipherEmpty).AuthCipherIc(authCipherIc).AuthCipherIe(authCipherIe).AuthCipherIew(authCipherIew).AuthCipherIsw(authCipherIsw).AuthCipherN(authCipherN).AuthCipherNic(authCipherNic).AuthCipherNie(authCipherNie).AuthCipherNiew(authCipherNiew).AuthCipherNisw(authCipherNisw).AuthPsk(authPsk).AuthPskEmpty(authPskEmpty).AuthPskIc(authPskIc).AuthPskIe(authPskIe).AuthPskIew(authPskIew).AuthPskIsw(authPskIsw).AuthPskN(authPskN).AuthPskNic(authPskNic).AuthPskNie(authPskNie).AuthPskNiew(authPskNiew).AuthPskNisw(authPskNisw).AuthType(authType).AuthTypeEmpty(authTypeEmpty).AuthTypeIc(authTypeIc).AuthTypeIe(authTypeIe).AuthTypeIew(authTypeIew).AuthTypeIsw(authTypeIsw).AuthTypeN(authTypeN).AuthTypeNic(authTypeNic).AuthTypeNie(authTypeNie).AuthTypeNiew(authTypeNiew).AuthTypeNisw(authTypeNisw).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).InterfaceId(interfaceId).InterfaceIdN(interfaceIdN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Offset(offset).Ordering(ordering).Q(q).Ssid(ssid).SsidEmpty(ssidEmpty).SsidIc(ssidIc).SsidIe(ssidIe).SsidIew(ssidIew).SsidIsw(ssidIsw).SsidN(ssidN).SsidNic(ssidNic).SsidNie(ssidNie).SsidNiew(ssidNiew).SsidNisw(ssidNisw).Status(status).StatusEmpty(statusEmpty).StatusIc(statusIc).StatusIe(statusIe).StatusIew(statusIew).StatusIsw(statusIsw).StatusN(statusN).StatusNic(statusNic).StatusNie(statusNie).StatusNiew(statusNiew).StatusNisw(statusNisw).Tag(tag).TagN(tagN).Tenant(tenant).TenantN(tenantN).TenantGroup(tenantGroup).TenantGroupN(tenantGroupN).TenantGroupId(tenantGroupId).TenantGroupIdN(tenantGroupIdN).TenantId(tenantId).TenantIdN(tenantIdN).UpdatedByRequest(updatedByRequest).VlanId(vlanId).VlanIdN(vlanIdN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLansList`: PaginatedWirelessLANList
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLansList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authCipher** | **[]string** |  | 
 **authCipherEmpty** | **bool** |  | 
 **authCipherIc** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherIe** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherIew** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherIsw** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherN** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNic** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNie** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNiew** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNisw** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authPsk** | **[]string** |  | 
 **authPskEmpty** | **bool** |  | 
 **authPskIc** | **[]string** |  | 
 **authPskIe** | **[]string** |  | 
 **authPskIew** | **[]string** |  | 
 **authPskIsw** | **[]string** |  | 
 **authPskN** | **[]string** |  | 
 **authPskNic** | **[]string** |  | 
 **authPskNie** | **[]string** |  | 
 **authPskNiew** | **[]string** |  | 
 **authPskNisw** | **[]string** |  | 
 **authType** | **[]string** |  | 
 **authTypeEmpty** | **bool** |  | 
 **authTypeIc** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeIe** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeIew** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeIsw** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeN** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNic** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNie** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNiew** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNisw** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdLt** | [**[]time.Time**](time.Time.md) |  | 
 **createdLte** | [**[]time.Time**](time.Time.md) |  | 
 **createdN** | [**[]time.Time**](time.Time.md) |  | 
 **createdByRequest** | **string** |  | 
 **description** | **[]string** |  | 
 **descriptionEmpty** | **bool** |  | 
 **descriptionIc** | **[]string** |  | 
 **descriptionIe** | **[]string** |  | 
 **descriptionIew** | **[]string** |  | 
 **descriptionIsw** | **[]string** |  | 
 **descriptionN** | **[]string** |  | 
 **descriptionNic** | **[]string** |  | 
 **descriptionNie** | **[]string** |  | 
 **descriptionNiew** | **[]string** |  | 
 **descriptionNisw** | **[]string** |  | 
 **group** | **[]string** |  | 
 **groupN** | **[]string** |  | 
 **groupId** | **[]string** |  | 
 **groupIdN** | **[]string** |  | 
 **id** | **[]int32** |  | 
 **idEmpty** | **bool** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **interfaceId** | **[]int32** |  | 
 **interfaceIdN** | **[]int32** |  | 
 **lastUpdated** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedN** | [**[]time.Time**](time.Time.md) |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedByRequest** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 
 **ssid** | **[]string** |  | 
 **ssidEmpty** | **bool** |  | 
 **ssidIc** | **[]string** |  | 
 **ssidIe** | **[]string** |  | 
 **ssidIew** | **[]string** |  | 
 **ssidIsw** | **[]string** |  | 
 **ssidN** | **[]string** |  | 
 **ssidNic** | **[]string** |  | 
 **ssidNie** | **[]string** |  | 
 **ssidNiew** | **[]string** |  | 
 **ssidNisw** | **[]string** |  | 
 **status** | **[]string** |  | 
 **statusEmpty** | **bool** |  | 
 **statusIc** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusIe** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusIew** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusIsw** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusN** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusNic** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusNie** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusNiew** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **statusNisw** | [**[]WirelessWirelessLansListStatusIcParameterInner**](WirelessWirelessLansListStatusIcParameterInner.md) | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;disabled&#x60; - Disabled * &#x60;deprecated&#x60; - Deprecated | 
 **tag** | **[]string** |  | 
 **tagN** | **[]string** |  | 
 **tenant** | **[]string** | Tenant (slug) | 
 **tenantN** | **[]string** | Tenant (slug) | 
 **tenantGroup** | **[]string** |  | 
 **tenantGroupN** | **[]string** |  | 
 **tenantGroupId** | **[]string** |  | 
 **tenantGroupIdN** | **[]string** |  | 
 **tenantId** | **[]int32** | Tenant (ID) | 
 **tenantIdN** | **[]int32** | Tenant (ID) | 
 **updatedByRequest** | **string** |  | 
 **vlanId** | **[]int32** |  | 
 **vlanIdN** | **[]int32** |  | 

### Return type

[**PaginatedWirelessLANList**](PaginatedWirelessLANList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansPartialUpdate

> WirelessLAN WirelessWirelessLansPartialUpdate(ctx, id).PatchedWritableWirelessLANRequest(patchedWritableWirelessLANRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN.
	patchedWritableWirelessLANRequest := *openapiclient.NewPatchedWritableWirelessLANRequest() // PatchedWritableWirelessLANRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLansPartialUpdate(context.Background(), id).PatchedWritableWirelessLANRequest(patchedWritableWirelessLANRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLansPartialUpdate`: WirelessLAN
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLansPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableWirelessLANRequest** | [**PatchedWritableWirelessLANRequest**](PatchedWritableWirelessLANRequest.md) |  | 

### Return type

[**WirelessLAN**](WirelessLAN.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansRetrieve

> WirelessLAN WirelessWirelessLansRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLansRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLansRetrieve`: WirelessLAN
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLansRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WirelessLAN**](WirelessLAN.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLansUpdate

> WirelessLAN WirelessWirelessLansUpdate(ctx, id).WritableWirelessLANRequest(writableWirelessLANRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless LAN.
	writableWirelessLANRequest := *openapiclient.NewWritableWirelessLANRequest("Ssid_example") // WritableWirelessLANRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLansUpdate(context.Background(), id).WritableWirelessLANRequest(writableWirelessLANRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLansUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLansUpdate`: WirelessLAN
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLansUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless LAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLansUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableWirelessLANRequest** | [**WritableWirelessLANRequest**](WritableWirelessLANRequest.md) |  | 

### Return type

[**WirelessLAN**](WirelessLAN.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksBulkDestroy

> WirelessWirelessLinksBulkDestroy(ctx).WirelessLinkRequest(wirelessLinkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLinkRequest := []openapiclient.WirelessLinkRequest{*openapiclient.NewWirelessLinkRequest(*openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"), *openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"))} // []WirelessLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WirelessAPI.WirelessWirelessLinksBulkDestroy(context.Background()).WirelessLinkRequest(wirelessLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLinkRequest** | [**[]WirelessLinkRequest**](WirelessLinkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksBulkPartialUpdate

> []WirelessLink WirelessWirelessLinksBulkPartialUpdate(ctx).WirelessLinkRequest(wirelessLinkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLinkRequest := []openapiclient.WirelessLinkRequest{*openapiclient.NewWirelessLinkRequest(*openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"), *openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"))} // []WirelessLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLinksBulkPartialUpdate(context.Background()).WirelessLinkRequest(wirelessLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLinksBulkPartialUpdate`: []WirelessLink
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLinksBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLinkRequest** | [**[]WirelessLinkRequest**](WirelessLinkRequest.md) |  | 

### Return type

[**[]WirelessLink**](WirelessLink.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksBulkUpdate

> []WirelessLink WirelessWirelessLinksBulkUpdate(ctx).WirelessLinkRequest(wirelessLinkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	wirelessLinkRequest := []openapiclient.WirelessLinkRequest{*openapiclient.NewWirelessLinkRequest(*openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"), *openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"))} // []WirelessLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLinksBulkUpdate(context.Background()).WirelessLinkRequest(wirelessLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLinksBulkUpdate`: []WirelessLink
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLinksBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wirelessLinkRequest** | [**[]WirelessLinkRequest**](WirelessLinkRequest.md) |  | 

### Return type

[**[]WirelessLink**](WirelessLink.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksCreate

> WirelessLink WirelessWirelessLinksCreate(ctx).WritableWirelessLinkRequest(writableWirelessLinkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	writableWirelessLinkRequest := *openapiclient.NewWritableWirelessLinkRequest(*openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"), *openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example")) // WritableWirelessLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLinksCreate(context.Background()).WritableWirelessLinkRequest(writableWirelessLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLinksCreate`: WirelessLink
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLinksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableWirelessLinkRequest** | [**WritableWirelessLinkRequest**](WritableWirelessLinkRequest.md) |  | 

### Return type

[**WirelessLink**](WirelessLink.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksDestroy

> WirelessWirelessLinksDestroy(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WirelessAPI.WirelessWirelessLinksDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksList

> PaginatedWirelessLinkList WirelessWirelessLinksList(ctx).AuthCipher(authCipher).AuthCipherEmpty(authCipherEmpty).AuthCipherIc(authCipherIc).AuthCipherIe(authCipherIe).AuthCipherIew(authCipherIew).AuthCipherIsw(authCipherIsw).AuthCipherN(authCipherN).AuthCipherNic(authCipherNic).AuthCipherNie(authCipherNie).AuthCipherNiew(authCipherNiew).AuthCipherNisw(authCipherNisw).AuthPsk(authPsk).AuthPskEmpty(authPskEmpty).AuthPskIc(authPskIc).AuthPskIe(authPskIe).AuthPskIew(authPskIew).AuthPskIsw(authPskIsw).AuthPskN(authPskN).AuthPskNic(authPskNic).AuthPskNie(authPskNie).AuthPskNiew(authPskNiew).AuthPskNisw(authPskNisw).AuthType(authType).AuthTypeEmpty(authTypeEmpty).AuthTypeIc(authTypeIc).AuthTypeIe(authTypeIe).AuthTypeIew(authTypeIew).AuthTypeIsw(authTypeIsw).AuthTypeN(authTypeN).AuthTypeNic(authTypeNic).AuthTypeNie(authTypeNie).AuthTypeNiew(authTypeNiew).AuthTypeNisw(authTypeNisw).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Distance(distance).DistanceEmpty(distanceEmpty).DistanceGt(distanceGt).DistanceGte(distanceGte).DistanceLt(distanceLt).DistanceLte(distanceLte).DistanceN(distanceN).DistanceUnit(distanceUnit).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).InterfaceAId(interfaceAId).InterfaceAIdN(interfaceAIdN).InterfaceBId(interfaceBId).InterfaceBIdN(interfaceBIdN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Offset(offset).Ordering(ordering).Q(q).Ssid(ssid).SsidEmpty(ssidEmpty).SsidIc(ssidIc).SsidIe(ssidIe).SsidIew(ssidIew).SsidIsw(ssidIsw).SsidN(ssidN).SsidNic(ssidNic).SsidNie(ssidNie).SsidNiew(ssidNiew).SsidNisw(ssidNisw).Status(status).StatusEmpty(statusEmpty).StatusIc(statusIc).StatusIe(statusIe).StatusIew(statusIew).StatusIsw(statusIsw).StatusN(statusN).StatusNic(statusNic).StatusNie(statusNie).StatusNiew(statusNiew).StatusNisw(statusNisw).Tag(tag).TagN(tagN).Tenant(tenant).TenantN(tenantN).TenantGroup(tenantGroup).TenantGroupN(tenantGroupN).TenantGroupId(tenantGroupId).TenantGroupIdN(tenantGroupIdN).TenantId(tenantId).TenantIdN(tenantIdN).UpdatedByRequest(updatedByRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	authCipher := []string{"Inner_example"} // []string |  (optional)
	authCipherEmpty := true // bool |  (optional)
	authCipherIc := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherIe := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherIew := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherIsw := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherN := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNic := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNie := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNiew := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authCipherNisw := []openapiclient.AuthenticationCipher{openapiclient.Authentication_cipher("")} // []AuthenticationCipher | * `auto` - Auto * `tkip` - TKIP * `aes` - AES (optional)
	authPsk := []string{"Inner_example"} // []string |  (optional)
	authPskEmpty := true // bool |  (optional)
	authPskIc := []string{"Inner_example"} // []string |  (optional)
	authPskIe := []string{"Inner_example"} // []string |  (optional)
	authPskIew := []string{"Inner_example"} // []string |  (optional)
	authPskIsw := []string{"Inner_example"} // []string |  (optional)
	authPskN := []string{"Inner_example"} // []string |  (optional)
	authPskNic := []string{"Inner_example"} // []string |  (optional)
	authPskNie := []string{"Inner_example"} // []string |  (optional)
	authPskNiew := []string{"Inner_example"} // []string |  (optional)
	authPskNisw := []string{"Inner_example"} // []string |  (optional)
	authType := []string{"Inner_example"} // []string |  (optional)
	authTypeEmpty := true // bool |  (optional)
	authTypeIc := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeIe := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeIew := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeIsw := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeN := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNic := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNie := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNiew := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	authTypeNisw := []openapiclient.AuthenticationType1{openapiclient.Authentication_type_1("")} // []AuthenticationType1 | * `open` - Open * `wep` - WEP * `wpa-personal` - WPA Personal (PSK) * `wpa-enterprise` - WPA Enterprise (optional)
	created := []time.Time{time.Now()} // []time.Time |  (optional)
	createdEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdN := []time.Time{time.Now()} // []time.Time |  (optional)
	createdByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	description := []string{"Inner_example"} // []string |  (optional)
	descriptionEmpty := true // bool |  (optional)
	descriptionIc := []string{"Inner_example"} // []string |  (optional)
	descriptionIe := []string{"Inner_example"} // []string |  (optional)
	descriptionIew := []string{"Inner_example"} // []string |  (optional)
	descriptionIsw := []string{"Inner_example"} // []string |  (optional)
	descriptionN := []string{"Inner_example"} // []string |  (optional)
	descriptionNic := []string{"Inner_example"} // []string |  (optional)
	descriptionNie := []string{"Inner_example"} // []string |  (optional)
	descriptionNiew := []string{"Inner_example"} // []string |  (optional)
	descriptionNisw := []string{"Inner_example"} // []string |  (optional)
	distance := []float64{float64(123)} // []float64 |  (optional)
	distanceEmpty := true // bool |  (optional)
	distanceGt := []float64{float64(123)} // []float64 |  (optional)
	distanceGte := []float64{float64(123)} // []float64 |  (optional)
	distanceLt := []float64{float64(123)} // []float64 |  (optional)
	distanceLte := []float64{float64(123)} // []float64 |  (optional)
	distanceN := []float64{float64(123)} // []float64 |  (optional)
	distanceUnit := openapiclient.wireless_wireless_links_list_distance_unit_parameter("ft") // WirelessWirelessLinksListDistanceUnitParameter | * `km` - Kilometers * `m` - Meters * `mi` - Miles * `ft` - Feet (optional)
	id := []int32{int32(123)} // []int32 |  (optional)
	idEmpty := true // bool |  (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	interfaceAId := []int32{int32(123)} // []int32 |  (optional)
	interfaceAIdN := []int32{int32(123)} // []int32 |  (optional)
	interfaceBId := []int32{int32(123)} // []int32 |  (optional)
	interfaceBIdN := []int32{int32(123)} // []int32 |  (optional)
	lastUpdated := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUpdatedN := []time.Time{time.Now()} // []time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	modifiedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	q := "q_example" // string | Search (optional)
	ssid := []string{"Inner_example"} // []string |  (optional)
	ssidEmpty := true // bool |  (optional)
	ssidIc := []string{"Inner_example"} // []string |  (optional)
	ssidIe := []string{"Inner_example"} // []string |  (optional)
	ssidIew := []string{"Inner_example"} // []string |  (optional)
	ssidIsw := []string{"Inner_example"} // []string |  (optional)
	ssidN := []string{"Inner_example"} // []string |  (optional)
	ssidNic := []string{"Inner_example"} // []string |  (optional)
	ssidNie := []string{"Inner_example"} // []string |  (optional)
	ssidNiew := []string{"Inner_example"} // []string |  (optional)
	ssidNisw := []string{"Inner_example"} // []string |  (optional)
	status := []string{"Inner_example"} // []string |  (optional)
	statusEmpty := true // bool |  (optional)
	statusIc := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusIe := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusIew := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusIsw := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusN := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusNic := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusNie := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusNiew := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	statusNisw := []openapiclient.DcimCablesListStatusIcParameterInner{openapiclient.dcim_cables_list_status__ic_parameter_inner("connected")} // []DcimCablesListStatusIcParameterInner | * `connected` - Connected * `planned` - Planned * `decommissioning` - Decommissioning (optional)
	tag := []string{"Inner_example"} // []string |  (optional)
	tagN := []string{"Inner_example"} // []string |  (optional)
	tenant := []string{"Inner_example"} // []string | Tenant (slug) (optional)
	tenantN := []string{"Inner_example"} // []string | Tenant (slug) (optional)
	tenantGroup := []string{"Inner_example"} // []string |  (optional)
	tenantGroupN := []string{"Inner_example"} // []string |  (optional)
	tenantGroupId := []string{"Inner_example"} // []string |  (optional)
	tenantGroupIdN := []string{"Inner_example"} // []string |  (optional)
	tenantId := []*int32{int32(123)} // []*int32 | Tenant (ID) (optional)
	tenantIdN := []*int32{int32(123)} // []*int32 | Tenant (ID) (optional)
	updatedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLinksList(context.Background()).AuthCipher(authCipher).AuthCipherEmpty(authCipherEmpty).AuthCipherIc(authCipherIc).AuthCipherIe(authCipherIe).AuthCipherIew(authCipherIew).AuthCipherIsw(authCipherIsw).AuthCipherN(authCipherN).AuthCipherNic(authCipherNic).AuthCipherNie(authCipherNie).AuthCipherNiew(authCipherNiew).AuthCipherNisw(authCipherNisw).AuthPsk(authPsk).AuthPskEmpty(authPskEmpty).AuthPskIc(authPskIc).AuthPskIe(authPskIe).AuthPskIew(authPskIew).AuthPskIsw(authPskIsw).AuthPskN(authPskN).AuthPskNic(authPskNic).AuthPskNie(authPskNie).AuthPskNiew(authPskNiew).AuthPskNisw(authPskNisw).AuthType(authType).AuthTypeEmpty(authTypeEmpty).AuthTypeIc(authTypeIc).AuthTypeIe(authTypeIe).AuthTypeIew(authTypeIew).AuthTypeIsw(authTypeIsw).AuthTypeN(authTypeN).AuthTypeNic(authTypeNic).AuthTypeNie(authTypeNie).AuthTypeNiew(authTypeNiew).AuthTypeNisw(authTypeNisw).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Distance(distance).DistanceEmpty(distanceEmpty).DistanceGt(distanceGt).DistanceGte(distanceGte).DistanceLt(distanceLt).DistanceLte(distanceLte).DistanceN(distanceN).DistanceUnit(distanceUnit).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).InterfaceAId(interfaceAId).InterfaceAIdN(interfaceAIdN).InterfaceBId(interfaceBId).InterfaceBIdN(interfaceBIdN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Offset(offset).Ordering(ordering).Q(q).Ssid(ssid).SsidEmpty(ssidEmpty).SsidIc(ssidIc).SsidIe(ssidIe).SsidIew(ssidIew).SsidIsw(ssidIsw).SsidN(ssidN).SsidNic(ssidNic).SsidNie(ssidNie).SsidNiew(ssidNiew).SsidNisw(ssidNisw).Status(status).StatusEmpty(statusEmpty).StatusIc(statusIc).StatusIe(statusIe).StatusIew(statusIew).StatusIsw(statusIsw).StatusN(statusN).StatusNic(statusNic).StatusNie(statusNie).StatusNiew(statusNiew).StatusNisw(statusNisw).Tag(tag).TagN(tagN).Tenant(tenant).TenantN(tenantN).TenantGroup(tenantGroup).TenantGroupN(tenantGroupN).TenantGroupId(tenantGroupId).TenantGroupIdN(tenantGroupIdN).TenantId(tenantId).TenantIdN(tenantIdN).UpdatedByRequest(updatedByRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLinksList`: PaginatedWirelessLinkList
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLinksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authCipher** | **[]string** |  | 
 **authCipherEmpty** | **bool** |  | 
 **authCipherIc** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherIe** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherIew** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherIsw** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherN** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNic** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNie** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNiew** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authCipherNisw** | [**[]AuthenticationCipher**](AuthenticationCipher.md) | * &#x60;auto&#x60; - Auto * &#x60;tkip&#x60; - TKIP * &#x60;aes&#x60; - AES | 
 **authPsk** | **[]string** |  | 
 **authPskEmpty** | **bool** |  | 
 **authPskIc** | **[]string** |  | 
 **authPskIe** | **[]string** |  | 
 **authPskIew** | **[]string** |  | 
 **authPskIsw** | **[]string** |  | 
 **authPskN** | **[]string** |  | 
 **authPskNic** | **[]string** |  | 
 **authPskNie** | **[]string** |  | 
 **authPskNiew** | **[]string** |  | 
 **authPskNisw** | **[]string** |  | 
 **authType** | **[]string** |  | 
 **authTypeEmpty** | **bool** |  | 
 **authTypeIc** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeIe** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeIew** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeIsw** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeN** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNic** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNie** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNiew** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **authTypeNisw** | [**[]AuthenticationType1**](AuthenticationType1.md) | * &#x60;open&#x60; - Open * &#x60;wep&#x60; - WEP * &#x60;wpa-personal&#x60; - WPA Personal (PSK) * &#x60;wpa-enterprise&#x60; - WPA Enterprise | 
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdLt** | [**[]time.Time**](time.Time.md) |  | 
 **createdLte** | [**[]time.Time**](time.Time.md) |  | 
 **createdN** | [**[]time.Time**](time.Time.md) |  | 
 **createdByRequest** | **string** |  | 
 **description** | **[]string** |  | 
 **descriptionEmpty** | **bool** |  | 
 **descriptionIc** | **[]string** |  | 
 **descriptionIe** | **[]string** |  | 
 **descriptionIew** | **[]string** |  | 
 **descriptionIsw** | **[]string** |  | 
 **descriptionN** | **[]string** |  | 
 **descriptionNic** | **[]string** |  | 
 **descriptionNie** | **[]string** |  | 
 **descriptionNiew** | **[]string** |  | 
 **descriptionNisw** | **[]string** |  | 
 **distance** | **[]float64** |  | 
 **distanceEmpty** | **bool** |  | 
 **distanceGt** | **[]float64** |  | 
 **distanceGte** | **[]float64** |  | 
 **distanceLt** | **[]float64** |  | 
 **distanceLte** | **[]float64** |  | 
 **distanceN** | **[]float64** |  | 
 **distanceUnit** | [**WirelessWirelessLinksListDistanceUnitParameter**](WirelessWirelessLinksListDistanceUnitParameter.md) | * &#x60;km&#x60; - Kilometers * &#x60;m&#x60; - Meters * &#x60;mi&#x60; - Miles * &#x60;ft&#x60; - Feet | 
 **id** | **[]int32** |  | 
 **idEmpty** | **bool** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **interfaceAId** | **[]int32** |  | 
 **interfaceAIdN** | **[]int32** |  | 
 **interfaceBId** | **[]int32** |  | 
 **interfaceBIdN** | **[]int32** |  | 
 **lastUpdated** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUpdatedN** | [**[]time.Time**](time.Time.md) |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **modifiedByRequest** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 
 **ssid** | **[]string** |  | 
 **ssidEmpty** | **bool** |  | 
 **ssidIc** | **[]string** |  | 
 **ssidIe** | **[]string** |  | 
 **ssidIew** | **[]string** |  | 
 **ssidIsw** | **[]string** |  | 
 **ssidN** | **[]string** |  | 
 **ssidNic** | **[]string** |  | 
 **ssidNie** | **[]string** |  | 
 **ssidNiew** | **[]string** |  | 
 **ssidNisw** | **[]string** |  | 
 **status** | **[]string** |  | 
 **statusEmpty** | **bool** |  | 
 **statusIc** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusIe** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusIew** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusIsw** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusN** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusNic** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusNie** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusNiew** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **statusNisw** | [**[]DcimCablesListStatusIcParameterInner**](DcimCablesListStatusIcParameterInner.md) | * &#x60;connected&#x60; - Connected * &#x60;planned&#x60; - Planned * &#x60;decommissioning&#x60; - Decommissioning | 
 **tag** | **[]string** |  | 
 **tagN** | **[]string** |  | 
 **tenant** | **[]string** | Tenant (slug) | 
 **tenantN** | **[]string** | Tenant (slug) | 
 **tenantGroup** | **[]string** |  | 
 **tenantGroupN** | **[]string** |  | 
 **tenantGroupId** | **[]string** |  | 
 **tenantGroupIdN** | **[]string** |  | 
 **tenantId** | **[]int32** | Tenant (ID) | 
 **tenantIdN** | **[]int32** | Tenant (ID) | 
 **updatedByRequest** | **string** |  | 

### Return type

[**PaginatedWirelessLinkList**](PaginatedWirelessLinkList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksPartialUpdate

> WirelessLink WirelessWirelessLinksPartialUpdate(ctx, id).PatchedWritableWirelessLinkRequest(patchedWritableWirelessLinkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless link.
	patchedWritableWirelessLinkRequest := *openapiclient.NewPatchedWritableWirelessLinkRequest() // PatchedWritableWirelessLinkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLinksPartialUpdate(context.Background(), id).PatchedWritableWirelessLinkRequest(patchedWritableWirelessLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLinksPartialUpdate`: WirelessLink
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLinksPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableWirelessLinkRequest** | [**PatchedWritableWirelessLinkRequest**](PatchedWritableWirelessLinkRequest.md) |  | 

### Return type

[**WirelessLink**](WirelessLink.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksRetrieve

> WirelessLink WirelessWirelessLinksRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLinksRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLinksRetrieve`: WirelessLink
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLinksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WirelessLink**](WirelessLink.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WirelessWirelessLinksUpdate

> WirelessLink WirelessWirelessLinksUpdate(ctx, id).WritableWirelessLinkRequest(writableWirelessLinkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this wireless link.
	writableWirelessLinkRequest := *openapiclient.NewWritableWirelessLinkRequest(*openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example"), *openapiclient.NewBriefInterfaceRequest(*openapiclient.NewBriefDeviceRequest(), "Name_example")) // WritableWirelessLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessAPI.WirelessWirelessLinksUpdate(context.Background(), id).WritableWirelessLinkRequest(writableWirelessLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessAPI.WirelessWirelessLinksUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessWirelessLinksUpdate`: WirelessLink
	fmt.Fprintf(os.Stdout, "Response from `WirelessAPI.WirelessWirelessLinksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this wireless link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWirelessWirelessLinksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableWirelessLinkRequest** | [**WritableWirelessLinkRequest**](WritableWirelessLinkRequest.md) |  | 

### Return type

[**WirelessLink**](WirelessLink.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

