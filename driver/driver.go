package driver

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

//ConnectDB will be an exported func
func ConnectDB() *sql.DB {

	var err error

	const (
		user     = "rodrigovalente"
		password = "Gustavo2012"
		// Quando rodar dentro do container Docker
		// host = "172.17.0.3"
		// Quando rodar sem Docker
		host   = "localhost"
		port   = 5432
		dbname = "vacaamarela"
	)

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db

}
