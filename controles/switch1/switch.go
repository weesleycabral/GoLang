package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch {
	case nota == 10 || nota == 9:
		return "Tirou A"
	case nota == 8 || nota == 7:
		return "Tirou B"
	case nota == 6 || nota == 5:
		return "Tirou C"
	case nota == 4 || nota == 3:
		return "Tirou D"
	case nota == 2 || nota == 1:
		return "Tirou E"
	case nota == 0:
		return "Tirou F"
	default:
		return "Nota invalida"
	}
	// switch nota {
	// case 10:
	// 	fallthrough
	// case 9:
	// 	return "A"
	// case 8, 7:
	// 	return "B"
	// case 6, 5:
	// 	return "C"
	// case 4, 3:
	// 	return "D"
	// case 2, 1:
	// 	return "E"
	// case 0:
	// 	return "F"
	// default:
	// 	return "Nota invalida"
	// }
}

func main() {
	fmt.Println(notaParaConceito(0))
}
