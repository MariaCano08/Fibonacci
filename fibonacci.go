package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 110

	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}

}


func main4() {
	var i, f int
	fmt.Scan(&i)
	f = fibonacci(i)
	fmt.Println(f)
}
