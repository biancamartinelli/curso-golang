package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()	

	erro := checkmail.ValidateFormat("buki.k3@gamil.com")
	fmt.Println(erro)
}
