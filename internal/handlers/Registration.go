package handlers

import (
    "fmt"
    "log"
    "net/http"
    "context"
    "database/sql"

    "internal/database"
    "internal/entities"

    _ "github.com/go-sql-driver/mysql"
)


func SignUp(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Method:", r.Method)

    // parse input values
    err := r.ParseForm()
    if err != nil {
        watswrong := "404\nCan't load page"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when parsing form:", err.Error())
        return
    }

    var user = entities.User{}
    user.Username = r.FormValue("username")
    user.Email = r.FormValue("email")
    user.Password = r.FormValue("password")
    password_conf := r.FormValue("password_confirmation")
    // print data
    fmt.Println(user.Username)
    fmt.Println(user.Email)
    fmt.Println(user.Password)
    fmt.Println(password_conf)

    // confirm password
    if user.Password != password_conf {
        watswrong := "Password don't equeal to confirmation!"
        StandardTemplate("something_wrong", w, r, watswrong)
        return
    } else if (user.Username == "") || (user.Email == "") || (user.Password == "") {
        watswrong := "Missing data: you have to enter username, email and password"
        StandardTemplate("something_wrong", w, r, watswrong)
        return
    }

    // parse db configs
    config, err := database.ParseConfig()
    if err != nil {
        watswrong := "Sonething went wrong..."
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when parsing db configs:", err.Error())
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

    // Insert data to table
    query := "INSERT INTO `users` (`username`, `email`, `password`) VALUES (?, ?, ?)"
    insert, err := db.ExecContext(
        context.Background(),
        query,
        user.Username,
        user.Email,
        user.Password,
    )
    if err != nil {
        watswrong := "Some data is wrong. Maybe, your email is currently in use."
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when inserting to db:", err.Error())
        return
    }
    insertId, err := insert.LastInsertId()
    if err != nil {
        watswrong := "Something went wrong..."
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when getting last id:", err.Error())
        return
    }
    fmt.Println("[+] Insert in users: success. Inserted id:", insertId)

    // Redirect: (response writer, request, page to redirect, response code)
    http.Redirect(w, r, "/signin/", 301) // http.StatusSeeOther() = 301
}
