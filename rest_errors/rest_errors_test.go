package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("invalid parameters")
	assert.Equal(t, "invalid parameters", err.Message)
	assert.Equal(t, "bad_request", err.Error)
	assert.Equal(t, http.StatusBadRequest, err.Status)
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("internal service error", errors.New("db error"))
	assert.Equal(t, "internal service error", err.Message)
	assert.Equal(t, 1, len(err.Causes))
	assert.Equal(t, "db error", err.Causes[0])
	assert.Equal(t, "internal_server_error", err.Error)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
}