package main

import (
    "fmt"
    "instant-messenger-revival/src/libpurple"
    "instant-messenger-revival/src/nina"
)

func main() {
    // Initialize libpurple
    lp := libpurple.LibPurple{}
    err := lp.Init()
    if err != nil {
        fmt.Println("Error initializing libpurple:", err)
        return
    }

    // Connect to Nina chat API
    ninaAPI := nina.NinaAPI{}
    err = ninaAPI.Authenticate()
    if err != nil {
        fmt.Println("Error authenticating with Nina API:", err)
        return
    }

    fmt.Println("Instant Messenger Revival is running...")
    
    // Example of sending a message
    err = lp.Connect()
    if err != nil {
        fmt.Println("Error connecting to libpurple:", err)
        return
    }

    err = lp.SendMessage("Hello from Instant Messenger Revival!")
    if err != nil {
        fmt.Println("Error sending message:", err)
    }
}