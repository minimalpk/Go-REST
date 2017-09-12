package main

// Import default packages

import (
	"database/sql"
	"net/http"
)

// Import middlewares packages

import (
	middlewareAccess "rest/middlewares/access"
	middlewareDatabaseConnect "rest/middlewares/database/connect"
	middlewareDatabaseDisconnect "rest/middlewares/database/disconnect"
	middlewareRequest "rest/middlewares/request"
	middlewareResponse "rest/middlewares/response"
	middlewareRoutes "rest/middlewares/routes"
)

const (
	HTTP_HOST = "localhost"
	HTTP_PORT = "8080"
)

// Execute action

func execute(request *http.Request, parameters map[string]interface{}, database *sql.DB, action func(*http.Request, map[string]interface{}, *sql.DB) (int, interface{})) (int, interface{}) {
	return action(request, parameters, database)
}

// Main

func main() {
	// Hanle actions

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// Base

		response.Header().Set("Content-Type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "*")

		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusNoContent)

			return
		}

		// Middleware database connect

		var database = middlewareDatabaseConnect.Run()

		// Middleware request

		var parameters, error = middlewareRequest.Run(response, request)

		if error != nil {
			// Middleware database dissconnect

			middlewareDatabaseDisconnect.Run(database)

			return
		}

		// Middleware access

		var status = middlewareAccess.Run(response, request, database)

		if status == false {
			// Middleware database dissconnect

			middlewareDatabaseDisconnect.Run(database)

			return
		}

		// Middleware routes

		var action = middlewareRoutes.Run(response, request, parameters, database)

		if action == nil {
			// Middleware database dissconnect

			middlewareDatabaseDisconnect.Run(database)

			return
		}

		// Execute

		var (
			code int
			body interface{}
		)

		code, body = execute(request, parameters, database, action)

		// Middleware response

		middlewareResponse.Run(response, code, body)

		// Middleware database dissconnect

		middlewareDatabaseDisconnect.Run(database)
	})

	// Run server

	http.ListenAndServe(HTTP_HOST+":"+HTTP_PORT, nil)
}