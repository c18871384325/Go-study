package main

import "fmt"

func main() {

	// 切片内元素的移除
	// 移除切片内的第一个或最后一个元素
	nums := []int{1, 2, 3, 4, 5}

	// start 为0  ， start可以省略
	// end = len(nums)  ， end可以省略
	nums = nums[1:]
	fmt.Println(nums)   // [2 3 4 5]


	// 移除切片的最后一个元素  
	nums = nums[:len(nums)-1]
	fmt.Println(nums)   // [2 3 4]


	// 移除切片中间的某个元素
	// 示例移除上面 nums的第二个元素，也就是值是 3 的元素
	copy(nums[1:], nums[2:])   
	nums = nums[:len(nums)-1]
	fmt.Println(nums)  // [2 4] 
}