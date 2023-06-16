package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
    "fmt"
)

type quote struct {
    Text string `json:"q"`
    Author string `json:"a"`
}

func GetInspirationQuote() string {
    c := http.Client{}

    r, err := c.Get("https://zenquotes.io/api/random")
    if err != nil {
        return ""
    }
    var a []quote 
    b, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(b, &a)
    fmt.Println(a)
    return "\"" + a[0].Text + "\", " + a[0].Author
    
}
