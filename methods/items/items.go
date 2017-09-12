package items

import (
	"database/sql"
	"net/http"
)

func Get(request *http.Request, parameters map[string]interface{}, database *sql.DB) (int, interface{}) {
	var name, number, month, year string
	var body []map[string]string

	var rows, error = database.Query("SELECT name, number, month, year FROM items")

	if error != nil {
		return 500, nil
	}

	for rows.Next() {
		rows.Scan(&name, &number, &month, &year)

		body = append(body, map[string]string{"name": name, "number": number, "month": month, "year": year})
	}

	return 200, body
}
