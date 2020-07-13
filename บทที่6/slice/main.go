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

	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(e)
	fmt.Println(y)

	//ความแตกต่างระหว่างสไล๢์กับอาร์เรย์คือไม่มีการระบุความยาวในวงเล็บและในกรณีนี้x มีความยาวเป็น0 !
	//สามารถสร้าง sliceจาก array ได้
	//สร้างแบบshort declarationได้
	//ถ้าใช้ make สร้าง slice ต้องกำหนดค่าในargumentด้วย
	//ถ้าใช้ := ต้องกำหนดค่าด้วย (ประกาศตัวแปรอย่างเดียวไม่ได้)

}
