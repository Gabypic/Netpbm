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
	pixel := Pixel{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//var data []Pixel
	var width, height, max int
	var magicnumber string
	//data = make([]Pixel, height)

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

		for i := 0; i < height; i++ {
			for j := 0; j < width; j += 3 {
				// Take data
				//lineData := make([]Pixel, width)

				var byteCase []int
				strCase := strings.Fields(line)

				for k := 0; k < width; k++ {
					temp, _ := strconv.Atoi(strCase[k])
					byteCase = append(byteCase, temp) //recup en 3/3
				}

				for l := 0; l < 3; l++ {
				 	ppm.data[i].R =
				}

				if len(strCase) < width {
					break
				}
			}
		}
	}

	return &PPM{}, nil
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
