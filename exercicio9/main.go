package main

import "fmt"

func main() {
	var alturas [10]float64
	soma := 0.0

	// Lê as alturas e calcula a média.
	for i := 0; i < len(alturas); i++ {
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}

	media := soma / float64(len(alturas))
	fmt.Printf("Média das alturas: %.2f\n", media)
	fmt.Println("Alturas acima da média:")

	// Mostra apenas as alturas maiores que a média.
	encontrou := false
	for _, altura := range alturas {
		if altura > media {
			fmt.Printf("%.2f ", altura)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nenhuma altura acima da média.")
	} else {
		fmt.Println()
	}
}
