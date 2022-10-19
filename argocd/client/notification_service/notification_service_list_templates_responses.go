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

// NotificationServiceListTemplatesReader is a Reader for the NotificationServiceListTemplates structure.
type NotificationServiceListTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationServiceListTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationServiceListTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNotificationServiceListTemplatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNotificationServiceListTemplatesOK creates a NotificationServiceListTemplatesOK with default headers values
func NewNotificationServiceListTemplatesOK() *NotificationServiceListTemplatesOK {
	return &NotificationServiceListTemplatesOK{}
}

/* NotificationServiceListTemplatesOK describes a response with status code 200, with default header values.

A successful response.
*/
type NotificationServiceListTemplatesOK struct {
	Payload *models.NotificationTemplateList
}

// IsSuccess returns true when this notification service list templates o k response has a 2xx status code
func (o *NotificationServiceListTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notification service list templates o k response has a 3xx status code
func (o *NotificationServiceListTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notification service list templates o k response has a 4xx status code
func (o *NotificationServiceListTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notification service list templates o k response has a 5xx status code
func (o *NotificationServiceListTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notification service list templates o k response a status code equal to that given
func (o *NotificationServiceListTemplatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *NotificationServiceListTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications/templates][%d] notificationServiceListTemplatesOK  %+v", 200, o.Payload)
}

func (o *NotificationServiceListTemplatesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications/templates][%d] notificationServiceListTemplatesOK  %+v", 200, o.Payload)
}

func (o *NotificationServiceListTemplatesOK) GetPayload() *models.NotificationTemplateList {
	return o.Payload
}

func (o *NotificationServiceListTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationTemplateList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationServiceListTemplatesDefault creates a NotificationServiceListTemplatesDefault with default headers values
func NewNotificationServiceListTemplatesDefault(code int) *NotificationServiceListTemplatesDefault {
	return &NotificationServiceListTemplatesDefault{
		_statusCode: code,
	}
}

/* NotificationServiceListTemplatesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type NotificationServiceListTemplatesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the notification service list templates default response
func (o *NotificationServiceListTemplatesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this notification service list templates default response has a 2xx status code
func (o *NotificationServiceListTemplatesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this notification service list templates default response has a 3xx status code
func (o *NotificationServiceListTemplatesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this notification service list templates default response has a 4xx status code
func (o *NotificationServiceListTemplatesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this notification service list templates default response has a 5xx status code
func (o *NotificationServiceListTemplatesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this notification service list templates default response a status code equal to that given
func (o *NotificationServiceListTemplatesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NotificationServiceListTemplatesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/notifications/templates][%d] NotificationService_ListTemplates default  %+v", o._statusCode, o.Payload)
}

func (o *NotificationServiceListTemplatesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/notifications/templates][%d] NotificationService_ListTemplates default  %+v", o._statusCode, o.Payload)
}

func (o *NotificationServiceListTemplatesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *NotificationServiceListTemplatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}