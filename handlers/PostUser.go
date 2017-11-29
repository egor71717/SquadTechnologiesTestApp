package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egor71717/SquadTechnologiesTestApp/database"
	"github.com/egor71717/SquadTechnologiesTestApp/model"
	"golang.org/x/crypto/bcrypt"
)

//PostUser handler used for creating new user accounts
func PostUser(db database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var auth = model.StandardAuth{}
		json.NewDecoder(r.Body).Decode(&auth)

		//encode password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//maybe some validation ? not today for sure)
		db.CreateUser(&model.User{Login: auth.Login, PasswordHash: string(hashedPassword)})
		w.WriteHeader(http.StatusOK)
	})
}
