package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/valentergs/vacaamarela/models"
	"github.com/valentergs/vacaamarela/utils"
)

//ControllerReserva será exportado
type ControllerReserva struct{}

// //SpotInserir será exportado ===========================================
// func (c ControllerReserva) SpotInserir(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var spot models.Spot

// 		json.NewDecoder(r.Body).Decode(&spot)

// 		expressaoSQL := `INSERT INTO spot (unidade, tipo, livre) values ($1,$2,$3);`
// 		_, err := db.Exec(expressaoSQL, spot.Unidade, spot.Tipo, spot.Livre)
// 		if err != nil {
// 			panic(err)
// 		}

// 		SuccessMessage := "Spot inserido com sucesso!"

// 		w.Header().Set("Content-Type", "application/json")

// 		utils.ResponseJSON(w, SuccessMessage)

// 	}
// }

//ReservaTodos será exportado =======================================
func (c ControllerReserva) ReservaTodos(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("select * from reserva")
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

// //SpotUnico será exportado ==================================
// func (c ControllerSpot) SpotUnico(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var error models.Error
// 		var spot models.Spot

// 		if r.Method != "GET" {
// 			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 			return
// 		}

// 		// Params são os valores informados pelo spot no URL
// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			error.Message = "Numero ID inválido"
// 		}

// 		// O ID usaso neste argumento traz o valor inserido no Params
// 		row := db.QueryRow("select * from spot where spot_id=$1;", id)

// 		err = row.Scan(&spot.ID, &spot.Unidade, &spot.Tipo, &spot.Livre)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				error.Message = "Spot inexistente"
// 				utils.RespondWithError(w, http.StatusBadRequest, error)
// 				return
// 			} else {
// 				log.Fatal(err)
// 			}
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		utils.ResponseJSON(w, spot)

// 	}
// }

// //SpotApagar será exportado =========================================
// func (c ControllerSpot) SpotApagar(db *sql.DB) http.HandlerFunc {

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

// 		db.QueryRow("DELETE FROM spot where spot_id=$1;", id)

// 		SuccessMessage := "Spot deletado com sucesso!"

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		utils.ResponseJSON(w, SuccessMessage)

// 	}
// }

// //SpotEditar será exportado =========================================
// func (c ControllerSpot) SpotEditar(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var spot models.Spot
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

// 		json.NewDecoder(r.Body).Decode(&spot)

// 		expressaoSQL := `UPDATE spot SET unidade=$1, tipo=$2, livre=$3 WHERE spot_id=$4;`
// 		_, err = db.Exec(expressaoSQL, spot.Unidade, spot.Tipo, spot.Livre, id)
// 		if err != nil {
// 			panic(err)
// 		}

// 		row := db.QueryRow("SELECT * FROM spot WHERE spot_id=$1;", id)
// 		err = row.Scan(&spot.ID, &spot.Unidade, &spot.Tipo, &spot.Livre)

// 		w.Header().Set("Content-Type", "application/json")

// 		utils.ResponseJSON(w, spot)

// 	}
// }