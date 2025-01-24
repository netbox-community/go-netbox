# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConfigRetrieve**](UsersAPI.md#UsersConfigRetrieve) | **Get** /api/users/config/ | 
[**UsersGroupsBulkDestroy**](UsersAPI.md#UsersGroupsBulkDestroy) | **Delete** /api/users/groups/ | 
[**UsersGroupsBulkPartialUpdate**](UsersAPI.md#UsersGroupsBulkPartialUpdate) | **Patch** /api/users/groups/ | 
[**UsersGroupsBulkUpdate**](UsersAPI.md#UsersGroupsBulkUpdate) | **Put** /api/users/groups/ | 
[**UsersGroupsCreate**](UsersAPI.md#UsersGroupsCreate) | **Post** /api/users/groups/ | 
[**UsersGroupsDestroy**](UsersAPI.md#UsersGroupsDestroy) | **Delete** /api/users/groups/{id}/ | 
[**UsersGroupsList**](UsersAPI.md#UsersGroupsList) | **Get** /api/users/groups/ | 
[**UsersGroupsPartialUpdate**](UsersAPI.md#UsersGroupsPartialUpdate) | **Patch** /api/users/groups/{id}/ | 
[**UsersGroupsRetrieve**](UsersAPI.md#UsersGroupsRetrieve) | **Get** /api/users/groups/{id}/ | 
[**UsersGroupsUpdate**](UsersAPI.md#UsersGroupsUpdate) | **Put** /api/users/groups/{id}/ | 
[**UsersPermissionsBulkDestroy**](UsersAPI.md#UsersPermissionsBulkDestroy) | **Delete** /api/users/permissions/ | 
[**UsersPermissionsBulkPartialUpdate**](UsersAPI.md#UsersPermissionsBulkPartialUpdate) | **Patch** /api/users/permissions/ | 
[**UsersPermissionsBulkUpdate**](UsersAPI.md#UsersPermissionsBulkUpdate) | **Put** /api/users/permissions/ | 
[**UsersPermissionsCreate**](UsersAPI.md#UsersPermissionsCreate) | **Post** /api/users/permissions/ | 
[**UsersPermissionsDestroy**](UsersAPI.md#UsersPermissionsDestroy) | **Delete** /api/users/permissions/{id}/ | 
[**UsersPermissionsList**](UsersAPI.md#UsersPermissionsList) | **Get** /api/users/permissions/ | 
[**UsersPermissionsPartialUpdate**](UsersAPI.md#UsersPermissionsPartialUpdate) | **Patch** /api/users/permissions/{id}/ | 
[**UsersPermissionsRetrieve**](UsersAPI.md#UsersPermissionsRetrieve) | **Get** /api/users/permissions/{id}/ | 
[**UsersPermissionsUpdate**](UsersAPI.md#UsersPermissionsUpdate) | **Put** /api/users/permissions/{id}/ | 
[**UsersTokensBulkDestroy**](UsersAPI.md#UsersTokensBulkDestroy) | **Delete** /api/users/tokens/ | 
[**UsersTokensBulkPartialUpdate**](UsersAPI.md#UsersTokensBulkPartialUpdate) | **Patch** /api/users/tokens/ | 
[**UsersTokensBulkUpdate**](UsersAPI.md#UsersTokensBulkUpdate) | **Put** /api/users/tokens/ | 
[**UsersTokensCreate**](UsersAPI.md#UsersTokensCreate) | **Post** /api/users/tokens/ | 
[**UsersTokensDestroy**](UsersAPI.md#UsersTokensDestroy) | **Delete** /api/users/tokens/{id}/ | 
[**UsersTokensList**](UsersAPI.md#UsersTokensList) | **Get** /api/users/tokens/ | 
[**UsersTokensPartialUpdate**](UsersAPI.md#UsersTokensPartialUpdate) | **Patch** /api/users/tokens/{id}/ | 
[**UsersTokensProvisionCreate**](UsersAPI.md#UsersTokensProvisionCreate) | **Post** /api/users/tokens/provision/ | 
[**UsersTokensRetrieve**](UsersAPI.md#UsersTokensRetrieve) | **Get** /api/users/tokens/{id}/ | 
[**UsersTokensUpdate**](UsersAPI.md#UsersTokensUpdate) | **Put** /api/users/tokens/{id}/ | 
[**UsersUsersBulkDestroy**](UsersAPI.md#UsersUsersBulkDestroy) | **Delete** /api/users/users/ | 
[**UsersUsersBulkPartialUpdate**](UsersAPI.md#UsersUsersBulkPartialUpdate) | **Patch** /api/users/users/ | 
[**UsersUsersBulkUpdate**](UsersAPI.md#UsersUsersBulkUpdate) | **Put** /api/users/users/ | 
[**UsersUsersCreate**](UsersAPI.md#UsersUsersCreate) | **Post** /api/users/users/ | 
[**UsersUsersDestroy**](UsersAPI.md#UsersUsersDestroy) | **Delete** /api/users/users/{id}/ | 
[**UsersUsersList**](UsersAPI.md#UsersUsersList) | **Get** /api/users/users/ | 
[**UsersUsersPartialUpdate**](UsersAPI.md#UsersUsersPartialUpdate) | **Patch** /api/users/users/{id}/ | 
[**UsersUsersRetrieve**](UsersAPI.md#UsersUsersRetrieve) | **Get** /api/users/users/{id}/ | 
[**UsersUsersUpdate**](UsersAPI.md#UsersUsersUpdate) | **Put** /api/users/users/{id}/ | 



## UsersConfigRetrieve

> map[string]interface{} UsersConfigRetrieve(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersConfigRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersConfigRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersConfigRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersConfigRetrieve`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	groupRequest := []openapiclient.GroupRequest{*openapiclient.NewGroupRequest("Name_example")} // []GroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersGroupsBulkDestroy(context.Background()).GroupRequest(groupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsBulkDestroy``: %v\n", err)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	groupRequest := []openapiclient.GroupRequest{*openapiclient.NewGroupRequest("Name_example")} // []GroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsBulkPartialUpdate(context.Background()).GroupRequest(groupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsBulkPartialUpdate`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsBulkPartialUpdate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	groupRequest := []openapiclient.GroupRequest{*openapiclient.NewGroupRequest("Name_example")} // []GroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsBulkUpdate(context.Background()).GroupRequest(groupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsBulkUpdate`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsBulkUpdate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsCreate(context.Background()).GroupRequest(groupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsCreate`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsCreate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersGroupsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsDestroy``: %v\n", err)
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

> PaginatedGroupList UsersGroupsList(ctx).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).NotificationGroupId(notificationGroupId).NotificationGroupIdN(notificationGroupIdN).Offset(offset).Ordering(ordering).PermissionId(permissionId).PermissionIdN(permissionIdN).Q(q).UserId(userId).UserIdN(userIdN).Execute()





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
	notificationGroupId := []int32{int32(123)} // []int32 | Notification group (ID) (optional)
	notificationGroupIdN := []int32{int32(123)} // []int32 | Notification group (ID) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	permissionId := []int32{int32(123)} // []int32 | Permission (ID) (optional)
	permissionIdN := []int32{int32(123)} // []int32 | Permission (ID) (optional)
	q := "q_example" // string | Search (optional)
	userId := []int32{int32(123)} // []int32 | User (ID) (optional)
	userIdN := []int32{int32(123)} // []int32 | User (ID) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsList(context.Background()).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).NotificationGroupId(notificationGroupId).NotificationGroupIdN(notificationGroupIdN).Offset(offset).Ordering(ordering).PermissionId(permissionId).PermissionIdN(permissionIdN).Q(q).UserId(userId).UserIdN(userIdN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsList`: PaginatedGroupList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGroupsListRequest struct via the builder pattern


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
 **id** | **[]int32** |  | 
 **idEmpty** | **bool** |  | 
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
 **notificationGroupId** | **[]int32** | Notification group (ID) | 
 **notificationGroupIdN** | **[]int32** | Notification group (ID) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **permissionId** | **[]int32** | Permission (ID) | 
 **permissionIdN** | **[]int32** | Permission (ID) | 
 **q** | **string** | Search | 
 **userId** | **[]int32** | User (ID) | 
 **userIdN** | **[]int32** | User (ID) | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this group.
	patchedGroupRequest := *openapiclient.NewPatchedGroupRequest() // PatchedGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsPartialUpdate(context.Background(), id).PatchedGroupRequest(patchedGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsPartialUpdate`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsPartialUpdate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsRetrieve`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsRetrieve`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this group.
	groupRequest := *openapiclient.NewGroupRequest("Name_example") // GroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGroupsUpdate(context.Background(), id).GroupRequest(groupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGroupsUpdate`: Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGroupsUpdate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	objectPermissionRequest := []openapiclient.ObjectPermissionRequest{*openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"})} // []ObjectPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersPermissionsBulkDestroy(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsBulkDestroy``: %v\n", err)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	objectPermissionRequest := []openapiclient.ObjectPermissionRequest{*openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"})} // []ObjectPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsBulkPartialUpdate(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsBulkPartialUpdate`: []ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsBulkPartialUpdate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	objectPermissionRequest := []openapiclient.ObjectPermissionRequest{*openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"})} // []ObjectPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsBulkUpdate(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsBulkUpdate`: []ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsBulkUpdate`: %v\n", resp)
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

> ObjectPermission UsersPermissionsCreate(ctx).ObjectPermissionRequest(objectPermissionRequest).Execute()





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
	objectPermissionRequest := *openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // ObjectPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsCreate(context.Background()).ObjectPermissionRequest(objectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsCreate`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectPermissionRequest** | [**ObjectPermissionRequest**](ObjectPermissionRequest.md) |  | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersPermissionsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsDestroy``: %v\n", err)
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

> PaginatedObjectPermissionList UsersPermissionsList(ctx).CanAdd(canAdd).CanChange(canChange).CanDelete(canDelete).CanView(canView).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Enabled(enabled).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).ObjectType(objectType).ObjectTypeIc(objectTypeIc).ObjectTypeIe(objectTypeIe).ObjectTypeIew(objectTypeIew).ObjectTypeIsw(objectTypeIsw).ObjectTypeN(objectTypeN).ObjectTypeNic(objectTypeNic).ObjectTypeNie(objectTypeNie).ObjectTypeNiew(objectTypeNiew).ObjectTypeNisw(objectTypeNisw).ObjectTypeId(objectTypeId).ObjectTypeIdN(objectTypeIdN).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).Execute()





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
	canAdd := true // bool |  (optional)
	canChange := true // bool |  (optional)
	canDelete := true // bool |  (optional)
	canView := true // bool |  (optional)
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
	idEmpty := true // bool |  (optional)
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
	objectType := "objectType_example" // string |  (optional)
	objectTypeIc := "objectTypeIc_example" // string |  (optional)
	objectTypeIe := "objectTypeIe_example" // string |  (optional)
	objectTypeIew := "objectTypeIew_example" // string |  (optional)
	objectTypeIsw := "objectTypeIsw_example" // string |  (optional)
	objectTypeN := "objectTypeN_example" // string |  (optional)
	objectTypeNic := "objectTypeNic_example" // string |  (optional)
	objectTypeNie := "objectTypeNie_example" // string |  (optional)
	objectTypeNiew := "objectTypeNiew_example" // string |  (optional)
	objectTypeNisw := "objectTypeNisw_example" // string |  (optional)
	objectTypeId := []int32{int32(123)} // []int32 |  (optional)
	objectTypeIdN := []int32{int32(123)} // []int32 |  (optional)
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
	resp, r, err := apiClient.UsersAPI.UsersPermissionsList(context.Background()).CanAdd(canAdd).CanChange(canChange).CanDelete(canDelete).CanView(canView).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Enabled(enabled).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Limit(limit).Name(name).NameEmpty(nameEmpty).NameIc(nameIc).NameIe(nameIe).NameIew(nameIew).NameIsw(nameIsw).NameN(nameN).NameNic(nameNic).NameNie(nameNie).NameNiew(nameNiew).NameNisw(nameNisw).ObjectType(objectType).ObjectTypeIc(objectTypeIc).ObjectTypeIe(objectTypeIe).ObjectTypeIew(objectTypeIew).ObjectTypeIsw(objectTypeIsw).ObjectTypeN(objectTypeN).ObjectTypeNic(objectTypeNic).ObjectTypeNie(objectTypeNie).ObjectTypeNiew(objectTypeNiew).ObjectTypeNisw(objectTypeNisw).ObjectTypeId(objectTypeId).ObjectTypeIdN(objectTypeIdN).ObjectTypes(objectTypes).ObjectTypesN(objectTypesN).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsList`: PaginatedObjectPermissionList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canAdd** | **bool** |  | 
 **canChange** | **bool** |  | 
 **canDelete** | **bool** |  | 
 **canView** | **bool** |  | 
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
 **idEmpty** | **bool** |  | 
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
 **objectType** | **string** |  | 
 **objectTypeIc** | **string** |  | 
 **objectTypeIe** | **string** |  | 
 **objectTypeIew** | **string** |  | 
 **objectTypeIsw** | **string** |  | 
 **objectTypeN** | **string** |  | 
 **objectTypeNic** | **string** |  | 
 **objectTypeNie** | **string** |  | 
 **objectTypeNiew** | **string** |  | 
 **objectTypeNisw** | **string** |  | 
 **objectTypeId** | **[]int32** |  | 
 **objectTypeIdN** | **[]int32** |  | 
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

> ObjectPermission UsersPermissionsPartialUpdate(ctx, id).PatchedObjectPermissionRequest(patchedObjectPermissionRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this permission.
	patchedObjectPermissionRequest := *openapiclient.NewPatchedObjectPermissionRequest() // PatchedObjectPermissionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsPartialUpdate(context.Background(), id).PatchedObjectPermissionRequest(patchedObjectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsPartialUpdate`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsPartialUpdate`: %v\n", resp)
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

 **patchedObjectPermissionRequest** | [**PatchedObjectPermissionRequest**](PatchedObjectPermissionRequest.md) |  | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsRetrieve`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsRetrieve`: %v\n", resp)
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

> ObjectPermission UsersPermissionsUpdate(ctx, id).ObjectPermissionRequest(objectPermissionRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this permission.
	objectPermissionRequest := *openapiclient.NewObjectPermissionRequest("Name_example", []string{"ObjectTypes_example"}, []string{"Actions_example"}) // ObjectPermissionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPermissionsUpdate(context.Background(), id).ObjectPermissionRequest(objectPermissionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPermissionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPermissionsUpdate`: ObjectPermission
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPermissionsUpdate`: %v\n", resp)
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

 **objectPermissionRequest** | [**ObjectPermissionRequest**](ObjectPermissionRequest.md) |  | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	tokenRequest := []openapiclient.TokenRequest{*openapiclient.NewTokenRequest(*openapiclient.NewBriefUserRequest("Username_example"))} // []TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTokensBulkDestroy(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensBulkDestroy``: %v\n", err)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	tokenRequest := []openapiclient.TokenRequest{*openapiclient.NewTokenRequest(*openapiclient.NewBriefUserRequest("Username_example"))} // []TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensBulkPartialUpdate(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensBulkPartialUpdate`: []Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensBulkPartialUpdate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	tokenRequest := []openapiclient.TokenRequest{*openapiclient.NewTokenRequest(*openapiclient.NewBriefUserRequest("Username_example"))} // []TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensBulkUpdate(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensBulkUpdate`: []Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensBulkUpdate`: %v\n", resp)
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

> Token UsersTokensCreate(ctx).TokenRequest(tokenRequest).Execute()





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
	tokenRequest := *openapiclient.NewTokenRequest(*openapiclient.NewBriefUserRequest("Username_example")) // TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensCreate(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensCreate`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTokensDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensDestroy``: %v\n", err)
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

> PaginatedTokenList UsersTokensList(ctx).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Expires(expires).ExpiresGte(expiresGte).ExpiresLte(expiresLte).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Key(key).KeyEmpty(keyEmpty).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNisw(keyNisw).LastUsed(lastUsed).LastUsedEmpty(lastUsedEmpty).LastUsedGt(lastUsedGt).LastUsedGte(lastUsedGte).LastUsedLt(lastUsedLt).LastUsedLte(lastUsedLte).LastUsedN(lastUsedN).Limit(limit).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).WriteEnabled(writeEnabled).Execute()





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
	idEmpty := true // bool |  (optional)
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
	lastUsed := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUsedEmpty := true // bool |  (optional)
	lastUsedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUsedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUsedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUsedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastUsedN := []time.Time{time.Now()} // []time.Time |  (optional)
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
	resp, r, err := apiClient.UsersAPI.UsersTokensList(context.Background()).Created(created).CreatedGte(createdGte).CreatedLte(createdLte).Description(description).DescriptionEmpty(descriptionEmpty).DescriptionIc(descriptionIc).DescriptionIe(descriptionIe).DescriptionIew(descriptionIew).DescriptionIsw(descriptionIsw).DescriptionN(descriptionN).DescriptionNic(descriptionNic).DescriptionNie(descriptionNie).DescriptionNiew(descriptionNiew).DescriptionNisw(descriptionNisw).Expires(expires).ExpiresGte(expiresGte).ExpiresLte(expiresLte).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).Key(key).KeyEmpty(keyEmpty).KeyIc(keyIc).KeyIe(keyIe).KeyIew(keyIew).KeyIsw(keyIsw).KeyN(keyN).KeyNic(keyNic).KeyNie(keyNie).KeyNiew(keyNiew).KeyNisw(keyNisw).LastUsed(lastUsed).LastUsedEmpty(lastUsedEmpty).LastUsedGt(lastUsedGt).LastUsedGte(lastUsedGte).LastUsedLt(lastUsedLt).LastUsedLte(lastUsedLte).LastUsedN(lastUsedN).Limit(limit).Offset(offset).Ordering(ordering).Q(q).User(user).UserN(userN).UserId(userId).UserIdN(userIdN).WriteEnabled(writeEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensList`: PaginatedTokenList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensList`: %v\n", resp)
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
 **idEmpty** | **bool** |  | 
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
 **lastUsed** | [**[]time.Time**](time.Time.md) |  | 
 **lastUsedEmpty** | **bool** |  | 
 **lastUsedGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUsedGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUsedLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastUsedLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastUsedN** | [**[]time.Time**](time.Time.md) |  | 
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

> Token UsersTokensPartialUpdate(ctx, id).PatchedTokenRequest(patchedTokenRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this token.
	patchedTokenRequest := *openapiclient.NewPatchedTokenRequest() // PatchedTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensPartialUpdate(context.Background(), id).PatchedTokenRequest(patchedTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensPartialUpdate`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensPartialUpdate`: %v\n", resp)
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

 **patchedTokenRequest** | [**PatchedTokenRequest**](PatchedTokenRequest.md) |  | 

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

> TokenProvision UsersTokensProvisionCreate(ctx).TokenProvisionRequest(tokenProvisionRequest).Execute()





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
	tokenProvisionRequest := *openapiclient.NewTokenProvisionRequest("Username_example", "Password_example") // TokenProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensProvisionCreate(context.Background()).TokenProvisionRequest(tokenProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensProvisionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensProvisionCreate`: TokenProvision
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensProvisionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersTokensProvisionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenProvisionRequest** | [**TokenProvisionRequest**](TokenProvisionRequest.md) |  | 

### Return type

[**TokenProvision**](TokenProvision.md)

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensRetrieve`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensRetrieve`: %v\n", resp)
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

> Token UsersTokensUpdate(ctx, id).TokenRequest(tokenRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this token.
	tokenRequest := *openapiclient.NewTokenRequest(*openapiclient.NewBriefUserRequest("Username_example")) // TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersTokensUpdate(context.Background(), id).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTokensUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersTokensUpdate`: Token
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersTokensUpdate`: %v\n", resp)
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

 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	userRequest := []openapiclient.UserRequest{*openapiclient.NewUserRequest("Username_example", "Password_example")} // []UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUsersBulkDestroy(context.Background()).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersBulkDestroy``: %v\n", err)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	userRequest := []openapiclient.UserRequest{*openapiclient.NewUserRequest("Username_example", "Password_example")} // []UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersBulkPartialUpdate(context.Background()).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersBulkPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersBulkPartialUpdate`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersBulkPartialUpdate`: %v\n", resp)
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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	userRequest := []openapiclient.UserRequest{*openapiclient.NewUserRequest("Username_example", "Password_example")} // []UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersBulkUpdate(context.Background()).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersBulkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersBulkUpdate`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersBulkUpdate`: %v\n", resp)
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

> User UsersUsersCreate(ctx).UserRequest(userRequest).Execute()





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
	userRequest := *openapiclient.NewUserRequest("Username_example", "Password_example") // UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersCreate(context.Background()).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersCreate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRequest** | [**UserRequest**](UserRequest.md) |  | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUsersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersDestroy``: %v\n", err)
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

> PaginatedUserList UsersUsersList(ctx).DateJoined(dateJoined).DateJoinedEmpty(dateJoinedEmpty).DateJoinedGt(dateJoinedGt).DateJoinedGte(dateJoinedGte).DateJoinedLt(dateJoinedLt).DateJoinedLte(dateJoinedLte).DateJoinedN(dateJoinedN).Email(email).EmailEmpty(emailEmpty).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNisw(emailNisw).FirstName(firstName).FirstNameEmpty(firstNameEmpty).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNisw(firstNameNisw).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).IsActive(isActive).IsStaff(isStaff).IsSuperuser(isSuperuser).LastLogin(lastLogin).LastLoginEmpty(lastLoginEmpty).LastLoginGt(lastLoginGt).LastLoginGte(lastLoginGte).LastLoginLt(lastLoginLt).LastLoginLte(lastLoginLte).LastLoginN(lastLoginN).LastName(lastName).LastNameEmpty(lastNameEmpty).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNisw(lastNameNisw).Limit(limit).NotificationGroupId(notificationGroupId).NotificationGroupIdN(notificationGroupIdN).Offset(offset).Ordering(ordering).PermissionId(permissionId).PermissionIdN(permissionIdN).Q(q).Username(username).UsernameEmpty(usernameEmpty).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNisw(usernameNisw).Execute()





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
	dateJoined := []time.Time{time.Now()} // []time.Time |  (optional)
	dateJoinedEmpty := true // bool |  (optional)
	dateJoinedGt := []time.Time{time.Now()} // []time.Time |  (optional)
	dateJoinedGte := []time.Time{time.Now()} // []time.Time |  (optional)
	dateJoinedLt := []time.Time{time.Now()} // []time.Time |  (optional)
	dateJoinedLte := []time.Time{time.Now()} // []time.Time |  (optional)
	dateJoinedN := []time.Time{time.Now()} // []time.Time |  (optional)
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
	idEmpty := true // bool |  (optional)
	idGt := []int32{int32(123)} // []int32 |  (optional)
	idGte := []int32{int32(123)} // []int32 |  (optional)
	idLt := []int32{int32(123)} // []int32 |  (optional)
	idLte := []int32{int32(123)} // []int32 |  (optional)
	idN := []int32{int32(123)} // []int32 |  (optional)
	isActive := true // bool |  (optional)
	isStaff := true // bool |  (optional)
	isSuperuser := true // bool |  (optional)
	lastLogin := []time.Time{time.Now()} // []time.Time |  (optional)
	lastLoginEmpty := true // bool |  (optional)
	lastLoginGt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastLoginGte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastLoginLt := []time.Time{time.Now()} // []time.Time |  (optional)
	lastLoginLte := []time.Time{time.Now()} // []time.Time |  (optional)
	lastLoginN := []time.Time{time.Now()} // []time.Time |  (optional)
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
	notificationGroupId := []int32{int32(123)} // []int32 | Notification group (ID) (optional)
	notificationGroupIdN := []int32{int32(123)} // []int32 | Notification group (ID) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
	permissionId := []int32{int32(123)} // []int32 | Permission (ID) (optional)
	permissionIdN := []int32{int32(123)} // []int32 | Permission (ID) (optional)
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
	resp, r, err := apiClient.UsersAPI.UsersUsersList(context.Background()).DateJoined(dateJoined).DateJoinedEmpty(dateJoinedEmpty).DateJoinedGt(dateJoinedGt).DateJoinedGte(dateJoinedGte).DateJoinedLt(dateJoinedLt).DateJoinedLte(dateJoinedLte).DateJoinedN(dateJoinedN).Email(email).EmailEmpty(emailEmpty).EmailIc(emailIc).EmailIe(emailIe).EmailIew(emailIew).EmailIsw(emailIsw).EmailN(emailN).EmailNic(emailNic).EmailNie(emailNie).EmailNiew(emailNiew).EmailNisw(emailNisw).FirstName(firstName).FirstNameEmpty(firstNameEmpty).FirstNameIc(firstNameIc).FirstNameIe(firstNameIe).FirstNameIew(firstNameIew).FirstNameIsw(firstNameIsw).FirstNameN(firstNameN).FirstNameNic(firstNameNic).FirstNameNie(firstNameNie).FirstNameNiew(firstNameNiew).FirstNameNisw(firstNameNisw).Group(group).GroupN(groupN).GroupId(groupId).GroupIdN(groupIdN).Id(id).IdEmpty(idEmpty).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdN(idN).IsActive(isActive).IsStaff(isStaff).IsSuperuser(isSuperuser).LastLogin(lastLogin).LastLoginEmpty(lastLoginEmpty).LastLoginGt(lastLoginGt).LastLoginGte(lastLoginGte).LastLoginLt(lastLoginLt).LastLoginLte(lastLoginLte).LastLoginN(lastLoginN).LastName(lastName).LastNameEmpty(lastNameEmpty).LastNameIc(lastNameIc).LastNameIe(lastNameIe).LastNameIew(lastNameIew).LastNameIsw(lastNameIsw).LastNameN(lastNameN).LastNameNic(lastNameNic).LastNameNie(lastNameNie).LastNameNiew(lastNameNiew).LastNameNisw(lastNameNisw).Limit(limit).NotificationGroupId(notificationGroupId).NotificationGroupIdN(notificationGroupIdN).Offset(offset).Ordering(ordering).PermissionId(permissionId).PermissionIdN(permissionIdN).Q(q).Username(username).UsernameEmpty(usernameEmpty).UsernameIc(usernameIc).UsernameIe(usernameIe).UsernameIew(usernameIew).UsernameIsw(usernameIsw).UsernameN(usernameN).UsernameNic(usernameNic).UsernameNie(usernameNie).UsernameNiew(usernameNiew).UsernameNisw(usernameNisw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersList`: PaginatedUserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateJoined** | [**[]time.Time**](time.Time.md) |  | 
 **dateJoinedEmpty** | **bool** |  | 
 **dateJoinedGt** | [**[]time.Time**](time.Time.md) |  | 
 **dateJoinedGte** | [**[]time.Time**](time.Time.md) |  | 
 **dateJoinedLt** | [**[]time.Time**](time.Time.md) |  | 
 **dateJoinedLte** | [**[]time.Time**](time.Time.md) |  | 
 **dateJoinedN** | [**[]time.Time**](time.Time.md) |  | 
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
 **idEmpty** | **bool** |  | 
 **idGt** | **[]int32** |  | 
 **idGte** | **[]int32** |  | 
 **idLt** | **[]int32** |  | 
 **idLte** | **[]int32** |  | 
 **idN** | **[]int32** |  | 
 **isActive** | **bool** |  | 
 **isStaff** | **bool** |  | 
 **isSuperuser** | **bool** |  | 
 **lastLogin** | [**[]time.Time**](time.Time.md) |  | 
 **lastLoginEmpty** | **bool** |  | 
 **lastLoginGt** | [**[]time.Time**](time.Time.md) |  | 
 **lastLoginGte** | [**[]time.Time**](time.Time.md) |  | 
 **lastLoginLt** | [**[]time.Time**](time.Time.md) |  | 
 **lastLoginLte** | [**[]time.Time**](time.Time.md) |  | 
 **lastLoginN** | [**[]time.Time**](time.Time.md) |  | 
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
 **notificationGroupId** | **[]int32** | Notification group (ID) | 
 **notificationGroupIdN** | **[]int32** | Notification group (ID) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **permissionId** | **[]int32** | Permission (ID) | 
 **permissionIdN** | **[]int32** | Permission (ID) | 
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

> User UsersUsersPartialUpdate(ctx, id).PatchedUserRequest(patchedUserRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this user.
	patchedUserRequest := *openapiclient.NewPatchedUserRequest() // PatchedUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersPartialUpdate(context.Background(), id).PatchedUserRequest(patchedUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersPartialUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersPartialUpdate`: %v\n", resp)
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

 **patchedUserRequest** | [**PatchedUserRequest**](PatchedUserRequest.md) |  | 

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
	openapiclient "github.com/netbox-community/go-netbox/v4"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersRetrieve`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersRetrieve`: %v\n", resp)
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

> User UsersUsersUpdate(ctx, id).UserRequest(userRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this user.
	userRequest := *openapiclient.NewUserRequest("Username_example", "Password_example") // UserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsersUpdate(context.Background(), id).UserRequest(userRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsersUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsersUpdate`: %v\n", resp)
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

 **userRequest** | [**UserRequest**](UserRequest.md) |  | 

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

