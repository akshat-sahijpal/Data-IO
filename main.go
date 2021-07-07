package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fileReader()
	stringReader()
}
func stringReader() {
	reader := strings.NewReader("Akshat Sahijpal")
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err == io.EOF {
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}
func fileReader() {
	file, error := os.Open("run.txt") // Open Returns pointer to the file
	if error != nil {
		panic(error)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error Closing the file ", err)
		}
	}(file)

	// Data is stored in a buffer
	buffer := make([]byte, 10) // buffer of size 10
	// func (f *File) Read(b []byte) (n int, err error)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buffer[:n]))
	}
}
