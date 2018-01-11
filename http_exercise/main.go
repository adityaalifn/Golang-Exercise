package main

import (
	"os"
	"fmt"
	"net/http"
	"bufio"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	fmt.Println(bs, resp)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}