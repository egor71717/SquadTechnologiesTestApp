package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/egor71717/SquadTechnologiesTestApp/config"
	"github.com/egor71717/SquadTechnologiesTestApp/database"
)

//VKGetToken get vk acccess_token
func VKGetToken(db database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		accessToken := r.URL.Query().Get("access_token")
		if accessToken != "" {
			w.Write([]byte(accessToken))
			return
		} else if code != "" {
			client := &http.Client{}
			//var url = "https://oauth.vk.com/access_token?client_id=6287513&client_secret=WHHZaX1wn7vvYwR6Je57&code=a4465cda3aaf468ccc&redirect_uri=http://localhost:8800/auth/gettoken/vk"
			var url = config.VKGetAccessTokenURL + "?client_id=" + config.VKClientID + "&redirect_uri=" + config.VKRedirectURI + "&code=" + code + "&client_secret=" + config.VKSecret
			log.Println("access_token_url", url)
			resp, err := client.Get(url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(data)
		} else {
			http.Error(w, "unexpected error", http.StatusInternalServerError)
		}
	})
}
