// Package routes - Check actions
package routes

// Import default packages

import (
	"database/sql"
	"net/http"
	"rest/methods/item"
	"rest/methods/items"
)

// Import action packages

// Actions list

var actions = map[string]map[string]func(*http.Request, map[string]interface{}, *sql.DB) (int, interface{}){
	"/items": {"GET": items.Get},
	"/item":  {"GET": item.Get, "POST": item.Create, "PUT": item.Edit, "DELETE": item.Delete},
}

// Run - Execute package for check package
func Run(response http.ResponseWriter, request *http.Request, parameters map[string]interface{}, database *sql.DB) func(*http.Request, map[string]interface{}, *sql.DB) (int, interface{}) {
	if actions[request.URL.Path] == nil {
		response.WriteHeader(http.StatusNotFound)

		return nil
	}

	if actions[request.URL.Path][request.Method] == nil {
		response.WriteHeader(http.StatusMethodNotAllowed)

		return nil
	}

	return actions[request.URL.Path][request.Method]
}
