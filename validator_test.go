package jsonschema

import "testing"

func TestBasicValidation(t *testing.T) {
	val := NewValidatorFromString(`
		{
			"type": "integer"
		}
	`)

	ok, err := val.Validate(1)
	if !ok {
		t.Errorf("Value should pass validation. Didn't.")
	}
	if err != nil {
		t.Errorf("Shouldn't have an error. Got one: %s", err)
	}

}

func TestCanValidate(t *testing.T) {
	// This isn't intended to be an exhaustive test of schema validation,
	// we're just checking the plumbing.
	var tests = []struct {
		schema string
		input  interface{}
		want   bool
	}{
		{`{"type":"string"}`, "foo", true},
		{`{"type":"string"}`, "8", true},
		{`{"type":"string"}`, 8, false},
	}
	for _, test := range tests {
		val := NewValidatorFromString(test.schema)
		if got, _ := val.Validate(test.input); got != test.want {
			t.Errorf("wanted %v, got %v", test.want, got)
		}
	}
}

func TestCanValidateString(t *testing.T) {
	// This isn't intended to be an exhaustive test of schema validation,
	// we're just checking the plumbing.
	var tests = []struct {
		schema string
		doc    string
		want   bool
	}{
		{`{"type":"string"}`, `"foo"`, true},
		{`{"type":"string"}`, `"8"`, true},
		{`{"type":"string"}`, `8`, false},
		{`{"type":"object","properties":{"foo":{"type":"integer"}}}`, `{"foo":666}`, true},
		{`{"type":"object","properties":{"foo":{"type":"integer"}}}`, `{"foo":"666"}`, false},
		{`{"type":"object","properties":{"foo":{"type":"integer"}}}`, `{"bar":123}`, true},
		{`{"type":"object","required":["foo"],"properties":{"foo":{"type":"integer"}}}`, `{"bar":123}`, false},
		{`{"type":"object","required":["foo"],"properties":{"foo":{"type":"integer"}}}`, `{"foo":666, "bar":123}`, true},
	}
	for _, test := range tests {
		val := NewValidatorFromString(test.schema)
		if got, _ := val.ValidateString(test.doc); got != test.want {
			t.Errorf("wanted %v, got %v", test.want, got)
		}
	}
}
