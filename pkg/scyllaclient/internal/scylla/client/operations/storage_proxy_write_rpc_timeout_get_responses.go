// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// StorageProxyWriteRPCTimeoutGetReader is a Reader for the StorageProxyWriteRPCTimeoutGet structure.
type StorageProxyWriteRPCTimeoutGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyWriteRPCTimeoutGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyWriteRPCTimeoutGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyWriteRPCTimeoutGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyWriteRPCTimeoutGetOK creates a StorageProxyWriteRPCTimeoutGetOK with default headers values
func NewStorageProxyWriteRPCTimeoutGetOK() *StorageProxyWriteRPCTimeoutGetOK {
	return &StorageProxyWriteRPCTimeoutGetOK{}
}

/*
StorageProxyWriteRPCTimeoutGetOK handles this case with default header values.

StorageProxyWriteRPCTimeoutGetOK storage proxy write Rpc timeout get o k
*/
type StorageProxyWriteRPCTimeoutGetOK struct {
	Payload interface{}
}

func (o *StorageProxyWriteRPCTimeoutGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StorageProxyWriteRPCTimeoutGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyWriteRPCTimeoutGetDefault creates a StorageProxyWriteRPCTimeoutGetDefault with default headers values
func NewStorageProxyWriteRPCTimeoutGetDefault(code int) *StorageProxyWriteRPCTimeoutGetDefault {
	return &StorageProxyWriteRPCTimeoutGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyWriteRPCTimeoutGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyWriteRPCTimeoutGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy write Rpc timeout get default response
func (o *StorageProxyWriteRPCTimeoutGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyWriteRPCTimeoutGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyWriteRPCTimeoutGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyWriteRPCTimeoutGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
