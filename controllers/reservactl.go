package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/valentergs/vacaamarela/models"
	"github.com/valentergs/vacaamarela/utils"
)

//ControllerReserva será exportado
type ControllerReserva struct{}

//ReservaInserir será exportado ===========================================
func (c ControllerReserva) ReservaInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var reserva models.Reserva

		json.NewDecoder(r.Body).Decode(&reserva)

		expressaoSQL := `INSERT INTO reserva (usuario, spot) values ($1,$2);`
		_, err := db.Exec(expressaoSQL, reserva.Usuario, reserva.Spot)
		if err != nil {
			panic(err)
		}

		SuccessMessage := "Reserva criada com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//ReservaTodos será exportado =======================================
func (c ControllerReserva) ReservaTodos(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		//rows, err := db.Query("select * from reserva")
		rows, err := db.Query(`SELECT reserva.reserva_id, usuario.nome, sobrenome, unidade.nome AS unidade, spot_id AS spot, tipo AS tipo_de_spot, hora_inicio, hora_fim
		FROM reserva
		INNER JOIN usuario ON usuario.usuario_id = reserva.usuario
		INNER JOIN spot ON spot.spot_id = reserva.spot
		INNER JOIN unidade ON unidade.unidade_id = spot.unidade`)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		// Usar models.Reserva com para pesquisa pura, sem INNER JOIN.
		// Usar models.ReservaJoin para incluir pesquisa com INNER JOIN, incluindo tabelas "usuario", "spot" e "unidade"
		clts := make([]models.ReservaJoin, 0)
		for rows.Next() {
			clt := models.ReservaJoin{}
			err := rows.Scan(&clt.ID, &clt.Nome, &clt.Sobrenome, &clt.Unidade, &clt.Spot, &clt.Tipo, &clt.HoraInicio, &clt.HoraFim)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				fmt.Println(err)
				return
			}
			clts = append(clts, clt)
		}
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Reserva inexistente"
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

//ReservaAberta será exportado =======================================
func (c ControllerReserva) ReservaAberta(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("SELECT * FROM reserva WHERE hora_fim is null;")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		defer rows.Close()

		clts := make([]models.Reserva, 0)
		for rows.Next() {
			clt := models.Reserva{}
			err := rows.Scan(&clt.ID, &clt.Usuario, &clt.Spot, &clt.HoraInicio, &clt.HoraFim)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				fmt.Println(err)
				return
			}
			clts = append(clts, clt)
		}
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Reserva inexistente"
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

//ReservaUnico será exportado ==================================
func (c ControllerReserva) ReservaUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error
		//var reserva models.Reserva

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo spot no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		// O ID usaso neste argumento traz o valor inserido no Params
		rows, err := db.Query("select * from reserva where reserva_id=$1;", id)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		// err = row.Scan(&reserva.ID, &reserva.Usuario, &reserva.Spot, &reserva.HoraInicio, &reserva.HoraFim)
		// if err != nil {
		// 	if err == sql.ErrNoRows {
		// 		error.Message = "Reserva inexistente"
		// 		utils.RespondWithError(w, http.StatusBadRequest, error)
		// 		return
		// 	} else {
		// 		log.Fatal(err)
		// 	}
		// }

		clts := make([]models.Reserva, 0)
		for rows.Next() {
			clt := models.Reserva{}
			err := rows.Scan(&clt.ID, &clt.Usuario, &clt.Spot, &clt.HoraInicio, &clt.HoraFim)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				fmt.Println(err)
				return
			}
			clts = append(clts, clt)
		}
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Reserva inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, clts)

	}
}

//ReservaApagar será exportado =========================================
func (c ControllerReserva) ReservaApagar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "DELETE" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo usuario no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		db.QueryRow("DELETE FROM reserva where reserva_id=$1;", id)

		SuccessMessage := "Reserva deletada com sucesso!"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		utils.ResponseJSON(w, SuccessMessage)

	}
}

// Criar RESER FECHAR
