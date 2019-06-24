package main

import "fmt"

func main() {
	nome := "Douglas"
	var versao float32 = 1.1
	fmt.Println("Bom dia srx: ", nome)
	fmt.Println("Eu sou a versão do seu sistema", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	// Neste caso estamos tipando a variavel pois queremos receber um int
	var comando int

	// Função para leitura de input de teclado
	/* O &comando serve para dizer que iremos guardar o valor dentro da váriavel comando
	dentro do seu caminho na memoria */
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
}
