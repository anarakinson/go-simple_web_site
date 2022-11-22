package handlers

import (
    "fmt"
    "net/http"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"

    "internal/database"
    "internal/articles"
)


func ShowArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK) // http.StatusOK -> 200

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

    query := fmt.Sprintf("SELECT `id`, `title`, `announce`, `text` FROM `articles` WHERE `id` = %s", vars["id"])
    fmt.Println(query)
    res, err := db.Query(query)
    if err != nil {
        panic(err.Error())
    }
    defer res.Close()

    // parse result
    var showPost = articles.Article{}
    for res.Next() {
        var post articles.Article
        err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
        if err != nil {
            panic(err.Error)
        }
        showPost = post
    }

    // display article
    StandardTemplate("show_article", w, r, showPost)
}
