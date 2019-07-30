package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/user/vacaamarela/models"
	"github.com/user/vacaamarela/utils"
	"golang.org/x/crypto/bcrypt"
)

//Controller será exportado
type Controller struct{}

//Login será exportado ============================================
func (c Controller) Login(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var usuario models.Usuario
		var jwt models.JWT
		var error models.Error

		json.NewDecoder(r.Body).Decode(&usuario)

		// Essa é a senhaa passada pelo usuário quando enviar o request
		senha := usuario.Senha

		// Verificar se o usuário existe no DB
		row := db.QueryRow("SELECT * FROM usuario where email=$1;", usuario.Email)

		err := row.Scan(&usuario.ID, &usuario.Nome, &usuario.Sobrenome, &usuario.Email, &usuario.Senha, &usuario.CPF, &usuario.CEP, &usuario.Endereco, &usuario.Cidade, &usuario.Estado, &usuario.Celular, &usuario.Superuser, &usuario.Ativo)
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

		// Encriptar a senha recebida do DB
		hashedPassword := usuario.Senha

		// Comparar senha enviada pelo usuário e a senha equivalente encontrada no DB
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(senha))
		if err != nil {
			error.Message = "Senha inválida"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		jwt.Token = token

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, jwt)

	}

}

//UsuarioInserir será exportado ===========================================
func (c Controller) UsuarioInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var usuario models.Usuario

		json.NewDecoder(r.Body).Decode(&usuario)

		// Hash usuario.Senha
		hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), 10)
		if err != nil {
			log.Fatal(err)
		}

		// Usar hash em usuario.Senha
		usuario.Senha = string(hash)

		expressaoSQL := `INSERT INTO usuario (nome, sobrenome, email, senha, cpf, cep, endereco, cidade, estado, celular, superuser, ativo) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);`
		_, err = db.Exec(expressaoSQL, usuario.Nome, usuario.Sobrenome, usuario.Email, usuario.Senha, usuario.CPF, usuario.CEP, usuario.Endereco, usuario.Cidade, usuario.Estado, usuario.Celular, usuario.Superuser, usuario.Ativo)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("SELECT * FROM usuario WHERE email=$1;", usuario.Email)
		err = row.Scan(&usuario.ID, &usuario.Nome, &usuario.Sobrenome, &usuario.Email, &usuario.Senha, &usuario.CPF, &usuario.CEP, &usuario.Endereco, &usuario.Cidade, &usuario.Estado, &usuario.Celular, &usuario.Superuser, &usuario.Ativo)
		if err != nil {
			panic(err)
		}

		// Esconder usuario.Senha
		usuario.Senha = "********"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, usuario)

	}
}

//UsuarioTodos será exportado =======================================
func (c Controller) UsuarioTodos(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("select * from usuario")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		clts := make([]models.Usuario, 0)
		for rows.Next() {
			clt := models.Usuario{}
			err := rows.Scan(&clt.ID, &clt.Nome, &clt.Sobrenome, &clt.Email, &clt.Senha, &clt.CPF, &clt.CEP, &clt.Endereco, &clt.Cidade, &clt.Estado, &clt.Celular, &clt.Superuser, &clt.Ativo)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			clts = append(clts, clt)
		}
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Usuário inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		utils.ResponseJSON(w, clts)
	}
}

//UsuarioUnico será exportado ==================================
func (c Controller) UsuarioUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error
		var usuario models.Usuario

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo usuario no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		// O ID usaso neste argumento traz o valor inserido no Params
		row := db.QueryRow("select * from usuario where usuario_id=$1;", id)

		err = row.Scan(&usuario.ID, &usuario.Nome, &usuario.Sobrenome, &usuario.Email, &usuario.Senha, &usuario.CPF, &usuario.CEP, &usuario.Endereco, &usuario.Cidade, &usuario.Estado, &usuario.Celular, &usuario.Superuser, &usuario.Ativo)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Usuário inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, usuario)

	}
}
