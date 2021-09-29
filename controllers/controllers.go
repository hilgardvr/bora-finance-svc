package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	// "github.com/hilgardvr/bora-finance-svc/models"
	"github.com/hilgardvr/bora-finance-svc/service"
)

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("url:%s", r.URL.Path)
    // p := "." + r.URL.Path
    p := "./temp-images/upload-983822252.png"
    http.ServeFile(w, r, p)
}

func HomePageController(w http.ResponseWriter, r *http.Request) {
	allProperties := service.ListProperties()
	urlProps := service.MakePropertyUrls(allProperties)

	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatalln("template parsing err", err)
	}

	err = t.Execute(w, urlProps)
	
	if err != nil {
		log.Fatalln("template executing err", err)
	}
}

func SubmitPropertyController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		service.AddProperty(r)
		ref := r.Header.Get("Referer")
		http.Redirect(w, r, ref, http.StatusSeeOther)
	} else {
		http.Error(w, "Expecting POST request", http.StatusMethodNotAllowed)
	}
}

