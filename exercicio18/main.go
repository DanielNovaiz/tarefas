package main

import "fmt"

func main() {
	var numeros [10]int

	// O enunciado já garante que os valores serão lidos em ordem crescente.
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
	}

	for _, numero := range numeros {
		fmt.Print(numero, " ")
	}
	fmt.Println()
}
