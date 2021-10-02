package main

import (
	"log"
	"net/http"
	"github.com/hilgardvr/bora-finance-svc/controllers"
)

const port = ":9000"


func main() {
	log.Println("server running on port", port)
	http.HandleFunc("/", controllers.HomePageController)
	http.HandleFunc("/home", controllers.HomePageController)
	http.HandleFunc("/api/logo", controllers.ServeFlavicon)
	http.HandleFunc("/api/images", controllers.ServeFiles)
	http.HandleFunc("/api/submitProperty", controllers.SubmitPropertyController)
	http.HandleFunc("/api/buy", controllers.BuyTokens)
	log.Fatal(http.ListenAndServe(port, nil))
}