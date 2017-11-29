package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/egor71717/SquadTechnologiesTestApp/config"
	"github.com/egor71717/SquadTechnologiesTestApp/database"
	"github.com/egor71717/SquadTechnologiesTestApp/model"
	"golang.org/x/crypto/bcrypt"
)

//PostLogin handler used for JWT authentication
func PostLogin(db database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//find user
		var auth = model.StandardAuth{}
		json.NewDecoder(r.Body).Decode(&auth)
		user, err := db.GetUserByLogin(auth.Login)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		//check pass
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(auth.Password))
		if err != nil {
			http.Error(w, "Invalid Password", http.StatusUnauthorized)
			return
		}

		// Create a new token object, specifying signing method and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": user.Login,
			"iat": time.Now(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(config.SigningKey)
		if err != nil {
			http.Error(w, "Signing Error", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(tokenString))
	})
}
