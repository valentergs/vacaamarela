package models

//Spot is an exportable type
type Spot struct {
	ID      int    `json:"spot_id"`
	Unidade int    `json:"unidade"`
	Tipo    string `json:"tipo"`
	Livre   bool   `json:"livre"`
}
