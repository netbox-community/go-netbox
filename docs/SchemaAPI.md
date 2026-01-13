# \SchemaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaRetrieve**](SchemaAPI.md#SchemaRetrieve) | **Get** /api/schema/ | 



## SchemaRetrieve

> map[string]interface{} SchemaRetrieve(ctx).Format(format).Lang(lang).Execute()





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
	format := openapiclient.schema_retrieve_format_parameter("json") // SchemaRetrieveFormatParameter |  (optional)
	lang := openapiclient.schema_retrieve_lang_parameter("cs") // SchemaRetrieveLangParameter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.SchemaRetrieve(context.Background()).Format(format).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.SchemaRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchemaRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.SchemaRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchemaRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**SchemaRetrieveFormatParameter**](SchemaRetrieveFormatParameter.md) |  | 
 **lang** | [**SchemaRetrieveLangParameter**](SchemaRetrieveLangParameter.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.oai.openapi, application/yaml, application/vnd.oai.openapi+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

