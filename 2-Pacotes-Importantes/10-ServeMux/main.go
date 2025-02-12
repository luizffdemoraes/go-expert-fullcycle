package main

import "net/http"

/***

Multiplexer - ServeMux
E um componente do pacote net/http que permite criar um servidor HTTP
com multiplos handlers para diferentes rotas.

1️⃣ http.ListenAndServe(":8080", mux)
Usa o mux como roteador, que foi criado com http.NewServeMux().
O mux é um multiplexador de requisições HTTP que permite definir rotas personalizadas.

2️⃣ http.ListenAndServe(":8080", nil)
Quando passamos nil, o servidor usa o DefaultServeMux, que é o roteador padrão do pacote net/http.
Ele só funciona se as rotas forem registradas usando http.HandleFunc() ou http.Handle().


***/
func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })

	mux.HandleFunc("/", HomeHandler)     // Registra a rota no DefaultServeMux
	mux.Handle("/blog", blog{"My Blog"}) // ServeHTTP
	http.ListenAndServe(":8080", mux)    // Usa o mux como roteador
	http.ListenAndServe(":8080", nil)    // Usa o DefaultServeMux como roteador
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
