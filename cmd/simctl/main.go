package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println(" simctl [command]")
		return
	}

	fmt.Println("simctl")
}
