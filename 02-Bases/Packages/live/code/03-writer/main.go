package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./code/03-writer/test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// write file
	data := []byte("Hello, World!\n")
	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}