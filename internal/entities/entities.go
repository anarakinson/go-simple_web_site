package entities


type Article struct {
    Id uint16 `json:id`
    Title string `json:title`
    Announce string `json:announce`
    Text string `json:text`
}

type User struct {
    Username string `json:username`
    Email string `json:email`
    Password string `json:password`
}
