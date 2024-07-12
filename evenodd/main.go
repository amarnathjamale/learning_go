package main

import "fmt"

func main() {
	for _, number := range makeRange(0, 10) {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
