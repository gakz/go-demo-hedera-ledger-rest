package model

type VerifiedServicer struct {
	Id            int16            `json:"id"`
	Name          string           `json:"name"`
	StreetAddress string           `json:"streetaddress"`
	City          string           `json:"city"`
	PostalCode    string           `json:"postalcode"`
	Country       string           `json:"country"`
	Services      [][2]interface{} `json:"services"`
	Technicians   [][2]interface{} `json:"technicians"`
}
