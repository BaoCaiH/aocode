package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Read(filepath string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return bufio.NewScanner(file), file
}
