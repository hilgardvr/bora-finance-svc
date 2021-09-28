package main

import (
	"log"
	"net/http"
	"github.com/hilgardvr/bora-finance-svc/controllers"
)

const port = ":9000"


func main() {
	// http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("server running on port", port)
	http.HandleFunc("/", controllers.HomePageController)
	http.HandleFunc("/images", controllers.ServeFiles)
	http.HandleFunc("/home", controllers.HomePageController)
	http.HandleFunc("/submitProperty", controllers.SubmitPropertyController)
	log.Fatal(http.ListenAndServe(port, nil))
}