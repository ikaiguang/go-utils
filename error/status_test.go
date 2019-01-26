package goerror

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(http.StatusNotFound, http.StatusText(http.StatusNotFound))

	t.Log(err)
}

func TestNewWithError(t *testing.T) {
	oldErr := New(http.StatusNotFound, http.StatusText(http.StatusNotFound))

	newErr := NewWithError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), oldErr)

	t.Log(newErr)

}

func TestFromError(t *testing.T) {
	err := New(http.StatusNotFound, http.StatusText(http.StatusNotFound))

	status, ok := FromError(err)
	if !ok {
		t.Errorf("FromError(err) error : err isn't produced from this package")
		return
	}

	if status.Code != http.StatusNotFound {
		t.Errorf("FromError(err) error : status.Code != input.Code")
	}
	t.Logf("\n%v", err)
}
