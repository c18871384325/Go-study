package main 

import (
	"os"
	//"time"
	"fmt"
	"bufio"
	"encoding/csv"
)


func main() {
	file, err := os.Open("csv.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	read := bufio.NewReader(file)

	CsvReader := csv.NewReader(read)

	var task [][]string
	task, err = CsvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", task)
}