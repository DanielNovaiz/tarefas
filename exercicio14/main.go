package main

import "fmt"

func main() {
	var vetor1 [10]int
	var vetor2 [10]int
	var intercalado [20]int

	// Lê os dois vetores de 10 posições.
	for i := 0; i < len(vetor1); i++ {
		fmt.Scan(&vetor1[i])
	}
	for i := 0; i < len(vetor2); i++ {
		fmt.Scan(&vetor2[i])
	}

	// Intercala os elementos, um de cada vetor por vez.
	for i := 0; i < len(vetor1); i++ {
		intercalado[2*i] = vetor1[i]
		intercalado[2*i+1] = vetor2[i]
	}

	for _, valor := range intercalado {
		fmt.Print(valor, " ")
	}
	fmt.Println()
}
