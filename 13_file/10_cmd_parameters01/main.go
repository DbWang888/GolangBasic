package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数有：", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}

}
