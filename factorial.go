package main

import "fmt"

func factorial(i uint) uint {
	if i == 0 {
		return 1
	}
	var rez, m uint = 1, 0
	for m = 2; m <= i; m++ {
		rez *= m
	}
	return rez
}
func main() {
	fmt.Println(factorial(6))
}
