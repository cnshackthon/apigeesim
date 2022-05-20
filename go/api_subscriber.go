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
	"encoding/json"
	"net/http"
	"strings"

	//"github.com/gorilla/mux"
)

// SubscriberApiController binds http requests to an api service and writes the service results to the http response
type SubscriberApiController struct {
	service SubscriberApiServicer
	errorHandler ErrorHandler
}

// SubscriberApiOption for how the controller is set up.
type SubscriberApiOption func(*SubscriberApiController)

// WithSubscriberApiErrorHandler inject ErrorHandler into controller
func WithSubscriberApiErrorHandler(h ErrorHandler) SubscriberApiOption {
	return func(c *SubscriberApiController) {
		c.errorHandler = h
	}
}

// NewSubscriberApiController creates a default api controller
func NewSubscriberApiController(s SubscriberApiServicer, opts ...SubscriberApiOption) Router {
	controller := &SubscriberApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the SubscriberApiController
func (c *SubscriberApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateTestUser",
			strings.ToUpper("Put"),
			"/nac/v2/subscriber/testuser",
			c.CreateTestUser,
		},
		{
			"DeleteTestuser",
			strings.ToUpper("Delete"),
			"/nac/v2/subscriber/testuser",
			c.DeleteTestuser,
		},
		{
			"GetCustomBandwidth",
			strings.ToUpper("Post"),
			"/nac/v2/subscriber/bandwidth/custom",
			c.GetCustomBandwidth,
		},
		{
			"GetSubscriber",
			strings.ToUpper("Post"),
			"/nac/v2/subscriber",
			c.GetSubscriber,
		},
		{
			"GetSubscriberBandwidth",
			strings.ToUpper("Post"),
			"/nac/v2/subscriber/bandwidth",
			c.GetSubscriberBandwidth,
		},
		{
			"GetSubscriberLocation",
			strings.ToUpper("Post"),
			"/nac/v2/subscriber/location",
			c.GetSubscriberLocation,
		},
		{
			"UpdateCustomBandwidth",
			strings.ToUpper("Patch"),
			"/nac/v2/subscriber/bandwidth/custom",
			c.UpdateCustomBandwidth,
		},
		{
			"UpdateSubscriberBandwidth",
			strings.ToUpper("Patch"),
			"/nac/v2/subscriber/bandwidth",
			c.UpdateSubscriberBandwidth,
		},
	}
}

// CreateTestUser - Create a new test user
func (c *SubscriberApiController) CreateTestUser(w http.ResponseWriter, r *http.Request) {
	subscriberDetailsParam := SubscriberDetails{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscriberDetailsParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscriberDetailsRequired(subscriberDetailsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.CreateTestUser(r.Context(), subscriberDetailsParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteTestuser - Delete an existing test user
func (c *SubscriberApiController) DeleteTestuser(w http.ResponseWriter, r *http.Request) {
	subscriberIdParam := SubscriberId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscriberIdParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscriberIdRequired(subscriberIdParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.DeleteTestuser(r.Context(), subscriberIdParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetCustomBandwidth - Get upload/download limit
func (c *SubscriberApiController) GetCustomBandwidth(w http.ResponseWriter, r *http.Request) {
	subscriberIdParam := SubscriberId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscriberIdParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscriberIdRequired(subscriberIdParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.GetCustomBandwidth(r.Context(), subscriberIdParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSubscriber - Get the subscriber SUPI/IMSI and MSISDN
func (c *SubscriberApiController) GetSubscriber(w http.ResponseWriter, r *http.Request) {
	subscriberIdParam := SubscriberId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscriberIdParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscriberIdRequired(subscriberIdParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.GetSubscriber(r.Context(), subscriberIdParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSubscriberBandwidth - Get the current subscriber bandwidth
func (c *SubscriberApiController) GetSubscriberBandwidth(w http.ResponseWriter, r *http.Request) {
	subscriberIdParam := SubscriberId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscriberIdParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscriberIdRequired(subscriberIdParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.GetSubscriberBandwidth(r.Context(), subscriberIdParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetSubscriberLocation - Get last reported location
func (c *SubscriberApiController) GetSubscriberLocation(w http.ResponseWriter, r *http.Request) {
	subscriberIdParam := SubscriberId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&subscriberIdParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSubscriberIdRequired(subscriberIdParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.GetSubscriberLocation(r.Context(), subscriberIdParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateCustomBandwidth - Set upload limit
func (c *SubscriberApiController) UpdateCustomBandwidth(w http.ResponseWriter, r *http.Request) {
	customLimitsParam := CustomLimits{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&customLimitsParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCustomLimitsRequired(customLimitsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.UpdateCustomBandwidth(r.Context(), customLimitsParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateSubscriberBandwidth - Update the bandwidth of the subscriber
func (c *SubscriberApiController) UpdateSubscriberBandwidth(w http.ResponseWriter, r *http.Request) {
	bandwidthUpdateParam := BandwidthUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bandwidthUpdateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBandwidthUpdateRequired(bandwidthUpdateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	xTestmodeParam := r.Header.Get("x-testmode")
	result, err := c.service.UpdateSubscriberBandwidth(r.Context(), bandwidthUpdateParam, xTestmodeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}