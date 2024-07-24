package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://dnbstatus.no",
		"http://github.com",
		"https://vg.no",
		"https://pkg.go.dev",
	}

	c := make(chan string)
	for _, link := range links {
		go isUp(link, c)
	}

	for l := range c {
		go (func(link string) {
			time.Sleep(5 * time.Second)
			isUp(link, c)
		})(l)
	}
}

func isUp(l string, c chan string) {
	u, urlerr := url.Parse(l)

	if urlerr != nil {
		fmt.Println("Error:", urlerr)
	}

	_, err := http.Get(l)
	if err != nil {
		fmt.Println(u.Host, "might be down")
		c <- l
		return
	}
	fmt.Println(u.Host, "is up")
	c <- l
}
