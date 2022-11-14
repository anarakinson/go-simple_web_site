package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type User struct {
    Name string
    Age uint16
    Money int
    Avg_grades float64
    Happiness float64
    Hobbies []string
}

func (u User) getInfo() (string) {
    return fmt.Sprintf(
        "User name is %s. He is %d y.o., and he have %d money.",
        u.Name, u.Age, u.Money,
    )
}

func (u *User) setName(new_name string) {
    u.Name = new_name
}


func handleRequest() {
    http.HandleFunc("/", home_page)
    http.HandleFunc("/contacts/", contacts_page)
    http.ListenAndServe(":8084", nil)
}

func home_page(w http.ResponseWriter, request *http.Request) {
    bobby := User{
        Name: "Bobby",
        Age: 25,
        Money: -50,
        Avg_grades: 4.5,
        Happiness: 0.8,
        Hobbies: []string{"Sex", "Drugs", "Rock-n-Roll"},
    }

    templ, err := template.ParseFiles("templates/homepage.html")
    if (err == nil) {
        templ.Execute(w, bobby)
    } else {
        fmt.Println(err)
    }
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Contacts")
}


func main() {
    handleRequest()
}
