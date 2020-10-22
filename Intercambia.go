package main

import (
	"fmt"
)

func intercambia(i, j *int) {
	aux := *i
	*i = *j
	*j = aux
}

func main3() {
	var a, b int = 0, 0

	fmt.Print("Ingrese el primer valor:")
	fmt.Scanln(&a)
	fmt.Print("Ingrese el primer valor:")
	fmt.Scanln(&b)

	intercambia(&a, &b)

	fmt.Printf("Valores intercambiados: %d, %d\n", a, b)
}
