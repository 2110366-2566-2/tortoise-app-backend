package models

// user model from user collection
type User struct {
	User_id     string `json:"user_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Image       string `json:"image"`
	Role        int    `json:"role"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     struct {
		Province    string `json:"province"`
		District    string `json:"district"`
		SubDistrict string `json:"sub_district"`
		PostalCode  string `json:"postal_code"`
		Street      string `json:"street"`
		Building    string `json:"building"`
		HouseNumber string `json:"house_number"`
	} `json:"address"`
}
