package main

import "fmt"

func main() {
	nome := "Lucas"
	versao := 1.1
	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O valor da variável comando é:", comando)
}
