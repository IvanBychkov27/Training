package main

import "fmt"

func main() {
	a := []int{0, 2, 4, 6, 8, 10, 12, 14}
	b := []int{1, 3, 5, 7, 9, 11, 13, 15}

	printSlice(sumSlice(a, b))
}

func sumSlice(x []int, y []int) []int {

	s := make([]int, len(x)+len(y))

	var kx, ky int

	if len(x) == 0 {
		s = s[:0]
		s = append(s, y[:]...)
		return s
	}

	if len(y) == 0 {
		s = s[:0]
		s = append(s, x[:]...)
		return s
	}

	for i := 0; i < len(s); i++ {
		if x[kx] <= y[ky] {
			s[i] = x[kx]
			if kx == len(x)-1 {
				s = s[:i+1]
				s = append(s, y[ky:]...)
				break
			}
			kx++
		} else {
			s[i] = y[ky]
			if ky == len(y)-1 {
				s = s[:i+1]
				s = append(s, x[kx:]...)
				break
			}
			ky++
		}
	}
	return s
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
