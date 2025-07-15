package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/trevisharp/celltomata/application/usecases"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Celltomata is running..."))
		if err != nil {
			log.Println(err)
		}
	})

	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		var body CreateUserRequest

		var err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request body.", http.StatusBadRequest)
			return
		}

		err = usecases.CreateUser(body.Username, body.Email, body.Password, body.RepeatPassord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Println("listening on localhost:3000/")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}
