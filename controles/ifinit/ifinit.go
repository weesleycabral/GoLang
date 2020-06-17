package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // Pegar os nanosegundo da hora atual
	r := rand.New(s)                           // Gerar um numero aleatorio entre os nanosegundos
	return r.Intn(10)                          // Limitar o numero ate 10
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tbm suportado no switch (se numero = numeroaleatorio e numero maior que 5)
		fmt.Println("Ganhou!! Numero: ", i)
	} else {
		fmt.Println("Perdeu !! Numero: ", i)
	}
}
