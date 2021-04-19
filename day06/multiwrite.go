package main

import (
	//"fmt"
	"io"
	"os"
)

func main() {
	file1, err := os.Create("./test/w1.go")
	file2, err := os.Create("./test/w2.go")
	if err != nil {
		return
	}

	filewrite := io.MultiWriter(file1, file2, os.Stdout)

	filewrite.Write([]byte("this is demo"))

	file1.Close()
	file2.Close()
}