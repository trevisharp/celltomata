package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Celltomata is running..."))
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("listening on localhost:3000/")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}
