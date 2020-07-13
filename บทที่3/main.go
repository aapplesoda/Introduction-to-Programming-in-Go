package main

import (
	"fmt"
)

//บทที่3 types
//int float64 string boolean
//เมื่อเราย้ายตัวแปรออกมาไว้ข้างนอกfunction main แล้วนั้นfunction อื่นๆก็สามารถที่จะเรียกใ๡้งานตัวแปรนั้นได้เช่นกัน

func main() {

	fmt.Println("1+1=", 1+1)
	fmt.Println("32132 x 42452=", 32132*42452)
	fmt.Println("---------")
	//
	fmt.Println((true && false) || (false && true) || !(false && false))
	fmt.Println("---------")
	//

	//“Hello World” ช่องว่างถือเป็นอักษรด้วยจึงมี11ตัว
	//\n ขึ้นบรรทัดใหม่

}
