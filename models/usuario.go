package models

//Usuario is an exportable type
type Usuario struct {
	ID        int    `json:"usuario_id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
	CPF       string `json:"cpf"`
	Endereco  string `json:"endereco"`
	Cidade    string `json:"cidade"`
	Estado    string `json:"estado"`
	CEP       string `json:"cep"`
	Celular   string `json:"celular"`
	Superuser bool   `json:"superuser"`
	Ativo     bool   `json:"ativo"`
}
