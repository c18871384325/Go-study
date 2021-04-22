package main

import (
	"bufio"
	"fmt"
	"os"
	"strconver"
	"strings"
	"time"
	"errors"
)


type Task struct {
	Id        int64
	User      string
	Name      string
	Starttime *time.Time
	Endtime   *time.Time
}


func Parsestr(line string) (*Task, error) {
	node := strings.Split(line, ",")
	if len(node) < 5 {
		return nil, errors.New("字符串数量不正确")
	}

	Id := strconver.Atoi(node[0])

	User := node[1]

	Name := node[2]

	
}


func main() {

	file, err := os.Open("task.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	fmt.Printf("%T, %#v\n", scanner.Text(), scanner.Text())

}