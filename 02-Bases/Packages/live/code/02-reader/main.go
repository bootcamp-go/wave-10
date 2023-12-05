package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// read file
	data, err := readFile("./code/02-reader/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func readFile(path string) (data string, err error) {
	// open file
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	// f := strings.NewReader("Hello, World!\n")

	// read file
	buf := make([]byte, 10)
	for {
		var n int
		n, err = f.Read(buf)
		if err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			return
		}

		fmt.Println(string(buf[:n]))
		data += string(buf[:n])
	}
}