package handlers

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "internal/database"
    "internal/entities"
)


func ListArticles(w http.ResponseWriter, r *http.Request) {
    // parse configs
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

    res, err := db.Query("SELECT `id`, `title`, `announce`, `text` FROM `articles`")
    if err != nil {
        watswrong := "Can't find any articles!"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when loading article:", err.Error())
        return
    }
    defer res.Close()

    // Load articles to page
    var posts = []entities.Article{}
    for res.Next() {
        var post entities.Article
        err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
        if err != nil {
            watswrong := "Can't find article!"
            StandardTemplate("something_wrong", w, r, watswrong)
            log.Println("[!] Error when loading article:", err.Error())
            return
        }
        // fmt.Printf("Id: %d\nTitle: %s\nAnnounce: %s\n", post.Id, post.Title, post.Announce)
        posts = append(posts, post)
    }

    // display page
    StandardTemplate("articles", w, r, posts)
}
