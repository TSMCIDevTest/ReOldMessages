package main

import (
    "fmt"
    "instantmessengerrevival/src/libpurple"
    "instantmessengerrevival/src/nina"
    "log"
)

func main() {
    // Initialize libpurple
    libPurple := libpurple.LibPurple{}
    err := libPurple.Init()
    if err != nil {
        fmt.Println("Error initializing libpurple:", err)
        return
    }

    // Connect to Nina chat API
    ninaChatAPI := nina.NinaAPI{}
    err = ninaChatAPI.Authenticate()
    if err != nil {
        fmt.Println("Error authenticating with Nina API:", err)
        return
    }

    log.Println("Instant Messenger Revival is running...")
    // Example of connecting to libpurple and sending a message
    // Example of sending a message
    err = libPurple.Connect()
    if err != nil {
        fmt.Println("Error connecting to libpurple:", err)
        return
    }

    err = libPurple.SendMessage("Hello from Instant Messenger Revival!")
    if err != nil {
        fmt.Println("Error sending message:", err)
    }
}