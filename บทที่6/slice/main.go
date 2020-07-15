package main

import (
	"fmt"
)

func main() {
	var x []float64
	//var declaration

	var z = []int{}
	//var declaration
	//"=" require expression "{....}"

	var e = make([]float64, 5)
	//var declaration
	//make
	//

	y := make([]float64, 5)
	//short declaration
	//make
	//

	arr := [5]float64{
		1,
		2,
		3,
		4,
		5,
	}

	arr2 := arr[0:5]
	//สร้างsliceจากarray

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	//append จะรวมsliceเข้าด้วยกัน

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	//copy จะคัดลอกจากslice1 ไปยัง slice2 ตามขนาดความจุ

	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(e)
	fmt.Println(y)
	fmt.Println(arr2)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println(slice1, slice2)

	//ความแตกต่างระหว่างสไล๢์กับอาร์เรย์คือไม่มีการระบุความยาวในวงเล็บและในกรณีนี้x มีความยาวเป็น0 !
	//สามารถสร้าง sliceจาก array ได้
	//สร้างแบบshort declarationได้
	//ถ้าใช้ make สร้าง slice ต้องกำหนดค่าในargumentด้วย
	//ถ้าใช้ := ต้องกำหนดค่าด้วย (ประกาศตัวแปรอย่างเดียวไม่ได้)

}
