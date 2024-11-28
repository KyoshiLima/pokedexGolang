package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitiDB(){
	var err error
	DB, err = sql.Open("postgres", "user=postgress dbname=pokedex sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}

	log.Println("Conectado ao banco de dados!")
}