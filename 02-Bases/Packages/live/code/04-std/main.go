package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// reader
	rd := strings.NewReader("messi")

	f, err := os.OpenFile("./code/04-std/test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// writer
	// io.Copy(os.Stdout, rd)
	io.Copy(f, rd)
}