package controllers

import (
	"html/template"
	"log"
	"strconv"
	"strings"
	"net/http"
	"io/ioutil"

	"github.com/hilgardvr/bora-finance-svc/service"
	"github.com/hilgardvr/bora-finance-svc/models"
)

func HomePageController(w http.ResponseWriter, r *http.Request) {
	allProperties := service.ListProperties()
	pageVars := models.PageVariables{
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

func checkParsingError(e error, w http.ResponseWriter, r *http.Request) {
	ref := r.Header.Get("Referer")
	if e != nil {
		log.Println("Error parsing value to int", e)
		http.Redirect(w, r, ref, http.StatusSeeOther)
	}
}

func SubmitPropertyController(w http.ResponseWriter, r *http.Request) {
	ref := r.Header.Get("Referer")
	if r.Method == "POST" {
		r.ParseMultipartForm(10 << 20)
		propName 	:= r.FormValue("propertyName")
		address 	:= r.FormValue("address")
		owner 		:= strings.Split(r.FormValue("owner"), ",")
		yield, err	:= strconv.Atoi(r.FormValue("yield"))
		checkParsingError(err, w, r)
		value, err	:= strconv.Atoi(r.FormValue("value"))
		checkParsingError(err, w, r)
		numNfts, err 	:= strconv.Atoi(r.FormValue("numnfts"))
		// service.CheckErr(err)
		checkParsingError(err, w, r)
		nfts := []string{"placeholder"}
		file, handler, err := r.FormFile("picture")
		if err != nil {
			log.Println("Error Retrieving the File")
        	log.Println(err)
        	return
		}
		log.Printf("Uploaded File: %+v\n", handler.Filename)
		log.Printf("File Size: %+v\n", handler.Size)
		log.Printf("MIME Header: %+v\n", handler.Header)
		// Create a temporary file within our temp-images directory that follows
    	// a particular naming pattern
		tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
		if err != nil {
			log.Println(err)
		}
		defer tempFile.Close()
		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
			fileBytes = nil
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		log.Printf("Successfully Uploaded File\n")
		propDetails := models.PropertyDetails{
			PropName		: propName,
			Address 		: address,
			Owners			: owner,
			ExpectedYield	: yield,
			Value			: value,
			NumNFTs			: numNfts,
			NFTs			: nfts,
			Picture			: fileBytes,
		}
		service.AddProperty(propDetails)
	} else {
		http.Redirect(w, r, ref, http.StatusSeeOther)
	}
}

