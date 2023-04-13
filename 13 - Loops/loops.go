package main

import (
	"fmt"
	"time"
)

func main() {
	 fmt.Println("Loops")
	 i := 0

	 for i < 10 {
	 	i++
		fmt.Println("incrementando i")
	 	time.Sleep(time.Second)
	 }

	 fmt.Println(i)

	 for j := 0; j < 10; j += 2 {
	 fmt.Println("Incrementando j", j)
	 time.Sleep(time.Second)
	 }

	nomes := []string{"Bianca", "Carina", "Vitor"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Bianca",
		"sobrenome": "M",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando inifinitamente")
		time.Sleep(time.Second)
	}
}
