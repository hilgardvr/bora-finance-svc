package models

type PropertyDetails struct {
	Id					string		`json:"id"`
	TokenName			string		`json:"tokenName"`
	Address 			string		`json:"address"`
	Owners				[]string	`json:"owners"`
	ExpectedYield		int			`json:"yield"`
	Value 				int			`json:"value"`
	NumTokens			int			`json:"numTokens"`
	PictureUrl			string		`json:"-"`
	Picture				[]byte		`json:"-"`
}

type TokenName struct {
	TokenName			string		`json:"unTokenName"`
}

type MintParams struct {
	MpTokenName			TokenName		`json:"mpTokenName"`
	MpAmount			int				`json:"mpAmount"`
}

type PageVariables struct {
	Properties 	[]PropertyDetails
}
