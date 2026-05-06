package main

import "fmt"

func main() {
	var numeros [10]int
	pares := make([]int, 0, 10)
	impares := make([]int, 0, 10)
	somaPares := 0

	// Lê os 10 números e separa pares e ímpares.
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
		if numeros[i]%2 == 0 {
			pares = append(pares, numeros[i])
			somaPares += numeros[i]
		} else {
			impares = append(impares, numeros[i])
		}
	}

	fmt.Println("Números pares digitados:")
	for _, numero := range pares {
		fmt.Print(numero, " ")
	}
	fmt.Println()

	fmt.Println("Soma dos números pares:", somaPares)

	fmt.Println("Números ímpares digitados:")
	for _, numero := range impares {
		fmt.Print(numero, " ")
	}
	fmt.Println()

	fmt.Println("Quantidade de números ímpares:", len(impares))
}
