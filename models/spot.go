package models

//Spot is an exportable type
type Spot struct {
	ID        int64   `json:"spot_id"`
	Unidade   int64   `json:"unidade"`
	Tipo      string  `json:"tipo"`
	Preco     float64 `json:"preco"`
	Livre     bool    `json:"livre"`
	Bloqueado bool    `json:"bloqueado"`
	Height    string  `json:"height"`
	Width     string  `json:"width"`
	Y         string  `json:"y"`
	X         string  `json:"x"`
	Transform string  `json:"transform"`
}
