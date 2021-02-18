package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var random = rand.Intn(10)
	
	var guess int

	for i := 1; i <= 6; i++ {
		var k int

		if i == 6 {
			fmt.Print("太笨了，是否继续：")
			var l string
			fmt.Scan(&l)
			switch l {
			case "yes", "y", "Y":
				i = 1
			default:
				k = 1
			}
			continue
		}

		fmt.Print("请输入随机数：")
		fmt.Scan(&guess)
		
		if guess == random {
			fmt.Printf("经过%d次对了，太聪明了\n", i)
			var j string
			fmt.Print("是否继续游戏：")
			fmt.Scan(&j)
			switch j {
			case "yes", "y", "Y":
				i = 1
				continue
			default:
				k = 1
			}
		} else if guess > random {
			fmt.Println("输入随机数太大，请重新输入")
		} else if guess < random {
			fmt.Println("输入随机数太小，请重新输入")
		} 

		if k == 1 {
			break
		}
	}

}