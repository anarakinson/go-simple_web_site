package handlers

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"

    "internal/database"
    "internal/entities"
    "internal/utils"

    _ "github.com/go-sql-driver/mysql"
)


func SignIn(w http.ResponseWriter, r *http.Request) {
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
    password := r.FormValue("password")

    fmt.Println(user.Username)
    fmt.Println(password)
    // confirm password
    if (user.Username == "") || (password == "") {
        watswrong := "Missing data: you have to enter username and password"
        StandardTemplate("something_wrong", w, r, watswrong)
        return
    }

    // hashing password
    user.Password = utils.GetPswrdHash(password)

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
    query := fmt.Sprintf("SELECT `username`, `email`,  `password` FROM `users` WHERE `username` = '%s'", user.Username)
    fmt.Println(query)
    res, err := db.Query(query)
    if err != nil {
        watswrong := "Can't find user!"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when logging user:", err.Error())
        return
    }
    defer res.Close()

    // parse result
    var logedUser = entities.User{}
    for res.Next() {
        var userInfo entities.User
        err = res.Scan(&userInfo.Username, &userInfo.Email, &userInfo.Password)
        if err != nil {
            watswrong := "Can't find user!"
            StandardTemplate("something_wrong", w, r, watswrong)
            log.Println("[!] Error when logging user:", err.Error())
            return
        }
        logedUser = userInfo
    }

    // display user
    if logedUser.Username == "" {
        watswrong := "Can't find user!"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when logging user: No such user in database!")
        return
    }

    // check password
    if logedUser.Password == user.Password {
        fmt.Println("SUCCESS")
    } else {
        watswrong := "Incorrect password!"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when logging user: Incorrect password!")
        return
    }

    // Redirect: (response writer, request, page to redirect, response code)
    http.Redirect(w, r, "/main/", 301) // http.StatusSeeOther() = 301
}
