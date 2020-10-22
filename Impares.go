package main

import (
	"fmt"
)

func generateOdd() func() uint {
	i := uint(1)

	return func() uint {
		var odd = i
		i += 2
		return odd
	}
}

func main3() {
	var nextOdd = generateOdd()

	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
