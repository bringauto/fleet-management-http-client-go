# \PlatformHWAPI

All URIs are relative to */v2/management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHws**](PlatformHWAPI.md#CreateHws) | **Post** /platformhw | Create new Platform HW objects.
[**DeleteHw**](PlatformHWAPI.md#DeleteHw) | **Delete** /platformhw/{platformHwId} | Delete Platform HW with the given ID.
[**GetHw**](PlatformHWAPI.md#GetHw) | **Get** /platformhw/{platformHwId} | Find Platform HW with the given ID.
[**GetHws**](PlatformHWAPI.md#GetHws) | **Get** /platformhw | Find and return all existing Platform HW.



## CreateHws

> []PlatformHW CreateHws(ctx).PlatformHW(platformHW).Execute()

Create new Platform HW objects.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	platformHW := []openapiclient.PlatformHW{*openapiclient.NewPlatformHW("ABCD1234EF56")} // []PlatformHW | A list of Platform HW models in JSON format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformHWAPI.CreateHws(context.Background()).PlatformHW(platformHW).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformHWAPI.CreateHws``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHws`: []PlatformHW
	fmt.Fprintf(os.Stdout, "Response from `PlatformHWAPI.CreateHws`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformHW** | [**[]PlatformHW**](PlatformHW.md) | A list of Platform HW models in JSON format. | 

### Return type

[**[]PlatformHW**](PlatformHW.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHw

> DeleteHw(ctx, platformHwId).Execute()

Delete Platform HW with the given ID.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	platformHwId := int32(56) // int32 | The Platform HW ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformHWAPI.DeleteHw(context.Background(), platformHwId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformHWAPI.DeleteHw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformHwId** | **int32** | The Platform HW ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHwRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHw

> PlatformHW GetHw(ctx, platformHwId).Execute()

Find Platform HW with the given ID.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	platformHwId := int32(56) // int32 | The Platform HW ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformHWAPI.GetHw(context.Background(), platformHwId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformHWAPI.GetHw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHw`: PlatformHW
	fmt.Fprintf(os.Stdout, "Response from `PlatformHWAPI.GetHw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformHwId** | **int32** | The Platform HW ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHwRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformHW**](PlatformHW.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHws

> []PlatformHW GetHws(ctx).Execute()

Find and return all existing Platform HW.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformHWAPI.GetHws(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformHWAPI.GetHws``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHws`: []PlatformHW
	fmt.Fprintf(os.Stdout, "Response from `PlatformHWAPI.GetHws`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHwsRequest struct via the builder pattern


### Return type

[**[]PlatformHW**](PlatformHW.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

