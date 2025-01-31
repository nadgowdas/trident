// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockLegalHoldBeginReader is a Reader for the SnaplockLegalHoldBegin structure.
type SnaplockLegalHoldBeginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldBeginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnaplockLegalHoldBeginCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldBeginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldBeginCreated creates a SnaplockLegalHoldBeginCreated with default headers values
func NewSnaplockLegalHoldBeginCreated() *SnaplockLegalHoldBeginCreated {
	return &SnaplockLegalHoldBeginCreated{}
}

/*
SnaplockLegalHoldBeginCreated describes a response with status code 201, with default header values.

Created
*/
type SnaplockLegalHoldBeginCreated struct {
	Payload *models.SnaplockLitigation
}

// IsSuccess returns true when this snaplock legal hold begin created response has a 2xx status code
func (o *SnaplockLegalHoldBeginCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock legal hold begin created response has a 3xx status code
func (o *SnaplockLegalHoldBeginCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock legal hold begin created response has a 4xx status code
func (o *SnaplockLegalHoldBeginCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock legal hold begin created response has a 5xx status code
func (o *SnaplockLegalHoldBeginCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock legal hold begin created response a status code equal to that given
func (o *SnaplockLegalHoldBeginCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SnaplockLegalHoldBeginCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations][%d] snaplockLegalHoldBeginCreated  %+v", 201, o.Payload)
}

func (o *SnaplockLegalHoldBeginCreated) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations][%d] snaplockLegalHoldBeginCreated  %+v", 201, o.Payload)
}

func (o *SnaplockLegalHoldBeginCreated) GetPayload() *models.SnaplockLitigation {
	return o.Payload
}

func (o *SnaplockLegalHoldBeginCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLitigation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLegalHoldBeginDefault creates a SnaplockLegalHoldBeginDefault with default headers values
func NewSnaplockLegalHoldBeginDefault(code int) *SnaplockLegalHoldBeginDefault {
	return &SnaplockLegalHoldBeginDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLegalHoldBeginDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
| 14090340    | {field} is a required field  |
| 14090343    | Invalid Field  |
| 14090641    | The specified volume.name and volume.uuid refer to different volumes  |
*/
type SnaplockLegalHoldBeginDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock legal hold begin default response
func (o *SnaplockLegalHoldBeginDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snaplock legal hold begin default response has a 2xx status code
func (o *SnaplockLegalHoldBeginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock legal hold begin default response has a 3xx status code
func (o *SnaplockLegalHoldBeginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock legal hold begin default response has a 4xx status code
func (o *SnaplockLegalHoldBeginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock legal hold begin default response has a 5xx status code
func (o *SnaplockLegalHoldBeginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock legal hold begin default response a status code equal to that given
func (o *SnaplockLegalHoldBeginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnaplockLegalHoldBeginDefault) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations][%d] snaplock_legal_hold_begin default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldBeginDefault) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations][%d] snaplock_legal_hold_begin default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldBeginDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldBeginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
