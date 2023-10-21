package main

import "fmt"

func main() {
	str := "Goland"
	var p *string = &str
	*p = "Hello"
	fmt.Println(str)
}
