package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/egor71717/SquadTechnologiesTestApp/database"
)

//GetUser handler
func GetUser(db database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id name in query string", http.StatusBadRequest)
			return
		}
		intid, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "wrong query param", http.StatusBadRequest)
			return
		}
		val, err := db.GetUser(intid)
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
