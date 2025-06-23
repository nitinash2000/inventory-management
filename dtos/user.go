package dtos

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Mobile  string  `json:"mobile"`
	Address Address `json:"address"`
	Role    string  `json:"role"`
}

type Address struct {
	AddressId string `json:"address_id"`
	Line1     string `json:"line1"`
	Line2     string `json:"line2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	ZipCode   string `json:"zip_code"`
}
