package main

import "fmt"

func main() {
	var f bool

	s := "[(())[]()[()]]"
	// s := "([])"

	f = scoba(s)

	fmt.Println("String: ", s)
	fmt.Println(f)

}

func scoba(st string) bool {
	var fl bool
	if st == "" {
		fmt.Println("the line is empty")
		return fl
	}

	var stack []byte

	for i := range st {

		if len(stack) == 0 {
			stack = append(stack, st[i])
		} else {
			switch st[i] {
			case 41:
				if stack[len(stack)-1] == 40 {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, st[i])
				}
			case 93:
				if stack[len(stack)-1] == 91 {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, st[i])
				}
			default:
				stack = append(stack, st[i])
			}

		}
	}

	if len(stack) == 0 {
		fl = true
	}
	return fl
}

// func printSlice(s []byte) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }
