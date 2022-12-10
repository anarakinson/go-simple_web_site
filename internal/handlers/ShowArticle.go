package handlers

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"

    "internal/database"
    "internal/entities"
)


func ShowArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK) // http.StatusOK -> 200

    // parse db configs
    config, err := database.ParseConfig()
    if err != nil {
        watswrong := "Something went wrong..."
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
        log.Println("[!] Error when connrcting to db:", err.Error())
        return
    }
    defer db.Close()

    query := fmt.Sprintf("SELECT `id`, `title`, `announce`, `text` FROM `articles` WHERE `id` = %s", vars["id"])
    fmt.Println(query)
    res, err := db.Query(query)
    if err != nil {
        watswrong := "Can't find article!"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when loading article:", err.Error())
        return
    }
    defer res.Close()

    // parse result
    var showPost = entities.Article{}
    for res.Next() {
        var post entities.Article
        err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
        if err != nil {
            watswrong := "Can't find article!"
            StandardTemplate("something_wrong", w, r, watswrong)
            log.Println("[!] Error when loading article:", err.Error())
            return
        }
        showPost = post
    }

    // display article
    StandardTemplate("show_article", w, r, showPost)
}
