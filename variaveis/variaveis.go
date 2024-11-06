package main

import (
	"fmt"
)

func main() {

	const constante string = "Constante"
	var nome string = "Lucas"
	idade := 21 //tipagem dinâmica (inferência de tipos)

	var (
		variavel1 string = "blabla"
		variavel2 string = "bleble"
	)

	fmt.Println(constante)
	fmt.Println(idade)
	fmt.Println(nome)
	fmt.Println(variavel1, variavel2)
}
