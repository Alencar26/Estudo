package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)   //linux
	fmt.Println(runtime.GOARCH) //amd64
}
