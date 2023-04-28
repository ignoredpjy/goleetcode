package sub

import (
	"fmt"
	"goleetcode/study/sub"
)

func Inner() {
	fmt.Println("test sub package public func")
}
func inner() {
	fmt.Println("test sub package private func")
	println(sub.OutVar)
}
