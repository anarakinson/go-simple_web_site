package main

import (
    "fmt"
    "net/http"
    "log"

    "internal/handlers"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
)


func handleFunc() {
    router := mux.NewRouter()
    port := viper.GetString("app.port")

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
    router.HandleFunc("/post/{id:[0-9]+}/", handlers.ShowArticle).Methods("GET")

    http.Handle("/", router)
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
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

func initConfig() error {
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}


func main() {
    // parse configs
    err := initConfig()
    if err != nil {
        log.Fatal("[!] Error when parsing configs: %s", err.Error())
    }
    // parse variables
    err = godotenv.Load()
    if err != nil {
        log.Fatal("[!] Error when parsing environment variables: %s", err.Error())
    }

    handleFunc()
}
