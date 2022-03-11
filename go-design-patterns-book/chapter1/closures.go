package main

import "fmt"

func addN(m int) func(int) int {
	return func(n int) int {
		return m + n
	}
}
func main() {
	result := addN(6)(5)

	fmt.Println("hello ", result)
}
