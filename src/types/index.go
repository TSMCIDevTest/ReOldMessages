package types

import "time"

type Message struct {
    ID      string `json:"id"`
    Content string `json:"content"`
    Sender  User   `json:"sender"`
    Time    time.Time `json:"time"`
}

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
type Chat struct {
    ID       string    `json:"id"`
    Users    []User    `json:"users"`
    Messages []Message `json:"messages"`
}