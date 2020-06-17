package main

import "fmt"

func main() {
	// array eh so do mesmo tipo e eh fixo

	var notas [3]float64 // exemplo de criação de array
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1 // adicionando valores em um array
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
