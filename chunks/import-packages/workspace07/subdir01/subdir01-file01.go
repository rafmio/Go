package subdir01

import "fmt"

var Sd01VarStr string = "SdVarStr: Hey!"

func Sd01Func() {
	fmt.Println("from subdir01, SdFunc():", Sd01VarStr)
}
