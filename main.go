package main

import (
	"adquirer/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/adquirer/valid", handler.ValidateCardHandler)

	log.Println("API iniciada na porta 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
