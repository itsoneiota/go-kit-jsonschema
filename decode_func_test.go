package jsonschema

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDecodeFuncPassThrough(t *testing.T) {
	funcCalled := false
	myFunc := func(req *http.Request) (request interface{}, err error) {
		funcCalled = true
		return "pass", fmt.Errorf("notanerror")
	}

	val := new(PassingValidator)
	wrapped := NewDecodeFunc(val, myFunc)

	req, _ := http.NewRequest("GET", "dummyURL", nil)
	result, err := wrapped(req)

	if !funcCalled {
		t.Errorf("My decode function should have been called. It wasn't.")
	}
	if result != "pass" {
		t.Errorf("Should have returned result from wrapped function. Didn't.")
	}
	if err == nil {
		t.Errorf("Should have retured the error from the wrapped function. Didn't.")
	} else if err.Error() != "notanerror" {
		t.Errorf("Should have retured the error from the wrapped function. Returned a different one: %s", err)
	}

}

func TestCanBlockDecodeOnFailure(t *testing.T) {
	funcCalled := false
	myFunc := func(req *http.Request) (request interface{}, err error) {
		funcCalled = true
		return "pass", nil
	}

	val := new(FailingValidator)
	wrapped := NewDecodeFunc(val, myFunc)

	req, _ := http.NewRequest("GET", "dummyURL", nil)
	result, err := wrapped(req)

	if funcCalled {
		t.Errorf("My decode function should not have been called. It was.")
	}
	if result != nil {
		t.Errorf("Should have returned nil. Didn't.")
	}
	if err == nil {
		t.Errorf("Should have retured a validation error. Didn't.")
	}

}
