# \OrderStateAPI

All URIs are relative to */v2/management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrderStates**](OrderStateAPI.md#CreateOrderStates) | **Post** /orderstate | Add new Order States.
[**GetAllOrderStates**](OrderStateAPI.md#GetAllOrderStates) | **Get** /orderstate | Find Order States for all existing Orders.
[**GetOrderStates**](OrderStateAPI.md#GetOrderStates) | **Get** /orderstate/{orderId} | Find all Order States for a particular Order specified by its ID.



## CreateOrderStates

> []OrderState CreateOrderStates(ctx).OrderState(orderState).Execute()

Add new Order States.

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
	orderState := []openapiclient.OrderState{*openapiclient.NewOrderState(openapiclient.OrderStatus("to_accept"), int32(1))} // []OrderState | A list of Order State models in JSON format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStateAPI.CreateOrderStates(context.Background()).OrderState(orderState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStateAPI.CreateOrderStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderStates`: []OrderState
	fmt.Fprintf(os.Stdout, "Response from `OrderStateAPI.CreateOrderStates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderState** | [**[]OrderState**](OrderState.md) | A list of Order State models in JSON format. | 

### Return type

[**[]OrderState**](OrderState.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllOrderStates

> []OrderState GetAllOrderStates(ctx).Wait(wait).Since(since).LastN(lastN).CarId(carId).Execute()

Find Order States for all existing Orders.

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
	wait := true // bool | Applies to GET methods when no objects would be returned at the moment of request. If wait=true, \\ the request will wait for the next object to be created and then returns it. If wait=False or unspecified, the request will return \\ an empty list. (optional) (default to false)
	since := int64(789) // int64 | A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case). (optional)
	lastN := int32(56) // int32 | If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored. (optional) (default to 0)
	carId := int32(56) // int32 | An optional parameter for filtering only objects related to a car with the specified ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStateAPI.GetAllOrderStates(context.Background()).Wait(wait).Since(since).LastN(lastN).CarId(carId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStateAPI.GetAllOrderStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllOrderStates`: []OrderState
	fmt.Fprintf(os.Stdout, "Response from `OrderStateAPI.GetAllOrderStates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllOrderStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wait** | **bool** | Applies to GET methods when no objects would be returned at the moment of request. If wait&#x3D;true, \\ the request will wait for the next object to be created and then returns it. If wait&#x3D;False or unspecified, the request will return \\ an empty list. | [default to false]
 **since** | **int64** | A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case). | 
 **lastN** | **int32** | If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored. | [default to 0]
 **carId** | **int32** | An optional parameter for filtering only objects related to a car with the specified ID. | 

### Return type

[**[]OrderState**](OrderState.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderStates

> []OrderState GetOrderStates(ctx, orderId).Wait(wait).Since(since).LastN(lastN).Execute()

Find all Order States for a particular Order specified by its ID.

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
	orderId := int32(56) // int32 | The order ID.
	wait := true // bool | Applies to GET methods when no objects would be returned at the moment of request. If wait=true, \\ the request will wait for the next object to be created and then returns it. If wait=False or unspecified, the request will return \\ an empty list. (optional) (default to false)
	since := int64(789) // int64 | A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case). (optional)
	lastN := int32(56) // int32 | If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStateAPI.GetOrderStates(context.Background(), orderId).Wait(wait).Since(since).LastN(lastN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStateAPI.GetOrderStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderStates`: []OrderState
	fmt.Fprintf(os.Stdout, "Response from `OrderStateAPI.GetOrderStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** | The order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wait** | **bool** | Applies to GET methods when no objects would be returned at the moment of request. If wait&#x3D;true, \\ the request will wait for the next object to be created and then returns it. If wait&#x3D;False or unspecified, the request will return \\ an empty list. | [default to false]
 **since** | **int64** | A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case). | 
 **lastN** | **int32** | If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored. | [default to 0]

### Return type

[**[]OrderState**](OrderState.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

