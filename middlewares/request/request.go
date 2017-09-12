// Package request - Prepare request
package request

// Import default packages

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Run - Execute package for prepare request
func Run(response http.ResponseWriter, request *http.Request) (map[string]interface{}, error) {
	var (
		body       []byte
		error      error
		parameters map[string]interface{}
	)

	body, error = ioutil.ReadAll(request.Body)

	if error != nil {
		response.WriteHeader(http.StatusBadRequest)

		return nil, errors.New("Error decode stream")
	}

	if len(string(body)) == 0 {
		return parameters, nil
	}

	error = json.Unmarshal(body, &parameters)

	if error != nil {
		response.WriteHeader(http.StatusBadRequest)

		return nil, errors.New("Error decode JSON")
	}

	return parameters, nil
}
