package main

import (
	"fmt"

	"github.com/brianlk/simple-go-service/dir1"
	"github.com/brianlk/simple-go-service/dir2"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(dir1.SayHello())
	fmt.Println(dir2.Test1234())
}
