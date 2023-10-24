package main

import (
	"fmt"

	"github.com/brianlk/simple-go-service/dir1"
	_ "github.com/brianlk/simple-go-service/dir2"
)

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Subtract(1, 2))
	fmt.Println(dir1.SayHello())
}
