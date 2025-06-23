package models

type User struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	AddressId string `json:"address_id"`
	Role      string `json:"role"`
}

type Address struct {
	AddressId string `json:"address_id" gorm:"primaryKey"`
	Line1     string `json:"line1"`
	Line2     string `json:"line2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	ZipCode   string `json:"zip_code"`
}
