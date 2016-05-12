package jsonschema

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// Validator validates JSON documents.
type Validator interface {
	// Validate validates a decoded JSON document.
	Validate(doc interface{}) (valid bool, err error)

	// ValidateString validates a JSON string.
	ValidateString(doc string) (valid bool, err error)
}

// SchemaValidator is a JSON validator fixed with a given schema.
// This effectively allows us to partially apply the gojsonschema.Validate()
// function with the schema.
type SchemaValidator struct {
	// This loader defines the schema to be used.
	schemaLoader    gojsonschema.JSONLoader
	validationError error
}

// Validate validates the given document against the schema.
func (val *SchemaValidator) Validate(doc interface{}) (valid bool, err error) {
	documentLoader := gojsonschema.NewGoLoader(doc)
	return val.validate(documentLoader)
}

// ValidateString validates the given string document against the schema.
func (val *SchemaValidator) ValidateString(doc string) (valid bool, err error) {
	documentLoader := gojsonschema.NewStringLoader(doc)
	return val.validate(documentLoader)
}

func (val *SchemaValidator) validate(documentLoader gojsonschema.JSONLoader) (valid bool, err error) {
	result, err := gojsonschema.Validate(val.schemaLoader, documentLoader)
	if err != nil {
		return false, err
	}

	valid = result.Valid()

	if !valid {
		err = fmt.Errorf("input of %v failed JSON schema validation: %v", documentLoader, result.Errors())
	}

	return
}

// NewValidatorFromFile returns a new validator for the given schema file.
func NewValidatorFromFile(schemaFile string) Validator {
	val := new(SchemaValidator)
	val.schemaLoader = gojsonschema.NewReferenceLoader(schemaFile)
	return val
}

// NewValidatorFromString returns a new validator for the given schema file.
func NewValidatorFromString(schema string) Validator {
	val := new(SchemaValidator)
	val.schemaLoader = gojsonschema.NewStringLoader(schema)
	return val
}
