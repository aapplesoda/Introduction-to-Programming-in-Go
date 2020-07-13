package main

import (
	"fmt"
)

func main() {
	//ตั่งชื่อให้สื่อถึงตัวแปรมีความหมาย
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	fmt.Println("---------")
	//
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
	fmt.Println("---------")

	//
	// ใช้เครื่องหมาย “==” เพื่อเทียบว่าเท่ากันหรือไม่โดยจะให้ค่าผลลัพธ์ออกมาเป็นBoolean
	var xq string = "hello"
	var yq string = "world"
	fmt.Println(xq == yq)
	fmt.Println("---------")

	//
	//๢ึ่งเมื่อเราประกาศตัวแปรใดๆเป็นค่าคงที่แล้วตัวแปรตัวนั้นจะไม่สามารถถูกเปลี่ยนแปลงค่าได้
	const yu string = "Hello World"
	fmt.Println("---------")

	//the := operator is not pure declare, but more like declare and assign
	//ใ๤่เครื่องหมาย : เข้าไปข้างหน้า = โดยไม่ต้องประกาศชนิดของตัวแปรก็ได้เ๡่นกันเพราะCompiler ของ Go สามารถที่จะระบุชนิดของตัวแปรได้จากค่าที่รับเข้ามาเก็บไว้
	// "=" assign
	//ถ้าประกาศตัวแปรพร้อมกับassign ก็ไม่ต้องระบุประเภทตัวแปร เพราะgoจะรู้เอง
	//แต่ถ้าประกาศตัวแปรแต่ยังไม่assignค่าต้องระบุประเภทตัวแปร
	qq := "qq := \n"
	var xx = "var xx =\n"
	var yy string = "var yy string =\n"
	fmt.Println(qq, xx, yy)
	fmt.Println("---------\n")
	const pp string = "const pp string=\n"
	fmt.Println(pp)
	//multiple var declaration
	/*
		var (
			a = 5
			b = 10
			c = 15
		)
	*/
	//scopeขอบเขต ของ variable จะอยู่ที่เครื่องหมาย {}
	//x += “second”, x = x + ”second”
}
