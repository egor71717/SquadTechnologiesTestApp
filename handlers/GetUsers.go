package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/egor71717/SquadTechnologiesTestApp/database"
)

//GetUsers handler
func GetUsers(db database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val, err := db.GetUsers()
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
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
