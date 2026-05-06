package main

import "fmt"

func main() {
	var numeros [10]int
	menor := 0
	posicaoMenor := 0

	// Lê o vetor e guarda o menor valor com sua posição.
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
		if i == 0 || numeros[i] < menor {
			menor = numeros[i]
			posicaoMenor = i
		}
	}

	fmt.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", menor, posicaoMenor+1)
}
