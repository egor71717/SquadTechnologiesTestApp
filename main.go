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
		entry   string = "./ui/index.html"
		assets  string = "./ui/dist"
		dialect string = "postgres"
	)
	connectionString := config.GetPSQLInfo()

	mux := http.NewServeMux()

	//proper way to do it
	// db, err := database.InitDb(connectionString)
	// if err != nil {
	// 	panic(err)
	// }

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
	mux.Handle("/api/users/getall", handlers.GetUsers(psqldb))
	//GET  get single user json; example: /api/users/get?id=1
	mux.Handle("/api/users/get", handlers.GetUser(psqldb))
	//POST create a new user
	mux.Handle("/api/users/create", handlers.PostUser(psqldb))
	//POST login to application and get the JWT as response body
	mux.Handle("/api/login", handlers.PostLogin(psqldb))
	//serve index.html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entry)
	})
	log.Println("serving on port ", port)
	err = http.ListenAndServe(port, mux)
	log.Fatal(err)
}
