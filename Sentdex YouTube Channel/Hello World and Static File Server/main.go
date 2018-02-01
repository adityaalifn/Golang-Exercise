package main

import (
	"flag"
	"net/http"
	"log"
)

func main() {
	//var input string
	//fmt.Print("Enter Your Name: ")
	//fmt.Scanln(&input)
	//fmt.Println("Hello, %s!", input)

	port := flag.String("p", "8000", "port")
	dir := flag.String("d", ".", "dir")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s on Http port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}