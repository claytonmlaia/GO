package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo")
	auxiliar.Escreve()
	auxiliar.Escrever2()

	erro := checkmail.ValidateFormat("claytonhotmail.com")
	fmt.Println(erro)
}
