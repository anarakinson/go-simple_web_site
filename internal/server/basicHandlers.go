package server

import (
    "net/http"
    "internal/handlers"
)


func notFound(w http.ResponseWriter, r *http.Request) {
    watswrong := "404\nCan't find page"
    handlers.StandardTemplate("something_wrong", w, r, watswrong)
}

func index(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("index", w, r)
}

func database_query(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("database_query", w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("create", w, r)
}

func signup(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("signup", w, r)
}

func signin(w http.ResponseWriter, r *http.Request) {
    handlers.StandardTemplate("signin", w, r)
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
