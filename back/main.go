package main

import (
	"log"
	"net/http"
)

type response struct {
    Message string `json:"message"`
}

type data struct {
    Data string `json:"data"`
}

//Serves an endpoint that expects a http request with a command.
//A command is a string that is mapped into an action in command.go
func main(){
    log.Println("Server running on port 8888")
    http.HandleFunc("/", CommandHandler)
    log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
