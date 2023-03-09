package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")

}

func funcao2() {
	fmt.Println("Executando a função 2")

}

func main() {
	defer funcao1()
	funcao2()

}