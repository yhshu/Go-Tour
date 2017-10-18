package main

import "fmt"

func sum(a []int, c chan int) { // chan 即 channel 管道
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将sum送入c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 中获取
	// x=-5 y=17

	fmt.Println(x, y, x+y)
}
