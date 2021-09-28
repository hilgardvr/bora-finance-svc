package service

import (
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

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

//private func to update the oracle
func updateOracle(properties []models.PropertyDetails) {

	propertiesLength := len(properties)

	reqBody, err := json.Marshal(propertiesLength)
	CheckErr(err)

    key, err := ioutil.ReadFile("/home/hilgard/workspace/bora-finance-plutus/bora-oracle/oracle.cid")
	CheckErr(err)
    log.Println(key)
	url := fmt.Sprintf("http://127.0.0.1:8080/api/new/contract/instance/%s/endpoint/update", key)

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

func AddProperty(propertyDetails models.PropertyDetails) {
	properties = append(properties, propertyDetails)
	updateOracle(properties)
}

func ListProperties() []models.PropertyDetails {
	tmp := make([]models.PropertyDetails, len(properties))
	copy(tmp, properties)
	return tmp
}


