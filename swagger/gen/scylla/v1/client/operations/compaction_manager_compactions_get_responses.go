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

// CompactionManagerCompactionsGetReader is a Reader for the CompactionManagerCompactionsGet structure.
type CompactionManagerCompactionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompactionManagerCompactionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompactionManagerCompactionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCompactionManagerCompactionsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompactionManagerCompactionsGetOK creates a CompactionManagerCompactionsGetOK with default headers values
func NewCompactionManagerCompactionsGetOK() *CompactionManagerCompactionsGetOK {
	return &CompactionManagerCompactionsGetOK{}
}

/*CompactionManagerCompactionsGetOK handles this case with default header values.

CompactionManagerCompactionsGetOK compaction manager compactions get o k
*/
type CompactionManagerCompactionsGetOK struct {
	Payload []*models.Summary
}

func (o *CompactionManagerCompactionsGetOK) GetPayload() []*models.Summary {
	return o.Payload
}

func (o *CompactionManagerCompactionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompactionManagerCompactionsGetDefault creates a CompactionManagerCompactionsGetDefault with default headers values
func NewCompactionManagerCompactionsGetDefault(code int) *CompactionManagerCompactionsGetDefault {
	return &CompactionManagerCompactionsGetDefault{
		_statusCode: code,
	}
}

/*CompactionManagerCompactionsGetDefault handles this case with default header values.

internal server error
*/
type CompactionManagerCompactionsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the compaction manager compactions get default response
func (o *CompactionManagerCompactionsGetDefault) Code() int {
	return o._statusCode
}

func (o *CompactionManagerCompactionsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CompactionManagerCompactionsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CompactionManagerCompactionsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
