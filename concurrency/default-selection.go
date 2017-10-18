package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			//当其他条件分支都未准备好时，default分支会被执行
			//为了非阻塞的发送或接收，可使用default分支
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
