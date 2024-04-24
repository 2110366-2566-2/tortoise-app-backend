package models

// master data model from master_data collection

type MasterData struct {
	Category     string   `json:"category" example:"Dog"`
	SpeciesCount int      `json:"species_count" bson:"species_count" example:"5"`
	Species      []string `json:"species" example:"Golden Retriever,Poodle,Bulldog,Pug,Chihuahua"`
}

type MasterDataCategory struct {
	Categories []string `json:"categories"`
}
