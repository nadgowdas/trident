// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ClusterCreateReader is a Reader for the ClusterCreate structure.
type ClusterCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewClusterCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterCreateAccepted creates a ClusterCreateAccepted with default headers values
func NewClusterCreateAccepted() *ClusterCreateAccepted {
	return &ClusterCreateAccepted{}
}

/*
ClusterCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ClusterCreateAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cluster create accepted response has a 2xx status code
func (o *ClusterCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster create accepted response has a 3xx status code
func (o *ClusterCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster create accepted response has a 4xx status code
func (o *ClusterCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster create accepted response has a 5xx status code
func (o *ClusterCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster create accepted response a status code equal to that given
func (o *ClusterCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ClusterCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /cluster][%d] clusterCreateAccepted  %+v", 202, o.Payload)
}

func (o *ClusterCreateAccepted) String() string {
	return fmt.Sprintf("[POST /cluster][%d] clusterCreateAccepted  %+v", 202, o.Payload)
}

func (o *ClusterCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ClusterCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterCreateDefault creates a ClusterCreateDefault with default headers values
func NewClusterCreateDefault(code int) *ClusterCreateDefault {
	return &ClusterCreateDefault{
		_statusCode: code,
	}
}

/*
	ClusterCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262245 | The value provided is invalid. |
| 1179813 | Fields set for one node must be set for all nodes. |
| 1179817 | The IP address, subnet mask, and gateway must all be provided for cluster manangement interface. |
| 1179818 | The IP address and gateway must be of the same family. |
| 1179821 | An IP address and subnet mask conflicts with an existing entry. |
| 1179824 | An invalid gateway was provided. |
| 1179825 | All management and cluster config IP addresses must belong to the same address family. |
| 2097165 | An NTP server could not be reached. |
| 8847361 | Too many DNS domains provided. |
| 8847362 | Too many name servers provided. |
| 8847394 | An invalid DNS domain was provided. |
| 8978433 | An invalid license key was provided. |
| 9240587 | A name must be provided. |
| 9240594 | An invalid name was provided. |
| 39387137 | The URL provided is invalid. |
| 131727360 | A node could not be added to the cluster. This is a generic code, see response message for details. |
| 131727388 | Hostnames for NTP servers cannot be used without DNS configured. |
| 131727389 | URL and username are required for configuration backup. |
*/
type ClusterCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster create default response
func (o *ClusterCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster create default response has a 2xx status code
func (o *ClusterCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster create default response has a 3xx status code
func (o *ClusterCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster create default response has a 4xx status code
func (o *ClusterCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster create default response has a 5xx status code
func (o *ClusterCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster create default response a status code equal to that given
func (o *ClusterCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterCreateDefault) Error() string {
	return fmt.Sprintf("[POST /cluster][%d] cluster_create default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterCreateDefault) String() string {
	return fmt.Sprintf("[POST /cluster][%d] cluster_create default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
