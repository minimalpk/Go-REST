// Package response - Prepare response
package response

// Import default packages

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Run - Execute package for prepare response
func Run(response http.ResponseWriter, code int, bodyObject interface{}) error {
	response.WriteHeader(code)

	if bodyObject == nil {
		return nil
	}

	var bodyByte, error = json.Marshal(bodyObject)

	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)

		return errors.New("Error encode JSON")
	}

	response.Write(bodyByte)

	return nil
}
