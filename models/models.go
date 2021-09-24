package models


type PropertyDetails struct {
	Name		string		`json: "name"`
	Address 	string		`json: "address"`
	Owners		[]string	`json: "owners"`
	Yield		int			`json: "yield"`
	Value 		int			`json: "value"`
	NFTs		int			`json: "nfts"`
}

type PageVariables struct {
	Date		string
	Properties 	[]PropertyDetails
}
