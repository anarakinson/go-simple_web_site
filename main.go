package main

import (
    "fmt"
    "net/http"
    "html/template"

    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/gorilla/mux"
)

const db_login = "root"
const db_passwrd = ""
const db_address = "127.0.0.1:3306"
const db_name = "simple_website"

type Article struct {
    Id uint16
    Title string
    Announce string
    Text string
}

func handleFunc() {
    router := mux.NewRouter()

    http.Handle(
        "/static/",
        http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))),
    ) // connect static objects, such as styles, pictures, etc.
    router.HandleFunc("/main/", index).Methods("GET")
    router.HandleFunc("/create/", create).Methods("GET")
    router.HandleFunc("/about/", about).Methods("GET")
    router.HandleFunc("/contacts/", contacts).Methods("GET")
    router.HandleFunc("/something_wrong/", something_wrong)
    router.HandleFunc("/save_article/", save_article).Methods("POST")
    router.HandleFunc("/articles/", list_articles).Methods("GET")
    router.HandleFunc("/post/{id:[0-9]+}/", show_article)

    http.Handle("/", router)
    http.ListenAndServe(":8084", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    StandardTemplate("index", w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
    StandardTemplate("create", w, r)
}

func about(w http.ResponseWriter, r *http.Request) {
    StandardTemplate("about", w, r)
}

func contacts(w http.ResponseWriter, r *http.Request) {
    StandardTemplate("contacts", w, r)
}

func something_wrong(w http.ResponseWriter, r *http.Request) {
    StandardTemplate("something_wrong", w, r)
}

func show_article(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK) // http.StatusOK -> 200

    // Connect to db
    db, err := sql.Open(
        "mysql",
        fmt.Sprintf(
            "%s:%s@tcp(%s)/%s",
            db_login,
            db_passwrd,
            db_address,
            db_name,
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
    var showPost = Article{}
    for res.Next() {
        var post Article
        err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
        if err != nil {
            panic(err.Error)
        }
        showPost = post
    }

    // display article
    page_name := "show_article"
    templ, err := template.ParseFiles(
        "templates/" + page_name + ".html",
        "templates/head.html",
        "templates/header.html",
        "templates/footer.html",
    )

    if (err == nil) {
        fmt.Println("[+] Go to the page:", page_name)
        templ.ExecuteTemplate(w, page_name, showPost)
    } else {
        fmt.Fprintf(w, err.Error())
    }
}

func list_articles(w http.ResponseWriter, r *http.Request) {
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
            db_login,
            db_passwrd,
            db_address,
            db_name,
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
    var posts = []Article{}
    for res.Next() {
        var post Article
        err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
        if err != nil {
            panic(err.Error())
        }
        // fmt.Printf("Id: %d\nTitle: %s\nAnnounce: %s\n", post.Id, post.Title, post.Announce)
        posts = append(posts, post)
    }

    templ.ExecuteTemplate(w, "list_articles", posts)
}

func save_article(w http.ResponseWriter, r *http.Request) {
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
        SomethingWentWrong(watswrong, w, r)
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
            db_login,
            db_passwrd,
            db_address,
            db_name,
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

func StandardTemplate(page_name string, w http.ResponseWriter, r *http.Request) {
    templ, err := template.ParseFiles(
        "templates/" + page_name + ".html",
        "templates/head.html",
        "templates/header.html",
        "templates/footer.html",
    )

    if (err == nil) {
        fmt.Println("[+] Go to the page:", page_name)
        templ.ExecuteTemplate(w, page_name, nil)
    } else {
        fmt.Fprintf(w, err.Error())
    }
}

func SomethingWentWrong(watswrong string, w http.ResponseWriter, r *http.Request) {
    templ, err := template.ParseFiles(
        "templates/" + "something_wrong" + ".html",
        "templates/head.html",
        "templates/header.html",
        "templates/footer.html",
    )

    if (err == nil) {
        fmt.Println("[+] Go to the page:", "something_wrong")
        templ.ExecuteTemplate(w, "something_wrong", watswrong)
    } else {
        fmt.Fprintf(w, err.Error())
    }
}


func main() {
    handleFunc()
}
