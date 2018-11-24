package main

import (
	"fmt"
	"github.com/bassaer/randgen"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(randgen.RandNum(1000000000, 9999999999))
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Println(randgen.RandStr(10))
	}
}
