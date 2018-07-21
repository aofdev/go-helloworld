package main

import (
	"fmt"

	"github.com/aofdev/gotools"
	"github.com/aofdev/gotools/stringhelper"
)

func init() {
	fmt.Println("Init!!")
}

func main() {
	fmt.Println("Hello Word")
	fmt.Println(gotools.Add(10, 2))
	fmt.Println(stringhelper.Upper("Dog"))
}
