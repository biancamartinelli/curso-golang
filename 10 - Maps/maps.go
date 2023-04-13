package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Bianca",
		"sobrenome": "M",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Carina",
			"ultimo":   "Almeida",
		},
	}

	fmt.Println(usuario2)
}
