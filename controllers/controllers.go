package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/hilgardvr/bora-finance-svc/models"
	"github.com/hilgardvr/bora-finance-svc/service"
)

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
    // p := "." + r.URL.Path
    p := "./temp-images/upload-983822252.png"
    http.ServeFile(w, r, p)
}

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
		imageName := "upload1.png"
		tempFile, err := ioutil.TempFile("temp-images", imageName)
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
		id := fmt.Sprintf("%s%s",propName, address)
		// fmt.Printf("Host: %s\n", r.Host)
		url:= fmt.Sprintf("%s/%s", r.Host, imageName)
		propDetails := models.PropertyDetails{
			Id				: id,
			PropName		: propName,
			Address 		: address,
			Owners			: owner,
			ExpectedYield	: yield,
			Value			: value,
			NumNFTs			: numNfts,
			NFTs			: nfts,
			Picture			: fileBytes,
			PictureUrl		: url,
		}
		service.AddProperty(propDetails)
	} else {
		http.Redirect(w, r, ref, http.StatusSeeOther)
	}
}

