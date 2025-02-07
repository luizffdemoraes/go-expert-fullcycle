package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	/***
	Passando o endereço da variável e alterando o valor dela
	Quando usar ponteiro? Quando quiser alterar o valor da variável para trabalhar com valores mutaveis
	Quando não usar ponteiro? Quando não quiser alterar o valor da variável como uma soma
	***/
	soma(&minhaVar1, &minhaVar2)
	println(minhaVar1)
}
