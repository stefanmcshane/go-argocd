// Code generated by go-swagger; DO NOT EDIT.

package cluster_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// ClusterServiceRotateAuthReader is a Reader for the ClusterServiceRotateAuth structure.
type ClusterServiceRotateAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterServiceRotateAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterServiceRotateAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterServiceRotateAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterServiceRotateAuthOK creates a ClusterServiceRotateAuthOK with default headers values
func NewClusterServiceRotateAuthOK() *ClusterServiceRotateAuthOK {
	return &ClusterServiceRotateAuthOK{}
}

/* ClusterServiceRotateAuthOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterServiceRotateAuthOK struct {
	Payload models.ClusterClusterResponse
}

// IsSuccess returns true when this cluster service rotate auth o k response has a 2xx status code
func (o *ClusterServiceRotateAuthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster service rotate auth o k response has a 3xx status code
func (o *ClusterServiceRotateAuthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster service rotate auth o k response has a 4xx status code
func (o *ClusterServiceRotateAuthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster service rotate auth o k response has a 5xx status code
func (o *ClusterServiceRotateAuthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster service rotate auth o k response a status code equal to that given
func (o *ClusterServiceRotateAuthOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterServiceRotateAuthOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/clusters/{id.value}/rotate-auth][%d] clusterServiceRotateAuthOK  %+v", 200, o.Payload)
}

func (o *ClusterServiceRotateAuthOK) String() string {
	return fmt.Sprintf("[POST /api/v1/clusters/{id.value}/rotate-auth][%d] clusterServiceRotateAuthOK  %+v", 200, o.Payload)
}

func (o *ClusterServiceRotateAuthOK) GetPayload() models.ClusterClusterResponse {
	return o.Payload
}

func (o *ClusterServiceRotateAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterServiceRotateAuthDefault creates a ClusterServiceRotateAuthDefault with default headers values
func NewClusterServiceRotateAuthDefault(code int) *ClusterServiceRotateAuthDefault {
	return &ClusterServiceRotateAuthDefault{
		_statusCode: code,
	}
}

/* ClusterServiceRotateAuthDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterServiceRotateAuthDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the cluster service rotate auth default response
func (o *ClusterServiceRotateAuthDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster service rotate auth default response has a 2xx status code
func (o *ClusterServiceRotateAuthDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster service rotate auth default response has a 3xx status code
func (o *ClusterServiceRotateAuthDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster service rotate auth default response has a 4xx status code
func (o *ClusterServiceRotateAuthDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster service rotate auth default response has a 5xx status code
func (o *ClusterServiceRotateAuthDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster service rotate auth default response a status code equal to that given
func (o *ClusterServiceRotateAuthDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterServiceRotateAuthDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/clusters/{id.value}/rotate-auth][%d] ClusterService_RotateAuth default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterServiceRotateAuthDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/clusters/{id.value}/rotate-auth][%d] ClusterService_RotateAuth default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterServiceRotateAuthDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ClusterServiceRotateAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
