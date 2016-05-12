Go kit JSON Schema
==================

A wrapper for [xeipuuv/gojsonschema](http://github.com/xeipuuv/gojsonschema) to ease integration into Go kit services.

Validator
---------
`Validator` provides a simplified interface to gojsonschema.Validate. A validator is created with a schema, either from a schema file:

    val := jsonschema.NewValidatorFromFile("file://path/to/mySchema.json")

or from a JSON string:

    val := jsonschema.NewValidatorFromString(`{"type":"boolean"}`)

Go kit Request Decode Function
------------------------------
To aid integration with Go kit services, JSON validation can decorate a go-kit/kit/transport/http/DecodeRequestFunc, in much the same way as a Middleware wraps an Endpoint.

    func decodeRequest(req *http.Request) (request interface{}, err error) {
        var request Query
    	return request, nil
    }
    decoder := jsonschema.NewDecodeFunc(val, decodeRequest)

If the request body passes against the schema, the validation function will pass transparently on to the decorated function. If validation fails, an error will be returned and the decorated function will not be called.
