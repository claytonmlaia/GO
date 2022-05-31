package main

import "fmt"

func somar(a, b int) int {
	return a + b
}

// funções em go podem ter multiplos retornos
func multiplosRetornos(a int, b int) (int, int) {
	soma := a + b
	subtracao := a - b
	return soma, subtracao
}

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	resultadoSoma, resultadoSubtracao := multiplosRetornos(10, 20)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	// caso eu queira chamar apenas um retorno, basta usar o _ (underline), na posição do retorno que eu quero ignorar	(resultadoSoma, _ := multiplosRetornos(10, 20)) por exemplo

	// pode-se colocar uma função dentro de uma variável
	var f = func(a, b int) int {
		return a + b
	}

	fmt.Println(f(15, 25))
}