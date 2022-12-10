package handlers

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"

    "internal/database"
)

func RunQuery(w http.ResponseWriter, r *http.Request) {
    // Get data
    fmt.Println("Method:", r.Method)
    err := r.ParseForm()
    if err != nil {
        watswrong := "Something went wrong..."
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when parsing form:", err.Error())
        return
    }

    // get query
    query := r.FormValue("query")

    // Print parsed data
    fmt.Println("query:", query)

    // parse db configs
    config, err := database.ParseConfig()
    if err != nil {
        watswrong := fmt.Sprintf("Something went wrong...\n%s", err.Error())
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when parsing configs:", err.Error())
        return
    }

    // Connect to db
    db, err := sql.Open(
        "mysql",
        fmt.Sprintf(
            "%s:%s@tcp(%s:%s)/%s",
            config.Login,
            config.Passwrd,
            config.Address,
            config.Port,
            config.Name,
        ),
    )
    if err != nil {
        watswrong := "Database don't work!"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when connecting to db:", err.Error())
        return
    }
    defer db.Close()

    // run query
    rows, err := db.Query(query)
    if err != nil {
        watswrong := fmt.Sprintf("Can't execute query \n%s", err.Error())
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when executing query:", err.Error())
        return
    }
    defer rows.Close()

    cols, err := rows.Columns()
    if err != nil {
        watswrong := fmt.Sprintf("Can't execute query \n%s", err.Error())
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when executing query:", err.Error())
        return
    }


    // parse result
    queryResult := make(map[int]map[string]string)
    var row_id int = 0
    for rows.Next() {
        data := make(map[string]string)
        columns := make([]string, len(cols))
        columnPointers := make([]interface{}, len(cols))
        for i, _ := range columns {
            columnPointers[i] = &columns[i]
        }

        rows.Scan(columnPointers...)

        for i, colName := range cols {
            data[colName] = columns[i]
        }
        queryResult[row_id] = data
        row_id += 1
    }

    // display data
    for _, d := range(queryResult) {
        fmt.Println(d)
    }


    // return
    StandardTemplate("database_query", w, r, queryResult)
}
