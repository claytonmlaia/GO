package main

import "fmt"

func main() {
	var variavel string = "valor" // variavel tipo string explicita
	variavel2 := "valor2"         // variavel tipo string implicita

	fmt.Println(variavel)
	fmt.Println(variavel2)

	var ( // declarar mais de uma variavel do mesmo tipo - outra forma de declarar variavel
		variavel3 = "valor3"
		variavel4 = "valor4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "valor5", "valor6" // declarar variavel e atribuir valor a mesma ao mesmo tempo
	fmt.Println(variavel5, variavel6)

	const constante1 = "constante1" // declarar constante
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // inverte  valor das variaveis - toogle

	fmt.Println(variavel5, variavel6)

}
