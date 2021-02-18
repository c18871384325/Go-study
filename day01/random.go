package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	// 种子，只需要设置一次
	rand.Seed(time.Now().Unix())
	// rand.Seed(1)
	
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int() % 101)  // 取随机数内101内的余数
	}

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))  // 取随机数内100内的余数
	}
}