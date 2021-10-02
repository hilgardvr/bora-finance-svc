package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/hilgardvr/bora-finance-svc/service"
)

//todo dynamic populate
func ServeFiles(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("\nurl:%s\n", r.URL.Path)
    p := "./uploads/mansion.jpg"
    http.ServeFile(w, r, p)
}
func ServeFlavicon(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/BoraLogo.png")
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

func BuyTokens(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		service.BuyTokens(r)
		ref := r.Header.Get("Referer")
		http.Redirect(w, r, ref, http.StatusSeeOther)
	} else {
		http.Error(w, "Expecting POST request", http.StatusMethodNotAllowed)
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

