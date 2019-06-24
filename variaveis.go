package main

import "fmt"

import "reflect"

// Aula 02 (Declaração de variavel)
func variaveis() {
	/* Primeiro declarar como váriavel, depois dar um nome para mesma e por último mostrar o tipo (Em caso de String, int e boolean não é necessário declarar o tipo,
	pois o Go consegue saber qual o tipo de váriavel pela sua sintaxe, com o float isso também é possivel mas ele sempre vai usar o float64) */

	var nome = "Douglas"

	// Caso você não declare um valor para váriavel ela vai ser inicializada automaticamente como vazio (Caso int = 0, caso String = "", Caso Float = 0.0)
	//var idadeVazio int

	// o Go também consegue reconhecer quando uma váriavel é int, sem que seja necessário declarar o tipo.
	var idade = 24

	// No go temos as versão float 32bits e 64 bits, elas varias pela quantidade de dados que sua variavel vai receber
	var versao float32 = 1.1
	fmt.Println("Olá, srx", nome, "Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	// esse método reflect.TypeOf retorna o tipo de dado da variavel passada como parametrô. Para utiliza-ló
	fmt.Println("o Tipo da variavel 'nome' é", reflect.TypeOf(nome))

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
