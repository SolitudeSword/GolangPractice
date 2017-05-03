package main

import (
	"flag"
	"fmt"
)

func main()  {
	flag.String("a *********************************************** common *********************************************** -a", "", "")
	// 操作界面
	uiDefault := "cmd"
	var uiflag *string
	uiflag = flag.String("_ui", uiDefault, "   <选择操作界面> [web] [gui] [cmd]")
	fmt.Println(*uiflag, flag.CommandLine)
}
