package main

import (
	"fmt"
	"math"
)

func main() {
	var numeros [15]int
	var resultados [15]float64

	// Para cada número, grava a raiz quadrada ou -1 se o valor for negativo.
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
		if numeros[i] < 0 {
			resultados[i] = -1
		} else {
			resultados[i] = math.Sqrt(float64(numeros[i]))
		}
	}

	for _, valor := range resultados {
		fmt.Printf("%.2f ", valor)
	}
	fmt.Println()
}
