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
	//"net/http"
	//"errors"
	"log"
	"time"
)

// LocationApiService is a service that implements the logic for the LocationApiServicer
// This service should implement the business logic for every endpoint for the LocationApi API.
// Include any external packages or services that will be required by this service.
type LocationApiService struct {
}

// NewLocationApiService creates a default api service
func NewLocationApiService() LocationApiServicer {
	return &LocationApiService{}
}

// GetSubscriberLocation - Get last reported location
func (s *LocationApiService) GetSubscriberLocation(ctx context.Context, subscriberId SubscriberId, xTestmode string) (ImplResponse, error) {
	log.Printf("GetSubscriberLocation function is called, sub id is %s", subscriberId.Id)
	// TODO - update GetSubscriberLocation with the required logic for this service method.
	// Add api_location_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	/*
	{
  "eventTime": "2022-03-23T15:20:50Z",
  "id": "123456789",
  "locationInfo": {
    "ageOfLocationInfo": "123",
    "trackingAreaId": "3100120016A8",
    "plmnId": "310012",
    "lat": "0.000000000001",
    "long": "90.000000000001",
    "elev": "456"
  }
}
	*/
	var locRsp  LocationResponse
	locRsp.EventTime = time.Now()
	locRsp.Id = subscriberId.Id
	loc :=LocationDetails{ AgeOfLocationInfo: "123",
    TrackingAreaId: "3100120016A8",
    PlmnId: "310012",
    Lat: "0.000000000001",
    Long: "90.000000000001",
    Elev: "456"}
	locRsp.LocationInfo = loc
	//TODO: Uncomment the next line to return response Response(200, LocationResponse{}) or use other options such as http.Ok ...
	return Response(200, locRsp), nil

	//TODO: Uncomment the next line to return response Response(0, ApiError{}) or use other options such as http.Ok ...
	//return Response(0, ApiError{}), nil

	//return Response(http.StatusNotImplemented, nil), errors.New("GetSubscriberLocation method not implemented")
}