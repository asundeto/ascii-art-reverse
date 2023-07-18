package main

import (
	ascart "ascart/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}
	fileName := os.Args[1]
	fileName = ascart.CheckFileName(fileName)
	if fileName == "" {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		return
	}
	ascart.Reverse(fileName)
}
