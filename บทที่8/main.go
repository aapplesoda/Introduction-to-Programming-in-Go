package main

import "fmt"

//เรื่องpointer

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	//ตำแหน่งaddress

	zero(&x)
	fmt.Println(x) // x is 0
}

//& เก็บตำแหน่งaddress
//* เก็บค่าในaddress
//new
//passby value
//passby pointer
