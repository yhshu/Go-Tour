package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int { //该函数中引用了sum，而sum并不是在该函数中定义的
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
