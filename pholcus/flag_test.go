package pholcus

import (
	"flag"
	"testing"
	"fmt"
)

func TestFlag(t *testing.T) {
	flag.String("a *********************************************** common *********************************************** -a", "", "")
	// 操作界面
	uiDefault := "gui"
	var uiflag *string
	uiflag = flag.String("_ui", uiDefault, "   <选择操作界面> [web] [gui] [cmd]")
	fmt.Println(*uiflag)
}
