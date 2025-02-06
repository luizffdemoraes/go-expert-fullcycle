package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// A interface no go só permite que seja informato métodos, não atributos
type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e *Empresa) Desativar() {
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado %t\n", c.Nome, c.Ativo)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

// main é a função principal que será executada ao rodar o programa utilizara de composição para criar um endereço para o cliente
func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}
	Desativacao(&wesley)
	Desativacao(&minhaEmpresa)

}
