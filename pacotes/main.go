package main

import (
    "fmt"
    "modulo/auxiliar"
    "github.com/badoux/checkmail"
)

func main() {
    fmt.Println("Ola teste")
    auxiliar.Escrever()

    erro := checkmail.ValidateFormat("devbook@gmail.com")
    vartst := checkmail.ValidateFormat("123")

    fmt.Println(vartst)
    fmt.Println(erro)
    
}
