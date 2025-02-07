package main

type Conta struct {
	saldo int
}

/*
**
NewConta cria uma nova conta apontando para o endereço de memória
para a struct Conta e retorna o endereço de memória e alterações feitas serão refletidas
**
*/
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// Ao passar o endereço de memória da struct Conta, é possível alterar o valor do saldo
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
