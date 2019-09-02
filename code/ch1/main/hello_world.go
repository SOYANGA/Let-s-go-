package main //包，表明代码所在模块(包)

import (
	"fmt"
	"os"
) //引入代码依赖

//功能实现
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World",os.Args[1])
	}
	os.Exit(0)
}
