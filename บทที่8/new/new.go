package main

import (
	"fmt"
)

func one(xPtr *int) {
	*xPtr = 1
}

func zero(x *int) {
	*x = 0
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1!

	x := 5
	zero(&x)
	fmt.Println(x)
}

//new จะคล้ายๆหรือเหมือน& stackoverflowบอกมา
