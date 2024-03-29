package Netpbm

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
			fmt.Println(data)
			continue
		}

		if magicnumber == "P1" {
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
			fmt.Println("error", err)
			return nil, err
		}

		if magicnumber == "P4" {
			var char int32
			var rowdata []bool
			for j := 0; j < width*2; j++ {
				if len(rowdata) == width {
					data[i] = rowdata
					i++
					for l := 0; l < len(rowdata); l++ {
						rowdata = []bool{}
					}
				}
				fmt.Printf("j %d\n", j)
				char = int32(line[j])
				fmt.Printf("char %d\n", char)
				binary := fmt.Sprintf("%08b", char)
				fmt.Println(binary)
				booled := tobool(binary)
				for k := 0; k < len(binary); k++ {
					rowdata = append(rowdata, booled[k])
					if len(rowdata) == width {
						k = len(binary)
					}
				}
				fmt.Println(rowdata)

			}
			if len(rowdata) == width {
				data[i] = rowdata
				i++
				for l := 0; l < len(rowdata); l++ {
					rowdata = []bool{}
				}
			}
		}
	}
	//fmt.Println(data)
	return &PBM{data, width, height, magicnumber}, nil
}

func tobool(tab string) []bool {
	var rowData []bool
	for _, char := range tab {
		if char == '1' {
			rowData = append(rowData, true)
		} else if char == '0' {
			rowData = append(rowData, false)
		}
	}
	//fmt.Println(rowData)
	return rowData
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

// Flip flips the PBM image horizontally.
func (pbm *PBM) Flip() {
	for y := 0; y < pbm.height; y++ {
		for x := 0; x < pbm.width/2; x++ {
			pbm.data[y][x], pbm.data[y][pbm.width-x-1] = pbm.data[y][pbm.width-x-1], pbm.data[y][x]
		}
	}
}

// Flop flops the PBM image vertically.
func (pbm *PBM) Flop() {
	for y := 0; y < pbm.height/2; y++ {
		pbm.data[y], pbm.data[pbm.height-y-1] = pbm.data[pbm.height-y-1], pbm.data[y]
	}
}

// SetMagicNumber sets the magic number of the PBM image.
func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
}
