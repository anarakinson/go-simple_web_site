package main

import (
    "fmt"
    "net/http"
)

func main() {
    handleRequest()
}

func handleRequest() {
    http.HandleFunc("/", home_page)
    http.HandleFunc("/contacts/", contacts_page)
    http.ListenAndServe(":8084", nil)
}

func home_page(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w, "Hello")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Contacts")
}
