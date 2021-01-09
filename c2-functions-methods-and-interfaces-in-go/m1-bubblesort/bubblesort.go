/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
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
	MAX_LEN := 10

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a sequence of up to 10 integers separated by a space: ")
	if scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if len(s) > MAX_LEN {
			s = s[:MAX_LEN]
		}

		x := make([]int, 0, MAX_LEN)
		for _, ss := range s {
			xx, _ := strconv.Atoi(ss)
			x = append(x, xx)
		}

		BubbleSort(&x)

		fmt.Println(x)
	}

}

func BubbleSort(x *[]int) {
	n := len(*x)

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if (*x)[i] > (*x)[i+1] {
				Swap(x, i)
				swapped = true
			}
		}
	}

}

func Swap(x *[]int, i int) {
	xi := (*x)[i]
	(*x)[i] = (*x)[i+1]
	(*x)[i+1] = xi
}
