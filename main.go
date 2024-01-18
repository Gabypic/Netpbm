package main

import (
	"fmt"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}
type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           uint8
}

type PPM struct {
	data          [][]Pixel
	width, height int
	magicNumber   string
	max           uint8
}

type Pixel struct {
	R, G, B uint8
}

func main() {
	pbm, err := ReadPBM("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	pgm, err := ReadPGM("PGM.txt")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	ppm, err := ReadPPM("PPM.txt")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	pbm.PrintData()
	pbm.Save("save.PBM")
	pgm.PrintData()
	pgm.Save("save.PGM")
	ppm.PrintData()
	ppm.Save("save.PPM")
	//fmt.Println('ง')
}
