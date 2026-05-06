package main

import "fmt"

func main() {
	var original [30]int
	var transformado [30]int

	// Lê o vetor original e gera o segundo vetor conforme a posição.
	for i := 0; i < len(original); i++ {
		fmt.Scan(&original[i])
		if i%2 == 0 {
			transformado[i] = original[i] * 2
		} else {
			transformado[i] = original[i] * 3
		}
	}

	for _, valor := range transformado {
		fmt.Print(valor, " ")
	}
	fmt.Println()
}
