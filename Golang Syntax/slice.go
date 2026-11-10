package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:6]
	fmt.Println(s)
	fmt.Println(primes)
}