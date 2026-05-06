package main

import "fmt"

func main() {
	var jogadas [20]int
	var frequencia [7]int

	// Lê os 20 resultados do dado e contabiliza as frequências.
	for i := 0; i < len(jogadas); i++ {
		fmt.Scan(&jogadas[i])
		if jogadas[i] >= 1 && jogadas[i] <= 6 {
			frequencia[jogadas[i]]++
		}
	}

	fmt.Println("Números sorteados:")
	for _, valor := range jogadas {
		fmt.Print(valor, " ")
	}
	fmt.Println()

	fmt.Println("Frequência de cada face:")
	for face := 1; face <= 6; face++ {
		fmt.Printf("%d: %d\n", face, frequencia[face])
	}
}
