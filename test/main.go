package main

import (
    // "encoding/json"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/users", usersHandler)

    log.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}