package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CommandHandler (w http.ResponseWriter, r *http.Request) {
    //CORS
    allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Expose-Headers", "Authorization")
    w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }
    
    //decode data
    var d data
    if err := json.NewDecoder(r.Body).Decode(&d); err != nil{ 
        fmt.Println(err.Error())
        return
    }
   
    //execute get command from response body "data" field and execute it
    j, err := json.Marshal(response{Message: ExecuteCommand(d.Data)})
    if err != nil {
        fmt.Println(err)
        return
    }

    w.Write(j)
}

