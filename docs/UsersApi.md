# \UsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConfigRetrieve**](UsersApi.md#UsersConfigRetrieve) | **Get** /api/users/config/ | 
[**UsersGroupsBulkDestroy**](UsersApi.md#UsersGroupsBulkDestroy) | **Delete** /api/users/groups/ | 
[**UsersGroupsBulkPartialUpdate**](UsersApi.md#UsersGroupsBulkPartialUpdate) | **Patch** /api/users/groups/ | 
[**UsersGroupsBulkUpdate**](UsersApi.md#UsersGroupsBulkUpdate) | **Put** /api/users/groups/ | 
[**UsersGroupsCreate**](UsersApi.md#UsersGroupsCreate) | **Post** /api/users/groups/ | 
[**UsersGroupsDestroy**](UsersApi.md#UsersGroupsDestroy) | **Delete** /api/users/groups/{id}/ | 
[**UsersGroupsList**](UsersApi.md#UsersGroupsList) | **Get** /api/users/groups/ | 
[**UsersGroupsPartialUpdate**](UsersApi.md#UsersGroupsPartialUpdate) | **Patch** /api/users/groups/{id}/ | 
[**UsersGroupsRetrieve**](UsersApi.md#UsersGroupsRetrieve) | **Get** /api/users/groups/{id}/ | 
[**UsersGroupsUpdate**](UsersApi.md#UsersGroupsUpdate) | **Put** /api/users/groups/{id}/ | 
[**UsersPermissionsBulkDestroy**](UsersApi.md#UsersPermissionsBulkDestroy) | **Delete** /api/users/permissions/ | 
[**UsersPermissionsBulkPartialUpdate**](UsersApi.md#UsersPermissionsBulkPartialUpdate) | **Patch** /api/users/permissions/ | 
[**UsersPermissionsBulkUpdate**](UsersApi.md#UsersPermissionsBulkUpdate) | **Put** /api/users/permissions/ | 
[**UsersPermissionsCreate**](UsersApi.md#UsersPermissionsCreate) | **Post** /api/users/permissions/ | 
[**UsersPermissionsDestroy**](UsersApi.md#UsersPermissionsDestroy) | **Delete** /api/users/permissions/{id}/ | 
[**UsersPermissionsList**](UsersApi.md#UsersPermissionsList) | **Get** /api/users/permissions/ | 
[**UsersPermissionsPartialUpdate**](UsersApi.md#UsersPermissionsPartialUpdate) | **Patch** /api/users/permissions/{id}/ | 
[**UsersPermissionsRetrieve**](UsersApi.md#UsersPermissionsRetrieve) | **Get** /api/users/permissions/{id}/ | 
[**UsersPermissionsUpdate**](UsersApi.md#UsersPermissionsUpdate) | **Put** /api/users/permissions/{id}/ | 
[**UsersTokensBulkDestroy**](UsersApi.md#UsersTokensBulkDestroy) | **Delete** /api/users/tokens/ | 
[**UsersTokensBulkPartialUpdate**](UsersApi.md#UsersTokensBulkPartialUpdate) | **Patch** /api/users/tokens/ | 
[**UsersTokensBulkUpdate**](UsersApi.md#UsersTokensBulkUpdate) | **Put** /api/users/tokens/ | 
[**UsersTokensCreate**](UsersApi.md#UsersTokensCreate) | **Post** /api/users/tokens/ | 
[**UsersTokensDestroy**](UsersApi.md#UsersTokensDestroy) | **Delete** /api/users/tokens/{id}/ | 
[**UsersTokensList**](UsersApi.md#UsersTokensList) | **Get** /api/users/tokens/ | 
[**UsersTokensPartialUpdate**](UsersApi.md#UsersTokensPartialUpdate) | **Patch** /api/users/tokens/{id}/ | 
[**UsersTokensProvisionCreate**](UsersApi.md#UsersTokensProvisionCreate) | **Post** /api/users/tokens/provision/ | 
[**UsersTokensRetrieve**](UsersApi.md#UsersTokensRetrieve) | **Get** /api/users/tokens/{id}/ | 
[**UsersTokensUpdate**](UsersApi.md#UsersTokensUpdate) | **Put** /api/users/tokens/{id}/ | 
[**UsersUsersBulkDestroy**](UsersApi.md#UsersUsersBulkDestroy) | **Delete** /api/users/users/ | 
[**UsersUsersBulkPartialUpdate**](UsersApi.md#UsersUsersBulkPartialUpdate) | **Patch** /api/users/users/ | 
[**UsersUsersBulkUpdate**](UsersApi.md#UsersUsersBulkUpdate) | **Put** /api/users/users/ | 
[**UsersUsersCreate**](UsersApi.md#UsersUsersCreate) | **Post** /api/users/users/ | 
[**UsersUsersDestroy**](UsersApi.md#UsersUsersDestroy) | **Delete** /api/users/users/{id}/ | 
[**UsersUsersList**](UsersApi.md#UsersUsersList) | **Get** /api/users/users/ | 
[**UsersUsersPartialUpdate**](UsersApi.md#UsersUsersPartialUpdate) | **Patch** /api/users/users/{id}/ | 
[**UsersUsersRetrieve**](UsersApi.md#UsersUsersRetrieve) | **Get** /api/users/users/{id}/ | 
[**UsersUsersUpdate**](UsersApi.md#UsersUsersUpdate) | **Put** /api/users/users/{id}/ | 



## UsersConfigRetrieve

> map[string]interface{} UsersConfigRetrieve(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersConfigRetrieve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersConfigRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersConfigRetrieve`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersConfigRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersConfigRetrieveRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkDestroy

> UsersGroupsBulkDestroy(ctx).GroupRequest(groupRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    groupRequest := []openapiclient.GroupRequest{*openapiclient.NewGroupRequest("Name_example")} // []GroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersGroupsBulkDestroy(context.Background()).GroupRequest(groupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRequest** | [**[]GroupRequest**](GroupRequest.md) |  | 

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


## UsersGroupsBulkPartialUpdate

> []Group UsersGroupsBulkPartialUpdate(ctx).GroupRequest(groupRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    groupRequest := []openapiclient.GroupRequest{*openapiclient.NewGroupRequest("Name_example")} // []GroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkPartialUpdate(context.Background()).GroupRequest(groupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsBulkPartialUpdate`: []Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRequest** | [**[]GroupRequest**](GroupRequest.md) |  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsBulkUpdate

> []Group UsersGroupsBulkUpdate(ctx).GroupRequest(groupRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    groupRequest := []openapiclient.GroupRequest{*openapiclient.NewGroupRequest("Name_example")} // []GroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsBulkUpdate(context.Background()).GroupRequest(groupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsBulkUpdate`: []Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRequest** | [**[]GroupRequest**](GroupRequest.md) |  | 

### Return type

[**[]Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsCreate

> Group UsersGroupsCreate(ctx).GroupRequest(groupRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsCreate(context.Background()).GroupRequest(groupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsCreate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRequest** | [**GroupRequest**](GroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsDestroy

> UsersGroupsDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersGroupsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsDestroyRequest struct via the builder pattern


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


## UsersGroupsList

> PaginatedGroupList UsersGroupsList(ctx).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).Offset(offset).Ordering(ordering).Q(q).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := []int32{int32(123)} // []int32 |  (optional)
    idGt := []int32{int32(123)} // []int32 |  (optional)
    idGte := []int32{int32(123)} // []int32 |  (optional)
    idLt := []int32{int32(123)} // []int32 |  (optional)
    idLte := []int32{int32(123)} // []int32 |  (optional)
    idN := []int32{int32(123)} // []int32 |  (optional)
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
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    q := "q_example" // string | Search (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsList(context.Background()).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).Offset(offset).Ordering(ordering).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsList`: PaginatedGroupList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]int32** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
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
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 

### Return type

[**PaginatedGroupList**](PaginatedGroupList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsPartialUpdate

> Group UsersGroupsPartialUpdate(ctx, id).PatchedGroupRequest(patchedGroupRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.
    patchedGroupRequest := *openapiclient.NewPatchedGroupRequest() // PatchedGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsPartialUpdate(context.Background(), id).PatchedGroupRequest(patchedGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsPartialUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedGroupRequest** | [**PatchedGroupRequest**](PatchedGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsRetrieve

> Group UsersGroupsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsRetrieve`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGroupsUpdate

> Group UsersGroupsUpdate(ctx, id).GroupRequest(groupRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this group.
    groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersGroupsUpdate(context.Background(), id).GroupRequest(groupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGroupsUpdate`: Group
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRequest** | [**GroupRequest**](GroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkDestroy

> UsersPermissionsBulkDestroy(ctx).ObjectPermissionRequest(objectPermissionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    objectPermissionRequest := []openapiclient.ObjectPermissionRequest{*openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"})} // []ObjectPermissionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersPermissionsBulkDestroy(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectPermissionRequest** | [**[]ObjectPermissionRequest**](ObjectPermissionRequest.md) |  | 

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


## UsersPermissionsBulkPartialUpdate

> []ObjectPermission UsersPermissionsBulkPartialUpdate(ctx).ObjectPermissionRequest(objectPermissionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    objectPermissionRequest := []openapiclient.ObjectPermissionRequest{*openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"})} // []ObjectPermissionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkPartialUpdate(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsBulkPartialUpdate`: []ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectPermissionRequest** | [**[]ObjectPermissionRequest**](ObjectPermissionRequest.md) |  | 

### Return type

[**[]ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsBulkUpdate

> []ObjectPermission UsersPermissionsBulkUpdate(ctx).ObjectPermissionRequest(objectPermissionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    objectPermissionRequest := []openapiclient.ObjectPermissionRequest{*openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"})} // []ObjectPermissionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsBulkUpdate(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsBulkUpdate`: []ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectPermissionRequest** | [**[]ObjectPermissionRequest**](ObjectPermissionRequest.md) |  | 

### Return type

[**[]ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsCreate

> ObjectPermission UsersPermissionsCreate(ctx).WritableObjectPermissionRequest(writableObjectPermissionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    writableObjectPermissionRequest := *openapiclient.NewWritableObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // WritableObjectPermissionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsCreate(context.Background()).WritableObjectPermissionRequest(writableObjectPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsCreate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableObjectPermissionRequest** | [**WritableObjectPermissionRequest**](WritableObjectPermissionRequest.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsDestroy

> UsersPermissionsDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersPermissionsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsDestroyRequest struct via the builder pattern


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


## UsersPermissionsList

> PaginatedObjectPermissionList UsersPermissionsList(ctx).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Enabled(enabled).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
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
    group := []string{"Inner_example"} // []string | Group (name) (optional)
    groupN := []string{"Inner_example"} // []string | Group (name) (optional)
    groupId := []int32{int32(123)} // []int32 | Group (optional)
    groupIdN := []int32{int32(123)} // []int32 | Group (optional)
    id := []int32{int32(123)} // []int32 |  (optional)
    idGt := []int32{int32(123)} // []int32 |  (optional)
    idGte := []int32{int32(123)} // []int32 |  (optional)
    idLt := []int32{int32(123)} // []int32 |  (optional)
    idLte := []int32{int32(123)} // []int32 |  (optional)
    idN := []int32{int32(123)} // []int32 |  (optional)
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
    objectTypes := []int32{int32(123)} // []int32 |  (optional)
    objectTypesN := []int32{int32(123)} // []int32 |  (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    q := "q_example" // string | Search (optional)
    user := []string{"Inner_example"} // []string | User (name) (optional)
    userN := []string{"Inner_example"} // []string | User (name) (optional)
    userId := []int32{int32(123)} // []int32 | User (optional)
    userIdN := []int32{int32(123)} // []int32 | User (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsList(context.Background()).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Enabled(enabled).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsList`: PaginatedObjectPermissionList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
 **group** | **[]string** | Group (name) | 
 **groupN** | **[]string** | Group (name) | 
 **groupId** | **[]int32** | Group | 
 **groupIdN** | **[]int32** | Group | 
 **id** | **[]int32** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
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
 **objectTypes** | **[]int32** |  | 
 **objectTypesN** | **[]int32** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 
 **user** | **[]string** | User (name) | 
 **userN** | **[]string** | User (name) | 
 **userId** | **[]int32** | User | 
 **userIdN** | **[]int32** | User | 

### Return type

[**PaginatedObjectPermissionList**](PaginatedObjectPermissionList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsPartialUpdate

> ObjectPermission UsersPermissionsPartialUpdate(ctx, id).PatchedWritableObjectPermissionRequest(patchedWritableObjectPermissionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.
    patchedWritableObjectPermissionRequest := *openapiclient.NewPatchedWritableObjectPermissionRequest() // PatchedWritableObjectPermissionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsPartialUpdate(context.Background(), id).PatchedWritableObjectPermissionRequest(patchedWritableObjectPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsPartialUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableObjectPermissionRequest** | [**PatchedWritableObjectPermissionRequest**](PatchedWritableObjectPermissionRequest.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsRetrieve

> ObjectPermission UsersPermissionsRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsRetrieve`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPermissionsUpdate

> ObjectPermission UsersPermissionsUpdate(ctx, id).WritableObjectPermissionRequest(writableObjectPermissionRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this permission.
    writableObjectPermissionRequest := *openapiclient.NewWritableObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // WritableObjectPermissionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersPermissionsUpdate(context.Background(), id).WritableObjectPermissionRequest(writableObjectPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPermissionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPermissionsUpdate`: ObjectPermission
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPermissionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableObjectPermissionRequest** | [**WritableObjectPermissionRequest**](WritableObjectPermissionRequest.md) |  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkDestroy

> UsersTokensBulkDestroy(ctx).TokenRequest(tokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    tokenRequest := []openapiclient.TokenRequest{*openapiclient.NewTokenRequest(*openapiclient.NewNestedUserRequest("Username_example"))} // []TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersTokensBulkDestroy(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**[]TokenRequest**](TokenRequest.md) |  | 

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


## UsersTokensBulkPartialUpdate

> []Token UsersTokensBulkPartialUpdate(ctx).TokenRequest(tokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    tokenRequest := []openapiclient.TokenRequest{*openapiclient.NewTokenRequest(*openapiclient.NewNestedUserRequest("Username_example"))} // []TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensBulkPartialUpdate(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensBulkPartialUpdate`: []Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**[]TokenRequest**](TokenRequest.md) |  | 

### Return type

[**[]Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensBulkUpdate

> []Token UsersTokensBulkUpdate(ctx).TokenRequest(tokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    tokenRequest := []openapiclient.TokenRequest{*openapiclient.NewTokenRequest(*openapiclient.NewNestedUserRequest("Username_example"))} // []TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensBulkUpdate(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensBulkUpdate`: []Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**[]TokenRequest**](TokenRequest.md) |  | 

### Return type

[**[]Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensCreate

> Token UsersTokensCreate(ctx).WritableTokenRequest(writableTokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    writableTokenRequest := *openapiclient.NewWritableTokenRequest(int32(123)) // WritableTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensCreate(context.Background()).WritableTokenRequest(writableTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensCreate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableTokenRequest** | [**WritableTokenRequest**](WritableTokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensDestroy

> UsersTokensDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersTokensDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensDestroyRequest struct via the builder pattern


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


## UsersTokensList

> PaginatedTokenList UsersTokensList(ctx).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Expires(expires).ExpiresGte(expiresGte).ExpiresLte(expiresLte).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Key(key).KeyEmpty(keyEmpty).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNisw(keyNisw).Limit(limit).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).WriteEnabled(writeEnabled).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    created := time.Now() // time.Time |  (optional)
    createdGte := time.Now() // time.Time |  (optional)
    createdLte := time.Now() // time.Time |  (optional)
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
    expires := time.Now() // time.Time |  (optional)
    expiresGte := time.Now() // time.Time |  (optional)
    expiresLte := time.Now() // time.Time |  (optional)
    id := []int32{int32(123)} // []int32 |  (optional)
    idGt := []int32{int32(123)} // []int32 |  (optional)
    idGte := []int32{int32(123)} // []int32 |  (optional)
    idLt := []int32{int32(123)} // []int32 |  (optional)
    idLte := []int32{int32(123)} // []int32 |  (optional)
    idN := []int32{int32(123)} // []int32 |  (optional)
    key := []string{"Inner_example"} // []string |  (optional)
    keyEmpty := true // bool |  (optional)
    keyIc := []string{"Inner_example"} // []string |  (optional)
    keyIe := []string{"Inner_example"} // []string |  (optional)
    keyIew := []string{"Inner_example"} // []string |  (optional)
    keyIsw := []string{"Inner_example"} // []string |  (optional)
    keyN := []string{"Inner_example"} // []string |  (optional)
    keyNic := []string{"Inner_example"} // []string |  (optional)
    keyNie := []string{"Inner_example"} // []string |  (optional)
    keyNiew := []string{"Inner_example"} // []string |  (optional)
    keyNisw := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    q := "q_example" // string | Search (optional)
    user := []string{"Inner_example"} // []string | User (name) (optional)
    userN := []string{"Inner_example"} // []string | User (name) (optional)
    userId := []int32{int32(123)} // []int32 | User (optional)
    userIdN := []int32{int32(123)} // []int32 | User (optional)
    writeEnabled := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensList(context.Background()).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Expires(expires).ExpiresGte(expiresGte).ExpiresLte(expiresLte).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Key(key).KeyEmpty(keyEmpty).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNisw(keyNisw).Limit(limit).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).WriteEnabled(writeEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensList`: PaginatedTokenList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **created** | **time.Time** |  | 
 **createdGte** | **time.Time** |  | 
 **createdLte** | **time.Time** |  | 
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
 **expires** | **time.Time** |  | 
 **expiresGte** | **time.Time** |  | 
 **expiresLte** | **time.Time** |  | 
 **id** | **[]int32** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **key** | **[]string** |  | 
 **keyEmpty** | **bool** |  | 
 **keyIc** | **[]string** |  | 
 **keyIe** | **[]string** |  | 
 **keyIew** | **[]string** |  | 
 **keyIsw** | **[]string** |  | 
 **keyN** | **[]string** |  | 
 **keyNic** | **[]string** |  | 
 **keyNie** | **[]string** |  | 
 **keyNiew** | **[]string** |  | 
 **keyNisw** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 
 **user** | **[]string** | User (name) | 
 **userN** | **[]string** | User (name) | 
 **userId** | **[]int32** | User | 
 **userIdN** | **[]int32** | User | 
 **writeEnabled** | **bool** |  | 

### Return type

[**PaginatedTokenList**](PaginatedTokenList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensPartialUpdate

> Token UsersTokensPartialUpdate(ctx, id).PatchedWritableTokenRequest(patchedWritableTokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this token.
    patchedWritableTokenRequest := *openapiclient.NewPatchedWritableTokenRequest() // PatchedWritableTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensPartialUpdate(context.Background(), id).PatchedWritableTokenRequest(patchedWritableTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensPartialUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableTokenRequest** | [**PatchedWritableTokenRequest**](PatchedWritableTokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensProvisionCreate

> Token UsersTokensProvisionCreate(ctx).WritableTokenRequest(writableTokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    writableTokenRequest := *openapiclient.NewWritableTokenRequest(int32(123)) // WritableTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensProvisionCreate(context.Background()).WritableTokenRequest(writableTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensProvisionCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensProvisionCreate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensProvisionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensProvisionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableTokenRequest** | [**WritableTokenRequest**](WritableTokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensRetrieve

> Token UsersTokensRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensRetrieve`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersTokensUpdate

> Token UsersTokensUpdate(ctx, id).WritableTokenRequest(writableTokenRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this token.
    writableTokenRequest := *openapiclient.NewWritableTokenRequest(int32(123)) // WritableTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersTokensUpdate(context.Background(), id).WritableTokenRequest(writableTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersTokensUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersTokensUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableTokenRequest** | [**WritableTokenRequest**](WritableTokenRequest.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkDestroy

> UsersUsersBulkDestroy(ctx).UserRequest(userRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    userRequest := []openapiclient.UserRequest{*openapiclient.NewUserRequest("Username_example", "Password_example")} // []UserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersUsersBulkDestroy(context.Background()).UserRequest(userRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRequest** | [**[]UserRequest**](UserRequest.md) |  | 

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


## UsersUsersBulkPartialUpdate

> []User UsersUsersBulkPartialUpdate(ctx).UserRequest(userRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    userRequest := []openapiclient.UserRequest{*openapiclient.NewUserRequest("Username_example", "Password_example")} // []UserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersBulkPartialUpdate(context.Background()).UserRequest(userRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersBulkPartialUpdate`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersBulkPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRequest** | [**[]UserRequest**](UserRequest.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersBulkUpdate

> []User UsersUsersBulkUpdate(ctx).UserRequest(userRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    userRequest := []openapiclient.UserRequest{*openapiclient.NewUserRequest("Username_example", "Password_example")} // []UserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersBulkUpdate(context.Background()).UserRequest(userRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersBulkUpdate`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRequest** | [**[]UserRequest**](UserRequest.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersCreate

> User UsersUsersCreate(ctx).WritableUserRequest(writableUserRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    writableUserRequest := *openapiclient.NewWritableUserRequest("Username_example", "Password_example") // WritableUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersCreate(context.Background()).WritableUserRequest(writableUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersCreate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writableUserRequest** | [**WritableUserRequest**](WritableUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersDestroy

> UsersUsersDestroy(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersApi.UsersUsersDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersDestroyRequest struct via the builder pattern


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


## UsersUsersList

> PaginatedUserList UsersUsersList(ctx).Email(email).EmailEmpty(emailEmpty).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNisw(emailNisw).FirstName(firstName).FirstNameEmpty(firstNameEmpty).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNisw(firstNameNisw).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameEmpty(lastNameEmpty).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNisw(lastNameNisw).Limit(limit).Offset(offset).Ordering(ordering).Q(q).Username(username).UsernameEmpty(usernameEmpty).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNisw(usernameNisw).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    email := []string{"Inner_example"} // []string |  (optional)
    emailEmpty := true // bool |  (optional)
    emailIc := []string{"Inner_example"} // []string |  (optional)
    emailIe := []string{"Inner_example"} // []string |  (optional)
    emailIew := []string{"Inner_example"} // []string |  (optional)
    emailIsw := []string{"Inner_example"} // []string |  (optional)
    emailN := []string{"Inner_example"} // []string |  (optional)
    emailNic := []string{"Inner_example"} // []string |  (optional)
    emailNie := []string{"Inner_example"} // []string |  (optional)
    emailNiew := []string{"Inner_example"} // []string |  (optional)
    emailNisw := []string{"Inner_example"} // []string |  (optional)
    firstName := []string{"Inner_example"} // []string |  (optional)
    firstNameEmpty := true // bool |  (optional)
    firstNameIc := []string{"Inner_example"} // []string |  (optional)
    firstNameIe := []string{"Inner_example"} // []string |  (optional)
    firstNameIew := []string{"Inner_example"} // []string |  (optional)
    firstNameIsw := []string{"Inner_example"} // []string |  (optional)
    firstNameN := []string{"Inner_example"} // []string |  (optional)
    firstNameNic := []string{"Inner_example"} // []string |  (optional)
    firstNameNie := []string{"Inner_example"} // []string |  (optional)
    firstNameNiew := []string{"Inner_example"} // []string |  (optional)
    firstNameNisw := []string{"Inner_example"} // []string |  (optional)
    group := []string{"Inner_example"} // []string | Group (name) (optional)
    groupN := []string{"Inner_example"} // []string | Group (name) (optional)
    groupId := []int32{int32(123)} // []int32 | Group (optional)
    groupIdN := []int32{int32(123)} // []int32 | Group (optional)
    id := []int32{int32(123)} // []int32 |  (optional)
    idGt := []int32{int32(123)} // []int32 |  (optional)
    idGte := []int32{int32(123)} // []int32 |  (optional)
    idLt := []int32{int32(123)} // []int32 |  (optional)
    idLte := []int32{int32(123)} // []int32 |  (optional)
    idN := []int32{int32(123)} // []int32 |  (optional)
    isActive := true // bool |  (optional)
    isStaff := true // bool |  (optional)
    lastName := []string{"Inner_example"} // []string |  (optional)
    lastNameEmpty := true // bool |  (optional)
    lastNameIc := []string{"Inner_example"} // []string |  (optional)
    lastNameIe := []string{"Inner_example"} // []string |  (optional)
    lastNameIew := []string{"Inner_example"} // []string |  (optional)
    lastNameIsw := []string{"Inner_example"} // []string |  (optional)
    lastNameN := []string{"Inner_example"} // []string |  (optional)
    lastNameNic := []string{"Inner_example"} // []string |  (optional)
    lastNameNie := []string{"Inner_example"} // []string |  (optional)
    lastNameNiew := []string{"Inner_example"} // []string |  (optional)
    lastNameNisw := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    q := "q_example" // string | Search (optional)
    username := []string{"Inner_example"} // []string |  (optional)
    usernameEmpty := true // bool |  (optional)
    usernameIc := []string{"Inner_example"} // []string |  (optional)
    usernameIe := []string{"Inner_example"} // []string |  (optional)
    usernameIew := []string{"Inner_example"} // []string |  (optional)
    usernameIsw := []string{"Inner_example"} // []string |  (optional)
    usernameN := []string{"Inner_example"} // []string |  (optional)
    usernameNic := []string{"Inner_example"} // []string |  (optional)
    usernameNie := []string{"Inner_example"} // []string |  (optional)
    usernameNiew := []string{"Inner_example"} // []string |  (optional)
    usernameNisw := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersList(context.Background()).Email(email).EmailEmpty(emailEmpty).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNisw(emailNisw).FirstName(firstName).FirstNameEmpty(firstNameEmpty).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNisw(firstNameNisw).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).IsActive(isActive).IsStaff(isStaff).LastName(lastName).LastNameEmpty(lastNameEmpty).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNisw(lastNameNisw).Limit(limit).Offset(offset).Ordering(ordering).Q(q).Username(username).UsernameEmpty(usernameEmpty).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNisw(usernameNisw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersList`: PaginatedUserList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **[]string** |  | 
 **emailEmpty** | **bool** |  | 
 **emailIc** | **[]string** |  | 
 **emailIe** | **[]string** |  | 
 **emailIew** | **[]string** |  | 
 **emailIsw** | **[]string** |  | 
 **emailN** | **[]string** |  | 
 **emailNic** | **[]string** |  | 
 **emailNie** | **[]string** |  | 
 **emailNiew** | **[]string** |  | 
 **emailNisw** | **[]string** |  | 
 **firstName** | **[]string** |  | 
 **firstNameEmpty** | **bool** |  | 
 **firstNameIc** | **[]string** |  | 
 **firstNameIe** | **[]string** |  | 
 **firstNameIew** | **[]string** |  | 
 **firstNameIsw** | **[]string** |  | 
 **firstNameN** | **[]string** |  | 
 **firstNameNic** | **[]string** |  | 
 **firstNameNie** | **[]string** |  | 
 **firstNameNiew** | **[]string** |  | 
 **firstNameNisw** | **[]string** |  | 
 **group** | **[]string** | Group (name) | 
 **groupN** | **[]string** | Group (name) | 
 **groupId** | **[]int32** | Group | 
 **groupIdN** | **[]int32** | Group | 
 **id** | **[]int32** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **isActive** | **bool** |  | 
 **isStaff** | **bool** |  | 
 **lastName** | **[]string** |  | 
 **lastNameEmpty** | **bool** |  | 
 **lastNameIc** | **[]string** |  | 
 **lastNameIe** | **[]string** |  | 
 **lastNameIew** | **[]string** |  | 
 **lastNameIsw** | **[]string** |  | 
 **lastNameN** | **[]string** |  | 
 **lastNameNic** | **[]string** |  | 
 **lastNameNie** | **[]string** |  | 
 **lastNameNiew** | **[]string** |  | 
 **lastNameNisw** | **[]string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **q** | **string** | Search | 
 **username** | **[]string** |  | 
 **usernameEmpty** | **bool** |  | 
 **usernameIc** | **[]string** |  | 
 **usernameIe** | **[]string** |  | 
 **usernameIew** | **[]string** |  | 
 **usernameIsw** | **[]string** |  | 
 **usernameN** | **[]string** |  | 
 **usernameNic** | **[]string** |  | 
 **usernameNie** | **[]string** |  | 
 **usernameNiew** | **[]string** |  | 
 **usernameNisw** | **[]string** |  | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersPartialUpdate

> User UsersUsersPartialUpdate(ctx, id).PatchedWritableUserRequest(patchedWritableUserRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.
    patchedWritableUserRequest := *openapiclient.NewPatchedWritableUserRequest() // PatchedWritableUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersPartialUpdate(context.Background(), id).PatchedWritableUserRequest(patchedWritableUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersPartialUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedWritableUserRequest** | [**PatchedWritableUserRequest**](PatchedWritableUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersRetrieve

> User UsersUsersRetrieve(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersRetrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersRetrieve`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsersUpdate

> User UsersUsersUpdate(ctx, id).WritableUserRequest(writableUserRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/netbox-community/go-netbox/v3"
)

func main() {
    id := int32(56) // int32 | A unique integer value identifying this user.
    writableUserRequest := *openapiclient.NewWritableUserRequest("Username_example", "Password_example") // WritableUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UsersUsersUpdate(context.Background(), id).WritableUserRequest(writableUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersUsersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUsersUpdate`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersUsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writableUserRequest** | [**WritableUserRequest**](WritableUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

