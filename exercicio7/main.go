package main

import "fmt"

func main() {
	var numeros [100]int

	// Armazena os 100 primeiros números ímpares.
	for i := 0; i < len(numeros); i++ {
		numeros[i] = 2*i + 1
	}

	// Imprime o vetor completo.
	for _, numero := range numeros {
		fmt.Print(numero, " ")
	}
	fmt.Println()
}
