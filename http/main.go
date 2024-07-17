package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatal(err)
	}
	// content := make([]byte, 999999)

	// resp.Body.Read(content)

	// fmt.Println(string(content))

	io.Copy(os.Stdout, resp.Body)
}
