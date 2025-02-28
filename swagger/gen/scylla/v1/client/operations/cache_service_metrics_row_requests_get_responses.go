// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// CacheServiceMetricsRowRequestsGetReader is a Reader for the CacheServiceMetricsRowRequestsGet structure.
type CacheServiceMetricsRowRequestsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsRowRequestsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsRowRequestsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsRowRequestsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsRowRequestsGetOK creates a CacheServiceMetricsRowRequestsGetOK with default headers values
func NewCacheServiceMetricsRowRequestsGetOK() *CacheServiceMetricsRowRequestsGetOK {
	return &CacheServiceMetricsRowRequestsGetOK{}
}

/*CacheServiceMetricsRowRequestsGetOK handles this case with default header values.

CacheServiceMetricsRowRequestsGetOK cache service metrics row requests get o k
*/
type CacheServiceMetricsRowRequestsGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsRowRequestsGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsRowRequestsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsRowRequestsGetDefault creates a CacheServiceMetricsRowRequestsGetDefault with default headers values
func NewCacheServiceMetricsRowRequestsGetDefault(code int) *CacheServiceMetricsRowRequestsGetDefault {
	return &CacheServiceMetricsRowRequestsGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsRowRequestsGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsRowRequestsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics row requests get default response
func (o *CacheServiceMetricsRowRequestsGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsRowRequestsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsRowRequestsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsRowRequestsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
