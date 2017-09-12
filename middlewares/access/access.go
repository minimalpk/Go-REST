// Package access - Check access in actions
package access

// Import default packages

import (
	"database/sql"
	"net/http"
)

// Actions list

var actions = map[string]map[string]bool{
	"/items": {"GET": true},
	"/item":  {"GET": true, "POST": true, "PUT": true, "DELETE": true},
}

// Run - Execute package for check access in actions
func Run(response http.ResponseWriter, request *http.Request, database *sql.DB) bool {
	if actions[request.URL.Path] == nil {
		return true
	}

	if actions[request.URL.Path][request.Method] {
		var id string

		database.QueryRow("SELECT true FROM sessions WHERE token = $1", request.Header.Get("Token")).Scan(&id)

		if len(id) == 0 {
			response.WriteHeader(http.StatusForbidden)

			return false
		}

		return true
	}

	return true
}
