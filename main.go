package main

import (
	// "bytes"
	// "encoding/json"
	// "html/template"
	// "io/ioutil"
	"log"
	"net/http"
	// "time"
	"github.com/hilgardvr/bora-finance-svc/controllers"
)

const port = ":9000"


func main() {
	// http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("server running on port", port)
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/home", controllers.HomePage)
	http.HandleFunc("/submitProperty", controllers.SubmitProperty)
	// http.HandleFunc("/test", controllers.Test)
	log.Fatal(http.ListenAndServe(port, nil))
}