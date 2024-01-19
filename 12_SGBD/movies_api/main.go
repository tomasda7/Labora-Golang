package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tomasda7/movies_api/controllers"
	"github.com/tomasda7/movies_api/services"
)

func main() {
	// db conn
	err := services.EstablishDbConn()

	if err != nil {
		log.Fatal(err)
	}

	// routes
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/movies", controllers.GetMovies).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/movies/create", controllers.CreateMovie).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/movies/update/{id:[0-9]+}", controllers.UpdateMovie).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/movies/delete/{id:[0-9]+}", controllers.DeleteMovie).Methods(http.MethodDelete)

	// CORS config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5433"},
		AllowedMethods: []string{"GET", "POST"},
	})

	// launching server
	handler := corsOptions.Handler(router)
	PORT := ":8080"

	if err := startServer(PORT, handler); err != nil {
		log.Fatalf("Error trying to launch server \n%v", err)
	}
}

func startServer(port string, router http.Handler) error {
	server := &http.Server{
		Handler: router,
		Addr: port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	fmt.Println("Server started at <http://localhost" + port + ">" + " ...")

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
