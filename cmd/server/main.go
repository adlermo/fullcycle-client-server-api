package main

import (
	"log"
	"net/http"

	"fullcycle-client-server-api/internal/database"
	"fullcycle-client-server-api/internal/handler"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.CreateTable(db)

	http.HandleFunc("/cotacao", handler.CotacaoHandler(db))

	log.Println("Server rodando na :8080")
	http.ListenAndServe(":8080", nil)
}
