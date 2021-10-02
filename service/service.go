package service

import (
	"os"
	"strconv"
	"strings"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hilgardvr/bora-finance-svc/models"
)

//todo replace with db
var properties []models.PropertyDetails
const URL = "localhost:9000/uploads/"
const BORA_CID_MINTER_FILE = "BORA_CID_MINTER_FILE"
const BORA_MINTER_CID = "BORA_MINTER_CID"
const BORA_PAB_URL = "BORA_PAB_URL"
var pabUrl = os.Getenv(BORA_PAB_URL)

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func checkParsingError(e error, w http.ResponseWriter, r *http.Request) {
	ref := r.Header.Get("Referer")
	if e != nil {
		log.Println("Error parsing value to int", e)
		http.Redirect(w, r, ref, http.StatusSeeOther)
	}
}


//private func to update the oracle
func updateOracle(prop models.PropertyDetails) {

	tokenName := models.TokenName{
		TokenName: prop.TokenName,
	}
	mintParams := models.MintParams{
		MpTokenName: tokenName,
		MpAmount: prop.NumTokens,
	}
	reqBody, err := json.Marshal(mintParams)
	CheckErr(err)
	fmt.Println(string(reqBody))

	minterCid := os.Getenv(BORA_MINTER_CID)
	if (minterCid == "") {
		minterCidFile := os.Getenv(BORA_CID_MINTER_FILE)
		fmt.Printf("minter file: %s", minterCidFile)
		key, err := ioutil.ReadFile(minterCidFile)
		CheckErr(err)
		log.Println(key)
		minterCid = string(key)
	}
	url := fmt.Sprintf("%s%s/endpoint/Mint", pabUrl, minterCid)
	resp, err := http.Post(
		url,
		"application/json", 
		bytes.NewBuffer(reqBody))

    CheckErr(err)

	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func uploadFile(r *http.Request) (string, []byte, error) {
	file, handler, err := r.FormFile("picture")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return "", []byte{}, err
	}
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		log.Println("Error creating directory", err)
		panic(err)
	}
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", handler.Filename))
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer dst.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		fileBytes = nil
	}
	// write this byte array to our temporary file
	dst.Write(fileBytes)
	// return that we have successfully uploaded our file!
	log.Printf("Successfully Uploaded File\n")
	return handler.Filename, fileBytes, err
}

//todo validation of form values eg duplicates etc
func parseForm(r *http.Request) (models.PropertyDetails, error) {
	err := r.ParseMultipartForm(10 << 20)
	properyDetails := models.PropertyDetails{}
	if err != nil {
		return properyDetails, err
	}
	propName 	:= r.FormValue("tokenName")
	address 	:= r.FormValue("address")
	owner 		:= strings.Split(r.FormValue("owner"), ",")
	yield, err	:= strconv.Atoi(r.FormValue("yield"))
	if err != nil {
		return properyDetails, err
	}
	value, err	:= strconv.Atoi(r.FormValue("value"))
	if err != nil {
		return properyDetails, err
	}
	numNfts, err 	:= strconv.Atoi(r.FormValue("numTokens"))
	if err != nil {
		return properyDetails, err
	}
	name, bytes, err := uploadFile(r)
	if err != nil {
		log.Println("Error uploading file", err)
	}
	propDetails := models.PropertyDetails{
		Id				: fmt.Sprintf("%s%s", propName, address),
		TokenName		: propName,
		Address 		: address,
		Owners			: owner,
		ExpectedYield	: yield,
		Value			: value,
		NumTokens		: numNfts,
		Picture			: bytes,
		PictureUrl		: name,
	}
	return propDetails, nil
}

func AddProperty(r *http.Request) error {
	propertyDetails, err := parseForm(r)
	if err != nil {
		log.Println("Error parsing form: ", err)
		return err
	}
	updateOracle(propertyDetails)
	properties = append(properties, propertyDetails)
	return nil
}

func ListProperties() []models.PropertyDetails {
	tmp := make([]models.PropertyDetails, len(properties))
	copy(tmp, properties)
	return tmp
}

func MakePropertyUrls(props []models.PropertyDetails) []models.PropertyDetails {
	var urlProps []models.PropertyDetails
	for _, prop := range props {
		url := fmt.Sprintf("%s%s", URL, prop.PictureUrl)
		urlProps = append(urlProps, models.PropertyDetails{
			Id : prop.Id,
			TokenName : prop.TokenName,
			Address : prop.Address,
			Owners : prop.Owners,
			ExpectedYield : prop.ExpectedYield,
			Value : prop.Value,
			NumTokens : prop.NumTokens,
			PictureUrl : url,
			Picture : prop.Picture,
		})
	}
	return urlProps
}