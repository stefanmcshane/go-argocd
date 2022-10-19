// Code generated by go-swagger; DO NOT EDIT.

package notification_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// NotificationServiceListTriggersReader is a Reader for the NotificationServiceListTriggers structure.
type NotificationServiceListTriggersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationServiceListTriggersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationServiceListTriggersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNotificationServiceListTriggersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNotificationServiceListTriggersOK creates a NotificationServiceListTriggersOK with default headers values
func NewNotificationServiceListTriggersOK() *NotificationServiceListTriggersOK {
	return &NotificationServiceListTriggersOK{}
}

/* NotificationServiceListTriggersOK describes a response with status code 200, with default header values.

A successful response.
*/
type NotificationServiceListTriggersOK struct {
	Payload *models.NotificationTriggerList
}

// IsSuccess returns true when this notification service list triggers o k response has a 2xx status code
func (o *NotificationServiceListTriggersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notification service list triggers o k response has a 3xx status code
func (o *NotificationServiceListTriggersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification service list triggers o k response has a 4xx status code
func (o *NotificationServiceListTriggersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification service list triggers o k response has a 5xx status code
func (o *NotificationServiceListTriggersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notification service list triggers o k response a status code equal to that given
func (o *NotificationServiceListTriggersOK) IsCode(code int) bool {
	return code == 200
}

func (o *NotificationServiceListTriggersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications/triggers][%d] notificationServiceListTriggersOK  %+v", 200, o.Payload)
}

func (o *NotificationServiceListTriggersOK) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications/triggers][%d] notificationServiceListTriggersOK  %+v", 200, o.Payload)
}

func (o *NotificationServiceListTriggersOK) GetPayload() *models.NotificationTriggerList {
	return o.Payload
}

func (o *NotificationServiceListTriggersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationTriggerList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationServiceListTriggersDefault creates a NotificationServiceListTriggersDefault with default headers values
func NewNotificationServiceListTriggersDefault(code int) *NotificationServiceListTriggersDefault {
	return &NotificationServiceListTriggersDefault{
		_statusCode: code,
	}
}

/* NotificationServiceListTriggersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type NotificationServiceListTriggersDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the notification service list triggers default response
func (o *NotificationServiceListTriggersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this notification service list triggers default response has a 2xx status code
func (o *NotificationServiceListTriggersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this notification service list triggers default response has a 3xx status code
func (o *NotificationServiceListTriggersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this notification service list triggers default response has a 4xx status code
func (o *NotificationServiceListTriggersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this notification service list triggers default response has a 5xx status code
func (o *NotificationServiceListTriggersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this notification service list triggers default response a status code equal to that given
func (o *NotificationServiceListTriggersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NotificationServiceListTriggersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications/triggers][%d] NotificationService_ListTriggers default  %+v", o._statusCode, o.Payload)
}

func (o *NotificationServiceListTriggersDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications/triggers][%d] NotificationService_ListTriggers default  %+v", o._statusCode, o.Payload)
}

func (o *NotificationServiceListTriggersDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *NotificationServiceListTriggersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}