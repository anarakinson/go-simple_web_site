package handlers

import (
    "fmt"
    "log"
    "net/http"
    "html/template"

    _ "github.com/go-sql-driver/mysql"
)

func StandardTemplate(page_name string, w http.ResponseWriter, r *http.Request, args ...interface{}) {
    // fmt.Println(r.URL.Path)
    templ, err := template.ParseFiles(
        "templates/" + page_name + ".html",
        "templates/head.html",
        "templates/header.html",
        "templates/footer.html",
    )


    if (err != nil) {
        watswrong := "Can't find page"
        StandardTemplate("something_wrong", w, r, watswrong)
        log.Println("[!] Error when executing template:", err.Error())
        return
    } else {
        fmt.Println("[+] Go to the page:", page_name)

        if (args == nil) {
            templ.ExecuteTemplate(w, page_name, nil)
        } else {
            templ.ExecuteTemplate(w, page_name, args[0])
        }
    }
}
