package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/valentergs/vacaamarela/controllers"
	"github.com/valentergs/vacaamarela/driver"
)

var db *sql.DB

func main() {

	db := driver.ConnectDB()
	usuarioctl := controllers.ControllerUsuario{}
	unidadectl := controllers.ControllerUnidade{}

	// gorilla.mux
	router := mux.NewRouter()

	// USUARIO URLs ===================================================

	router.HandleFunc("/login", usuarioctl.Login(db)).Methods("POST")
	router.HandleFunc("/logado", usuarioctl.Logado(db)).Methods("GET")
	router.HandleFunc("/usuario", usuarioctl.UsuarioTodos(db)).Methods("GET")
	router.HandleFunc("/usuario/inserir", usuarioctl.UsuarioInserir(db)).Methods("POST")
	router.HandleFunc("/usuario/{id}", usuarioctl.UsuarioUnico(db)).Methods("GET")
	router.HandleFunc("/usuario/apagar/{id}", usuarioctl.UsuarioApagar(db)).Methods("DELETE")
	router.HandleFunc("/usuario/editar/{id}", usuarioctl.UsuarioEditar(db)).Methods("PUT")

	// USUARIO com TokenVerification ===================================

	// router.HandleFunc("/logged", controller.Logged(db)).Methods("GET")
	// //router.HandleFunc("/usuario", middlewares.TokenVerifyMiddleware(controller.UsuarioGetAll(db))).Methods("GET")
	// //router.HandleFunc("/usuario/{id}", middlewares.TokenVerifyMiddleware(controller.UsuarioGetOne(db))).Methods("GET")
	// router.HandleFunc("/search/usuario", controller.Search(db)).Methods("GET").Queries("q", "{q}")

	// UNIDADE URLs ===================================================

	router.HandleFunc("/unidade", unidadectl.UnidadeTodos(db)).Methods("GET")
	router.HandleFunc("/unidade/inserir", unidadectl.UnidadeInserir(db)).Methods("POST")
	router.HandleFunc("/unidade/{id}", unidadectl.UnidadeUnico(db)).Methods("GET")
	router.HandleFunc("/unidade/apagar/{id}", unidadectl.UnidadeApagar(db)).Methods("DELETE")
	router.HandleFunc("/unidade/editar/{id}", unidadectl.UnidadeEditar(db)).Methods("PUT")

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
