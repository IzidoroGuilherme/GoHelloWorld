package main

import "fmt"

import "os"

func testeIF() {

	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	// No Go não é necessário utilizar parentes no if
	/*
		if comando == 1 {
			fmt.Println("Monitorando")
		} else if comando == 2 {
			fmt.Println("Exibindo Logs ...")
		} else if comando == 0 {
			fmt.Println("Saindo do programa")
		} else {
			fmt.Println("Não conheço esse comando")
		}
	*/

	// Operador switch: O Go não necessita do colocar o break no seu switch
	switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Exibindo Logs ...")
	case 0:
		fmt.Println("Saindo do programa")
		// Método para sair do programa, e dizer que não ouve erro
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		// Método para sair do programa, e dizer que ouve algum erro na execução da aplicação
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	var versao float32 = 1.1
	fmt.Println("Bom dia srx: ", nome)
	fmt.Println("Eu sou a versão do seu sistema", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	// Neste caso estamos tipando a variavel pois queremos receber um int
	var comandoLido int

	// Função para leitura de input de teclado
	/* O &comando serve para dizer que iremos guardar o valor dentro da váriavel comando
	dentro do seu caminho na memoria */
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
