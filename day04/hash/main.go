package main

import (
	"fmt"
	"crypto/md5"
)


func main() {
	mirror := md5.New([]byte(This is 徐青 and 程哲))
	fmt.Println(mirror)
}