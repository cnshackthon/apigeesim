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

// AdminApiController binds http requests to an api service and writes the service results to the http response
type AdminApiController struct {
	service AdminApiServicer
	errorHandler ErrorHandler
}

// AdminApiOption for how the controller is set up.
type AdminApiOption func(*AdminApiController)

// WithAdminApiErrorHandler inject ErrorHandler into controller
func WithAdminApiErrorHandler(h ErrorHandler) AdminApiOption {
	return func(c *AdminApiController) {
		c.errorHandler = h
	}
}

// NewAdminApiController creates a default api controller
func NewAdminApiController(s AdminApiServicer, opts ...AdminApiOption) Router {
	controller := &AdminApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the AdminApiController
func (c *AdminApiController) Routes() Routes {
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
			"GetApiStatus",
			strings.ToUpper("Get"),
			"/nac/v2/hello",
			c.GetApiStatus,
		},
	}
}

// CreateTestUser - Create a new test user
func (c *AdminApiController) CreateTestUser(w http.ResponseWriter, r *http.Request) {
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
func (c *AdminApiController) DeleteTestuser(w http.ResponseWriter, r *http.Request) {
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

// GetApiStatus - Is API accessible
func (c *AdminApiController) GetApiStatus(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetApiStatus(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
