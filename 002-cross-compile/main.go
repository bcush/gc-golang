package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("operating system: %s\narchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
