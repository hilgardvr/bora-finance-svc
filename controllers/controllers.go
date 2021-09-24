package controllers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date	string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	pageVars := PageVariables{
		Date: now.Format("02-01-2006"),
	}

	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Println("tempalte parsing err", err)
	}

	err = t.Execute(w, pageVars)
	
	if err != nil {
		log.Println("tempalte executing err", err)
	}
}

func SubmitProperty(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		log.Println(r.FormValue("propertyName"))
		ref := r.Header.Get("Referer")
		log.Println(ref)
		http.Redirect(w, r, ref, http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {

	reqBody, err := json.Marshal(17)

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
