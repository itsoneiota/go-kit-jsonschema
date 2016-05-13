package jsonschema

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewDecodeFunc returns the given DecodeFunc wrapped with a
// JSON Schema validation check.
func NewDecodeFunc(val Validator, df httptransport.DecodeRequestFunc) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (request interface{}, err error) {
		// Read the content from the request, then push it back in so the next
		// func can read it again.
		var bodyBytes []byte
		if req.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(req.Body)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyString := string(bodyBytes)

		if valid, err := val.ValidateString(bodyString); !valid {
			return nil, err
		}
		return df(ctx, req)
	}
}
