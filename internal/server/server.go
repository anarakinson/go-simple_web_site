package server

import (
    "fmt"
    "net/http"

    "internal/handlers"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/spf13/viper"
)


func RunServer() {
    router := mux.NewRouter()
    port := viper.GetString("app.port")

    // connect static objects, such as styles, pictures, etc.
    prefix := http.StripPrefix("/static/", http.FileServer(secureFileServer{http.Dir("./static/")}))
    http.Handle("/static/", prefix)

    // site map
    router.HandleFunc("/main/", index).Methods("GET")
    // authentification
    router.HandleFunc("/signup/", signup).Methods("GET")
    router.HandleFunc("/signin/", signin).Methods("GET")
    router.HandleFunc("/signup_success/", handlers.SignUp).Methods("POST")
    // database administration
    router.HandleFunc("/database_query/", database_query).Methods("GET")
    router.HandleFunc("/database_query/", handlers.RunQuery).Methods("POST")
    // articles
    router.HandleFunc("/create/", create).Methods("GET")
    router.HandleFunc("/save_article/", handlers.SaveArticle).Methods("POST")
    router.HandleFunc("/articles/", handlers.ListArticles).Methods("GET")
    router.HandleFunc("/post/{id:[0-9]+}/", handlers.ShowArticle).Methods("GET")
    // other
    router.HandleFunc("/about/", about).Methods("GET")
    router.HandleFunc("/contacts/", contacts).Methods("GET")
    router.HandleFunc("/something_wrong/", something_wrong)

    // set errors
    nah := http.HandlerFunc(notFound)
    router.NotFoundHandler = nah

    // run server
    http.Handle("/", router)
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}


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
