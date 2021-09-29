package models

type PropertyDetails struct {
	Id					string		`json: "id"`
	PropName			string		`json: "name"`
	Address 			string		`json: "address"`
	Owners				[]string	`json: "owners"`
	ExpectedYield		int			`json: "yield"`
	Value 				int			`json: "value"`
	NumNFTs				int			`json: "numNFTs"`
	PictureUrl			string		`json: "-"`
	Picture				[]byte		`json: "-"`
}

type PageVariables struct {
	Properties 	[]PropertyDetails
}
