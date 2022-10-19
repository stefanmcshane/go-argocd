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

// ClusterServiceUpdateReader is a Reader for the ClusterServiceUpdate structure.
type ClusterServiceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterServiceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterServiceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterServiceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterServiceUpdateOK creates a ClusterServiceUpdateOK with default headers values
func NewClusterServiceUpdateOK() *ClusterServiceUpdateOK {
	return &ClusterServiceUpdateOK{}
}

/* ClusterServiceUpdateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterServiceUpdateOK struct {
	Payload *models.V1alpha1Cluster
}

// IsSuccess returns true when this cluster service update o k response has a 2xx status code
func (o *ClusterServiceUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster service update o k response has a 3xx status code
func (o *ClusterServiceUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster service update o k response has a 4xx status code
func (o *ClusterServiceUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster service update o k response has a 5xx status code
func (o *ClusterServiceUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster service update o k response a status code equal to that given
func (o *ClusterServiceUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterServiceUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/clusters/{id.value}][%d] clusterServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *ClusterServiceUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/clusters/{id.value}][%d] clusterServiceUpdateOK  %+v", 200, o.Payload)
}

func (o *ClusterServiceUpdateOK) GetPayload() *models.V1alpha1Cluster {
	return o.Payload
}

func (o *ClusterServiceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterServiceUpdateDefault creates a ClusterServiceUpdateDefault with default headers values
func NewClusterServiceUpdateDefault(code int) *ClusterServiceUpdateDefault {
	return &ClusterServiceUpdateDefault{
		_statusCode: code,
	}
}

/* ClusterServiceUpdateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterServiceUpdateDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the cluster service update default response
func (o *ClusterServiceUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster service update default response has a 2xx status code
func (o *ClusterServiceUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster service update default response has a 3xx status code
func (o *ClusterServiceUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster service update default response has a 4xx status code
func (o *ClusterServiceUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster service update default response has a 5xx status code
func (o *ClusterServiceUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster service update default response a status code equal to that given
func (o *ClusterServiceUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterServiceUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/clusters/{id.value}][%d] ClusterService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterServiceUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/clusters/{id.value}][%d] ClusterService_Update default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterServiceUpdateDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *ClusterServiceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
