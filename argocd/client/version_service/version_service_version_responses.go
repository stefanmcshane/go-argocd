// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// VersionServiceVersionReader is a Reader for the VersionServiceVersion structure.
type VersionServiceVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionServiceVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionServiceVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionServiceVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionServiceVersionOK creates a VersionServiceVersionOK with default headers values
func NewVersionServiceVersionOK() *VersionServiceVersionOK {
	return &VersionServiceVersionOK{}
}

/* VersionServiceVersionOK describes a response with status code 200, with default header values.

A successful response.
*/
type VersionServiceVersionOK struct {
	Payload *models.VersionVersionMessage
}

// IsSuccess returns true when this version service version o k response has a 2xx status code
func (o *VersionServiceVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this version service version o k response has a 3xx status code
func (o *VersionServiceVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this version service version o k response has a 4xx status code
func (o *VersionServiceVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this version service version o k response has a 5xx status code
func (o *VersionServiceVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this version service version o k response a status code equal to that given
func (o *VersionServiceVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *VersionServiceVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/version][%d] versionServiceVersionOK  %+v", 200, o.Payload)
}

func (o *VersionServiceVersionOK) String() string {
	return fmt.Sprintf("[GET /api/version][%d] versionServiceVersionOK  %+v", 200, o.Payload)
}

func (o *VersionServiceVersionOK) GetPayload() *models.VersionVersionMessage {
	return o.Payload
}

func (o *VersionServiceVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionVersionMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionServiceVersionDefault creates a VersionServiceVersionDefault with default headers values
func NewVersionServiceVersionDefault(code int) *VersionServiceVersionDefault {
	return &VersionServiceVersionDefault{
		_statusCode: code,
	}
}

/* VersionServiceVersionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VersionServiceVersionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the version service version default response
func (o *VersionServiceVersionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this version service version default response has a 2xx status code
func (o *VersionServiceVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this version service version default response has a 3xx status code
func (o *VersionServiceVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this version service version default response has a 4xx status code
func (o *VersionServiceVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this version service version default response has a 5xx status code
func (o *VersionServiceVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this version service version default response a status code equal to that given
func (o *VersionServiceVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VersionServiceVersionDefault) Error() string {
	return fmt.Sprintf("[GET /api/version][%d] VersionService_Version default  %+v", o._statusCode, o.Payload)
}

func (o *VersionServiceVersionDefault) String() string {
	return fmt.Sprintf("[GET /api/version][%d] VersionService_Version default  %+v", o._statusCode, o.Payload)
}

func (o *VersionServiceVersionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *VersionServiceVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
