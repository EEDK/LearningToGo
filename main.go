package main

import (
	"fmt"
)

func stringSwap(x , y string) (string, string){
	return y, x
}

func main() {
	a, b := stringSwap("hello", "world")
	fmt.Println(a, b)
}

