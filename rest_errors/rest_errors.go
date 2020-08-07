package rest_errors

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - error: %s - causes [%v]",
		e.Message(),
		e.Status(),
		e.Error(),
		e.Causes(),
	)
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) Status() int {
	return e.status
}

func (e restErr) Causes() []interface{} {
	return e.causes
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFounfError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}

func NewExternalServiceError(message string, err error) RestErr {
	e := restErr{
		message: message,
		status:  http.StatusFailedDependency,
		error:   "failed_dependency",
	}
	if err != nil {
		e.causes = append(e.causes, err.Error())
	}
	return e
}

func NewNotImplementedError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotImplemented,
		error:   "not_implemented",
	}
}
