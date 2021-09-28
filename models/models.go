package models

type PropertyDetails struct {
	PropName			string		`json: "name"`
	Address 			string		`json: "address"`
	Owners				[]string	`json: "owners"`
	ExpectedYield		int			`json: "yield"`
	Value 				int			`json: "value"`
	NumNFTs				int			`json: "numNFTs"`
	NFTs				[]string	`json: "nfts"`
	Picture				[]byte		`json: "-"`
}

type PageVariables struct {
	Properties 	[]PropertyDetails
}
