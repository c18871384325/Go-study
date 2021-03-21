package main

import "fmt"

func main() {

	A := 12
	B := 23
	C := 34

	fmt.Println(A)
	goto END //  goto 出现时，会将代码的执行直接跳转指定的label代码段处并往后执行，此处label为 END， 代码执行会跳过B 的打印，直接从 END: 往后代码开始执行
	fmt.Println(B)
END:
	fmt.Println(C)

	// 使用 goto 打印 1+2+3+。。100

	var (
		i     = 0
		total = 0
	)

START:
	i += 1
	total += i
	if i == 100 {
		goto OVER
	}
	goto START // 代码执行到该行后，goto会将代码跳回到 START: 处，再次执行形成循环的结果，直到 i等于100 ，通过 goto OVER 将代码跳转到下面的代码开始执行

OVER:
	fmt.Println(total)
}
