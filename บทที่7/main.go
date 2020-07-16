package main

import (
	"fmt"
)

//บทที่8 function

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
		fmt.Println(total)
	}
	return total / float64(len(xs))
}

//cap,len,range
//cap ความจุ
//len จำนวนที่มีอยู่

//stack function
func f1() int {
	fmt.Println("f1 called")
	return f2()
}

//สามารถกําหนด๡ื่อให้ตัวแปรให้กับค่าที่ต้องการ๤่งกลับได้
//โดยเราสามารถกําหนดค่าให้ตัวแปรนี้แทนการ return ค่าตรงๆได้
func f2() (r int) {
	fmt.Println("f2 called")
	r = 1
	return
}

//การ๤่งค่ากลับหลายค่า
func f3() (int, int) {
	fmt.Println("f3 called")
	return 5, 6
}

//ฟังก์๡ันแบบรับพารามิเตอร์ได้โดยไม่จํากัดจํานวน (Variadic Functions)
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(len(xs))
	fmt.Println(cap(xs))

	fmt.Println(f1())
	fmt.Println(f3())

	x, y := f3()
	fmt.Println(x, y)

	fmt.Println(add(1, 2, 3))

}

//ดูความแตกต่างระหว่าง parameter กับ return typeให้ดี
