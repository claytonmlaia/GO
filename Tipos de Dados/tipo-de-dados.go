package main

import (
	"errors"
	"fmt"
)

func main() {
	var x int
	// variaveis inte possuiem variações, como int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, uintptr
	// uint é para inteiros positivos apenas
	var y float64
	// existem float32, float64
	var z string
	// não possui tipo CHAR
	// se colocar algo entre chaves simples, '', ele mostra o valor da variavel na tabela ASCII
	var w bool
	// bool é para verdadeiro ou falso

	var erros error

	x = 1000000000000000000
	y = 10.5
	z = "Hello World"
	w = true
	erros = nil
	// nil é para nulo, e tambem pode conter valor (0x0,0x0)
	// para atribuir um valor de erro, utiliza-se uma biblioteca do go chamada errors como no exemplo abaixo
	// errors.New("Mensagem de erro que desejar")

	var erroos2 error = errors.New("Mensagem de erro que desejar")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
	fmt.Println(erros)
	fmt.Println(erroos2)
}
