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
	var width, height, max int
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
