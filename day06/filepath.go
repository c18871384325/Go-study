package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs(".filepath.go") // 获取指定的绝对路径，. 为获取当前路径的绝对路径 D:\src\Go-study\day06
	fmt.Println(path)
	fmt.Println(filepath.Base(path))                       // 获取路径的基名，最后一个文件或文件夹名称
	fmt.Println(filepath.Dir(path))                        // 获取文件的目录名，最后一个文件或文件夹之前的路径
	fmt.Println(filepath.Ext(path))                        // 获取路径上最后一个带 . 后缀的内容  ，此处获取 .go
	fmt.Println(filepath.FromSlash("/123/436/////1/3//2")) // 替换斜线为当前系统的分隔符 \123\436\\\\\1\3\\2
	fmt.Println(filepath.ToSlash("\\3\\436\\///1/3//2"))   // 用 / 替换分隔符的结果， /3/436////1/3//2
	path2, _ := filepath.Abs(".filepath")
	fmt.Println(filepath.HasPrefix(path, path2)) // true path2指定的前缀是否包含在path内，包含则为true,否则为false
	path3 := "./filepath.go"
	fmt.Println(filepath.IsAbs(path3)) // false 判断路径是否为绝对路径，真则为true
	fmt.Println(filepath.IsAbs(path))  // true

	dir, _ := filepath.Abs("c://opt//cmdb")
	fmt.Println(filepath.Join(dir, "etc", "demo")) // Join 将给定的参数添加分隔符拼接为实际路径   c:\opt\cmdb\etc\demo

	fmt.Println(filepath.Split(path)) // 将文件路径与文件名分开，相当于单独打印Base,Dir  D:\src\Go-study\day06\ .filepath.go
	paths := "c://test/a;c://test/b;c://test/c"
	fmt.Println(filepath.SplitList(paths)) //  [c://test/a c://test/b c://test/c]  ， 将路径分隔符替换为空格，windows中path变量分隔符为；Linux中为：

	fmt.Println(filepath.Glob("./test/a*")) // 查找，通配符匹配， 指定路径下是否有 a 开头的文件 * 为通配符
	fmt.Println(filepath.Glob("./test/*.txt"))

	fmt.Println(filepath.Match("./test/a.go", "./test/a.go")) //true <nil>
	fmt.Println(filepath.Match("./test/a.*", "./test/a.go"))  // true <nil>
	fmt.Println(filepath.Match("./test/b.*", "./test/a.go"))  // name是否匹配shell模式，  ./test/b.*  为模式， ./test/a.go 为shell文件名

}
