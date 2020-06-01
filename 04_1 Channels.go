package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 2)
	d := make(chan int)

	go pin(c)

	process(f, c, d, 5)

	go printer(d)

	var input string
	fmt.Scanln(&input)

}

func process(g func(int) int, in <-chan int, out chan<- int, n int) {

	go func() {
		for i := 0; i < n; i++ {
			sum1 := <-in
			sum2 := <-in
			sum := g(sum1) + g(sum2)
			out <- sum
		}
	}()
}

func f(n int) int {
	return n * 2
}

func pin(c chan int) {
	for i := 1; ; i++ {
		c <- i
	}
}

func printer(c chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
