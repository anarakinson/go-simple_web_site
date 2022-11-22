package main

import (
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"

    "internal/handlers"
)


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
    router.HandleFunc("/save_article/", handlers.SaveArticle).Methods("POST")
    router.HandleFunc("/articles/", handlers.ListArticles).Methods("GET")
    router.HandleFunc("/post/{id:[0-9]+}/", handlers.ShowArticle)

    http.Handle("/", router)
    http.ListenAndServe(":8084", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("index", w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("create", w, r)
}

func about(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("about", w, r)
}

func contacts(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("contacts", w, r)
}

func something_wrong(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("something_wrong", w, r)
}



func main() {
    handleFunc()
}
