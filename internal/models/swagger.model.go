package models

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"some error message"`
}

type PetResponse struct {
	Success bool `json:"success" example:"true"`
	Data    Pet  `json:"data"`
}

type PetRequest struct {
	Name            string           `json:"name" example:"Fluffy"`
	Age             int32            `json:"age" example:"3"`
	Price           int32            `json:"price" example:"500"`
	Is_sold         bool             `json:"is_sold" example:"false"`
	Description     string           `json:"description" example:"A friendly and playful dog"`
	Weight          float64          `json:"weight" example:"25.5"`
	Sex             string           `json:"sex" example:"male"`
	Category        string           `json:"category" example:"Dog"`
	Species         string           `json:"species" example:"Golden Retriever"`
	Behavior        string           `json:"behavior" example:"Friendly"`
	Media           string           `json:"media" example:"https://example.com/fluffy.jpg"`
	Medical_records []Medical_record `json:"medical_records" bson:"medical_records"`
}

type PetCardResponse struct {
	Success bool      `json:"success" example:"true"`
	Count   int       `json:"count" example:"1"`
	Data    []PetCard `json:"data"`
}

type DeletePetResponse struct {
	Success     bool `json:"success" example:"true"`
	DeleteCount int  `json:"delete_count" example:"1"`
}
