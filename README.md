# Go API client for openapi

Specification for BringAuto fleet backend HTTP API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.1.0
- Package version: 2.0.0
- Generator version: 7.6.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://bringauto.com](https://bringauto.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/bringauto/fleet-management-http-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */v2/management*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiAPI* | [**CheckApiIsAlive**](docs/ApiAPI.md#checkapiisalive) | **Head** /apialive | Check HTTP server is running and accessible.
*CarAPI* | [**CreateCar**](docs/CarAPI.md#createcar) | **Post** /car | Create a new Car object.
*CarAPI* | [**DeleteCar**](docs/CarAPI.md#deletecar) | **Delete** /car/{carId} | Delete a Car identified by its ID.
*CarAPI* | [**GetCar**](docs/CarAPI.md#getcar) | **Get** /car/{carId} | Find a single Car by its ID.
*CarAPI* | [**GetCars**](docs/CarAPI.md#getcars) | **Get** /car | Find and return all existing Cars.
*CarAPI* | [**UpdateCar**](docs/CarAPI.md#updatecar) | **Put** /car | Update already existing Car.
*CarStateAPI* | [**AddCarState**](docs/CarStateAPI.md#addcarstate) | **Post** /carstate | Add a new Car State.
*CarStateAPI* | [**GetAllCarStates**](docs/CarStateAPI.md#getallcarstates) | **Get** /carstate | Find one or all Car States for all existing Cars.
*CarStateAPI* | [**GetCarStates**](docs/CarStateAPI.md#getcarstates) | **Get** /carstate/{carId} | Find one or all Car States for a Car with given ID.
*OrderAPI* | [**CreateOrder**](docs/OrderAPI.md#createorder) | **Post** /order | Create a new Order.
*OrderAPI* | [**DeleteOrder**](docs/OrderAPI.md#deleteorder) | **Delete** /order/{carId}/{orderId} | Delete an Order identified by its ID and ID of a car to which it is assigned.
*OrderAPI* | [**GetCarOrders**](docs/OrderAPI.md#getcarorders) | **Get** /order/{carId} | Find existing Orders by the corresponding Car ID and return them.
*OrderAPI* | [**GetOrder**](docs/OrderAPI.md#getorder) | **Get** /order/{carId}/{orderId} | Find an existing Order by the car ID and the order ID and return it.
*OrderAPI* | [**GetOrders**](docs/OrderAPI.md#getorders) | **Get** /order | Find all currently existing Orders.
*OrderStateAPI* | [**CreateOrderState**](docs/OrderStateAPI.md#createorderstate) | **Post** /orderstate | Add a new Order State.
*OrderStateAPI* | [**GetAllOrderStates**](docs/OrderStateAPI.md#getallorderstates) | **Get** /orderstate | Find Order States for all existing Orders.
*OrderStateAPI* | [**GetOrderStates**](docs/OrderStateAPI.md#getorderstates) | **Get** /orderstate/{orderId} | Find all Order States for a particular Order specified by its ID.
*PlatformHWAPI* | [**CreateHw**](docs/PlatformHWAPI.md#createhw) | **Post** /platformhw | Create a new Platform HW object.
*PlatformHWAPI* | [**DeleteHw**](docs/PlatformHWAPI.md#deletehw) | **Delete** /platformhw/{platformHwId} | Delete Platform HW with the given ID.
*PlatformHWAPI* | [**GetHw**](docs/PlatformHWAPI.md#gethw) | **Get** /platformhw/{platformHwId} | Find Platform HW with the given ID.
*PlatformHWAPI* | [**GetHws**](docs/PlatformHWAPI.md#gethws) | **Get** /platformhw | Find and return all existing Platform HW.
*RouteAPI* | [**CreateRoute**](docs/RouteAPI.md#createroute) | **Post** /route | Create a new Route.
*RouteAPI* | [**DeleteRoute**](docs/RouteAPI.md#deleteroute) | **Delete** /route/{routeId} | Delete a Route with the specified ID.
*RouteAPI* | [**GetRoute**](docs/RouteAPI.md#getroute) | **Get** /route/{routeId} | Find a single Route with the specified ID.
*RouteAPI* | [**GetRouteVisualization**](docs/RouteAPI.md#getroutevisualization) | **Get** /route-visualization/{routeId} | Find Route Visualization for a Route identified by the route&#39;s ID.
*RouteAPI* | [**GetRoutes**](docs/RouteAPI.md#getroutes) | **Get** /route | Find and return all existing Routes.
*RouteAPI* | [**RedefineRouteVisualization**](docs/RouteAPI.md#redefineroutevisualization) | **Post** /route-visualization | Redefine Route Visualization for an existing Route.
*RouteAPI* | [**UpdateRoute**](docs/RouteAPI.md#updateroute) | **Put** /route | Update already existing Route.
*SecurityAPI* | [**Login**](docs/SecurityAPI.md#login) | **Get** /login | 
*SecurityAPI* | [**TokenGet**](docs/SecurityAPI.md#tokenget) | **Get** /token_get | 
*SecurityAPI* | [**TokenRefresh**](docs/SecurityAPI.md#tokenrefresh) | **Get** /token_refresh | 
*StopAPI* | [**CreateStop**](docs/StopAPI.md#createstop) | **Post** /stop | Create a new Stop.
*StopAPI* | [**DeleteStop**](docs/StopAPI.md#deletestop) | **Delete** /stop/{stopId} | Delete a Stop with the specified ID.
*StopAPI* | [**GetStop**](docs/StopAPI.md#getstop) | **Get** /stop/{stopId} | Find and return a single Stop by its ID.
*StopAPI* | [**GetStops**](docs/StopAPI.md#getstops) | **Get** /stop | Find and return all existing Stops.
*StopAPI* | [**UpdateStop**](docs/StopAPI.md#updatestop) | **Put** /stop | Update already existing Stop.


## Documentation For Models

 - [Car](docs/Car.md)
 - [CarState](docs/CarState.md)
 - [CarStatus](docs/CarStatus.md)
 - [Error](docs/Error.md)
 - [GNSSPosition](docs/GNSSPosition.md)
 - [MobilePhone](docs/MobilePhone.md)
 - [Order](docs/Order.md)
 - [OrderState](docs/OrderState.md)
 - [OrderStatus](docs/OrderStatus.md)
 - [PlatformHW](docs/PlatformHW.md)
 - [Route](docs/Route.md)
 - [RouteVisualization](docs/RouteVisualization.md)
 - [Stop](docs/Stop.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### APIKeyAuth

- **Type**: API key
- **API key parameter name**: api_key
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api_key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"api_key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### oAuth2AuthCode


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://keycloak.bringauto.com/realms/bringauto/protocol/openid-connect/auth
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, openapi.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

fleet@bringauto.com

