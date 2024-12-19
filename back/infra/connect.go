package infra

import (
	"os"

	"github.com/ei-sugimoto/pokepoke/back/ent"
)

var host string
var port string
var user string
var password string
var dbname string

func init() {
	if os.Getenv("PROD") == "true" {
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
		if host == "" || port == "" || user == "" || password == "" || dbname == "" {
			panic("env vars not set")
		}
	} else {
		host = "db"
		port = "5432"
		user = "postgres"
		password = "postgres"
		dbname = "pokepoke"
	}
}

func Conn() *ent.Client {
	client, err := ent.Open("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		panic("db connecting error: " + err.Error())
	}
	return client
}
