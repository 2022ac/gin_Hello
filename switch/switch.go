package main

import "fmt"

func main() {
	type Gender int8
	const (
		Male   Gender = 1
		Female Gender = 2
	)
	var gender Gender = 3
	switch gender {
	case Female:
		fmt.Println("female")
	case Male:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
}
