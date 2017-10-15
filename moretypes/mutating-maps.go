package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42                       //插入元素
	fmt.Println("The value:", m["Answer"]) //获得元素

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") //删除元素
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] //若Answer在m中，ok为true，v为对应值；否则，ok为false，v为相应类型的零值
	fmt.Println("The value:", v, "Present?", ok)
}
