package main

import "fmt"

func main() {
	var a = [][]int{
		{1},
		{1, 2},
		{4, 8, 3},
		{5, 4, 3, 4},
		{7, 8, 7, 8, 7},
		{7, 8, 9, 9, 7, 8},
	}

	printMatrix(a)

	fmt.Println("Min way  =", minWay(a))

	// fmt.Println("Min way2 =", minWay2(a))
}

func minWay(m [][]int) int {
	s := make([]int, len(m))
	r := make([]int, len(m))

	for i := range m {

		for j := range m[i] {
			switch j {
			case 0:
				r[j] = s[j] + m[i][j]
			case len(m[i]) - 1:
				r[j] = s[j-1] + m[i][j]
			default:
				r[j] = minZ((s[j-1] + m[i][j]), (s[j] + m[i][j]))
			}
		}
		for j := range r {
			s[j] = r[j]
		}

	}
	// printSlice(s)
	return pMin(s)
}

// func minWay2(m [][]int) int {
// 	var r []int

// 	r = append(r, m[0][0]+m[1][0]+m[2][0]+m[3][0])
// 	r = append(r, m[0][0]+m[1][0]+m[2][0]+m[3][1])
// 	r = append(r, m[0][0]+m[1][0]+m[2][1]+m[3][1])
// 	r = append(r, m[0][0]+m[1][0]+m[2][1]+m[3][2])

// 	r = append(r, m[0][0]+m[1][1]+m[2][1]+m[3][1])
// 	r = append(r, m[0][0]+m[1][1]+m[2][1]+m[3][2])
// 	r = append(r, m[0][0]+m[1][1]+m[2][2]+m[3][2])
// 	r = append(r, m[0][0]+m[1][1]+m[2][2]+m[3][3])

// 	// printSlice(r)
// 	return pMin(r)
// }

func printMas(m [][]int) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("a[%d][%d] = %d\n", i, j, m[i][j])
		}
	}
}

func minZ(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func printMatrix(m [][]int) {
	for i := range m {
		fmt.Println(m[i])
	}
}

func pMin(s []int) int {
	var min int
	if len(s) > 0 {
		min = s[0]
	} else {
		return -1
	}
	for i := range s {
		if min > s[i] {
			min = s[i]
		}
	}
	return min
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
