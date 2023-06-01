package main

import (
	"blog_server/global"
	"blog_server/initialize"
	"fmt"
)

func main() {
	initialize.InitComnf()
	fmt.Println(global.Config)
}
