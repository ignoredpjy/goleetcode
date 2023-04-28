/*
1.子包、包嵌套，不能越过私有规定
2.子包甚至可以与父包同名
*/
package sub

import (
	"fmt"
	"goleetcode/study/sub/inner"
)

var OutVar int

var outVar int

func Out() {
	fmt.Println("test sub package public func")
}
func out() {
	fmt.Println("test sub package private func")
	sub.Inner()
}
