package main

import "fmt"

//struct

func main() {
	type Circle struct {
		x float64
		y float64
		r float64
	}
	fmt.Println(Circle.x)
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c)
	//คีย์เวิร์ดtype เป็นตัวที่ใ๡้บอกว่าเรามีไทป์ใหม่
	//จากนั้นเราจะใ๤่๡ื่อไทป์ลงไปโดยตัวอย่างนี้เราใ๡้Circle
	//คีย์เวิร์ดstructมีไว้เพื่อบอกว่าไทป์ใหม่นี้เป็นstruct
	//ภายในเครื่องหมายวงเล็บปีกกาจะเป็นบริเวณที่เรามีไว้เพื่อประกาศฟิลด์ที่อยู่ในสตรัคท์
	//ถ้าเรามีฟิลด์ที่เป็นประเภทเดียวกันหลายๆตัวเราสามารถรวมเป็นกลุ่มไว้เป็นบรรทัดเดียวกันได้
	//x, y, r float64
}

//ใช้รวมเก็บข้อมูลหลายๆประเภท
