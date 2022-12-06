package entities


type Article struct {
    Id uint16
    Title string
    Announce string
    Text string
}

type User struct {
    Username string
    Email string
    Password string
}
