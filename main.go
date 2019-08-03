package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/user/vacaamarela/controllers"
	"github.com/user/vacaamarela/driver"
)

var db *sql.DB

func main() {

	db := driver.ConnectDB()
	controller := controllers.Controller{}

	// gorilla.mux
	router := mux.NewRouter()

	// USUARIO URLs ===================================================

	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/logado", controller.Logado(db)).Methods("GET")
	router.HandleFunc("/usuario", controller.UsuarioTodos(db)).Methods("GET")
	router.HandleFunc("/usuario/inserir", controller.UsuarioInserir(db)).Methods("POST")
	router.HandleFunc("/usuario/{id}", controller.UsuarioUnico(db)).Methods("GET")
	router.HandleFunc("/usuario/apagar/{id}", controller.UsuarioApagar(db)).Methods("DELETE")
	router.HandleFunc("/usuario/editar/{id}", controller.UsuarioEditar(db)).Methods("PUT")

	// USUARIO com TokenVerification ===================================

	// router.HandleFunc("/logged", controller.Logged(db)).Methods("GET")
	// //router.HandleFunc("/usuario", middlewares.TokenVerifyMiddleware(controller.UsuarioGetAll(db))).Methods("GET")
	// //router.HandleFunc("/usuario/{id}", middlewares.TokenVerifyMiddleware(controller.UsuarioGetOne(db))).Methods("GET")
	// router.HandleFunc("/search/usuario", controller.Search(db)).Methods("GET").Queries("q", "{q}")

	// UNIDADE URLs ===================================================

	router.HandleFunc("/unidade", controller.UnidadeTodos(db)).Methods("GET")
	router.HandleFunc("/unidade/inserir", controller.UnidadeInserir(db)).Methods("POST")
	router.HandleFunc("/unidade/{id}", controller.UnidadeUnico(db)).Methods("GET")
	// router.HandleFunc("/unidade/apagar/{id}", controller.UnidadeApagar(db)).Methods("DELETE")
	// router.HandleFunc("/unidade/editar/{id}", controller.UnidadeEditar(db)).Methods("PUT")

	// CORS ==========================================================

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})

	// Insert the middleware
	handler := c.Handler(router)

	log.Println("Listen on port 8080...")
	// Quando  usar o CORS colocar "handler" no lugar do "router"
	// Quando usar ROUTER usar "router"
	log.Fatal(http.ListenAndServe(":8080", handler))
}
