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

type Bandwidth string

// List of Bandwidth
const (
	BANDWIDTH_STREAMING Bandwidth = "uav_streaming"
	BANDWIDTH_LOWPOWERMODE Bandwidth = "uav_lowpowermode"
)

// AssertBandwidthRequired checks if the required fields are not zero-ed
func AssertBandwidthRequired(obj Bandwidth) error {
	return nil
}

// AssertRecurseBandwidthRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Bandwidth (e.g. [][]Bandwidth), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBandwidthRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBandwidth, ok := obj.(Bandwidth)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBandwidthRequired(aBandwidth)
	})
}
