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

// ApplicationServiceManagedResourcesReader is a Reader for the ApplicationServiceManagedResources structure.
type ApplicationServiceManagedResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceManagedResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceManagedResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceManagedResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceManagedResourcesOK creates a ApplicationServiceManagedResourcesOK with default headers values
func NewApplicationServiceManagedResourcesOK() *ApplicationServiceManagedResourcesOK {
	return &ApplicationServiceManagedResourcesOK{}
}

/* ApplicationServiceManagedResourcesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceManagedResourcesOK struct {
	Payload *models.ApplicationManagedResourcesResponse
}

// IsSuccess returns true when this application service managed resources o k response has a 2xx status code
func (o *ApplicationServiceManagedResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application service managed resources o k response has a 3xx status code
func (o *ApplicationServiceManagedResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application service managed resources o k response has a 4xx status code
func (o *ApplicationServiceManagedResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application service managed resources o k response has a 5xx status code
func (o *ApplicationServiceManagedResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application service managed resources o k response a status code equal to that given
func (o *ApplicationServiceManagedResourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplicationServiceManagedResourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/managed-resources][%d] applicationServiceManagedResourcesOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceManagedResourcesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/managed-resources][%d] applicationServiceManagedResourcesOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceManagedResourcesOK) GetPayload() *models.ApplicationManagedResourcesResponse {
	return o.Payload
}

func (o *ApplicationServiceManagedResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationManagedResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceManagedResourcesDefault creates a ApplicationServiceManagedResourcesDefault with default headers values
func NewApplicationServiceManagedResourcesDefault(code int) *ApplicationServiceManagedResourcesDefault {
	return &ApplicationServiceManagedResourcesDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceManagedResourcesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceManagedResourcesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service managed resources default response
func (o *ApplicationServiceManagedResourcesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application service managed resources default response has a 2xx status code
func (o *ApplicationServiceManagedResourcesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application service managed resources default response has a 3xx status code
func (o *ApplicationServiceManagedResourcesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application service managed resources default response has a 4xx status code
func (o *ApplicationServiceManagedResourcesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application service managed resources default response has a 5xx status code
func (o *ApplicationServiceManagedResourcesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application service managed resources default response a status code equal to that given
func (o *ApplicationServiceManagedResourcesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationServiceManagedResourcesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/managed-resources][%d] ApplicationService_ManagedResources default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceManagedResourcesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/managed-resources][%d] ApplicationService_ManagedResources default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceManagedResourcesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceManagedResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
