package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	os.Stdout.Write(data)
}
