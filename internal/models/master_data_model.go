package models

// master data model from master_data collection

type MasterData struct {
	Category     string   `json:"category"`
	SpeciesCount int      `json:"species_count" bson:"species_count"`
	Species      []string `json:"species"`
}

type MasterDataCategory struct {
	Categories []string `json:"categories"`
}
