package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cria um arquivo chamado "arquivo.txt"
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	/***
	 Escreve no arquivo
	 tamanho, err := f.WriteString("Hello, World!")
	***/
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	if err != nil {
		panic(err)
	}

	/***
	 Exibe o tamanho do arquivo
	 cat arquivo.txt -> Hello, World!
	***/
	fmt.Printf("Arquivo criado com sucesso. Tamanho: %d bytes\n", tamanho)
	f.Close()

	// Leitura do arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Esse código abre o arquivo arquivo.txt para leitura.
	arquivoTwo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	/***
		Cria um leitor bufferizado para facilitar a leitura do arquivo.
		Isso permite ler o arquivo linha por linha ou em blocos, melhorando a eficiência.

		O que são Chunks?
		O termo chunk significa "pedaço" ou "bloco" de dados.
		Em programação e computação, chunks são usados para processar grandes quantidades de dados em partes menores, melhorando a eficiência e reduzindo o uso de memória.
	***/
	reader := bufio.NewReader(arquivoTwo)

	// Esse código cria um buffer de bytes com tamanho 10.
	buffer := make([]byte, 10)

	// Lê um arquivo em pedaços de 10 bytes e imprime até o final.
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
