package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/trevisharp/celltomata/api/application/services"
	"github.com/trevisharp/celltomata/api/application/usecases"
	"github.com/trevisharp/celltomata/api/infrastructure"
)

func main() {
	router := chi.NewRouter()
	var mailService = infrastructure.GomailEmailService{}
	var validAccountService = services.ValidateAccountService{
		EmailService: mailService,
	}
	var cryptoService = infrastructure.BCryptService{}
	var userRepo = infrastructure.SupabaseUserRepository{}

	usecases.CreateUserUseCase(router, cryptoService, userRepo, validAccountService)

	log.Println("listening on localhost:3000/")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}
