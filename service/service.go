package service

import (
	"log"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/hilgardvr/bora-finance-svc/models"
)

//todo replace with db
var properties []models.PropertyDetails

//private func to update the oracle
func updateOracle(properties []models.PropertyDetails) {

	propertiesLength := len(properties)

	reqBody, err := json.Marshal(propertiesLength)

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

func AddProperty(propertyDetails models.PropertyDetails) {
	properties = append(properties, propertyDetails)
	updateOracle(properties)
}

func ListProperties() []models.PropertyDetails {
	tmp := make([]models.PropertyDetails, len(properties))
	copy(tmp, properties)
	return tmp
}


