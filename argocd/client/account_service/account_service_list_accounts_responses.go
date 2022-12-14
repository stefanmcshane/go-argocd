// Code generated by go-swagger; DO NOT EDIT.

package account_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stefanmcshane/go-argocd/argocd/models"
)

// AccountServiceListAccountsReader is a Reader for the AccountServiceListAccounts structure.
type AccountServiceListAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountServiceListAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountServiceListAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountServiceListAccountsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountServiceListAccountsOK creates a AccountServiceListAccountsOK with default headers values
func NewAccountServiceListAccountsOK() *AccountServiceListAccountsOK {
	return &AccountServiceListAccountsOK{}
}

/* AccountServiceListAccountsOK describes a response with status code 200, with default header values.

A successful response.
*/
type AccountServiceListAccountsOK struct {
	Payload *models.AccountAccountsList
}

// IsSuccess returns true when this account service list accounts o k response has a 2xx status code
func (o *AccountServiceListAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account service list accounts o k response has a 3xx status code
func (o *AccountServiceListAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account service list accounts o k response has a 4xx status code
func (o *AccountServiceListAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account service list accounts o k response has a 5xx status code
func (o *AccountServiceListAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account service list accounts o k response a status code equal to that given
func (o *AccountServiceListAccountsOK) IsCode(code int) bool {
	return code == 200
}

func (o *AccountServiceListAccountsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/account][%d] accountServiceListAccountsOK  %+v", 200, o.Payload)
}

func (o *AccountServiceListAccountsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/account][%d] accountServiceListAccountsOK  %+v", 200, o.Payload)
}

func (o *AccountServiceListAccountsOK) GetPayload() *models.AccountAccountsList {
	return o.Payload
}

func (o *AccountServiceListAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountAccountsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountServiceListAccountsDefault creates a AccountServiceListAccountsDefault with default headers values
func NewAccountServiceListAccountsDefault(code int) *AccountServiceListAccountsDefault {
	return &AccountServiceListAccountsDefault{
		_statusCode: code,
	}
}

/* AccountServiceListAccountsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AccountServiceListAccountsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the account service list accounts default response
func (o *AccountServiceListAccountsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this account service list accounts default response has a 2xx status code
func (o *AccountServiceListAccountsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account service list accounts default response has a 3xx status code
func (o *AccountServiceListAccountsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account service list accounts default response has a 4xx status code
func (o *AccountServiceListAccountsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account service list accounts default response has a 5xx status code
func (o *AccountServiceListAccountsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account service list accounts default response a status code equal to that given
func (o *AccountServiceListAccountsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AccountServiceListAccountsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/account][%d] AccountService_ListAccounts default  %+v", o._statusCode, o.Payload)
}

func (o *AccountServiceListAccountsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/account][%d] AccountService_ListAccounts default  %+v", o._statusCode, o.Payload)
}

func (o *AccountServiceListAccountsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AccountServiceListAccountsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
