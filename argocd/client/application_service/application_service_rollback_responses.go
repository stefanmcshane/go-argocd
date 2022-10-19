// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// ApplicationServiceRollbackReader is a Reader for the ApplicationServiceRollback structure.
type ApplicationServiceRollbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceRollbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceRollbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceRollbackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceRollbackOK creates a ApplicationServiceRollbackOK with default headers values
func NewApplicationServiceRollbackOK() *ApplicationServiceRollbackOK {
	return &ApplicationServiceRollbackOK{}
}

/* ApplicationServiceRollbackOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceRollbackOK struct {
	Payload *models.V1alpha1Application
}

// IsSuccess returns true when this application service rollback o k response has a 2xx status code
func (o *ApplicationServiceRollbackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application service rollback o k response has a 3xx status code
func (o *ApplicationServiceRollbackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application service rollback o k response has a 4xx status code
func (o *ApplicationServiceRollbackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application service rollback o k response has a 5xx status code
func (o *ApplicationServiceRollbackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application service rollback o k response a status code equal to that given
func (o *ApplicationServiceRollbackOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplicationServiceRollbackOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/rollback][%d] applicationServiceRollbackOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceRollbackOK) String() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/rollback][%d] applicationServiceRollbackOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceRollbackOK) GetPayload() *models.V1alpha1Application {
	return o.Payload
}

func (o *ApplicationServiceRollbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceRollbackDefault creates a ApplicationServiceRollbackDefault with default headers values
func NewApplicationServiceRollbackDefault(code int) *ApplicationServiceRollbackDefault {
	return &ApplicationServiceRollbackDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceRollbackDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceRollbackDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service rollback default response
func (o *ApplicationServiceRollbackDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application service rollback default response has a 2xx status code
func (o *ApplicationServiceRollbackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application service rollback default response has a 3xx status code
func (o *ApplicationServiceRollbackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application service rollback default response has a 4xx status code
func (o *ApplicationServiceRollbackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application service rollback default response has a 5xx status code
func (o *ApplicationServiceRollbackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application service rollback default response a status code equal to that given
func (o *ApplicationServiceRollbackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationServiceRollbackDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/rollback][%d] ApplicationService_Rollback default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceRollbackDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/rollback][%d] ApplicationService_Rollback default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceRollbackDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceRollbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
