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

// ClusterServiceGetReader is a Reader for the ClusterServiceGet structure.
type ClusterServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterServiceGetOK creates a ClusterServiceGetOK with default headers values
func NewClusterServiceGetOK() *ClusterServiceGetOK {
	return &ClusterServiceGetOK{}
}

/* ClusterServiceGetOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterServiceGetOK struct {
	Payload *models.V1alpha1Cluster
}

// IsSuccess returns true when this cluster service get o k response has a 2xx status code
func (o *ClusterServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster service get o k response has a 3xx status code
func (o *ClusterServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster service get o k response has a 4xx status code
func (o *ClusterServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster service get o k response has a 5xx status code
func (o *ClusterServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster service get o k response a status code equal to that given
func (o *ClusterServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{id.value}][%d] clusterServiceGetOK  %+v", 200, o.Payload)
}

func (o *ClusterServiceGetOK) String() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{id.value}][%d] clusterServiceGetOK  %+v", 200, o.Payload)
}

func (o *ClusterServiceGetOK) GetPayload() *models.V1alpha1Cluster {
	return o.Payload
}

func (o *ClusterServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterServiceGetDefault creates a ClusterServiceGetDefault with default headers values
func NewClusterServiceGetDefault(code int) *ClusterServiceGetDefault {
	return &ClusterServiceGetDefault{
		_statusCode: code,
	}
}

/* ClusterServiceGetDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterServiceGetDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the cluster service get default response
func (o *ClusterServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster service get default response has a 2xx status code
func (o *ClusterServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster service get default response has a 3xx status code
func (o *ClusterServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster service get default response has a 4xx status code
func (o *ClusterServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster service get default response has a 5xx status code
func (o *ClusterServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster service get default response a status code equal to that given
func (o *ClusterServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{id.value}][%d] ClusterService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/clusters/{id.value}][%d] ClusterService_Get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterServiceGetDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ClusterServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
