package models

//Unidade is an exportable type
type Unidade struct {
	ID       int    `json:"unidade_id"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
	Cidade   string `json:"cidade"`
	Estado   string `json:"estado"`
	CEP      string `json:"cep"`
	Ativa    bool   `json:"ativa"`
}
