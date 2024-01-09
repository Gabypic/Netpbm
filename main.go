package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

func main() {
	pbm, err := ReadPBM("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pbm.PrintData()
	fmt.Println(pbm.Size())
	fmt.Println(pbm.At(1, 2))
	pbm.Set(1, 2, false)
	pbm.PrintData()
	pbm.Save("save.PBM")
}

func ReadPBM(filename string) (*PBM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]bool
	var width, height int
	var magicnumber string
	var row []bool
	i := 0

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
			data = make([][]bool, height)
			continue
		}

		// Take binary data
		var rowData []bool
		for _, char := range line {
			if char == '1' {
				rowData = append(rowData, true)
			} else if char == '0' {
				rowData = append(rowData, false)
			}
		}

		// Ajouter rowData à data
		data[i] = rowData
		i++

		// compile
		if len(row) == width*8 {
			fmt.Println("compile")
			data = append(data, row)
			row = []bool{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error")
		return nil, err
	}
	return &PBM{data, width, height, magicnumber}, nil
}

func (p *PBM) PrintData() {
	fmt.Printf("Magic Number: %s\n", p.magicNumber)
	fmt.Printf("Width: %d\n", p.width)
	fmt.Printf("Height: %d\n", p.height)
	fmt.Println("Data:")
	for _, row := range p.data {
		for _, val := range row {
			if val {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (pbm *PBM) Size() (int, int) {
	return pbm.width, pbm.height
}

// At returns the value of the pixel at (x, y).
func (pbm *PBM) At(x, y int) bool {
	return pbm.data[x][y]
}

// Set sets the value of the pixel at (x, y).
func (pbm *PBM) Set(x, y int, value bool) {
	pbm.data[x][y] = value
	fmt.Print("Nouvelle valeur du pixel ", x, y, " ")
	fmt.Println(pbm.data[x][y])
}

// Save saves the PBM image to a file and returns an error if there was a problem.
func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Write magic number, width, and height
	fmt.Fprintf(writer, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)

	// Write image data
	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Fprintf(writer, "1 ")
			} else {
				fmt.Fprintf(writer, "0 ")
			}
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
	return nil
}

// Invert inverts the colors of the PBM image.
func (pbm *PBM) Invert() {
	for y := 0; y < pbm.height; y++ {
		for x := 0; x < pbm.width; x++ {
			pbm.data[y][x] = !pbm.data[y][x]
		}
	}
}
