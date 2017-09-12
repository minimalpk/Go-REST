// Package connect - Connect in database
package connect

// Import default packages

import "database/sql"

// Import other packages

import _ "github.com/lib/pq"

// Configuration

const (
    DATABASE_USER     = "postgres"
    DATABASE_PASSWORD = "12345"
    DATABASE_NAME     = "development"
)

// Run - Execute package for connect in database
func Run() *sql.DB {
    var database, _ = sql.Open("postgres", "user="+DATABASE_USER+" password="+DATABASE_PASSWORD+" dbname="+DATABASE_NAME+" sslmode=disable")

    return database
}
