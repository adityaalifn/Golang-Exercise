package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {
	links := []string{
		"http://a.c",
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://tokopedia.com",
		"http://bukalapak.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c{
		go func(link string) {
			time.Sleep(5 *time.Second)
			checkLink(link,c)
		}(l)
		//fmt.Println()
	}
}


func checkLink(link string, c chan string)  {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "might be down!")
		fmt.Println(link, "is down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
	//fmt.Println(link, "is up!")
}