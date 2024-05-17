/*
BringAuto Fleet Management v2 API

Testing OrderAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_OrderAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrderAPIService CreateOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrderAPI.CreateOrder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAPIService DeleteOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var carId int32
		var orderId int32

		httpRes, err := apiClient.OrderAPI.DeleteOrder(context.Background(), carId, orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAPIService GetCarOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var carId int32

		resp, httpRes, err := apiClient.OrderAPI.GetCarOrders(context.Background(), carId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAPIService GetOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var carId int32
		var orderId int32

		resp, httpRes, err := apiClient.OrderAPI.GetOrder(context.Background(), carId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAPIService GetOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrderAPI.GetOrders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}