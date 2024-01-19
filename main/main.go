package main

import (
	"Netpbm"
	"fmt"
)

func main() {
	pbm, err := Netpbm.ReadPBM("testImages/pbm/testP4.pbm")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//pgm, err := Netpbm.ReadPGM("PGM.txt")
	//if err != nil {
	//	fmt.Println("Error : ", err)
	//	return
	//}
	//ppm, err := Netpbm.ReadPPM("PPM.txt")
	//if err != nil {
	//	fmt.Println("Error : ", err)
	//	return
	//}

	pbm.PrintData()
	//pbm.Save("save.PBM")
	//pgm.PrintData()
	//pgm.Save("save.PGM")
	//ppm.PrintData()
	//ppm.Save("save.PPM")
	//fmt.Println('ง')
}
