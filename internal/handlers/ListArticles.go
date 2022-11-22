package handlers

import (
    "fmt"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "internal/database"
    "internal/articles"
)


func ListArticles(w http.ResponseWriter, r *http.Request) {
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

    res, err := db.Query("SELECT `id`, `title`, `announce`, `text` FROM `articles`")
    if err != nil {
        panic(err.Error())
    }
    defer res.Close()

    // Load articles to page
    var posts = []articles.Article{}
    for res.Next() {
        var post articles.Article
        err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
        if err != nil {
            panic(err.Error())
        }
        // fmt.Printf("Id: %d\nTitle: %s\nAnnounce: %s\n", post.Id, post.Title, post.Announce)
        posts = append(posts, post)
    }

    // display page
    StandardTemplate("list_articles", w, r, posts)
}
