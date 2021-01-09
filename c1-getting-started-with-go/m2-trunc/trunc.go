package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter float: ")

	var f float64
	fmt.Scanf("%f", &f)
	x := int(f)

	fmt.Println(x)
}
