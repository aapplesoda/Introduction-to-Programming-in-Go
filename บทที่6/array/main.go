package main

import "fmt"

//Arrays, Slices และMap

func main() {

	z := [5]float64{
		98,
		93,
		77,
		82,
		83,
		//99,
	}

	y := [5]float64{98, 93, 77, 82, 83}

	var q = [5]int{1, 2, 3, 4}
	fmt.Println(q)

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	fmt.Println(len(y))
	fmt.Println(len(z))

	/*
		var total float64 = 0
		for i, value := range x {
			total += value
		}
		fmt.Println(total / float64(len(x)))
	*/

	/*
		var total float64 = 0
		for _, value := range x {
			total += value
		}
		fmt.Println(total / float64(len(x)))
	*/

	// "_" ถูกใช้เป็น๡ื่อตัวแปรเพื่อบอกคอมไพเลอร์ว่าเราไม่ต้องการใช้ตัวแปรนั้น

	//-------------------------------------------------------------

}
