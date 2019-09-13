package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/valentergs/vacaamarela/models"
	"github.com/valentergs/vacaamarela/utils"
)

//ControllerSpot será exportado
type ControllerSpot struct{}

// //UnidadeInserir será exportado ===========================================
// func (c ControllerUnidade) UnidadeInserir(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var unidade models.Unidade

// 		json.NewDecoder(r.Body).Decode(&unidade)

// 		expressaoSQL := `INSERT INTO unidade (nome, endereco, cidade, estado, cep, ativa, novo) values ($1,$2,$3,$4,$5,$6,$7);`
// 		_, err := db.Exec(expressaoSQL, unidade.Nome, unidade.Endereco, unidade.Cidade, unidade.Estado, unidade.CEP, unidade.Ativa)
// 		if err != nil {
// 			panic(err)
// 		}

// 		row := db.QueryRow("SELECT * FROM unidade WHERE nome=$1;", unidade.Nome)
// 		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Ativa)
// 		if err != nil {
// 			panic(err)
// 		}

// 		w.Header().Set("Content-Type", "application/json")

// 		utils.ResponseJSON(w, unidade)

// 	}
// }

//SpotTodos será exportado =======================================
func (c ControllerSpot) SpotTodos(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("select * from spot")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		clts := make([]models.Spot, 0)
		for rows.Next() {
			clt := models.Spot{}
			err := rows.Scan(&clt.ID, &clt.Unidade, &clt.Tipo, &clt.Livre)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			clts = append(clts, clt)
		}
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Spot inexistente"
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

// //UnidadeUnico será exportado ==================================
// func (c ControllerUnidade) UnidadeUnico(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var error models.Error
// 		var unidade models.Unidade

// 		if r.Method != "GET" {
// 			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 			return
// 		}

// 		// Params são os valores informados pelo unidade no URL
// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			error.Message = "Numero ID inválido"
// 		}

// 		// O ID usaso neste argumento traz o valor inserido no Params
// 		row := db.QueryRow("select * from unidade where unidade_id=$1;", id)

// 		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Ativa)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				error.Message = "Unidade inexistente"
// 				utils.RespondWithError(w, http.StatusBadRequest, error)
// 				return
// 			} else {
// 				log.Fatal(err)
// 			}
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		utils.ResponseJSON(w, unidade)

// 	}
// }

// //UnidadeApagar será exportado =========================================
// func (c ControllerUnidade) UnidadeApagar(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var error models.Error

// 		if r.Method != "DELETE" {
// 			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 			return
// 		}

// 		// Params são os valores informados pelo usuario no URL
// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			error.Message = "Numero ID inválido"
// 		}

// 		db.QueryRow("DELETE FROM unidade where unidade_id=$1;", id)

// 		SuccessMessage := "Unidade deletada com sucesso!"

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		utils.ResponseJSON(w, SuccessMessage)

// 	}
// }

// //UnidadeEditar será exportado =========================================
// func (c ControllerUnidade) UnidadeEditar(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var unidade models.Unidade
// 		var error models.Error

// 		if r.Method != "PUT" {
// 			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 			return
// 		}

// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			error.Message = "Numero ID inválido"
// 		}

// 		json.NewDecoder(r.Body).Decode(&unidade)

// 		expressaoSQL := `UPDATE unidade SET nome=$1, endereco=$2, cidade=$3, estado=$4, cep=$5, ativa=$6 WHERE unidade_id=$7;`
// 		_, err = db.Exec(expressaoSQL, unidade.Nome, unidade.Endereco, unidade.Cidade, unidade.Estado, unidade.CEP, unidade.Ativa, id)
// 		if err != nil {
// 			panic(err)
// 		}

// 		row := db.QueryRow("select * from unidade where unidade_id=$1;", id)
// 		err = row.Scan(&unidade.ID, &unidade.Nome, &unidade.Endereco, &unidade.Cidade, &unidade.Estado, &unidade.CEP, &unidade.Ativa)

// 		w.Header().Set("Content-Type", "application/json")

// 		utils.ResponseJSON(w, unidade)

// 	}
// }
