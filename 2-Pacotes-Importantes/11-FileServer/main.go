package main

import (
	"log"
	"net/http"
)

// Criar um servidor de arquivos
func main() {

	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog!"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
