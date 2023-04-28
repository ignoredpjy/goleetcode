package main

import (
	/*
		父子包同名是可以的
	*/
	subOut "goleetcode/study/sub"
	subInner "goleetcode/study/sub/inner"
)

func main() {
	subOut.Out()
	subInner.Inner()
}
