package main

import (
	"encoding/json"
	// "fmt"
	"bytes"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("MAIN Started")
	for {
		doRequest()
		time.Sleep(time.Second * 5)
	}
}

func doRequest() {
	values := map[string]string{"username": "username", "password": "password"}
	jsonValue, _ := json.Marshal(values)
	response, err := http.Post("http://localhost:8080/check", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}
	log.Println(response)
}
