package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 2005

func Notes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset-utf-8")
	fmt.Fprintln(w, "<h1>Lista de notas...</h1>")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", Notes)

	fmt.Printf("Servidor rodando na porta :%d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
