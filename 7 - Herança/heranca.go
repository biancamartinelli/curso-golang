package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 170}
	fmt.Println(p1)

	e1 := estudante {p1, "ADM", "USP"}
	fmt.Println(e1)

}
