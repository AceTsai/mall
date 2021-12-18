package main

import (
	"fmt"
	"mall.com/global"
	"mall.com/initialize"
)

func main() {
	initialize.MainPath()
	fmt.Println(global.MainPath)
}
