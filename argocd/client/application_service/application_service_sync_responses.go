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

// ApplicationServiceSyncReader is a Reader for the ApplicationServiceSync structure.
type ApplicationServiceSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceSyncDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceSyncOK creates a ApplicationServiceSyncOK with default headers values
func NewApplicationServiceSyncOK() *ApplicationServiceSyncOK {
	return &ApplicationServiceSyncOK{}
}

/* ApplicationServiceSyncOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceSyncOK struct {
	Payload *models.V1alpha1Application
}

// IsSuccess returns true when this application service sync o k response has a 2xx status code
func (o *ApplicationServiceSyncOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application service sync o k response has a 3xx status code
func (o *ApplicationServiceSyncOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application service sync o k response has a 4xx status code
func (o *ApplicationServiceSyncOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application service sync o k response has a 5xx status code
func (o *ApplicationServiceSyncOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application service sync o k response a status code equal to that given
func (o *ApplicationServiceSyncOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplicationServiceSyncOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/sync][%d] applicationServiceSyncOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceSyncOK) String() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/sync][%d] applicationServiceSyncOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceSyncOK) GetPayload() *models.V1alpha1Application {
	return o.Payload
}

func (o *ApplicationServiceSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceSyncDefault creates a ApplicationServiceSyncDefault with default headers values
func NewApplicationServiceSyncDefault(code int) *ApplicationServiceSyncDefault {
	return &ApplicationServiceSyncDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceSyncDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceSyncDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service sync default response
func (o *ApplicationServiceSyncDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application service sync default response has a 2xx status code
func (o *ApplicationServiceSyncDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application service sync default response has a 3xx status code
func (o *ApplicationServiceSyncDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application service sync default response has a 4xx status code
func (o *ApplicationServiceSyncDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application service sync default response has a 5xx status code
func (o *ApplicationServiceSyncDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application service sync default response a status code equal to that given
func (o *ApplicationServiceSyncDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationServiceSyncDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/sync][%d] ApplicationService_Sync default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceSyncDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/applications/{name}/sync][%d] ApplicationService_Sync default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceSyncDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceSyncDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}