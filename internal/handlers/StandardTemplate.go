package handlers

import (
    "fmt"
    "net/http"
    "html/template"

    _ "github.com/go-sql-driver/mysql"
)

func StandardTemplate(page_name string, w http.ResponseWriter, r *http.Request, args ...interface{}) {
    templ, err := template.ParseFiles(
        "templates/" + page_name + ".html",
        "templates/head.html",
        "templates/header.html",
        "templates/footer.html",
    )


    if (err != nil) {
        fmt.Fprintf(w, err.Error())
    } else {
        fmt.Println("[+] Go to the page:", page_name)

        if (args == nil) {
            templ.ExecuteTemplate(w, page_name, nil)
        } else {
            templ.ExecuteTemplate(w, page_name, args[0])
        }
    }
}
