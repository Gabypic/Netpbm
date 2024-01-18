package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadPGM(filename string) (*PGM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]uint8
	var width, height int
	var max uint8
	var magicnumber string
	data = make([][]uint8, height)

	for scanner.Scan() {
		line := scanner.Text()

		//check comments
		if strings.HasPrefix(line, "#") {
			continue
		}

		//takeMagic
		if magicnumber == "" {
			magicnumber = line
			continue
		}

		if width == 0 && height == 0 {
			_, err := fmt.Sscanf(line, "%d %d", &width, &height)
			if err != nil {
				return nil, err
			}
			continue
		}

		if max == 0 {
			_, err := fmt.Sscanf(line, "%d", &max)
			if err != nil {
				return nil, err
			}
			continue
		}

		// Take data
		lineData := make([]uint8, width)

		byteCase := strings.Fields(line)

		if len(byteCase) < width {
			break
		}

		for j := 0; j < width; j++ {
			intData, _ := strconv.Atoi(byteCase[j])
			lineData[j] = uint8(intData)
		}
		data = append(data, lineData)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error")
		return nil, err
	}
	return &PGM{data, width, height, magicnumber, max}, nil
}

func (pgm *PGM) Size() (int, int) {
	return pgm.width, pgm.height
}

func (pgm *PGM) At(x, y int) uint8 {
	return pgm.data[y][x]
}

func (pgm *PGM) Set(x, y int, value uint8) {
	pgm.data[y][x] = value
}

func (pgm *PGM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintf(writer, "%s\n%d %d\n%d\n", pgm.magicNumber, pgm.width, pgm.height, pgm.max)

	// Write pixel data
	for _, row := range pgm.data {
		for _, value := range row {
			fmt.Fprintf(writer, "%d ", value)
		}
		fmt.Fprintln(writer)
	}

	return writer.Flush()
}

func (pgm *PGM) Invert() {
	for i := range pgm.data {
		for j := range pgm.data[i] {
			pgm.data[i][j] = uint8(pgm.max) - pgm.data[i][j]
		}
	}
}

func (pgm *PGM) Flip() {
	for i := range pgm.data {
		for j, k := 0, pgm.width-1; j < k; j, k = j+1, k-1 {
			pgm.data[i][j], pgm.data[i][k] = pgm.data[i][k], pgm.data[i][j]
		}
	}
}

func (pgm *PGM) Flop() {
	for i, j := 0, pgm.height-1; i < j; i, j = i+1, j-1 {
		pgm.data[i], pgm.data[j] = pgm.data[j], pgm.data[i]
	}
}

func (pgm *PGM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
}

func (pgm *PGM) SetMaxValue(maxValue uint8) {

	for i := 0; i < pgm.height; i++ {
		for j := 0; j < pgm.width; j++ {
			pgm.data[i][j] = uint8(float64(pgm.data[i][j]) * float64(maxValue) / float64(pgm.max))
		}
	}
	pgm.max = maxValue
}

func (pgm *PGM) Rotate90CW() {
	rotatedData := make([][]uint8, pgm.width)
	for i := range rotatedData {
		rotatedData[i] = make([]uint8, pgm.height)
	}

	for i := range pgm.data {
		for j := range pgm.data[i] {
			rotatedData[j][pgm.height-i-1] = pgm.data[i][j]
		}
	}

	pgm.data = rotatedData
	pgm.width, pgm.height = pgm.height, pgm.width
}

func (pgm *PGM) ToPBM() *PBM {
	pbmData := make([][]bool, pgm.height)
	for i := range pbmData {
		pbmData[i] = make([]bool, pgm.width)
	}

	threshold := uint8(pgm.max / 2)

	for i := range pgm.data {
		for j := range pgm.data[i] {
			pbmData[i][j] = pgm.data[i][j] > threshold
		}
	}

	return &PBM{
		data:        pbmData,
		width:       pgm.width,
		height:      pgm.height,
		magicNumber: "P1",
	}
}

func (g *PGM) PrintData() {
	fmt.Printf("Magic Number: %s\n", g.magicNumber)
	fmt.Printf("Width: %d\n", g.width)
	fmt.Printf("Height: %d\n", g.height)
	fmt.Printf("max : %d\n", g.max)
	fmt.Println("Data:", g.data)
	for _, row := range g.data {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}
