// Code generated by go-swagger; DO NOT EDIT.

package session_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// SessionServiceCreateReader is a Reader for the SessionServiceCreate structure.
type SessionServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSessionServiceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSessionServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSessionServiceCreateOK creates a SessionServiceCreateOK with default headers values
func NewSessionServiceCreateOK() *SessionServiceCreateOK {
	return &SessionServiceCreateOK{}
}

/* SessionServiceCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type SessionServiceCreateOK struct {
	Payload *models.SessionSessionResponse
}

// IsSuccess returns true when this session service create o k response has a 2xx status code
func (o *SessionServiceCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this session service create o k response has a 3xx status code
func (o *SessionServiceCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this session service create o k response has a 4xx status code
func (o *SessionServiceCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this session service create o k response has a 5xx status code
func (o *SessionServiceCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this session service create o k response a status code equal to that given
func (o *SessionServiceCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SessionServiceCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/session][%d] sessionServiceCreateOK  %+v", 200, o.Payload)
}

func (o *SessionServiceCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v1/session][%d] sessionServiceCreateOK  %+v", 200, o.Payload)
}

func (o *SessionServiceCreateOK) GetPayload() *models.SessionSessionResponse {
	return o.Payload
}

func (o *SessionServiceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSessionServiceCreateDefault creates a SessionServiceCreateDefault with default headers values
func NewSessionServiceCreateDefault(code int) *SessionServiceCreateDefault {
	return &SessionServiceCreateDefault{
		_statusCode: code,
	}
}

/* SessionServiceCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SessionServiceCreateDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the session service create default response
func (o *SessionServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this session service create default response has a 2xx status code
func (o *SessionServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this session service create default response has a 3xx status code
func (o *SessionServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this session service create default response has a 4xx status code
func (o *SessionServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this session service create default response has a 5xx status code
func (o *SessionServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this session service create default response a status code equal to that given
func (o *SessionServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SessionServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/session][%d] SessionService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *SessionServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/session][%d] SessionService_Create default  %+v", o._statusCode, o.Payload)
}

func (o *SessionServiceCreateDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *SessionServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
