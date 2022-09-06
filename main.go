package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	data := []string{"foo", "bar"}
	data = slices.Clip(data)

	fmt.Println("Hello world")
}
