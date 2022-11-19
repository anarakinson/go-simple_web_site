package main

import (
    "fmt"
    "net/http"
    "html/template"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func handleFunc() {
    http.Handle(
        "/static/",
        http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))),
    )
    http.HandleFunc("/main/", index)
    http.HandleFunc("/create/", create)
    http.HandleFunc("/about/", about)
    http.HandleFunc("/contacts/", contacts)
    http.HandleFunc("/something_wrong/", something_wrong)
    http.HandleFunc("/save_article/", save_article)
    http.ListenAndServe(":8084", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    ExecuteTemplate("index", w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
    ExecuteTemplate("create", w, r)
}

func about(w http.ResponseWriter, r *http.Request) {
    ExecuteTemplate("about", w, r)
}

func contacts(w http.ResponseWriter, r *http.Request) {
    ExecuteTemplate("contacts", w, r)
}

func something_wrong(w http.ResponseWriter, r *http.Request) {
    ExecuteTemplate("something_wrong", w, r)
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
    anounce := r.FormValue("anounce")
    article_text := r.FormValue("article_text")

    if (title == "") || (article_text == "") {
        watswrong := "Title or text are not found"
        // http.Redirect(w, r, "/something_wrong/", 301)
        SomethingWentWrong(watswrong, w, r)
        return
    }

    // Print parsed data
    fmt.Println("title:", title)
    fmt.Println("anounce:", anounce)
    rune_text := string([]rune(article_text))
    if (len(rune_text) < 100) {
        fmt.Println("article_text:", rune_text)
    } else {
        fmt.Println("article_text:", rune_text[:100])
    }

    // Connect to db
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/simple_website")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Insert data to table
    insert, err := db.Query(
        fmt.Sprintf(
            "INSERT INTO articles (title, anounce, text) VALUES ('%s', '%s', '%s')",
            title,
            anounce,
            article_text,
        ),
    )
    if err != nil {
        panic(err)
    }
    defer insert.Close()

    // Redirect: (response writer, request, page to redirect, response code)
    http.Redirect(w, r, "/main/", 301) // http.StatusSeeOther() = 301

}

func ExecuteTemplate(page_name string, w http.ResponseWriter, r *http.Request) {
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
        panic(err)
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
        panic(err)
    }
}


func main() {
    handleFunc()
}
