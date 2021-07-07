package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ReadD()
}
func ReadD() {
	r := strings.NewReader("Reading....")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		fmt.Println("Error")
	}
	_, err = io.CopyBuffer(os.Stdout, strings.NewReader("Akshat Sahijpal"), make([]byte, 8))
	if err != nil {
		return
	}
}
