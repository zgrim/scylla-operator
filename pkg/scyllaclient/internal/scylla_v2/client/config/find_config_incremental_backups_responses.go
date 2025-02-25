// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigIncrementalBackupsReader is a Reader for the FindConfigIncrementalBackups structure.
type FindConfigIncrementalBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigIncrementalBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigIncrementalBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigIncrementalBackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigIncrementalBackupsOK creates a FindConfigIncrementalBackupsOK with default headers values
func NewFindConfigIncrementalBackupsOK() *FindConfigIncrementalBackupsOK {
	return &FindConfigIncrementalBackupsOK{}
}

/*
FindConfigIncrementalBackupsOK handles this case with default header values.

Config value
*/
type FindConfigIncrementalBackupsOK struct {
	Payload bool
}

func (o *FindConfigIncrementalBackupsOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigIncrementalBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigIncrementalBackupsDefault creates a FindConfigIncrementalBackupsDefault with default headers values
func NewFindConfigIncrementalBackupsDefault(code int) *FindConfigIncrementalBackupsDefault {
	return &FindConfigIncrementalBackupsDefault{
		_statusCode: code,
	}
}

/*
FindConfigIncrementalBackupsDefault handles this case with default header values.

unexpected error
*/
type FindConfigIncrementalBackupsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config incremental backups default response
func (o *FindConfigIncrementalBackupsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigIncrementalBackupsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigIncrementalBackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigIncrementalBackupsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
