package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
**

	Estrutura para armazenar a resposta da API do ViaCEP
	https://mholt.github.io/json-to-go/
	https://viacep.com.br/ws/01001000/json/

**
*/
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	/***
	Imprimindo os argumentos recebidos via linha de comando
	Exemplo de execução: go run .\main.go 01001000
	go build -o cep.exe main.go
	cep.exe 01001000
	***/
	for _, cep := range os.Args[1:] {
		// Fazendo uma requisição HTTP GET para a URL informada
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			// Se houver erro na requisição, imprime no stderr e continua
			fmt.Fprintf(os.Stderr, "Erro ao buscar o CEP: %v\n", err)
			continue
		}
		// Garante que o corpo da resposta será fechado ao final da execução
		defer req.Body.Close()

		// Lendo o corpo da resposta HTTP
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o corpo da resposta: %v\n", err)
			continue
		}

		// Criando uma variável para armazenar os dados do JSON
		var data ViaCEP

		// Convertendo o JSON recebido para a estrutura ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter o JSON: %v\n", err)
		}
		// Criando o arquivo para armazenar os dados convertidos
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
		}
		defer file.Close()

		// Escrevendo os dados do ViaCEP no arquivo
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s\n", data.Cep, data.Localidade, data.Uf))
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade: ", data.Localidade)
	}
}
