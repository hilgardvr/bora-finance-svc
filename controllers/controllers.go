package controllers

import (
	"html/template"
	"log"
	"strconv"
	"strings"
	"time"
	"net/http"

	"github.com/hilgardvr/bora-finance-svc/service"
	"github.com/hilgardvr/bora-finance-svc/models"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	allProperties := service.ListProperties()
	pageVars := models.PageVariables{
		Date: now.Format("02-01-2006"),
		Properties: allProperties,
	}

	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatalln("template parsing err", err)
	}

	err = t.Execute(w, pageVars)
	
	if err != nil {
		log.Fatalln("template executing err", err)
	}
}

func SubmitProperty(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		propName 	:= r.FormValue("propertyName")
		address 	:= r.FormValue("address")
		owner 		:= strings.Split(r.FormValue("owner"), ",")
		yield, err	:= strconv.Atoi(r.FormValue("yield"))
		if err != nil {
			log.Fatalln(err)
		}
		value, err	:= strconv.Atoi(r.FormValue("value"))
		if err != nil {
			log.Fatalln(err)
		}
		nfts, err 	:= strconv.Atoi(r.FormValue("nfts"))
		if err != nil {
			log.Fatalln(err)
		}
		ref 		:= r.Header.Get("Referer")
		propDetails := models.PropertyDetails{
			Name	: propName,
			Address : address,
			Owners	: owner,
			Yield	: yield,
			Value	: value,
			NFTs	: nfts,
		}
		// log.Println(ref)
		service.AddProperty(propDetails)
		http.Redirect(w, r, ref, http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

