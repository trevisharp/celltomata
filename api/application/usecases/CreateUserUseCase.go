package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/trevisharp/celltomata/api/application/payloads"
	appServices "github.com/trevisharp/celltomata/api/application/services"
	"github.com/trevisharp/celltomata/api/domain/services"
	domServices "github.com/trevisharp/celltomata/api/domain/services"
)

func CreateUserUseCase(
	router *chi.Mux,
	cryptoService appServices.CryptoService,
	userRepo domServices.UserRepository,
	valAccount appServices.ValidateAccountService) {

	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		var body payloads.CreateUserRequest

		var err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request body.", http.StatusBadRequest)
			return
		}

		var user services.UserData
		user.Username = body.Username
		user.Password = body.Password
		user.Email = body.Email
		user.Validated = false

		err = user.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user.Password, err = cryptoService.EncryptPassword(user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var findedUser, _ = userRepo.Find(user.Username)
		if findedUser != nil {
			http.Error(w, "username already is in use.", http.StatusBadRequest)
			return
		}

		err = userRepo.Create(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		valAccount.SendEmail(user.Username, user.Email)

		w.WriteHeader(http.StatusOK)
	})
}
