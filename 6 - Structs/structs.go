package main

import "fmt"

type usuario struct {
	nome string
	idade uint8


}

func main() {
	fmt.Println("Arquivo structs")	

	var u usuario
	u.nome = "Bianca"
	u.idade = 21
	fmt.Println(u)

	usuario2 := usuario{"Bianca", 21}
	fmt.Println(usuario2)

}


