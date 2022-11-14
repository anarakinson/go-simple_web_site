package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", home_page)
    http.ListenAndServe(":8084", nil)
}

func home_page(page http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(page, "Hello")
}
