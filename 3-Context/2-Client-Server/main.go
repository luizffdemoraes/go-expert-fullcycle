package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/***
🔥 Casos de Uso
✅ Simular requisições demoradas (exemplo: processamento pesado ou consulta a APIs externas).
✅ Gerenciar cancelamento de requisições (quando o cliente fecha o navegador antes da resposta).
✅ Evitar sobrecarga do servidor (cancelando processos desnecessários).
***/

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciado")
	defer log.Println("Request finalizado")
	select {
	case <-time.After(5 * time.Second):
		// Imprime no comand line stdout
		log.Println("Request processada com sucesso")
		// Imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)

	}
}
