package models

// pet model from pet collection
type Pet struct {
	ID            string  `json:"pet_id,omitempty"`
	Name          string  `json:"name"`
	Age           int     `json:"age"`
	Price         int     `json:"price"`
	IsSold        bool    `json:"is_sold"`
	Description   string  `json:"description"`
	Weight        float64 `json:"weight"`
	Sex           string  `json:"sex"`
	Species       string  `json:"species"`
	Type          string  `json:"type"`
	Behavior      string  `json:"behavior"`
	Media         string  `json:"media"`
	MedicalRecord struct {
		MedicalID   int    `json:"medical_id"`
		Date        string `json:"date"`
		Description string `json:"description"`
	} `json:"medical_record"`
	SellerID string `json:"seller_id,omitempty"`
}

// pet card model from pet collection
type PetCard struct {
	ID             string `json:"pet_id,omitempty"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Type           string `json:"type"`
	Media          string `json:"media"`
	SellerID       string `json:"seller_id"`
	Seller_name    string `json:"seller_name"`
	Seller_surname string `json:"seller_surname"`
}
