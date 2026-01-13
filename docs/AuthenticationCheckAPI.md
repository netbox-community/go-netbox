# \AuthenticationCheckAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationCheckRetrieve**](AuthenticationCheckAPI.md#AuthenticationCheckRetrieve) | **Get** /api/authentication-check/ | 



## AuthenticationCheckRetrieve

> map[string]interface{} AuthenticationCheckRetrieve(ctx).Execute()





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
	resp, r, err := apiClient.AuthenticationCheckAPI.AuthenticationCheckRetrieve(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationCheckAPI.AuthenticationCheckRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticationCheckRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationCheckAPI.AuthenticationCheckRetrieve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationCheckRetrieveRequest struct via the builder pattern


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

