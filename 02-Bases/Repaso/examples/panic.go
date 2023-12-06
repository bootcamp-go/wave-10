package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := doOperation()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("operation completed successfully")
}

func doOperation() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error: %v", r)
			return
		}
	}()

	// file 1
	f1, err := os.Open("./examples/file1.txt")
	if err != nil {
		return
	}
	defer func() {
		fmt.Println("closing file 1")
		f1.Close()
	}()

	bytes, err := io.ReadAll(f1)
	if err != nil {
		return
	}

	// file 2
	f2, err := os.OpenFile("./examples/file2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer func() {
		fmt.Println("closing file 2")
		f2.Close()
	}()

	panic("something went wrong")

	defer func() {
		fmt.Println("closing file 3")
	}()

	// transfer data from file 1 to file 2
	_, err = f2.Write(bytes)
	if err != nil {
		return
	}
	return
}