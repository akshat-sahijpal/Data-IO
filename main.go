package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("run.txt", os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("Error Loading File!!")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error Closing the File!!")
		}
	}(file)
	// Write the data
	buffer := []byte("Buffered data 1")
	_, err1 := file.Write(buffer)
	if err1 != nil {
		return
	} else {
		fmt.Println("Done!!!")
	}
}
