package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64

	// forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da circunferencia eh", area)

	// outras formas de declarar const/var

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// varias variaveis na mesma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	// modo reduzido
	g, h, i := 1, "oi", false
	fmt.Println(g, h, i)

	nome, idade := "Wesley", 19
	fmt.Println(nome, idade)

}
