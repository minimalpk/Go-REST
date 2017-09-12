package item

import (
    "database/sql"
    "net/http"
)

func Get(request *http.Request, parameters map[string]interface{}, database *sql.DB) (int, interface{}) {
    var id = request.URL.Query().Get("id")

    if len(id) == 0 {
        return 400, nil
    }

    var rows, error = database.Query("SELECT name, number, month, year FROM items WHERE id = $1", id)

    if error != nil {
        return 500, nil
    }

    if rows.Next() == false {
        return 404, nil
    }

    var name, number, month, year string

    rows.Scan(&name, &number, &month, &year)

    return 200, map[string]string{"name": name, "number": number, "month": month, "year": year}
}

func Create(request *http.Request, parameters map[string]interface{}, database *sql.DB) (int, interface{}) {
    // Status fieds check input parameters type

    var (
        name,   statusName   = parameters["name"].(string)
        number, statusNumber = parameters["number"].(string)
        month,  statusMonth  = parameters["month"].(string)
        year,   statusYear   = parameters["year"].(string)
    )

    if !statusName || !statusNumber || !statusMonth || !statusYear {
        return 400, nil
    }

    if len(name) == 0 || len(number) != 16 || len(month) != 2 || len(year) != 2 {
        return 400, nil
    }

    database.Query("INSERT INTO items (name, number, month, year) VALUES($1, $2, $3, $4)", name, number, month, year)

    return 201, nil
}

func Edit(request *http.Request, parameters map[string]interface{}, database *sql.DB) (int, interface{}) {
    // Status fieds check input parameters type

    var (
        id,   statusID   = parameters["id"].(float64)
        name, statusName = parameters["name"].(string)
    )

    if !statusID || !statusName {
        return 400, nil
    }

    if id == 0 || len(name) == 0 {
        return 400, nil
    }

    var status string

    database.QueryRow("UPDATE items SET name = $1 WHERE id = $2 RETURNING true", name, id).Scan(&status)

    if len(status) == 0 {
        return 404, nil
    }

    return 204, nil
}

func Delete(request *http.Request, parameters map[string]interface{}, database *sql.DB) (int, interface{}) {
    var id = request.URL.Query().Get("id")

    if len(id) == 0 {
        return 400, nil
    }

    var status string

    database.QueryRow("DELETE FROM items WHERE id = $1 RETURNING true", id).Scan(&status)

    if len(status) == 0 {
        return 404, nil
    }

    return 204, nil
}
