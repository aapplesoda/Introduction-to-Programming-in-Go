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
//พารามิเตอร์แบบนี้ต้องเป็นลําดับสุดท้ายของฟังก์๡ันเท่านั้นเอาไปอยู่ก่อนหน้าพารามิเตอร์อื่นไม่ได้
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
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

	//Closure คือฟังก์๡ันที่สร้างขึ้นมาโดยไม่จําเป็นต้องมี๡ื่อและสามารถกําหนดให้กับตัวแปรได้หรือจะ๤่งเป็นข้อมูลเข้าออกจากฟังก์๡ันอื่นได้
	xq := 0
	increment := func() int {
		xq++
		return xq
	}
	fmt.Println(increment())
	fmt.Println(increment())

	//อีกวิธีหนึ่งในการใ๡้closure คือเราจะเขียนฟังก์๡ันที่๤่งค่าออกมาได้เป็นฟังก์๡ันอื่น๢ึ่งเมื่อเอามาเรียกใ๡้จะยังคงรักษาค่าของตัวแปรโลคอลของฟังก์๡ันที่สร้างมันมา
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

}

//ดูความแตกต่างระหว่าง parameter กับ return typeให้ดี
//๡ื่อของตัวแปรที่๤่งให้กับฟังก์๡ันไม่จําเป็นต้องเป็น๡ื่อเหมือนกับพารามิเตอร์ตอนที่เราสร้างฟังก์๡ัน
//โค้ดในตัวฟังก์๡ันเองไม่สามารถเข้าใ๡้งานตัวแปรที่ถูกสร้างในฟังก์๡ันต้นทางที่เรียกใ๡้ฟังก์๡ันได้ ต้องผ่านparameter
//สามารถกําหนด๡ื่อให้ตัวแปรให้กับค่าที่ต้องการ๤่งกลับได้เ๡่น ex. func f2() (r int) {
