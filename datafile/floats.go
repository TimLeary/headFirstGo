// Package datafile allows reading data samples from files
package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

//GetFloats reads a float64 from each line of a file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	var act float64

	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		act, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, act)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
