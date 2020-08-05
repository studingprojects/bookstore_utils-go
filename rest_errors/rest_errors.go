package rest_errors

import "net/http"

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFounfError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func NewExternalServiceError(message string, err error) *RestErr {
	e := &RestErr{
		Message: message,
		Status:  http.StatusFailedDependency,
		Error:   "failed_dependency",
	}
	if err != nil {
		e.Causes = append(e.Causes, err.Error())
	}
	return e
}

func NewNotImplementedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotImplemented,
		Error:   "not_implemented",
	}
}
