// Package dissconnect - Disconnect in database
package dissconnect

// Import default packages

import "database/sql"

// Run - Execute package for disconnect in database
func Run(database *sql.DB) {
    database.Close()
}
