package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
