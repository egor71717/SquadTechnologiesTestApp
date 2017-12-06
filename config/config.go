package config

import "fmt"

const (
	//db configuration
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Epgbw2gc"
	dbname   = "gorm"
	//vk auth
	VKAuthURL           string = "https://oauth.vk.com/authorize"
	VKGetAccessTokenURL string = "https://oauth.vk.com/access_token"
	VKClientID          string = "6287513"
	VKRedirectURI       string = "http://localhost:8800/auth/gettoken/vk"
	VKScope             string = "friends"
	VKSecret            string = "WHHZaX1wn7vvYwR6Je57"
)

//SigningKey is a key used for hashing in application
var SigningKey = []byte("secret")

//GetPSQLInfo returns PostgreDB connection info as connectionString
func GetPSQLInfo() string {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlinfo
}
