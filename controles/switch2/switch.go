package main

import "fmt"

// Switch para identificar tipos , ex: saber se eh string , float e etc ...

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case string:
		return "string"
	case float32, float64:
		return "valor real"
	case func():
		return "função"
	default:
		return "tipo desconhecido"

	}
}

func main() {
	fmt.Println(tipo("a"))
}
