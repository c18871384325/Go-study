package main

import "fmt"

func main() {

	// 冒泡排序
	// 变量交换方法

	// 第一种
	// a := 1
	// b := 2

	// // 将a和b的值互换
	// tmp := a
	// a = b
	// b = tmp 
	// fmt.Println(a, b)


	// 其他方式
	// a, b := 1, 2
	// a, b = b, a
	// fmt.Println(a, b)

	// 游戏规则，有5个学生排队，要求按身高顺序，最高的学生排在最后，依次类推进行排序
	// 1.先定义切片，将学生身高写入切片
	height := []int{156, 177, 168, 170, 182}   //  5个值分别对应身高

	// 从第一个开始往后比较，当左侧大于右侧时，进行位置交换
	/*
		比较示例
		A: 156
		B: 177
		C: 168
		D: 170
		E: 182

		第一次： A, B   比较后排序结果  ABCDE
		第二次： B，C   比较后排序结果  ACBDE
		第三次： C, D   比较厚排序结果  ACDBE
		。。。。 
		通过四次比较即可得到身高最高的将排序到最后
		按照该种方式实现排序
	*/
	for j := 0 ; j < len(height) -1; j++ {
		for i := 0; i < len(height) -1; i++ {
			if height[i] > height[i+1] {
				height[i], height[i+1] = height[i+1], height[i]
			}
			fmt.Println(i, height)
		}
	}
	
	fmt.Println(height)
	
}