package main

import "fmt"

type lang interface {
	getGreetings() string
}

type english struct{}
type spanish struct{}

func main() {
	eng := english{}
	esp := spanish{}
	printGreetings(eng)
	printGreetings(esp)

}

func printGreetings(l lang) {
	fmt.Println(l.getGreetings())
}

func (english) getGreetings() string {
	return "Hello"
}

func (spanish) getGreetings() string {
	return "Hola"
}
