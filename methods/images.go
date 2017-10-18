package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())        //边界点
	fmt.Println(m.At(0, 0).RGBA()) //某点RGBA信息
}
