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
		"http://tokopedia.com",
		"http://bukalapak.com",
	}

	c := make(chan string, len(links))

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}


func checkLink(link string, c chan string)  {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "might be down"
		return
	}
	//fmt.Println(link, "is up!")
	c <- link + " is up"
}