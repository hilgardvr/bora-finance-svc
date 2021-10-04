package service

import (
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"github.com/hilgardvr/bora-finance-svc/models"
	"os"
)

const BORA_CID_MINTER_FILE = "BORA_CID_MINTER_FILE"
const BORA_MINTER_CID = "BORA_MINTER_CID"
const BORA_PAB_URL = "BORA_PAB_URL"
const BORA_CID_SELLER_FILE = "BORA_CID_SELLER_FILE"
const BORA_CID_BUYER2_FILE = "BORA_CID_BUYER2_FILE"
var pabUrl = os.Getenv(BORA_PAB_URL)
var minterCid = os.Getenv(BORA_MINTER_CID)
var minterCidFile = os.Getenv(BORA_CID_MINTER_FILE)
var sellerCidFile = os.Getenv(BORA_CID_SELLER_FILE)
var buyer2CidFile = os.Getenv(BORA_CID_BUYER2_FILE)


func buyTokens(amount int) *http.Response {
	key, err := ioutil.ReadFile(buyer2CidFile)
	CheckErr(err)
	buyer2Cid := string(key)
	// fmt.Println("buyer2Cid", buyer2Cid)
	url := fmt.Sprintf("%s%s/endpoint/Buy Tokens", pabUrl, buyer2Cid)
	reqBody, err := json.Marshal(amount)
	CheckErr(err)
	// log.Println("Request body: ", string(reqBody))
	resp, err := http.Post(
		url,
		"application/json", 
		bytes.NewBuffer(reqBody))

    CheckErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	log.Printf("Response status %s with body %s", resp.Status, string(body))

	return resp
}

func withdrawFunds(amount int) *http.Response {
	key, err := ioutil.ReadFile(sellerCidFile)
	CheckErr(err)
	sellerCid := string(key)
	// fmt.Println("sellerCid", sellerCid)
	url := fmt.Sprintf("%s%s/endpoint/Withdraw Funds", pabUrl, sellerCid)
	reqBody, err := json.Marshal(amount)
	CheckErr(err)
	// log.Println("Request body: ", string(reqBody))
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

	log.Printf("Withdraw funds response status %s with body %s", resp.Status, string(body))

	return resp

}

func withdrawTokens(amount int) *http.Response {
	key, err := ioutil.ReadFile(sellerCidFile)
	CheckErr(err)
	sellerCid := string(key)
	// fmt.Println("sellerCid", sellerCid)
	url := fmt.Sprintf("%s%s/endpoint/Withdraw Tokens", pabUrl, sellerCid)
	reqBody, err := json.Marshal(amount)
	CheckErr(err)
	// log.Println("Request body: ", string(reqBody))
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

	log.Printf("Withdraw tokens response status %s with body %s", resp.Status, string(body))

	return resp

}

func close() *http.Response {
	key, err := ioutil.ReadFile(sellerCidFile)
	CheckErr(err)
	sellerCid := string(key)
	reqBody, err := json.Marshal([]models.PropertyDetails{})
	CheckErr(err)
	// fmt.Println("sellerCid", sellerCid)
	url := fmt.Sprintf("%s%s/endpoint/Close", pabUrl, sellerCid)
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

	log.Printf("Close response status %s with body %s", resp.Status, string(body))

	return resp

}

func listProperty(amount int) *http.Response {
	key, err := ioutil.ReadFile(sellerCidFile)
	CheckErr(err)
	sellerCid := string(key)
	// fmt.Println("sellerCid", sellerCid)
	url := fmt.Sprintf("%s%s/endpoint/List Property", pabUrl, sellerCid)
	reqBody, err := json.Marshal(amount)
	CheckErr(err)
	resp, err := http.Post(
		url,
		"application/json", 
		bytes.NewBuffer(reqBody))

    CheckErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	log.Printf("List response status %d with body %s", resp.StatusCode, string(body))

	return resp
}

//private func to mint tokens via pab
func mintTokens(prop models.PropertyDetails) *http.Response {

	tokenName := models.TokenName{
		TokenName: prop.TokenName,
	}
	mintParams := models.MintParams{
		MpTokenName: tokenName,
		MpAmount: prop.NumTokens,
	}
	reqBody, err := json.Marshal(mintParams)
	CheckErr(err)
	log.Println("Request body: ", string(reqBody))

	if (minterCid == "") {
		log.Printf("Minter file used: %s", minterCidFile)
		key, err := ioutil.ReadFile(minterCidFile)
		CheckErr(err)
		minterCid = string(key)
	}
	url := fmt.Sprintf("%s%s/endpoint/Mint", pabUrl, minterCid)
	resp, err := http.Post(
		url,
		"application/json", 
		bytes.NewBuffer(reqBody))

    CheckErr(err)

	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		CheckErr(err)
		log.Printf("Mint response status %s with body %s", resp.Status, string(body))
	}

	return resp

}
