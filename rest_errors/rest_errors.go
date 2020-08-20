package rest_errors

import (
	"fmt"
	"net/http"
)

// RestErr rest error interface
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - error: %s - causes [%v]",
		e.ErrMessage,
		e.ErrStatus,
		e.ErrError,
		e.ErrCauses,
	)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

// NewRestError new custom rest error
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

// NewBadRequestError bad request error status:400
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

// NewNotFounfError not found error - status: 404
func NewNotFounfError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

// NewUnauthorizedError unauthorized error - status: 401
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

// NewInternalServerError internal error - status: 500
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}

// NewExternalServiceError external call failed - status: 424
func NewExternalServiceError(message string, err error) RestErr {
	e := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusFailedDependency,
		ErrError:   "failed_dependency",
	}
	if err != nil {
		e.ErrCauses = append(e.ErrCauses, err.Error())
	}
	return e
}

// NewNotImplementedError service not implemented yet - status: 501
func NewNotImplementedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotImplemented,
		ErrError:   "not_implemented",
	}
}
