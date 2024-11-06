package main

import (
	"errors"
	"fmt"
)

func main() {

	//INT
	var numero int = 100
	idade := 100

	//FLOAT
	var num float64 = 100.100
	real := 100.100

	//STRING
	var nome string = "Lucas"
	user := "Usuario"

	//BOOL
	var boleano bool = true
	boole := false

	//error
	var erro error = errors.New("Erro interno")

	fmt.Println(numero, idade)
	fmt.Println(num, real)
	fmt.Println(nome, user)
	fmt.Println(boleano, boole)
	fmt.Println(erro)
}
