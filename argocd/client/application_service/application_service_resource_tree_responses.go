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

// ApplicationServiceResourceTreeReader is a Reader for the ApplicationServiceResourceTree structure.
type ApplicationServiceResourceTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationServiceResourceTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationServiceResourceTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationServiceResourceTreeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationServiceResourceTreeOK creates a ApplicationServiceResourceTreeOK with default headers values
func NewApplicationServiceResourceTreeOK() *ApplicationServiceResourceTreeOK {
	return &ApplicationServiceResourceTreeOK{}
}

/* ApplicationServiceResourceTreeOK describes a response with status code 200, with default header values.

A successful response.
*/
type ApplicationServiceResourceTreeOK struct {
	Payload *models.V1alpha1ApplicationTree
}

// IsSuccess returns true when this application service resource tree o k response has a 2xx status code
func (o *ApplicationServiceResourceTreeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application service resource tree o k response has a 3xx status code
func (o *ApplicationServiceResourceTreeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application service resource tree o k response has a 4xx status code
func (o *ApplicationServiceResourceTreeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application service resource tree o k response has a 5xx status code
func (o *ApplicationServiceResourceTreeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application service resource tree o k response a status code equal to that given
func (o *ApplicationServiceResourceTreeOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplicationServiceResourceTreeOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/resource-tree][%d] applicationServiceResourceTreeOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceResourceTreeOK) String() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/resource-tree][%d] applicationServiceResourceTreeOK  %+v", 200, o.Payload)
}

func (o *ApplicationServiceResourceTreeOK) GetPayload() *models.V1alpha1ApplicationTree {
	return o.Payload
}

func (o *ApplicationServiceResourceTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1ApplicationTree)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationServiceResourceTreeDefault creates a ApplicationServiceResourceTreeDefault with default headers values
func NewApplicationServiceResourceTreeDefault(code int) *ApplicationServiceResourceTreeDefault {
	return &ApplicationServiceResourceTreeDefault{
		_statusCode: code,
	}
}

/* ApplicationServiceResourceTreeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ApplicationServiceResourceTreeDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the application service resource tree default response
func (o *ApplicationServiceResourceTreeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application service resource tree default response has a 2xx status code
func (o *ApplicationServiceResourceTreeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application service resource tree default response has a 3xx status code
func (o *ApplicationServiceResourceTreeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application service resource tree default response has a 4xx status code
func (o *ApplicationServiceResourceTreeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application service resource tree default response has a 5xx status code
func (o *ApplicationServiceResourceTreeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application service resource tree default response a status code equal to that given
func (o *ApplicationServiceResourceTreeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationServiceResourceTreeDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/resource-tree][%d] ApplicationService_ResourceTree default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceResourceTreeDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/applications/{applicationName}/resource-tree][%d] ApplicationService_ResourceTree default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationServiceResourceTreeDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ApplicationServiceResourceTreeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}