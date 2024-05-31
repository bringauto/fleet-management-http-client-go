# \StopAPI

All URIs are relative to */v2/management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStops**](StopAPI.md#CreateStops) | **Post** /stop | Create new Stops.
[**DeleteStop**](StopAPI.md#DeleteStop) | **Delete** /stop/{stopId} | Delete a Stop with the specified ID.
[**GetStop**](StopAPI.md#GetStop) | **Get** /stop/{stopId} | Find and return a single Stop by its ID.
[**GetStops**](StopAPI.md#GetStops) | **Get** /stop | Find and return all existing Stops.
[**UpdateStops**](StopAPI.md#UpdateStops) | **Put** /stop | Update already existing Stops.



## CreateStops

> []Stop CreateStops(ctx).Stop(stop).Execute()

Create new Stops.

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
	stop := []openapiclient.Stop{*openapiclient.NewStop("Lidická", *openapiclient.NewGNSSPosition())} // []Stop | A list of Stop models in JSON format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StopAPI.CreateStops(context.Background()).Stop(stop).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StopAPI.CreateStops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStops`: []Stop
	fmt.Fprintf(os.Stdout, "Response from `StopAPI.CreateStops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stop** | [**[]Stop**](Stop.md) | A list of Stop models in JSON format. | 

### Return type

[**[]Stop**](Stop.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStop

> DeleteStop(ctx, stopId).Execute()

Delete a Stop with the specified ID.

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
	stopId := int32(1) // int32 | ID of the Stop to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StopAPI.DeleteStop(context.Background(), stopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StopAPI.DeleteStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stopId** | **int32** | ID of the Stop to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStopRequest struct via the builder pattern


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


## GetStop

> Stop GetStop(ctx, stopId).Execute()

Find and return a single Stop by its ID.

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
	stopId := int32(1) // int32 | ID of Stop to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StopAPI.GetStop(context.Background(), stopId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StopAPI.GetStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStop`: Stop
	fmt.Fprintf(os.Stdout, "Response from `StopAPI.GetStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stopId** | **int32** | ID of Stop to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stop**](Stop.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStops

> []Stop GetStops(ctx).Execute()

Find and return all existing Stops.

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
	resp, r, err := apiClient.StopAPI.GetStops(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StopAPI.GetStops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStops`: []Stop
	fmt.Fprintf(os.Stdout, "Response from `StopAPI.GetStops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStopsRequest struct via the builder pattern


### Return type

[**[]Stop**](Stop.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStops

> UpdateStops(ctx).Stop(stop).Execute()

Update already existing Stops.

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
	stop := []openapiclient.Stop{*openapiclient.NewStop("Lidická", *openapiclient.NewGNSSPosition())} // []Stop | JSON representation of a list of the Stops with updated data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StopAPI.UpdateStops(context.Background()).Stop(stop).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StopAPI.UpdateStops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stop** | [**[]Stop**](Stop.md) | JSON representation of a list of the Stops with updated data. | 

### Return type

 (empty response body)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

