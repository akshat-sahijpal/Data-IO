package main

import (
	"fmt"
	"io"
	"strings"
)

type AlphaReader struct {
	src io.Reader
}

func NewAlphaReader(src io.Reader) *AlphaReader {
	return &AlphaReader{src: src}
}
func (reader *AlphaReader) Read(p *[]byte) (s0 int, err error) {
	fmt.Println(*reader)
	for i := 0; i < len(*p); i++ {
		if ((*reader)[i] >= 'A' && (*reader)[i] <= 'Z') || ((*reader)[i] >= 'a' && (*reader)[i] <= 'z') {
			(*p)[i] = (*reader)[i]
			if (*p)[i+1] == 0 {
				return 0, io.EOF
			}
		}
		s0++
	}

	return s0, nil
}

func main() {}
