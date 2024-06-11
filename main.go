package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "5432"
	dbname   = "sobrevidas-acs"
)

func main() {
	// Constrói uma string de conexão
	psqlInfoBanco := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("postgres", psqlInfoBanco)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verifica a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// Print de confirmação para ser retirado
	fmt.Println("Conexão bem-sucedida!")
}
