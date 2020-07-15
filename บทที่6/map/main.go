package main

import (
	"fmt"
)

func main() {
	var x map[string]int
	fmt.Println(x)
	//

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])
	//y[1] ไม่ได้หมายถึงสมาชิกตัวที่สองเหมือนarray
	//mapไม่มีการเรียงลำดีบสมาชิก
	//indexในmapไม่มีการเรียงลำดับ
	//indexในmapเป็นstring,intอื่นๆได้
	//

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	//

	/*
		elements := map[string]string{
			"H":  "Hydrogen",
			"He": "Helium",
			"Li": "Lithium",
			"Be": "Beryllium",
			"B":  "Boron",
			"C":  "Carbon",
			"N":  "Nitrogen",
			"O":  "Oxygen",
			"F":  "Fluorine",
			"Ne": "Neon",
		}//short
	*/

	elements3 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	fmt.Println(elements3)

	/*
		๤ังเกตว่าเราเปลี่ยนจากmap[string]stringเป็นmap[string]map[string]stringตอนนี้เรามีแมปของสตริงไปยังแมปของสตริงไปยังสติงแมปที่อยู่ด้าน
		นอกถูกใ๡้เป็นตารางค้นหาโดยใ๡้๤ัญลักษณ์ของธาตุขณะที่แมปที่อยู่ด้านในถูกใ๡้เพื่อเก็บข้อมูลทั่วไปของธาตุนั้น
	*/

}
