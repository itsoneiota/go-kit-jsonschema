package jsonschema

import "fmt"

// PassingValidator passes for everything.
type PassingValidator bool

// Validate passes. Always
func (val *PassingValidator) Validate(doc interface{}) (valid bool, err error) {
	return true, nil
}

// ValidateString passes. Always
func (val *PassingValidator) ValidateString(doc string) (valid bool, err error) {
	return true, nil
}

// FailingValidator fails for everything.
type FailingValidator bool

// Validate fails. Always
func (val *FailingValidator) Validate(doc interface{}) (valid bool, err error) {
	return false, fmt.Errorf("JSON schema validation failure")
}

// ValidateString fails. Always
func (val *FailingValidator) ValidateString(doc string) (valid bool, err error) {
	return false, fmt.Errorf("JSON schema validation failure")
}
