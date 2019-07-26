package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/user/vacaamarela/models"
	"github.com/user/vacaamarela/utils"
)

//Controller will be exported
type Controller struct{}

//Login will be exported
func (c Controller) Login(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var usuario models.Usuario
		var jwt models.JWT
		var error models.Error

		json.NewDecoder(r.Body).Decode(&usuario)

		// Essa é a senhaa passada pelo usuário quando enviar o request
		// senha := usuario.Senha

		// Verificar se o usuário existe no DB
		row := db.QueryRow("SELECT * FROM usuario where email=$1;", usuario.Email)

		err := row.Scan(&usuario.ID, &usuario.Nome, &usuario.Sobrenome, &usuario.Email, &usuario.Senha, &usuario.Endereco, &usuario.Cidade, &usuario.Estado, &usuario.CEP, &usuario.Celular, &usuario.Superuser, &usuario.Ativo)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Usuário inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		// Gerar token de validação para enviar ao client
		token, err := utils.GenerateToken(usuario)
		if err != nil {
			log.Fatal(err)
		}

		// // Encriptar a senha recebida do DB
		// hashedPassword := usuario.Senha

		// // Comparar senha enviada pelo usuário e a senha equivalente encontrada no DB
		// err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(senha))
		// if err != nil {
		// 	error.Message = "Senha inválida"
		// 	utils.RespondWithError(w, http.StatusUnauthorized, error)
		// 	return
		// }

		jwt.Token = token
		utils.ResponseJSON(w, jwt)

	}

}
