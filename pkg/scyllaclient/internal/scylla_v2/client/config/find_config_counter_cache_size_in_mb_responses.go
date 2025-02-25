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

// FindConfigCounterCacheSizeInMbReader is a Reader for the FindConfigCounterCacheSizeInMb structure.
type FindConfigCounterCacheSizeInMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCounterCacheSizeInMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCounterCacheSizeInMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCounterCacheSizeInMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCounterCacheSizeInMbOK creates a FindConfigCounterCacheSizeInMbOK with default headers values
func NewFindConfigCounterCacheSizeInMbOK() *FindConfigCounterCacheSizeInMbOK {
	return &FindConfigCounterCacheSizeInMbOK{}
}

/*
FindConfigCounterCacheSizeInMbOK handles this case with default header values.

Config value
*/
type FindConfigCounterCacheSizeInMbOK struct {
	Payload int64
}

func (o *FindConfigCounterCacheSizeInMbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCounterCacheSizeInMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCounterCacheSizeInMbDefault creates a FindConfigCounterCacheSizeInMbDefault with default headers values
func NewFindConfigCounterCacheSizeInMbDefault(code int) *FindConfigCounterCacheSizeInMbDefault {
	return &FindConfigCounterCacheSizeInMbDefault{
		_statusCode: code,
	}
}

/*
FindConfigCounterCacheSizeInMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigCounterCacheSizeInMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config counter cache size in mb default response
func (o *FindConfigCounterCacheSizeInMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCounterCacheSizeInMbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCounterCacheSizeInMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCounterCacheSizeInMbDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
