package main

import (
	"fmt"
	"time"
)

func main() {
	exibeNomes()
}

func exibeNomes() {
	// Esse é um slice, um array, do qual você não define um tamnho máximo
	nomes := []string{"Douglas", "Izidoro", "Teste"}

	// Com a propriedade Len, eu consigo saber o tamanho do slice, além disse tem o cap para saber a capacidade desse array
	fmt.Println("O meu slice tem um tamanho de ", len(nomes), " e capacidade ", cap(nomes))

	// Com a função append, conseguimos adicionar dados a um slice
	nomes = append(nomes, "Guilherme")

	// Quando adicionado novos itens ao slice, ele aumenta a capacidade além do necessário
	fmt.Println("O meu slice tem um tamanho de ", len(nomes), " e capacidade ", cap(nomes))

	// for melhorado para percorrer um slice, e pegar item a item
	for i, nome := range nomes {
		fmt.Println("Estou na posição: ", i, " e nela tenho o site: ", nome)
	}

	for i := 0; i < 3; i++ {
		// for melhorado para percorrer um slice, ele pega a item a item
		for i, nome := range nomes {
			fmt.Println("printando o nome ", i, ": ", nome)
		}
		// Usando a propriedade Time, para mandar ele rodar a cada 10 segundos
		time.Sleep(5 * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}
