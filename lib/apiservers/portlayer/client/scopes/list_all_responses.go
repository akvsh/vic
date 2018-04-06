package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// ListAllReader is a Reader for the ListAll structure.
type ListAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAllOK creates a ListAllOK with default headers values
func NewListAllOK() *ListAllOK {
	return &ListAllOK{}
}

/*ListAllOK handles this case with default header values.

OK
*/
type ListAllOK struct {
	Payload []*models.ScopeConfig
}

func (o *ListAllOK) Error() string {
	return fmt.Sprintf("[GET /scopes][%d] listAllOK  %+v", 200, o.Payload)
}

func (o *ListAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllDefault creates a ListAllDefault with default headers values
func NewListAllDefault(code int) *ListAllDefault {
	return &ListAllDefault{
		_statusCode: code,
	}
}

/*ListAllDefault handles this case with default header values.

error
*/
type ListAllDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list all default response
func (o *ListAllDefault) Code() int {
	return o._statusCode
}

func (o *ListAllDefault) Error() string {
	return fmt.Sprintf("[GET /scopes][%d] ListAll default  %+v", o._statusCode, o.Payload)
}

func (o *ListAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
