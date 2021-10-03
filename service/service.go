package service

import (
	"os"
	"strconv"
	"strings"

	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hilgardvr/bora-finance-svc/models"
)

//todo replace with db
var properties []models.PropertyDetails
const URL = "localhost:9000/uploads/"

func CheckErr(e error) {
	if e != nil {
		log.Println("Error", e)
	}
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
func parseTokenizeForm(r *http.Request) (models.PropertyDetails, error) {
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

func Mint(r *http.Request) error {
	propertyDetails, err := parseTokenizeForm(r)
	if err != nil {
		log.Println("Error parsing form: ", err)
		return err
	}
	resp := mintTokens(propertyDetails)
	// err = ListProperty(r)
	// if err != nil {
	// 	log.Println("Error listing prop from mint", err)
	// }
	if resp.StatusCode == 200 {
		properties = append(properties, propertyDetails)
		return nil
	} else {
		msg := fmt.Sprintf("unsuccessfull call to mint pab - reponse %d: %s", resp.StatusCode, resp.Body)
		log.Println(msg)
		return errors.New(msg)
	}
}

func ListProperty(r *http.Request) error {
	propertyDetails, err := parseTokenizeForm(r)
	if err != nil {
		log.Println("Error parsing form: ", err)
		return err
	}
	resp := listProperty(propertyDetails.NumTokens)
	if resp.StatusCode == 200 {
		properties = append(properties, propertyDetails)
		return nil
	} else {
		msg := fmt.Sprintf("unsuccessfull call to list pab - reponse %d: %s", resp.StatusCode, resp.Body)
		log.Println(msg)
		return errors.New(msg)
	}
}

func BuyTokens(r *http.Request) error {
	amount, err	:= strconv.Atoi(r.FormValue("buyAmount"))
	if err != nil {
		log.Println("could not parse buy amount to int", err)
		return err
	}
	 //todo
	// tokenName 	:= r.FormValue("tokenName")
	// buyer 		:= r.FormValue("buyer")
	buyTokens(amount)
	return nil
}

func WithdrawTokens(r *http.Request) error {
	amount, err	:= strconv.Atoi(r.FormValue("withdrawAmount"))
	if err != nil {
		log.Println("could not parse withdraw tokens to int", err)
		return err
	}
	withdrawTokens(amount)
	return nil
}

func WithdrawFund(r *http.Request) error {
	amount, err	:= strconv.Atoi(r.FormValue("withdrawAmount"))
	if err != nil {
		log.Println("could not parse withdraw tokens to int", err)
		return err
	}
	withdrawFunds(amount)
	return nil
}

func Close() error {
	close()
	return nil
}

func GetProperties() []models.PropertyDetails {
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