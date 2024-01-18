package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadPPM reads a PPM image from a file and returns a struct that represents the image.
func ReadPPM(filename string) (*PPM, error) {

	ppm := PPM{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var width, height int
	var max uint8
	var magicnumber string
	j := 0
	dataCreated := false

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
		CurrentLine := strings.Split(scanner.Text(), " ")

		if height != 0 && width != 0 && dataCreated == false {
			ppm.data = make([][]Pixel, height)
			for k := range ppm.data {
				ppm.data[k] = make([]Pixel, width)
			}
			dataCreated = true
		}
		compt := 0
		for i := 0; i < width; i++ {
			pixel := Pixel{}
			nb, err := strconv.Atoi(CurrentLine[compt])
			if err != nil {
				fmt.Println(err)
			}
			pixel.R = uint8(nb)
			compt++

			nb, err = strconv.Atoi(CurrentLine[compt])
			if err != nil {
				fmt.Println(err)
			}
			pixel.G = uint8(nb)
			compt++

			nb, err = strconv.Atoi(CurrentLine[compt])
			if err != nil {
				fmt.Println(err)
			}
			pixel.B = uint8(nb)
			compt++

			ppm.data[j][i] = pixel
		}
		j++
	}

	return &PPM{ppm.data, width, height, magicnumber, max}, nil
}

func (ppm *PPM) Size() (int, int) {
	return ppm.width, ppm.height
}

func (ppm *PPM) At(x, y int) Pixel {
	return ppm.data[y][x]
}

func (ppm *PPM) Set(x, y int, R, G, B uint8) {
	ppm.data[y][x].R = R
	ppm.data[y][x].G = G
	ppm.data[y][x].B = B
}

func (ppm *PPM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintf(writer, "%s\n%d %d\n%d\n", ppm.magicNumber, ppm.width, ppm.height, ppm.max)

	// Write pixel data
	for _, row := range ppm.data {
		for _, value := range row {
			fmt.Fprintf(writer, "%d ", value.R)
			fmt.Fprintf(writer, "%d ", value.G)
			fmt.Fprintf(writer, "%d ", value.B)
		}
		fmt.Fprintln(writer)
	}

	return writer.Flush()
}

func (ppm *PPM) Invert() {
	for i := range ppm.data {
		for j := range ppm.data[i] {
			ppm.data[i][j].R = uint8(ppm.max) - ppm.data[i][j].R
			ppm.data[i][j].G = uint8(ppm.max) - ppm.data[i][j].G
			ppm.data[i][j].B = uint8(ppm.max) - ppm.data[i][j].B
		}
	}
}

func (ppm *PPM) Flip() {
	for i := range ppm.data {
		for j, k := 0, ppm.width-1; j < k; j, k = j+1, k-1 {
			ppm.data[i][j], ppm.data[i][k] = ppm.data[i][k], ppm.data[i][j]
		}
	}
}

func (ppm *PPM) Flop() {
	for i, j := 0, ppm.height-1; i < j; i, j = i+1, j-1 {
		ppm.data[i], ppm.data[j] = ppm.data[j], ppm.data[i]
	}
}

func (ppm *PPM) SetMagicNumber(magicNumber string) {
	ppm.magicNumber = magicNumber
}

func (ppm *PPM) SetMaxValue(maxValue uint8) {

	for i := 0; i < ppm.height; i++ {
		for j := 0; j < ppm.width; j++ {
			ppm.data[i][j].R = uint8(float64(ppm.data[i][j].R) * float64(maxValue) / float64(ppm.max))
			ppm.data[i][j].G = uint8(float64(ppm.data[i][j].G) * float64(maxValue) / float64(ppm.max))
			ppm.data[i][j].B = uint8(float64(ppm.data[i][j].B) * float64(maxValue) / float64(ppm.max))
		}
	}
	ppm.max = maxValue
}

func (ppm *PPM) Rotate90CW() {
	rotatedData := make([][]Pixel, ppm.width)
	for i := range rotatedData {
		rotatedData[i] = make([]Pixel, ppm.height)
	}

	for i := range ppm.data {
		for j := range ppm.data[i] {
			rotatedData[j][ppm.height-i-1] = ppm.data[i][j]
		}
	}

	ppm.data = rotatedData
	ppm.width, ppm.height = ppm.height, ppm.width
}

func (pp *PPM) PrintData() {
	fmt.Printf("Magic Number: %s\n", pp.magicNumber)
	fmt.Printf("Width: %d\n", pp.width)
	fmt.Printf("Height: %d\n", pp.height)
	fmt.Printf("max : %d\n", pp.max)
	fmt.Println("Data:", pp.data)
	for _, row := range pp.data {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}
