package main

import (
	"log"
	"net/http"

	"adquirer/db"
	"adquirer/handler"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Falha na conexão: %v", err)
	}
	defer conn.Close()

	http.HandleFunc("/adquirer/valid", handler.ValidateCardHandler)

	log.Println("API iniciada na porta 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
