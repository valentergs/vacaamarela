package models

//Spot is an exportable type
type Spot struct {
	ID        int    `json:"spot_id"`
	Unidade   int    `json:"unidade"`
	Tipo      string `json:"tipo"`
	Livre     bool   `json:"livre"`
	Bloqueado bool   `json:"bloqueado"`
	Height    string `json:"height"`
	Width     string `json:"width"`
	Y         string `json:"y"`
	X         string `json:"x"`
}
