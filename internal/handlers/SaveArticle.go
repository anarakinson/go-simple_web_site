package handlers

import (
    "fmt"
    "net/http"
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "internal/database"
)


func SaveArticle(w http.ResponseWriter, r *http.Request) {
    // Get data
    fmt.Println("Method:", r.Method)
    err := r.ParseForm()
    if err != nil {
        panic(err)
        return
    }

    title := r.FormValue("title")
    announce := r.FormValue("announce")
    article_text := r.FormValue("article_text")

    if (title == "") || (article_text == "") {
        watswrong := "Title or text are not found"
        // http.Redirect(w, r, "/something_wrong/", 301)
        StandardTemplate("something_wrong", w, r, watswrong)
        return
    }

    // Print parsed data
    fmt.Println("title:", title)
    fmt.Println("announce:", announce)
    rune_text := string([]rune(article_text))
    if (len(rune_text) < 100) {
        fmt.Println("article_text:", rune_text)
    } else {
        fmt.Println("article_text:", rune_text[:100])
    }

    // Connect to db
    db, err := sql.Open(
        "mysql",
        fmt.Sprintf(
            "%s:%s@tcp(%s)/%s",
            database.Login,
            database.Passwrd,
            database.Address,
            database.Name,
        ),
    )
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Insert data to table
    query := "INSERT INTO `articles` (`title`, `announce`, `text`) VALUES (?, ?, ?)"
    insert, err := db.ExecContext(
        context.Background(),
        query,
        title,
        announce,
        article_text,
    )
    if err != nil {
        panic(err.Error())
    }
    insertId, err := insert.LastInsertId()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("inserted id:", insertId)

    // Redirect: (response writer, request, page to redirect, response code)
    http.Redirect(w, r, "/main/", 301) // http.StatusSeeOther() = 301

}
