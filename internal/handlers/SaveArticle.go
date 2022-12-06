package handlers

import (
    "fmt"
    "log"
    "net/http"
    "context"
    "database/sql"

    "internal/database"

    _ "github.com/go-sql-driver/mysql"
)


func SaveArticle(w http.ResponseWriter, r *http.Request) {
    // Get data
    fmt.Println("Method:", r.Method)
    err := r.ParseForm()
    if err != nil {
        watswrong := "Something went wrong..."
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when parsing form:", err.Error())
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

    // parse db configs
    config, err := database.ParseConfig()
    if err != nil {
        watswrong := "Something went wrong..."
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
        watswrong := "Something went wrong..."
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
    fmt.Println("[+] Insert in articles: success. Inserted id:", insertId)

    // Redirect: (response writer, request, page to redirect, response code)
    http.Redirect(w, r, "/articles/", 301) // http.StatusSeeOther() = 301

}
