package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/user/vacaamarela/models"
	"github.com/user/vacaamarela/utils"
	"golang.org/x/crypto/bcrypt"
)

//Controller será exportado
type Controller struct{}

//Claims será exportado
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//Login será exportado ============================================
func (c Controller) Login(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		var unidade models.Usuario
		var jwt models.JWT
		var error models.Error

		json.NewDecoder(r.Body).Decode(&unidade)

		// Essa é a senhaa passada pelo usuário quando enviar o request
		senha := unidade.Senha

		// Verificar se o usuário existe no DB
		row := db.QueryRow("SELECT * FROM unidade where email=$1;", unidade.Email)

		err := row.Scan(&unidade.ID, &unidade.Nome, &unidade.Sobrenome, &unidade.Email, &unidade.Senha, &unidade.CPF, &unidade.CEP, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.Celular, &unidade.Superuser, &unidade.Ativo)
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
		token, err := utils.GenerateToken(unidade)
		if err != nil {
			log.Fatal(err)
		}

		// Encriptar a senha recebida do DB
		hashedPassword := unidade.Senha

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

//Logado será exportado =====================================
func (c Controller) Logado(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var unidade models.Usuario
		var jwtKey = []byte("secret")

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// this header should have a key/value pair called "Authorization". "authHeader" will grab the key
		authHeader := r.Header.Get("Authorization")
		// bearerToken will remove the empty space found on the value
		bearerToken := strings.Split(authHeader, " ")
		// here we catch the value of bearerToken[1] leaving the word "bearer" out.
		authToken := bearerToken[1]

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		row := db.QueryRow("select * from unidade where email=$1;", claims.Email)
		// err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Sobrenome, &unidade.Senha, &unidade.Email, &unidade.Celular, &unidade.Superuser, &unidade.Ativo, &unidade.Departamento)
		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Sobrenome, &unidade.Email, &unidade.Senha, &unidade.CPF, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Celular, &unidade.Superuser, &unidade.Ativo)
		if err != nil {
			fmt.Println(err)
			if err == sql.ErrNoRows {
				erro.Message = "Usuário inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, unidade)
		fmt.Println(unidade)

	}
}

//UsuarioInserir será exportado ===========================================
func (c Controller) UsuarioInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var unidade models.Usuario

		json.NewDecoder(r.Body).Decode(&unidade)

		// Hash unidade.Senha
		hash, err := bcrypt.GenerateFromPassword([]byte(unidade.Senha), 10)
		if err != nil {
			log.Fatal(err)
		}

		// Usar hash em unidade.Senha
		unidade.Senha = string(hash)

		expressaoSQL := `INSERT INTO unidade (nome, sobrenome, email, senha, cpf, endereco, cidade, estado, cep, celular, superuser, ativo) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);`
		_, err = db.Exec(expressaoSQL, unidade.Nome, unidade.Sobrenome, unidade.Email, unidade.Senha, unidade.CPF, unidade.Endereco, unidade.Cidade, unidade.Estado, unidade.CEP, unidade.Celular, unidade.Superuser, unidade.Ativo)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("SELECT * FROM unidade WHERE email=$1;", unidade.Email)
		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Sobrenome, &unidade.Email, &unidade.Senha, &unidade.CPF, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, unidade.CEP, &unidade.Celular, &unidade.Superuser, &unidade.Ativo)
		if err != nil {
			panic(err)
		}

		// Esconder unidade.Senha
		unidade.Senha = "********"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, unidade)

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

		rows, err := db.Query("select * from unidade")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		clts := make([]models.Usuario, 0)
		for rows.Next() {
			clt := models.Usuario{}
			err := rows.Scan(&clt.ID, &clt.Nome, &clt.Sobrenome, &clt.Email, &clt.Senha, &clt.CPF, &clt.Endereco, &clt.Cidade, &clt.Estado, &clt.CEP, &clt.Celular, &clt.Superuser, &clt.Ativo)
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
		var unidade models.Usuario

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo unidade no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		// O ID usaso neste argumento traz o valor inserido no Params
		row := db.QueryRow("select * from unidade where unidade_id=$1;", id)

		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Sobrenome, &unidade.Email, &unidade.Senha, &unidade.CPF, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Celular, &unidade.Superuser, &unidade.Ativo)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Usuário inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		// Esconder unidade.Senha
		unidade.Senha = "********"

		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, unidade)

	}
}

//UsuarioApagar será exportado =========================================
func (c Controller) UsuarioApagar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "DELETE" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo unidade no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		db.QueryRow("DELETE FROM unidade where unidade_id=$1;", id)

		SuccessMessage := "Usuário deletado com sucesso!"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		utils.ResponseJSON(w, SuccessMessage)

	}
}

//UsuarioEditar será exportado =========================================
func (c Controller) UsuarioEditar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var unidade models.Usuario
		var error models.Error

		if r.Method != "PUT" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		json.NewDecoder(r.Body).Decode(&unidade)

		expressaoSQL := `UPDATE unidade SET nome=$1, sobrenome=$2, email=$3, cpf=$4, endereco=$5, cidade=$6, estado=$7, cep=$8, celular=$9, superuser=$10, ativo=$11 WHERE unidade_id=$12;`
		_, err = db.Exec(expressaoSQL, unidade.Nome, unidade.Sobrenome, unidade.Email, unidade.CPF, unidade.Endereco, unidade.Cidade, unidade.Estado, unidade.CEP, unidade.Celular, unidade.Superuser, unidade.Ativo, id)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("select * from unidade where email=$1;", unidade.Email)
		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Sobrenome, &unidade.Email, &unidade.Senha, &unidade.CPF, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Celular, &unidade.Superuser, &unidade.Ativo)

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, unidade)

	}
}

//UnidadeInserir será exportado ===========================================
func (c Controller) UnidadeInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var unidade models.Unidade

		json.NewDecoder(r.Body).Decode(&unidade)

		expressaoSQL := `INSERT INTO unidade (nome, endereco, cidade, estado, cep, ativa) values ($1,$2,$3,$4,$5,$6);`
		_, err := db.Exec(expressaoSQL, unidade.Nome, unidade.Endereco, unidade.Cidade, unidade.Estado, unidade.CEP, unidade.Ativa)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("SELECT * FROM unidade WHERE nome=$1;", unidade.Nome)
		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Ativa)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, unidade)

	}
}

//UnidadeTodos será exportado =======================================
func (c Controller) UnidadeTodos(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("select * from unidade")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		clts := make([]models.Unidade, 0)
		for rows.Next() {
			clt := models.Unidade{}
			err := rows.Scan(&clt.ID, &clt.Nome, &clt.Endereco, &clt.Cidade, &clt.Estado, &clt.CEP, &clt.Ativa)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			clts = append(clts, clt)
		}
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Unidade inexistente"
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

//UnidadeUnico será exportado ==================================
func (c Controller) UnidadeUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error
		var unidade models.Unidade

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo unidade no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		// O ID usaso neste argumento traz o valor inserido no Params
		row := db.QueryRow("select * from unidade where unidade_id=$1;", id)

		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Ativa)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Unidade inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, unidade)

	}
}
