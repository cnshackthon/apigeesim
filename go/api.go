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
)



// AdminApiRouter defines the required methods for binding the api requests to a responses for the AdminApi
// The AdminApiRouter implementation should parse necessary information from the http request,
// pass the data to a AdminApiServicer to perform the required actions, then write the service results to the http response.
type AdminApiRouter interface { 
	CreateTestUser(http.ResponseWriter, *http.Request)
	DeleteTestuser(http.ResponseWriter, *http.Request)
	GetApiStatus(http.ResponseWriter, *http.Request)
}
// BandwidthApiRouter defines the required methods for binding the api requests to a responses for the BandwidthApi
// The BandwidthApiRouter implementation should parse necessary information from the http request,
// pass the data to a BandwidthApiServicer to perform the required actions, then write the service results to the http response.
type BandwidthApiRouter interface { 
	GetCustomBandwidth(http.ResponseWriter, *http.Request)
	GetSubscriberBandwidth(http.ResponseWriter, *http.Request)
	UpdateCustomBandwidth(http.ResponseWriter, *http.Request)
	UpdateSubscriberBandwidth(http.ResponseWriter, *http.Request)
}
// LocationApiRouter defines the required methods for binding the api requests to a responses for the LocationApi
// The LocationApiRouter implementation should parse necessary information from the http request,
// pass the data to a LocationApiServicer to perform the required actions, then write the service results to the http response.
type LocationApiRouter interface { 
	GetSubscriberLocation(http.ResponseWriter, *http.Request)
}
// SubscriberApiRouter defines the required methods for binding the api requests to a responses for the SubscriberApi
// The SubscriberApiRouter implementation should parse necessary information from the http request,
// pass the data to a SubscriberApiServicer to perform the required actions, then write the service results to the http response.
type SubscriberApiRouter interface { 
	CreateTestUser(http.ResponseWriter, *http.Request)
	DeleteTestuser(http.ResponseWriter, *http.Request)
	GetCustomBandwidth(http.ResponseWriter, *http.Request)
	GetSubscriber(http.ResponseWriter, *http.Request)
	GetSubscriberBandwidth(http.ResponseWriter, *http.Request)
	GetSubscriberLocation(http.ResponseWriter, *http.Request)
	UpdateCustomBandwidth(http.ResponseWriter, *http.Request)
	UpdateSubscriberBandwidth(http.ResponseWriter, *http.Request)
}


// AdminApiServicer defines the api actions for the AdminApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AdminApiServicer interface { 
	CreateTestUser(context.Context, SubscriberDetails, string) (ImplResponse, error)
	DeleteTestuser(context.Context, SubscriberId, string) (ImplResponse, error)
	GetApiStatus(context.Context) (ImplResponse, error)
}


// BandwidthApiServicer defines the api actions for the BandwidthApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type BandwidthApiServicer interface { 
	GetCustomBandwidth(context.Context, SubscriberId, string) (ImplResponse, error)
	GetSubscriberBandwidth(context.Context, SubscriberId, string) (ImplResponse, error)
	UpdateCustomBandwidth(context.Context, CustomLimits, string) (ImplResponse, error)
	UpdateSubscriberBandwidth(context.Context, BandwidthUpdate, string) (ImplResponse, error)
}


// LocationApiServicer defines the api actions for the LocationApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type LocationApiServicer interface { 
	GetSubscriberLocation(context.Context, SubscriberId, string) (ImplResponse, error)
}


// SubscriberApiServicer defines the api actions for the SubscriberApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type SubscriberApiServicer interface { 
	CreateTestUser(context.Context, SubscriberDetails, string) (ImplResponse, error)
	DeleteTestuser(context.Context, SubscriberId, string) (ImplResponse, error)
	GetCustomBandwidth(context.Context, SubscriberId, string) (ImplResponse, error)
	GetSubscriber(context.Context, SubscriberId, string) (ImplResponse, error)
	GetSubscriberBandwidth(context.Context, SubscriberId, string) (ImplResponse, error)
	GetSubscriberLocation(context.Context, SubscriberId, string) (ImplResponse, error)
	UpdateCustomBandwidth(context.Context, CustomLimits, string) (ImplResponse, error)
	UpdateSubscriberBandwidth(context.Context, BandwidthUpdate, string) (ImplResponse, error)
}
