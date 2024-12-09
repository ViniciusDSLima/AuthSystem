package entity

type Address struct {
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	Complement   string `json:"complement"`
}
