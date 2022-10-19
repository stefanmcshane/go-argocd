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

// ApplicationServiceDeleteReader is a Reader for the ApplicationServiceDelete structure.
type ApplicationServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceDeleteOK creates a ApplicationServiceDeleteOK with default headers values
func NewApplicationServiceDeleteOK() *ApplicationServiceDeleteOK {
	return &ApplicationServiceDeleteOK{}
}

/* ApplicationServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceDeleteOK struct {
	Payload models.ApplicationApplicationResponse
}

// IsSuccess returns true when this application service delete o k response has a 2xx status code
func (o *ApplicationServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application service delete o k response has a 3xx status code
func (o *ApplicationServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application service delete o k response has a 4xx status code
func (o *ApplicationServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application service delete o k response has a 5xx status code
func (o *ApplicationServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application service delete o k response a status code equal to that given
func (o *ApplicationServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplicationServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/applications/{name}][%d] applicationServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/applications/{name}][%d] applicationServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceDeleteOK) GetPayload() models.ApplicationApplicationResponse {
	return o.Payload
}

func (o *ApplicationServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceDeleteDefault creates a ApplicationServiceDeleteDefault with default headers values
func NewApplicationServiceDeleteDefault(code int) *ApplicationServiceDeleteDefault {
	return &ApplicationServiceDeleteDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceDeleteDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service delete default response
func (o *ApplicationServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application service delete default response has a 2xx status code
func (o *ApplicationServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application service delete default response has a 3xx status code
func (o *ApplicationServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application service delete default response has a 4xx status code
func (o *ApplicationServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application service delete default response has a 5xx status code
func (o *ApplicationServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application service delete default response a status code equal to that given
func (o *ApplicationServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/applications/{name}][%d] ApplicationService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/applications/{name}][%d] ApplicationService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceDeleteDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}