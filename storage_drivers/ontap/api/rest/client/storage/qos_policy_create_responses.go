// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// QosPolicyCreateReader is a Reader for the QosPolicyCreate structure.
type QosPolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewQosPolicyCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyCreateAccepted creates a QosPolicyCreateAccepted with default headers values
func NewQosPolicyCreateAccepted() *QosPolicyCreateAccepted {
	return &QosPolicyCreateAccepted{}
}

/*
QosPolicyCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QosPolicyCreateAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this qos policy create accepted response has a 2xx status code
func (o *QosPolicyCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy create accepted response has a 3xx status code
func (o *QosPolicyCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy create accepted response has a 4xx status code
func (o *QosPolicyCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy create accepted response has a 5xx status code
func (o *QosPolicyCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy create accepted response a status code equal to that given
func (o *QosPolicyCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *QosPolicyCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qosPolicyCreateAccepted  %+v", 202, o.Payload)
}

func (o *QosPolicyCreateAccepted) String() string {
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qosPolicyCreateAccepted  %+v", 202, o.Payload)
}

func (o *QosPolicyCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QosPolicyCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyCreateDefault creates a QosPolicyCreateDefault with default headers values
func NewQosPolicyCreateDefault(code int) *QosPolicyCreateDefault {
	return &QosPolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	QosPolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8454147 | The maximum limit for QoS policies has been reached. |
| 8454154 | The name specified for creating conflicts with an existing QoS policy name. |
| 8454260 | Invalid value for maximum and minimum fields. Valid values for max_throughput_iops and max_throughput_mbps combination is for the ratio of max_throughput_mbps and max_throughput_iops to be within 1 to 4096. |
| 8454273 | Invalid value for an adaptive field. Value should be non-zero. |
| 8454277 | The name specified for creating an adaptive QoS policy conflicts with an existing fixed QoS policy name. |
| 8454278 | The name specified for creating a fixed QoS policy conflicts with an existing adaptive QoS policy name. |
*/
type QosPolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the qos policy create default response
func (o *QosPolicyCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this qos policy create default response has a 2xx status code
func (o *QosPolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qos policy create default response has a 3xx status code
func (o *QosPolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qos policy create default response has a 4xx status code
func (o *QosPolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qos policy create default response has a 5xx status code
func (o *QosPolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qos policy create default response a status code equal to that given
func (o *QosPolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QosPolicyCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qos_policy_create default  %+v", o._statusCode, o.Payload)
}

func (o *QosPolicyCreateDefault) String() string {
	return fmt.Sprintf("[POST /storage/qos/policies][%d] qos_policy_create default  %+v", o._statusCode, o.Payload)
}

func (o *QosPolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
