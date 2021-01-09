/*

 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var displaceFn func(float64) float64

	fmt.Println("Enter the acceleration, initial velocity, and initial displacement as space-separated numbers:")
	if scanner.Scan() {
		s := scanner.Text()
		sarr := strings.Split(s, " ")

		a, _ := strconv.ParseFloat(sarr[0], 64)
		v0, _ := strconv.ParseFloat(sarr[1], 64)
		s0, _ := strconv.ParseFloat(sarr[2], 64)

		displaceFn = GenDisplaceFn(a, v0, s0)
	}

	fmt.Println("Enter the time to calculate the displacement:")
	if scanner.Scan() {
		s := scanner.Text()
		t, _ := strconv.ParseFloat(s, 64)
		fmt.Println("Displacement: ", displaceFn(t), " m")
	}

}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}
