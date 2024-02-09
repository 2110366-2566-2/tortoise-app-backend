package models

// user model from user collection
type User struct {
	ID          string `json:"_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	Image       string `json:"image"`
	Role        int    `json:"role"`
	Address     struct {
		Province    string `json:"province"`
		District    string `json:"district"`
		SubDistrict string `json:"subdistrict"`
		PostalCode  string `json:"postalCode"`
		Street      string `json:"street"`
		Building    string `json:"building"`
		HouseNumber string `json:"houseNumber"`
	} `json:"address"`
}
