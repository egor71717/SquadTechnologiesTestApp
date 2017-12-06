package main

import (
	"log"
	"net/http"

	"github.com/egor71717/SquadTechnologiesTestApp/config"
	"github.com/egor71717/SquadTechnologiesTestApp/database"
	"github.com/egor71717/SquadTechnologiesTestApp/handlers"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	const (
		port    string = ":8800"
		auth    string = "./ui/auth"
		entry   string = "./ui/index.html"
		assets  string = "./ui/dist"
		dialect string = "postgres"
	)
	connectionString := config.GetPSQLInfo()

	mux := http.NewServeMux()

	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// db.DropTableIfExists(&model.User{})
	// db.CreateTable(&model.User{})
	psqldb := database.PostgreSQLDB{Db: db}

	//serve static assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assets))))
	//GET  get user by login (authentication required)
	mux.Handle("/api/userdata", handlers.AuthJWTMiddleware.Handler(handlers.GetUserByLogin(psqldb)))
	//GET  get all users json
	mux.Handle("/api/users/create", handlers.PostUser(psqldb))
	//POST login to application and get the JWT as response body
	mux.Handle("/api/login", handlers.PostLogin(psqldb))
	//serve index.html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entry)
	})

	//Auth
	mux.Handle("/auth/", http.StripPrefix("/auth/", http.FileServer(http.Dir(auth))))
	//redirect get vk code
	mux.HandleFunc("/auth/getcode/vk", func(w http.ResponseWriter, r *http.Request) {
		//var url = "https://oauth.vk.com/authorize?client_id=6287513&scope=friends&redirect_uri=http://localhost:8800/auth/gettoken/vk"
		var url = config.VKAuthURL + "?client_id=" + config.VKClientID + "&scope=" + config.VKScope + "&redirect_uri=" + config.VKRedirectURI
		http.Redirect(w, r, url, http.StatusFound)
	})
	//get vk token
	mux.Handle("/auth/gettoken/vk", handlers.VKGetToken(psqldb))

	log.Println("serving on port ", port)
	err = http.ListenAndServe(port, mux)
	//err = http.ListenAndServeTLS(port, "certificate/cert.pem", "certificate/key.pem", mux)
	log.Fatal(err)
}
