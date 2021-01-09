package main

import "fmt"

var x = 0

func g1() {
	x += 1
	x += 2
	x += 3
}

func g2() {
	x -= 2
	x -= 4
	x -= 6
}

func main() {

	for i := 0; i < 1000; i++ {
		x = 0
		go g1()
		go g2()
		fmt.Println("Trial ", i, ": x = ", x)
	}

	fmt.Println("A race condition is an undesirable situation that occurs when a device or system attempts to perform two or more operations at the same time, but because of the nature of the device or system, the operations must be done in the proper sequence to be done correctly. Race condition can occur when two subroutines are modifying or accessing a variable concurrently without a deterministic order of execution.")

}
