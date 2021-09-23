package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"net/http"
)

const port = ":8090"

func test(w http.ResponseWriter, r *http.Request) {
	reqBody, err := json.Marshal(5)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(
		"http://127.0.0.1:8080/api/new/contract/instance/a8718d09-bd90-4caa-b70d-7d3c8a21023b/endpoint/update", 
		"application/json", 
		bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}


func main() {
	fmt.Println("Hi")
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(port, nil))
}