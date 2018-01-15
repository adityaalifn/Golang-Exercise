package main

import (
	"net/http"
	"fmt"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://bukalapak.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string)  {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}