package subdir02

import (
	"fmt"
	"ltr/subdir01"
)

var StrToPassToFunc string = subdir01.Sd01VarStr + " " + "Hobana!"

func Sd02Func(s string) {
	fmt.Println("form subdir02, Sd02Func(): ", s)
}
