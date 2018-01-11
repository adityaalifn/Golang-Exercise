package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)

	bs := make([]byte, 999999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}