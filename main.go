package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/valentergs/vacaamarela/controllers"
	"github.com/valentergs/vacaamarela/driver"
)

var db *sql.DB

func main() {

	db := driver.ConnectDB()
	controller := controllers.Controller{}

	// gorilla.mux
	router := mux.NewRouter()
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	// router.HandleFunc("/logged", controller.Logged(db)).Methods("GET")
	// router.HandleFunc("/usuario/add", controller.UsuarioAdd(db)).Methods("POST")
	// //router.HandleFunc("/usuario", middlewares.TokenVerifyMiddleware(controller.UsuarioGetAll(db))).Methods("GET")
	// router.HandleFunc("/usuario", controller.UsuarioGetAll(db)).Methods("GET")
	// //router.HandleFunc("/usuario/{id}", middlewares.TokenVerifyMiddleware(controller.UsuarioGetOne(db))).Methods("GET")
	// router.HandleFunc("/usuario/{id}", controller.UsuarioGetOne(db)).Methods("GET")
	// router.HandleFunc("/search/usuario", controller.Search(db)).Methods("GET").Queries("q", "{q}")
	// router.HandleFunc("/usuario/delete/{id}", controller.UsuarioDeleteOne(db)).Methods("DELETE")
	// router.HandleFunc("/usuario/edit/{id}", controller.UsuarioUpdate(db)).Methods("PUT")

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
	// 	AllowCredentials: true,
	// 	// Enable Debugging for testing, consider disabling in production
	// 	AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
	// 	AllowedHeaders: []string{"Authorization", "Content-Type"},
	// 	Debug:          true,
	// })

	// // Insert the middleware
	// handler := c.Handler(router)

	log.Println("Listen on port 8080...")
	// Quando for usar o CORS colocar "handler" no lugar do "router"
	log.Fatal(http.ListenAndServe(":8080", router))
}
