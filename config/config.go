package config

import "fmt"

//db configuration
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Epgbw2gc"
	dbname   = "gorm"
)

//SigningKey is a key used for hashing in application
var SigningKey = []byte("AllYourBase")

//GetPSQLInfo returns PostgreDB connection info as connectionString
func GetPSQLInfo() string {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlinfo
}
