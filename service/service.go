package service

import (
	"log"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
)

type PropertyDetails struct {
	Name		string		`json: name`
	Address 	string		`json: address`
	Owners		[]string	`json: owners`
	Yield		int			`json: yield`
	Value 		int			`json: value`
	NFTs		int			`json: nfts`
}

//todo replace with db
var properties []PropertyDetails

//private func to update the oracle
func updateOracle(properties []PropertyDetails) {

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

func AddProperty(propertyDetails PropertyDetails) {
	properties = append(properties, propertyDetails)
	updateOracle(properties)
}

func ListProperties() []PropertyDetails {
	tmp := make([]PropertyDetails, len(properties))
	copy(tmp, properties)
	return tmp
}


