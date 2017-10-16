package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	prev, cur := 0, 1
	return func() int {
		cur_old := cur
		cur += prev
		prev = cur_old
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
