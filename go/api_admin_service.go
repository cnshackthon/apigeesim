/*
 * Network as Code
 *
 * Manipulate network conditions via simplified REST calls
 *
 * API version: 2
 * Contact: todd.levi@nokia.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"context"
	"net/http"
	"errors"
)

// AdminApiService is a service that implements the logic for the AdminApiServicer
// This service should implement the business logic for every endpoint for the AdminApi API.
// Include any external packages or services that will be required by this service.
type AdminApiService struct {
}

// NewAdminApiService creates a default api service
func NewAdminApiService() AdminApiServicer {
	return &AdminApiService{}
}

// CreateTestUser - Create a new test user
func (s *AdminApiService) CreateTestUser(ctx context.Context, subscriberDetails SubscriberDetails, xTestmode string) (ImplResponse, error) {
	// TODO - update CreateTestUser with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, SubscriberDetails{}) or use other options such as http.Ok ...
	//return Response(201, SubscriberDetails{}), nil

	//TODO: Uncomment the next line to return response Response(0, ApiError{}) or use other options such as http.Ok ...
	//return Response(0, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateTestUser method not implemented")
}

// DeleteTestuser - Delete an existing test user
func (s *AdminApiService) DeleteTestuser(ctx context.Context, subscriberId SubscriberId, xTestmode string) (ImplResponse, error) {
	// TODO - update DeleteTestuser with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(0, ApiError{}) or use other options such as http.Ok ...
	//return Response(0, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteTestuser method not implemented")
}

// GetApiStatus - Is API accessible
func (s *AdminApiService) GetApiStatus(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetApiStatus with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, HelloResponses{}) or use other options such as http.Ok ...
	//return Response(200, HelloResponses{}), nil

	//TODO: Uncomment the next line to return response Response(0, ApiError{}) or use other options such as http.Ok ...
	//return Response(0, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetApiStatus method not implemented")
}
