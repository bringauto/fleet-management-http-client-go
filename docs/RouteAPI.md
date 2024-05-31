# \RouteAPI

All URIs are relative to */v2/management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoutes**](RouteAPI.md#CreateRoutes) | **Post** /route | Create new Routes.
[**DeleteRoute**](RouteAPI.md#DeleteRoute) | **Delete** /route/{routeId} | Delete a Route with the specified ID.
[**GetRoute**](RouteAPI.md#GetRoute) | **Get** /route/{routeId} | Find a single Route with the specified ID.
[**GetRouteVisualization**](RouteAPI.md#GetRouteVisualization) | **Get** /route-visualization/{routeId} | Find Route Visualization for a Route identified by the route&#39;s ID.
[**GetRoutes**](RouteAPI.md#GetRoutes) | **Get** /route | Find and return all existing Routes.
[**RedefineRouteVisualizations**](RouteAPI.md#RedefineRouteVisualizations) | **Post** /route-visualization | Redefine Route Visualizations for existing Routes.
[**UpdateRoutes**](RouteAPI.md#UpdateRoutes) | **Put** /route | Update already existing Routes.



## CreateRoutes

> []Route CreateRoutes(ctx).Route(route).Execute()

Create new Routes.

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
	route := []openapiclient.Route{*openapiclient.NewRoute("Lužánky")} // []Route | A list of Route models in JSON format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAPI.CreateRoutes(context.Background()).Route(route).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAPI.CreateRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoutes`: []Route
	fmt.Fprintf(os.Stdout, "Response from `RouteAPI.CreateRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | [**[]Route**](Route.md) | A list of Route models in JSON format. | 

### Return type

[**[]Route**](Route.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteRoute(ctx, routeId).Execute()

Delete a Route with the specified ID.

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
	routeId := int32(1) // int32 | An ID a the Route

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RouteAPI.DeleteRoute(context.Background(), routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int32** | An ID a the Route | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


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


## GetRoute

> Route GetRoute(ctx, routeId).Execute()

Find a single Route with the specified ID.

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
	routeId := int32(1) // int32 | An ID a the Route

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAPI.GetRoute(context.Background(), routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAPI.GetRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoute`: Route
	fmt.Fprintf(os.Stdout, "Response from `RouteAPI.GetRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int32** | An ID a the Route | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Route**](Route.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteVisualization

> RouteVisualization GetRouteVisualization(ctx, routeId).Execute()

Find Route Visualization for a Route identified by the route's ID.

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
	routeId := int32(1) // int32 | An ID a the Route

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAPI.GetRouteVisualization(context.Background(), routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAPI.GetRouteVisualization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteVisualization`: RouteVisualization
	fmt.Fprintf(os.Stdout, "Response from `RouteAPI.GetRouteVisualization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int32** | An ID a the Route | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteVisualizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteVisualization**](RouteVisualization.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutes

> []Route GetRoutes(ctx).Execute()

Find and return all existing Routes.

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
	resp, r, err := apiClient.RouteAPI.GetRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAPI.GetRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutes`: []Route
	fmt.Fprintf(os.Stdout, "Response from `RouteAPI.GetRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesRequest struct via the builder pattern


### Return type

[**[]Route**](Route.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedefineRouteVisualizations

> []RouteVisualization RedefineRouteVisualizations(ctx).RouteVisualization(routeVisualization).Execute()

Redefine Route Visualizations for existing Routes.

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
	routeVisualization := []openapiclient.RouteVisualization{*openapiclient.NewRouteVisualization(int32(1), []openapiclient.GNSSPosition{*openapiclient.NewGNSSPosition()})} // []RouteVisualization | A list of Route Visualization models in JSON format.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteAPI.RedefineRouteVisualizations(context.Background()).RouteVisualization(routeVisualization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAPI.RedefineRouteVisualizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedefineRouteVisualizations`: []RouteVisualization
	fmt.Fprintf(os.Stdout, "Response from `RouteAPI.RedefineRouteVisualizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedefineRouteVisualizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeVisualization** | [**[]RouteVisualization**](RouteVisualization.md) | A list of Route Visualization models in JSON format. | 

### Return type

[**[]RouteVisualization**](RouteVisualization.md)

### Authorization

[oAuth2AuthCode](../README.md#oAuth2AuthCode), [APIKeyAuth](../README.md#APIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoutes

> UpdateRoutes(ctx).Route(route).Execute()

Update already existing Routes.

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
	route := []openapiclient.Route{*openapiclient.NewRoute("Lužánky")} // []Route | JSON representation of a list of the Routes with updated data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RouteAPI.UpdateRoutes(context.Background()).Route(route).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteAPI.UpdateRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | [**[]Route**](Route.md) | JSON representation of a list of the Routes with updated data. | 

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

