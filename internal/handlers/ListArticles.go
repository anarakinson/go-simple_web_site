package handlers

import (
    "fmt"
    "net/http"
    "html/template"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "internal/database"
    "internal/articles"
)


func ListArticles(w http.ResponseWriter, r *http.Request) {
    templ, err := template.ParseFiles(
        "templates/" + "list_articles" + ".html",
        "templates/head.html",
        "templates/header.html",
        "templates/footer.html",
    )

    if err!= nil {
        fmt.Fprintf(w, err.Error())
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

    templ.ExecuteTemplate(w, "list_articles", posts)
}
