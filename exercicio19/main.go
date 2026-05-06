package main

import "fmt"

func main() {
	var numeros [10]int
	var divisores [5]int

	// Lê os dois vetores pedidos no enunciado.
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
	}
	for i := 0; i < len(divisores); i++ {
		fmt.Scan(&divisores[i])
	}

	// Para cada número do primeiro vetor, mostra os divisores encontrados no segundo.
	for i, numero := range numeros {
		fmt.Printf("Número %d:\n", numero)
		encontrou := false
		for j, divisor := range divisores {
			if divisor != 0 && numero%divisor == 0 {
				fmt.Printf("Divisível por %d na posição %d\n", divisor, j)
				encontrou = true
			}
		}
		if !encontrou {
			fmt.Println("Nenhum divisor encontrado no segundo vetor.")
		}
		if i < len(numeros)-1 {
			fmt.Println()
		}
	}
}
