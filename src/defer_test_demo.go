package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("pre print") }()
	defer func() { fmt.Println("printing") }()
	defer func() { fmt.Println("post print") }()

	panic("trap execption!!!")
}
