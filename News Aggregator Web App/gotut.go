package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}


func main() {
	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)
	for _, Location := range s.Locations{
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		fmt.Println(n)
	}
}
