package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {

	// Criando um objeto
	conta := Conta{Numero: 1, Saldo: 100}
	// Convertendo conta para JSON
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	// Convertendo JSON para Conta
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	// Convertendo JSON puro
	jsonPuro := []byte(`{"n": 1, "s": 100}`)
	var contaPadrao Conta
	// Convertendo JSON puro para Conta
	err = json.Unmarshal(jsonPuro, &contaPadrao)
	if err != nil {
		panic(err)
	}
	println(contaPadrao.Saldo)

}
