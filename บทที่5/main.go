package main

import (
	"fmt"
)

//บทที่5 Control Structures
//if, if else , if else if
//switch
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	i := 0
	if i%2 == 0 {
		// even
	} else {
		// odd
	}

	if i%2 == 0 {
		// divisible by 2
	} else if i%3 == 0 {
		// divisible by 3
	} else if i%4 == 0 {
		// divisible by 4
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
