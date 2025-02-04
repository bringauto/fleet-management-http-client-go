# \CarActionAPI

All URIs are relative to */v2/management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCarActionStates**](CarActionAPI.md#GetCarActionStates) | **Get** /action/car/{carId} | Finds car action states for a Car with given carId.
[**PauseCar**](CarActionAPI.md#PauseCar) | **Post** /action/car/{carId}/pause | Finds and pauses a Car with given carId, if not already paused. Sets car action status to PAUSED if it is not in PAUSED action status already.
[**UnpauseCar**](CarActionAPI.md#UnpauseCar) | **Post** /action/car/{carId}/unpause | Finds and unpauses a Car with given carId, if paused. Sets car action status to NORMAL only if it is in PAUSED action status.



## GetCarActionStates

> []CarActionState GetCarActionStates(ctx, carId).Wait(wait).Since(since).LastN(lastN).Execute()

Finds car action states for a Car with given carId.

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
	carId := int32(56) // int32 | The car ID.
	wait := true // bool | Applies to GET methods when no objects would be returned at the moment of request. If wait=true, \\ the request will wait for the next object to be created and then returns it. If wait=False or unspecified, the request will return \\ an empty list. (optional) (default to false)
	since := int64(789) // int64 | A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case). (optional)
	lastN := int32(56) // int32 | If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CarActionAPI.GetCarActionStates(context.Background(), carId).Wait(wait).Since(since).LastN(lastN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CarActionAPI.GetCarActionStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCarActionStates`: []CarActionState
	fmt.Fprintf(os.Stdout, "Response from `CarActionAPI.GetCarActionStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carId** | **int32** | The car ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCarActionStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wait** | **bool** | Applies to GET methods when no objects would be returned at the moment of request. If wait&#x3D;true, \\ the request will wait for the next object to be created and then returns it. If wait&#x3D;False or unspecified, the request will return \\ an empty list. | [default to false]
 **since** | **int64** | A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case). | 
 **lastN** | **int32** | If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored. | [default to 0]

### Return type

[**[]CarActionState**](CarActionState.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseCar

> []CarActionState PauseCar(ctx, carId).Execute()

Finds and pauses a Car with given carId, if not already paused. Sets car action status to PAUSED if it is not in PAUSED action status already.

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
	carId := int32(56) // int32 | The car ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CarActionAPI.PauseCar(context.Background(), carId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CarActionAPI.PauseCar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseCar`: []CarActionState
	fmt.Fprintf(os.Stdout, "Response from `CarActionAPI.PauseCar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carId** | **int32** | The car ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseCarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CarActionState**](CarActionState.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpauseCar

> []CarActionState UnpauseCar(ctx, carId).Execute()

Finds and unpauses a Car with given carId, if paused. Sets car action status to NORMAL only if it is in PAUSED action status.

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
	carId := int32(56) // int32 | The car ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CarActionAPI.UnpauseCar(context.Background(), carId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CarActionAPI.UnpauseCar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpauseCar`: []CarActionState
	fmt.Fprintf(os.Stdout, "Response from `CarActionAPI.UnpauseCar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**carId** | **int32** | The car ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpauseCarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CarActionState**](CarActionState.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

