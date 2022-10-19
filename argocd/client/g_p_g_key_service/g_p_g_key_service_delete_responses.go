// Code generated by go-swagger; DO NOT EDIT.

package g_p_g_key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// GPGKeyServiceDeleteReader is a Reader for the GPGKeyServiceDelete structure.
type GPGKeyServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GPGKeyServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGPGKeyServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGPGKeyServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGPGKeyServiceDeleteOK creates a GPGKeyServiceDeleteOK with default headers values
func NewGPGKeyServiceDeleteOK() *GPGKeyServiceDeleteOK {
	return &GPGKeyServiceDeleteOK{}
}

/* GPGKeyServiceDeleteOK describes a response with status code 200, with default header values.

A successful response.
*/
type GPGKeyServiceDeleteOK struct {
	Payload models.GpgkeyGnuPGPublicKeyResponse
}

// IsSuccess returns true when this g p g key service delete o k response has a 2xx status code
func (o *GPGKeyServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this g p g key service delete o k response has a 3xx status code
func (o *GPGKeyServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this g p g key service delete o k response has a 4xx status code
func (o *GPGKeyServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this g p g key service delete o k response has a 5xx status code
func (o *GPGKeyServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this g p g key service delete o k response a status code equal to that given
func (o *GPGKeyServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *GPGKeyServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/gpgkeys][%d] gPGKeyServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *GPGKeyServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/gpgkeys][%d] gPGKeyServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *GPGKeyServiceDeleteOK) GetPayload() models.GpgkeyGnuPGPublicKeyResponse {
	return o.Payload
}

func (o *GPGKeyServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGPGKeyServiceDeleteDefault creates a GPGKeyServiceDeleteDefault with default headers values
func NewGPGKeyServiceDeleteDefault(code int) *GPGKeyServiceDeleteDefault {
	return &GPGKeyServiceDeleteDefault{
		_statusCode: code,
	}
}

/* GPGKeyServiceDeleteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GPGKeyServiceDeleteDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the g p g key service delete default response
func (o *GPGKeyServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this g p g key service delete default response has a 2xx status code
func (o *GPGKeyServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this g p g key service delete default response has a 3xx status code
func (o *GPGKeyServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this g p g key service delete default response has a 4xx status code
func (o *GPGKeyServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this g p g key service delete default response has a 5xx status code
func (o *GPGKeyServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this g p g key service delete default response a status code equal to that given
func (o *GPGKeyServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GPGKeyServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/gpgkeys][%d] GPGKeyService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *GPGKeyServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/gpgkeys][%d] GPGKeyService_Delete default  %+v", o._statusCode, o.Payload)
}

func (o *GPGKeyServiceDeleteDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *GPGKeyServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}