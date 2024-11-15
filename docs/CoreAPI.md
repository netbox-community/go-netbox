# \CoreAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreDataFilesList**](CoreAPI.md#CoreDataFilesList) | **Get** /api/core/data-files/ | 
[**CoreDataFilesRetrieve**](CoreAPI.md#CoreDataFilesRetrieve) | **Get** /api/core/data-files/{id}/ | 
[**CoreDataSourcesBulkDestroy**](CoreAPI.md#CoreDataSourcesBulkDestroy) | **Delete** /api/core/data-sources/ | 
[**CoreDataSourcesBulkPartialUpdate**](CoreAPI.md#CoreDataSourcesBulkPartialUpdate) | **Patch** /api/core/data-sources/ | 
[**CoreDataSourcesBulkUpdate**](CoreAPI.md#CoreDataSourcesBulkUpdate) | **Put** /api/core/data-sources/ | 
[**CoreDataSourcesCreate**](CoreAPI.md#CoreDataSourcesCreate) | **Post** /api/core/data-sources/ | 
[**CoreDataSourcesDestroy**](CoreAPI.md#CoreDataSourcesDestroy) | **Delete** /api/core/data-sources/{id}/ | 
[**CoreDataSourcesList**](CoreAPI.md#CoreDataSourcesList) | **Get** /api/core/data-sources/ | 
[**CoreDataSourcesPartialUpdate**](CoreAPI.md#CoreDataSourcesPartialUpdate) | **Patch** /api/core/data-sources/{id}/ | 
[**CoreDataSourcesRetrieve**](CoreAPI.md#CoreDataSourcesRetrieve) | **Get** /api/core/data-sources/{id}/ | 
[**CoreDataSourcesSyncCreate**](CoreAPI.md#CoreDataSourcesSyncCreate) | **Post** /api/core/data-sources/{id}/sync/ | 
[**CoreDataSourcesUpdate**](CoreAPI.md#CoreDataSourcesUpdate) | **Put** /api/core/data-sources/{id}/ | 
[**CoreJobsList**](CoreAPI.md#CoreJobsList) | **Get** /api/core/jobs/ | 
[**CoreJobsRetrieve**](CoreAPI.md#CoreJobsRetrieve) | **Get** /api/core/jobs/{id}/ | 



## CoreDataFilesList

> PaginatedDataFileList CoreDataFilesList(ctx).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Hash(hash).HashEmpty(hashEmpty).HashIc(hashIc).HashIe(hashIe).HashIew(hashIew).HashIsw(hashIsw).HashN(hashN).HashNic(hashNic).HashNie(hashNie).HashNiew(hashNiew).HashNisw(hashNisw).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Offset(offset).Ordering(ordering).Path(path).PathEmpty(pathEmpty).PathIc(pathIc).PathIe(pathIe).PathIew(pathIew).PathIsw(pathIsw).PathN(pathN).PathNic(pathNic).PathNie(pathNie).PathNiew(pathNiew).PathNisw(pathNisw).Q(q).Size(size).SizeEmpty(sizeEmpty).SizeGt(sizeGt).SizeGte(sizeGte).SizeLt(sizeLt).SizeLte(sizeLte).SizeN(sizeN).Source(source).SourceN(sourceN).SourceId(sourceId).SourceIdN(sourceIdN).UpdatedByRequest(updatedByRequest).Execute()





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
	created := []time.Time{time.Now()} // []time.Time |  (optional)
	createdEmpty := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdGte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLt := []time.Time{time.Now()} // []time.Time |  (optional)
	createdLte := []time.Time{time.Now()} // []time.Time |  (optional)
	createdN := []time.Time{time.Now()} // []time.Time |  (optional)
	createdByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	hash := []string{"Inner_example"} // []string |  (optional)
	hashEmpty := true // bool |  (optional)
	hashIc := []string{"Inner_example"} // []string |  (optional)
	hashIe := []string{"Inner_example"} // []string |  (optional)
	hashIew := []string{"Inner_example"} // []string |  (optional)
	hashIsw := []string{"Inner_example"} // []string |  (optional)
	hashN := []string{"Inner_example"} // []string |  (optional)
	hashNic := []string{"Inner_example"} // []string |  (optional)
	hashNie := []string{"Inner_example"} // []string |  (optional)
	hashNiew := []string{"Inner_example"} // []string |  (optional)
	hashNisw := []string{"Inner_example"} // []string |  (optional)
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
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	path := []string{"Inner_example"} // []string |  (optional)
	pathEmpty := true // bool |  (optional)
	pathIc := []string{"Inner_example"} // []string |  (optional)
	pathIe := []string{"Inner_example"} // []string |  (optional)
	pathIew := []string{"Inner_example"} // []string |  (optional)
	pathIsw := []string{"Inner_example"} // []string |  (optional)
	pathN := []string{"Inner_example"} // []string |  (optional)
	pathNic := []string{"Inner_example"} // []string |  (optional)
	pathNie := []string{"Inner_example"} // []string |  (optional)
	pathNiew := []string{"Inner_example"} // []string |  (optional)
	pathNisw := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	size := []int32{int32(123)} // []int32 |  (optional)
	sizeEmpty := true // bool |  (optional)
	sizeGt := []int32{int32(123)} // []int32 |  (optional)
	sizeGte := []int32{int32(123)} // []int32 |  (optional)
	sizeLt := []int32{int32(123)} // []int32 |  (optional)
	sizeLte := []int32{int32(123)} // []int32 |  (optional)
	sizeN := []int32{int32(123)} // []int32 |  (optional)
	source := []string{"Inner_example"} // []string | Data source (name) (optional)
	sourceN := []string{"Inner_example"} // []string | Data source (name) (optional)
	sourceId := []int32{int32(123)} // []int32 | Data source (ID) (optional)
	sourceIdN := []int32{int32(123)} // []int32 | Data source (ID) (optional)
	updatedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataFilesList(context.Background()).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Hash(hash).HashEmpty(hashEmpty).HashIc(hashIc).HashIe(hashIe).HashIew(hashIew).HashIsw(hashIsw).HashN(hashN).HashNic(hashNic).HashNie(hashNie).HashNiew(hashNiew).HashNisw(hashNisw).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Offset(offset).Ordering(ordering).Path(path).PathEmpty(pathEmpty).PathIc(pathIc).PathIe(pathIe).PathIew(pathIew).PathIsw(pathIsw).PathN(pathN).PathNic(pathNic).PathNie(pathNie).PathNiew(pathNiew).PathNisw(pathNisw).Q(q).Size(size).SizeEmpty(sizeEmpty).SizeGt(sizeGt).SizeGte(sizeGte).SizeLt(sizeLt).SizeLte(sizeLte).SizeN(sizeN).Source(source).SourceN(sourceN).SourceId(sourceId).SourceIdN(sourceIdN).UpdatedByRequest(updatedByRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataFilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataFilesList`: PaginatedDataFileList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataFilesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | [**[]time.Time**](time.Time.md) |  | 
 **createdEmpty** | [**[]time.Time**](time.Time.md) |  | 
 **createdGt** | [**[]time.Time**](time.Time.md) |  | 
 **createdGte** | [**[]time.Time**](time.Time.md) |  | 
 **createdLt** | [**[]time.Time**](time.Time.md) |  | 
 **createdLte** | [**[]time.Time**](time.Time.md) |  | 
 **createdN** | [**[]time.Time**](time.Time.md) |  | 
 **createdByRequest** | **string** |  | 
 **hash** | **[]string** |  | 
 **hashEmpty** | **bool** |  | 
 **hashIc** | **[]string** |  | 
 **hashIe** | **[]string** |  | 
 **hashIew** | **[]string** |  | 
 **hashIsw** | **[]string** |  | 
 **hashN** | **[]string** |  | 
 **hashNic** | **[]string** |  | 
 **hashNie** | **[]string** |  | 
 **hashNiew** | **[]string** |  | 
 **hashNisw** | **[]string** |  | 
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
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **path** | **[]string** |  | 
 **pathEmpty** | **bool** |  | 
 **pathIc** | **[]string** |  | 
 **pathIe** | **[]string** |  | 
 **pathIew** | **[]string** |  | 
 **pathIsw** | **[]string** |  | 
 **pathN** | **[]string** |  | 
 **pathNic** | **[]string** |  | 
 **pathNie** | **[]string** |  | 
 **pathNiew** | **[]string** |  | 
 **pathNisw** | **[]string** |  | 
 **q** | **string** |  | 
 **size** | **[]int32** |  | 
 **sizeEmpty** | **bool** |  | 
 **sizeGt** | **[]int32** |  | 
 **sizeGte** | **[]int32** |  | 
 **sizeLt** | **[]int32** |  | 
 **sizeLte** | **[]int32** |  | 
 **sizeN** | **[]int32** |  | 
 **source** | **[]string** | Data source (name) | 
 **sourceN** | **[]string** | Data source (name) | 
 **sourceId** | **[]int32** | Data source (ID) | 
 **sourceIdN** | **[]int32** | Data source (ID) | 
 **updatedByRequest** | **string** |  | 

### Return type

[**PaginatedDataFileList**](PaginatedDataFileList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataFilesRetrieve

> DataFile CoreDataFilesRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this data file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataFilesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataFilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataFilesRetrieve`: DataFile
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataFilesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataFilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataFile**](DataFile.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesBulkDestroy

> CoreDataSourcesBulkDestroy(ctx).DataSourceRequest(dataSourceRequest).Execute()





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
	dataSourceRequest := []openapiclient.DataSourceRequest{*openapiclient.NewDataSourceRequest("Name_example", openapiclient.DataSource_type_value("local"), "SourceUrl_example")} // []DataSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreDataSourcesBulkDestroy(context.Background()).DataSourceRequest(dataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesBulkDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceRequest** | [**[]DataSourceRequest**](DataSourceRequest.md) |  | 

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


## CoreDataSourcesBulkPartialUpdate

> []DataSource CoreDataSourcesBulkPartialUpdate(ctx).DataSourceRequest(dataSourceRequest).Execute()





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
	dataSourceRequest := []openapiclient.DataSourceRequest{*openapiclient.NewDataSourceRequest("Name_example", openapiclient.DataSource_type_value("local"), "SourceUrl_example")} // []DataSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesBulkPartialUpdate(context.Background()).DataSourceRequest(dataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesBulkPartialUpdate`: []DataSource
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceRequest** | [**[]DataSourceRequest**](DataSourceRequest.md) |  | 

### Return type

[**[]DataSource**](DataSource.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesBulkUpdate

> []DataSource CoreDataSourcesBulkUpdate(ctx).DataSourceRequest(dataSourceRequest).Execute()





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
	dataSourceRequest := []openapiclient.DataSourceRequest{*openapiclient.NewDataSourceRequest("Name_example", openapiclient.DataSource_type_value("local"), "SourceUrl_example")} // []DataSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesBulkUpdate(context.Background()).DataSourceRequest(dataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesBulkUpdate`: []DataSource
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceRequest** | [**[]DataSourceRequest**](DataSourceRequest.md) |  | 

### Return type

[**[]DataSource**](DataSource.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesCreate

> DataSource CoreDataSourcesCreate(ctx).WritableDataSourceRequest(writableDataSourceRequest).Execute()





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
	writableDataSourceRequest := *openapiclient.NewWritableDataSourceRequest("Name_example", "Type_example", "SourceUrl_example") // WritableDataSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesCreate(context.Background()).WritableDataSourceRequest(writableDataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesCreate`: DataSource
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableDataSourceRequest** | [**WritableDataSourceRequest**](WritableDataSourceRequest.md) |  | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesDestroy

> CoreDataSourcesDestroy(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this data source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.CoreDataSourcesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesDestroyRequest struct via the builder pattern


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


## CoreDataSourcesList

> PaginatedDataSourceList CoreDataSourcesList(ctx).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Enabled(enabled).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastSynced(lastSynced).LastSyncedEmpty(lastSyncedEmpty).LastSyncedGt(lastSyncedGt).LastSyncedGte(lastSyncedGte).LastSyncedLt(lastSyncedLt).LastSyncedLte(lastSyncedLte).LastSyncedN(lastSyncedN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).Offset(offset).Ordering(ordering).Q(q).SourceUrl(sourceUrl).SourceUrlEmpty(sourceUrlEmpty).SourceUrlIc(sourceUrlIc).SourceUrlIe(sourceUrlIe).SourceUrlIew(sourceUrlIew).SourceUrlIsw(sourceUrlIsw).SourceUrlN(sourceUrlN).SourceUrlNic(sourceUrlNic).SourceUrlNie(sourceUrlNie).SourceUrlNiew(sourceUrlNiew).SourceUrlNisw(sourceUrlNisw).Status(status).StatusN(statusN).Tag(tag).TagN(tagN).Type_(type_).TypeN(typeN).UpdatedByRequest(updatedByRequest).Execute()





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
	enabled := true // bool |  (optional)
	id := []int32{int32(123)} // []int32 |  (optional)
	idEmpty := true // bool |  (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	lastSynced := []time.Time{time.Now()} // []time.Time |  (optional)
	lastSyncedEmpty := true // bool |  (optional)
	lastSyncedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastSyncedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastSyncedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastSyncedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastSyncedN := []time.Time{time.Now()} // []time.Time |  (optional)
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
	q := "q_example" // string | Search (optional)
	sourceUrl := []string{"Inner_example"} // []string |  (optional)
	sourceUrlEmpty := true // bool |  (optional)
	sourceUrlIc := []string{"Inner_example"} // []string |  (optional)
	sourceUrlIe := []string{"Inner_example"} // []string |  (optional)
	sourceUrlIew := []string{"Inner_example"} // []string |  (optional)
	sourceUrlIsw := []string{"Inner_example"} // []string |  (optional)
	sourceUrlN := []string{"Inner_example"} // []string |  (optional)
	sourceUrlNic := []string{"Inner_example"} // []string |  (optional)
	sourceUrlNie := []string{"Inner_example"} // []string |  (optional)
	sourceUrlNiew := []string{"Inner_example"} // []string |  (optional)
	sourceUrlNisw := []string{"Inner_example"} // []string |  (optional)
	status := []string{"Inner_example"} // []string |  (optional)
	statusN := []string{"Inner_example"} // []string |  (optional)
	tag := []string{"Inner_example"} // []string |  (optional)
	tagN := []string{"Inner_example"} // []string |  (optional)
	type_ := []string{"Inner_example"} // []string |  (optional)
	typeN := []string{"Inner_example"} // []string |  (optional)
	updatedByRequest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesList(context.Background()).Created(created).CreatedEmpty(createdEmpty).CreatedGt(createdGt).CreatedGte(createdGte).CreatedLt(createdLt).CreatedLte(createdLte).CreatedN(createdN).CreatedByRequest(createdByRequest).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Enabled(enabled).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).LastSynced(lastSynced).LastSyncedEmpty(lastSyncedEmpty).LastSyncedGt(lastSyncedGt).LastSyncedGte(lastSyncedGte).LastSyncedLt(lastSyncedLt).LastSyncedLte(lastSyncedLte).LastSyncedN(lastSyncedN).LastUpdated(lastUpdated).LastUpdatedEmpty(lastUpdatedEmpty).LastUpdatedGt(lastUpdatedGt).LastUpdatedGte(lastUpdatedGte).LastUpdatedLt(lastUpdatedLt).LastUpdatedLte(lastUpdatedLte).LastUpdatedN(lastUpdatedN).Limit(limit).ModifiedByRequest(modifiedByRequest).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).Offset(offset).Ordering(ordering).Q(q).SourceUrl(sourceUrl).SourceUrlEmpty(sourceUrlEmpty).SourceUrlIc(sourceUrlIc).SourceUrlIe(sourceUrlIe).SourceUrlIew(sourceUrlIew).SourceUrlIsw(sourceUrlIsw).SourceUrlN(sourceUrlN).SourceUrlNic(sourceUrlNic).SourceUrlNie(sourceUrlNie).SourceUrlNiew(sourceUrlNiew).SourceUrlNisw(sourceUrlNisw).Status(status).StatusN(statusN).Tag(tag).TagN(tagN).Type_(type_).TypeN(typeN).UpdatedByRequest(updatedByRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesList`: PaginatedDataSourceList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
 **enabled** | **bool** |  | 
 **id** | **[]int32** |  | 
 **idEmpty** | **bool** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **lastSynced** | [**[]time.Time**](time.Time.md) |  | 
 **lastSyncedEmpty** | **bool** |  | 
 **lastSyncedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastSyncedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastSyncedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastSyncedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastSyncedN** | [**[]time.Time**](time.Time.md) |  | 
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
 **q** | **string** | Search | 
 **sourceUrl** | **[]string** |  | 
 **sourceUrlEmpty** | **bool** |  | 
 **sourceUrlIc** | **[]string** |  | 
 **sourceUrlIe** | **[]string** |  | 
 **sourceUrlIew** | **[]string** |  | 
 **sourceUrlIsw** | **[]string** |  | 
 **sourceUrlN** | **[]string** |  | 
 **sourceUrlNic** | **[]string** |  | 
 **sourceUrlNie** | **[]string** |  | 
 **sourceUrlNiew** | **[]string** |  | 
 **sourceUrlNisw** | **[]string** |  | 
 **status** | **[]string** |  | 
 **statusN** | **[]string** |  | 
 **tag** | **[]string** |  | 
 **tagN** | **[]string** |  | 
 **type_** | **[]string** |  | 
 **typeN** | **[]string** |  | 
 **updatedByRequest** | **string** |  | 

### Return type

[**PaginatedDataSourceList**](PaginatedDataSourceList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesPartialUpdate

> DataSource CoreDataSourcesPartialUpdate(ctx, id).PatchedWritableDataSourceRequest(patchedWritableDataSourceRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this data source.
	patchedWritableDataSourceRequest := *openapiclient.NewPatchedWritableDataSourceRequest() // PatchedWritableDataSourceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesPartialUpdate(context.Background(), id).PatchedWritableDataSourceRequest(patchedWritableDataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesPartialUpdate`: DataSource
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableDataSourceRequest** | [**PatchedWritableDataSourceRequest**](PatchedWritableDataSourceRequest.md) |  | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesRetrieve

> DataSource CoreDataSourcesRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this data source.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesRetrieve`: DataSource
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataSource**](DataSource.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesSyncCreate

> DataSource CoreDataSourcesSyncCreate(ctx, id).WritableDataSourceRequest(writableDataSourceRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this data source.
	writableDataSourceRequest := *openapiclient.NewWritableDataSourceRequest("Name_example", "Type_example", "SourceUrl_example") // WritableDataSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesSyncCreate(context.Background(), id).WritableDataSourceRequest(writableDataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesSyncCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesSyncCreate`: DataSource
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesSyncCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesSyncCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableDataSourceRequest** | [**WritableDataSourceRequest**](WritableDataSourceRequest.md) |  | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreDataSourcesUpdate

> DataSource CoreDataSourcesUpdate(ctx, id).WritableDataSourceRequest(writableDataSourceRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this data source.
	writableDataSourceRequest := *openapiclient.NewWritableDataSourceRequest("Name_example", "Type_example", "SourceUrl_example") // WritableDataSourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreDataSourcesUpdate(context.Background(), id).WritableDataSourceRequest(writableDataSourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreDataSourcesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreDataSourcesUpdate`: DataSource
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreDataSourcesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreDataSourcesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableDataSourceRequest** | [**WritableDataSourceRequest**](WritableDataSourceRequest.md) |  | 

### Return type

[**DataSource**](DataSource.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreJobsList

> PaginatedJobList CoreJobsList(ctx).Completed(completed).CompletedAfter(completedAfter).CompletedBefore(completedBefore).Created(created).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Interval(interval).IntervalEmpty(intervalEmpty).IntervalGt(intervalGt).IntervalGte(intervalGte).IntervalLt(intervalLt).IntervalLte(intervalLte).IntervalN(intervalN).JobId(jobId).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).ObjectId(objectId).ObjectIdEmpty(objectIdEmpty).ObjectIdGt(objectIdGt).ObjectIdGte(objectIdGte).ObjectIdLt(objectIdLt).ObjectIdLte(objectIdLte).ObjectIdN(objectIdN).ObjectType(objectType).ObjectTypeN(objectTypeN).Offset(offset).Ordering(ordering).Q(q).Scheduled(scheduled).ScheduledAfter(scheduledAfter).ScheduledBefore(scheduledBefore).Started(started).StartedAfter(startedAfter).StartedBefore(startedBefore).Status(status).StatusN(statusN).User(user).UserN(userN).Execute()





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
	completed := time.Now() // time.Time |  (optional)
	completedAfter := time.Now() // time.Time |  (optional)
	completedBefore := time.Now() // time.Time |  (optional)
	created := time.Now() // time.Time |  (optional)
	createdAfter := time.Now() // time.Time |  (optional)
	createdBefore := time.Now() // time.Time |  (optional)
	id := []int32{int32(123)} // []int32 |  (optional)
	idEmpty := true // bool |  (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	interval := []int32{int32(123)} // []int32 |  (optional)
	intervalEmpty := true // bool |  (optional)
	intervalGt := []int32{int32(123)} // []int32 |  (optional)
	intervalGte := []int32{int32(123)} // []int32 |  (optional)
	intervalLt := []int32{int32(123)} // []int32 |  (optional)
	intervalLte := []int32{int32(123)} // []int32 |  (optional)
	intervalN := []int32{int32(123)} // []int32 |  (optional)
	jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
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
	objectId := []int32{int32(123)} // []int32 |  (optional)
	objectIdEmpty := true // bool |  (optional)
	objectIdGt := []int32{int32(123)} // []int32 |  (optional)
	objectIdGte := []int32{int32(123)} // []int32 |  (optional)
	objectIdLt := []int32{int32(123)} // []int32 |  (optional)
	objectIdLte := []int32{int32(123)} // []int32 |  (optional)
	objectIdN := []int32{int32(123)} // []int32 |  (optional)
	objectType := int32(56) // int32 |  (optional)
	objectTypeN := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	q := "q_example" // string | Search (optional)
	scheduled := time.Now() // time.Time |  (optional)
	scheduledAfter := time.Now() // time.Time |  (optional)
	scheduledBefore := time.Now() // time.Time |  (optional)
	started := time.Now() // time.Time |  (optional)
	startedAfter := time.Now() // time.Time |  (optional)
	startedBefore := time.Now() // time.Time |  (optional)
	status := []string{"Inner_example"} // []string |  (optional)
	statusN := []string{"Inner_example"} // []string |  (optional)
	user := int32(56) // int32 |  (optional)
	userN := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreJobsList(context.Background()).Completed(completed).CompletedAfter(completedAfter).CompletedBefore(completedBefore).Created(created).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Interval(interval).IntervalEmpty(intervalEmpty).IntervalGt(intervalGt).IntervalGte(intervalGte).IntervalLt(intervalLt).IntervalLte(intervalLte).IntervalN(intervalN).JobId(jobId).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).ObjectId(objectId).ObjectIdEmpty(objectIdEmpty).ObjectIdGt(objectIdGt).ObjectIdGte(objectIdGte).ObjectIdLt(objectIdLt).ObjectIdLte(objectIdLte).ObjectIdN(objectIdN).ObjectType(objectType).ObjectTypeN(objectTypeN).Offset(offset).Ordering(ordering).Q(q).Scheduled(scheduled).ScheduledAfter(scheduledAfter).ScheduledBefore(scheduledBefore).Started(started).StartedAfter(startedAfter).StartedBefore(startedBefore).Status(status).StatusN(statusN).User(user).UserN(userN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreJobsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreJobsList`: PaginatedJobList
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreJobsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreJobsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completed** | **time.Time** |  | 
 **completedAfter** | **time.Time** |  | 
 **completedBefore** | **time.Time** |  | 
 **created** | **time.Time** |  | 
 **createdAfter** | **time.Time** |  | 
 **createdBefore** | **time.Time** |  | 
 **id** | **[]int32** |  | 
 **idEmpty** | **bool** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **interval** | **[]int32** |  | 
 **intervalEmpty** | **bool** |  | 
 **intervalGt** | **[]int32** |  | 
 **intervalGte** | **[]int32** |  | 
 **intervalLt** | **[]int32** |  | 
 **intervalLte** | **[]int32** |  | 
 **intervalN** | **[]int32** |  | 
 **jobId** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
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
 **objectId** | **[]int32** |  | 
 **objectIdEmpty** | **bool** |  | 
 **objectIdGt** | **[]int32** |  | 
 **objectIdGte** | **[]int32** |  | 
 **objectIdLt** | **[]int32** |  | 
 **objectIdLte** | **[]int32** |  | 
 **objectIdN** | **[]int32** |  | 
 **objectType** | **int32** |  | 
 **objectTypeN** | **int32** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 
 **scheduled** | **time.Time** |  | 
 **scheduledAfter** | **time.Time** |  | 
 **scheduledBefore** | **time.Time** |  | 
 **started** | **time.Time** |  | 
 **startedAfter** | **time.Time** |  | 
 **startedBefore** | **time.Time** |  | 
 **status** | **[]string** |  | 
 **statusN** | **[]string** |  | 
 **user** | **int32** |  | 
 **userN** | **int32** |  | 

### Return type

[**PaginatedJobList**](PaginatedJobList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreJobsRetrieve

> Job CoreJobsRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.CoreJobsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.CoreJobsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CoreJobsRetrieve`: Job
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.CoreJobsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreJobsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

