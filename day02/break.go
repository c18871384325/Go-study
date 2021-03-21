// break 通过label,可以跳出多层循环

package main 

import "fmt"

func main() {

	// for j := 0; j < 3; j++ {
	// 	for i := 0; i < 5; i++ {
	// 		if i == 4 {
	// 			break     //  此时的break 仅跳出for i 层级的循环，for j 的循环还是会继续执行，如果需要跳出for j 层级的循环，可以通过 label的方式
	// 		}
	// 		fmt.Println(i)
	// 	}
	// }

	END:
	for j := 0; j < 3; j++ {
		for i := 0; i < 5; i++ {
			if i == 4 {
				break END    // break 后面加 label， 退出时，会退出与定义的label同级的循环，此处 END: 定义在 for j 层级，因此会直接退出for j 层级的循环
			}
			fmt.Println(i)
		}
	}

	// 注意： continue 也支持labels的方式实现跳出上层循环，用法与break相同

}