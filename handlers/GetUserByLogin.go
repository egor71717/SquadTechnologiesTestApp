package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egor71717/SquadTechnologiesTestApp/database"
)

//GetUserByLogin handler
func GetUserByLogin(db database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login := r.URL.Query().Get("login")
		if login == "" {
			http.Error(w, "missing login name in query string", http.StatusBadRequest)
			return
		}
		val, err := db.GetUserByLogin(login)
		if err == database.ErrorNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		data, err := json.Marshal(val)
		if err != nil {
			http.Error(w, "JSON serializing error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
}
