package main

import (
	"log"
	"net/http"
	"os"

	email "github.com/Gabnatali/reto-tecnico/Email"
	authmiddleware "github.com/Gabnatali/reto-tecnico/authMiddleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(authmiddleware.Cors)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env:", err)
	}

	HTTP_API := os.Getenv("HTTP_API")
	ORG_ID := os.Getenv("ORG_ID")
	STREAM := os.Getenv("STREAM")
	USER := os.Getenv("ZO_ROOT_USER_EMAIL")
	PASSWORD := os.Getenv("ZO_ROOT_USER_PASSWORD")

	openObserverClient := email.NewOpenObserverClient(
		email.OpenObserverOptions{
			OrgID:      ORG_ID,
			StreamName: STREAM,
			User:       USER,
			Password:   PASSWORD,

			Http_api: HTTP_API,
		},
	)

	emailHandler := email.NewEmailHandler(*openObserverClient)

	r.Get("/email/{id}", emailHandler.GetEmailById)
	r.Get("/email", emailHandler.GetAllEmails)

	log.Println("Servidor corriendo en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
