package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: 1, Name: "Budi"},
    {ID: 2, Name: "Ani"},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(users)
    case http.MethodPost:
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }

        var newUser User
        if err := json.Unmarshal(body, &newUser); err != nil {
            http.Error(w, "Bad JSON", http.StatusBadRequest)
            return
        }

        newUser.ID = len(users) + 1
        users = append(users, newUser)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newUser)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}