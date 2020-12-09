// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// UpdateClusterInstallProgressReader is a Reader for the UpdateClusterInstallProgress structure.
type UpdateClusterInstallProgressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterInstallProgressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateClusterInstallProgressNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateClusterInstallProgressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateClusterInstallProgressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateClusterInstallProgressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateClusterInstallProgressMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateClusterInstallProgressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateClusterInstallProgressServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterInstallProgressNoContent creates a UpdateClusterInstallProgressNoContent with default headers values
func NewUpdateClusterInstallProgressNoContent() *UpdateClusterInstallProgressNoContent {
	return &UpdateClusterInstallProgressNoContent{}
}

/*UpdateClusterInstallProgressNoContent handles this case with default header values.

Update cluster install progress.
*/
type UpdateClusterInstallProgressNoContent struct {
}

func (o *UpdateClusterInstallProgressNoContent) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/progress][%d] updateClusterInstallProgressNoContent ", 204)
}

func (o *UpdateClusterInstallProgressNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClusterInstallProgressUnauthorized creates a UpdateClusterInstallProgressUnauthorized with default headers values
func NewUpdateClusterInstallProgressUnauthorized() *UpdateClusterInstallProgressUnauthorized {
	return &UpdateClusterInstallProgressUnauthorized{}
}

/*UpdateClusterInstallProgressUnauthorized handles this case with default header values.

Unauthorized.
*/
type UpdateClusterInstallProgressUnauthorized struct {
	Payload *models.InfraError
}

func (o *UpdateClusterInstallProgressUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/progress][%d] updateClusterInstallProgressUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateClusterInstallProgressUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateClusterInstallProgressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInstallProgressForbidden creates a UpdateClusterInstallProgressForbidden with default headers values
func NewUpdateClusterInstallProgressForbidden() *UpdateClusterInstallProgressForbidden {
	return &UpdateClusterInstallProgressForbidden{}
}

/*UpdateClusterInstallProgressForbidden handles this case with default header values.

Forbidden.
*/
type UpdateClusterInstallProgressForbidden struct {
	Payload *models.InfraError
}

func (o *UpdateClusterInstallProgressForbidden) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/progress][%d] updateClusterInstallProgressForbidden  %+v", 403, o.Payload)
}

func (o *UpdateClusterInstallProgressForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateClusterInstallProgressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInstallProgressNotFound creates a UpdateClusterInstallProgressNotFound with default headers values
func NewUpdateClusterInstallProgressNotFound() *UpdateClusterInstallProgressNotFound {
	return &UpdateClusterInstallProgressNotFound{}
}

/*UpdateClusterInstallProgressNotFound handles this case with default header values.

Error.
*/
type UpdateClusterInstallProgressNotFound struct {
	Payload *models.Error
}

func (o *UpdateClusterInstallProgressNotFound) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/progress][%d] updateClusterInstallProgressNotFound  %+v", 404, o.Payload)
}

func (o *UpdateClusterInstallProgressNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterInstallProgressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInstallProgressMethodNotAllowed creates a UpdateClusterInstallProgressMethodNotAllowed with default headers values
func NewUpdateClusterInstallProgressMethodNotAllowed() *UpdateClusterInstallProgressMethodNotAllowed {
	return &UpdateClusterInstallProgressMethodNotAllowed{}
}

/*UpdateClusterInstallProgressMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type UpdateClusterInstallProgressMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateClusterInstallProgressMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/progress][%d] updateClusterInstallProgressMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateClusterInstallProgressMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterInstallProgressMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInstallProgressInternalServerError creates a UpdateClusterInstallProgressInternalServerError with default headers values
func NewUpdateClusterInstallProgressInternalServerError() *UpdateClusterInstallProgressInternalServerError {
	return &UpdateClusterInstallProgressInternalServerError{}
}

/*UpdateClusterInstallProgressInternalServerError handles this case with default header values.

Error.
*/
type UpdateClusterInstallProgressInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateClusterInstallProgressInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/progress][%d] updateClusterInstallProgressInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateClusterInstallProgressInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterInstallProgressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInstallProgressServiceUnavailable creates a UpdateClusterInstallProgressServiceUnavailable with default headers values
func NewUpdateClusterInstallProgressServiceUnavailable() *UpdateClusterInstallProgressServiceUnavailable {
	return &UpdateClusterInstallProgressServiceUnavailable{}
}

/*UpdateClusterInstallProgressServiceUnavailable handles this case with default header values.

Unavailable.
*/
type UpdateClusterInstallProgressServiceUnavailable struct {
	Payload *models.Error
}

func (o *UpdateClusterInstallProgressServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /clusters/{cluster_id}/progress][%d] updateClusterInstallProgressServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateClusterInstallProgressServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterInstallProgressServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
