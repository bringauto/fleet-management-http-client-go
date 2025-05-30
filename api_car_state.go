/*
BringAuto Fleet Management v2 API

Specification for BringAuto fleet backend HTTP API

API version: 4.1.0
Contact: fleet@bringauto.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CarStateAPIService CarStateAPI service
type CarStateAPIService service

type ApiCreateCarStatesRequest struct {
	ctx context.Context
	ApiService *CarStateAPIService
	carState *[]CarState
}

// A list of Car State model in JSON format.
func (r ApiCreateCarStatesRequest) CarState(carState []CarState) ApiCreateCarStatesRequest {
	r.carState = &carState
	return r
}

func (r ApiCreateCarStatesRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateCarStatesExecute(r)
}

/*
CreateCarStates Add new Car States.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCarStatesRequest
*/
func (a *CarStateAPIService) CreateCarStates(ctx context.Context) ApiCreateCarStatesRequest {
	return ApiCreateCarStatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CarStateAPIService) CreateCarStatesExecute(r ApiCreateCarStatesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CarStateAPIService.CreateCarStates")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/carstate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.carState == nil {
		return nil, reportError("carState is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.carState
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAllCarStatesRequest struct {
	ctx context.Context
	ApiService *CarStateAPIService
	wait *bool
	since *int64
	lastN *int32
}

// Applies to GET methods when no objects would be returned at the moment of request. If wait&#x3D;true, \\ the request will wait for the next object to be created and then returns it. If wait&#x3D;False or unspecified, the request will return \\ an empty list.
func (r ApiGetAllCarStatesRequest) Wait(wait bool) ApiGetAllCarStatesRequest {
	r.wait = &wait
	return r
}

// A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case).
func (r ApiGetAllCarStatesRequest) Since(since int64) ApiGetAllCarStatesRequest {
	r.since = &since
	return r
}

// If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored.
func (r ApiGetAllCarStatesRequest) LastN(lastN int32) ApiGetAllCarStatesRequest {
	r.lastN = &lastN
	return r
}

func (r ApiGetAllCarStatesRequest) Execute() ([]CarState, *http.Response, error) {
	return r.ApiService.GetAllCarStatesExecute(r)
}

/*
GetAllCarStates Find one or all Car States for all existing Cars.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllCarStatesRequest
*/
func (a *CarStateAPIService) GetAllCarStates(ctx context.Context) ApiGetAllCarStatesRequest {
	return ApiGetAllCarStatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []CarState
func (a *CarStateAPIService) GetAllCarStatesExecute(r ApiGetAllCarStatesRequest) ([]CarState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CarState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CarStateAPIService.GetAllCarStates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/carstate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.wait != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "wait", r.wait, "form", "")
	} else {
		var defaultValue bool = false
		r.wait = &defaultValue
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	}
	if r.lastN != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastN", r.lastN, "form", "")
	} else {
		var defaultValue int32 = 0
		r.lastN = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCarStatesRequest struct {
	ctx context.Context
	ApiService *CarStateAPIService
	carId int32
	wait *bool
	since *int64
	lastN *int32
}

// Applies to GET methods when no objects would be returned at the moment of request. If wait&#x3D;true, \\ the request will wait for the next object to be created and then returns it. If wait&#x3D;False or unspecified, the request will return \\ an empty list.
func (r ApiGetCarStatesRequest) Wait(wait bool) ApiGetCarStatesRequest {
	r.wait = &wait
	return r
}

// A Unix timestamp in milliseconds. If specified, only objects created at the time or later will be returned. If unspecified, all objects are returned (since is set to 0 in that case).
func (r ApiGetCarStatesRequest) Since(since int64) ApiGetCarStatesRequest {
	r.since = &since
	return r
}

// If specified, only the last N objects will be returned. If unspecified, all objects are returned. \\ If some states have identical timestamps and they all do not fit into the maximum N states, only those with higher IDs are returned. If value smaller than 1 is provided, this filtering is ignored.
func (r ApiGetCarStatesRequest) LastN(lastN int32) ApiGetCarStatesRequest {
	r.lastN = &lastN
	return r
}

func (r ApiGetCarStatesRequest) Execute() ([]CarState, *http.Response, error) {
	return r.ApiService.GetCarStatesExecute(r)
}

/*
GetCarStates Find one or all Car States for a Car with given ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param carId The car ID.
 @return ApiGetCarStatesRequest
*/
func (a *CarStateAPIService) GetCarStates(ctx context.Context, carId int32) ApiGetCarStatesRequest {
	return ApiGetCarStatesRequest{
		ApiService: a,
		ctx: ctx,
		carId: carId,
	}
}

// Execute executes the request
//  @return []CarState
func (a *CarStateAPIService) GetCarStatesExecute(r ApiGetCarStatesRequest) ([]CarState, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CarState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CarStateAPIService.GetCarStates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/carstate/{carId}"
	localVarPath = strings.Replace(localVarPath, "{"+"carId"+"}", url.PathEscape(parameterValueToString(r.carId, "carId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.wait != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "wait", r.wait, "form", "")
	} else {
		var defaultValue bool = false
		r.wait = &defaultValue
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	}
	if r.lastN != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastN", r.lastN, "form", "")
	} else {
		var defaultValue int32 = 0
		r.lastN = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
