package driver

import "database/sql"

var db *sql.DB

//ConnectDB will be an exported func
func ConnectDB() *sql.DB {

	var err error

	db, err = sql.Open("postgres", "user=rodrigovalente password=Gustavo2012 host=localhost port=5432 dbname=vacaamarela sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db

}
