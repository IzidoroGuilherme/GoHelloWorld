package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3

const tempoDelay = 5

func main() {
	exibeIntroducao()

	// For do Go, quando não tem nenhum parêmtro, roda em looping até um break
	for {

		exibeMenu()

		comando := leComando()

		// Operador switch: O Go não necessita do colocar o break no seu switch
		switch comando {
		case 1:
			iniciarMonitoramento()
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando")

	// Slice criado com os sites que serão monitorados
	sites := []string{"https://random-status-code.herokuapp.com/", "https://alura.com.br/", "https://my.4bee.com.br/"}

	// Criando o for, para ele testar mais de uma vez
	for i := 0; i < monitoramentos; i++ {
		// for melhorado para percorrer um slice, ele pega a item a item
		for i, site := range sites {
			fmt.Println("Testando o site ", i, ": ", site)
			testaSite(site)
		}
		time.Sleep(tempoDelay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	// No Go algumas funcções pode retornar mais de um valor como a http.get, que devolve resposta e erro
	resp, _ := http.Get(site)

	// Usando as respostas dos objetos
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "esta com problemas. Status Code: ", resp.StatusCode)
	}
}
