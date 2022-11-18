package main

import (
    "fmt"
    "net/http"
    "html/template"
)

func handleFunc() {
    http.Handle(
        "/static/",
        http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))),
    )
    http.HandleFunc("/main/", index)
    http.ListenAndServe(":8084", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    templ, err := template.ParseFiles(
        "templates/index.html",
        "templates/header.html",
        "templates/footer.html",
    )

    if (err == nil) {
        fmt.Println("[+] OK")
        templ.ExecuteTemplate(w, "index", nil)
    } else {
        fmt.Println(err)
    }
}


func main() {
    handleFunc()
}
