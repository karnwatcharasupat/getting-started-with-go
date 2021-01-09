package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func minisort(x []int, wg *sync.WaitGroup, i int) {
	sort.Ints(x)
	fmt.Println("Subarray ", i, ": ", x)
	wg.Done()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minIdx(x1, x2, x3, x4 int) (int, int) {
	if x1 <= x2 && x1 <= x3 && x1 <= x4 {
		return x1, 0
	} else if x2 <= x1 && x2 <= x3 && x2 <= x4 {
		return x2, 1
	} else if x3 <= x1 && x3 <= x2 && x3 <= x4 {
		return x3, 2
	} else {
		return x4, 3
	}
}

func merge(x1, x2, x3, x4 []int) []int {
	nx := [4]int{len(x1), len(x2), len(x3), len(x4)}
	n := len(x1) + len(x2) + len(x3) + len(x4)
	x := make([]int, 0, n)
	i := make([]int, 4)
	completed := make([]bool, 4)

	idx := 0

	for idx < n {

		xxs := [4]int{x1[i[0]], x2[i[1]], x3[i[2]], x4[i[3]]}

		for ix := 0; ix < 4; ix++ {
			if completed[ix] {
				xxs[ix] = math.MaxInt32
			}
		}

		xx, ii := minIdx(xxs[0], xxs[1], xxs[2], xxs[3])
		x = append(x, xx)

		if i[ii] < nx[ii]-1 {
			i[ii]++
		} else {
			completed[ii] = true
		}

		idx++
	}

	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup

	fmt.Println("Enter a space-separated list of integers (at least 4 numbers)")

	if scanner.Scan() {
		st := scanner.Text()
		s := strings.Fields(st)
		x := make([]int, 0, len(s))

		for _, ss := range s {
			xx, _ := strconv.Atoi(ss)
			x = append(x, xx)
		}

		n := int(math.Ceil(float64(len(x)) / 4.0))

		for i := 0; i < 4; i++ {
			wg.Add(1)
			go minisort(x[i*n:Min((i+1)*n, len(x))], &wg, i)
		}

		wg.Wait()

		xsort := merge(x[0:n], x[n:2*n], x[2*n:3*n], x[3*n:])

		fmt.Println("Main", xsort)
	}
}
