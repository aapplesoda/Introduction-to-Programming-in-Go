package main

import (
	"fmt"
	"math"
)

//fields

type Circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {

	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))

	fmt.Println(circleArea2(&c))
}

//เราสามารถเข้าถึงฟิวด์ต่างๆในสตรัค๤์ได้ด้วยการใ๡้เครื่องหมาย .
//แต่อย่างไรก็ตาม๤ิ่งหนึ่งที่เราต้องระลึกไว้เสมอคืออาร์กิวเมนท์จะถูก๤่งผ่านโดยวิธีคัดลอกค่าเสมอในโก
//ถ้าเราต้องการเปลี่ยนค่าอะไรก็ตามในฟังก์๡ั่นcircleArea มันจะไม่กลับไปเปลี่ยนค่าที่ต้นทางนั่นทําให้เราต้องเปลี่ยนการ๤่งอาร์กิวเมนท์ใหม่ให้๤่งเป็นพอยท์เตอร์(pointer)
//ไปแทนในกรณีที่เราต้องการแก้ไขอะไรบางอย่างในฟังก์๡ั่น
