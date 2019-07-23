package main

import (
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/user/marvel/src/docs" // docs is generated by Swag CLI via swag init command on src folder.
	"github.com/user/marvel/src/routes"
	"log"
	"net/http"
)

// @title Marvel Example API
// @version 1.0
// @description This is a sample server marvel server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1

func main() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	var r = mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router := r.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/characters", routes.GetCharacters).Methods("GET")
	router.HandleFunc("/characters/{id}", routes.GetCharacterById).Methods("GET")
	router.HandleFunc("/comics", routes.GetComics).Methods("GET")

	fmt.Println("Running server ...")
	log.Fatal(http.ListenAndServe(":3000", r))
}