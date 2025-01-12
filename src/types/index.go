package types

type Message struct {
    ID      string `json:"id"`
    Content string `json:"content"`
    Sender  User   `json:"sender"`
    Time    string `json:"time"`
}

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

type ChatSession struct {
    ID      string   `json:"id"`
    Users    []User   `json:"users"`
    Messages []Message `json:"messages"`
}