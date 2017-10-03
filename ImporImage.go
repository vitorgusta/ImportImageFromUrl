package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "http://i.imgur.com/m1UIjW1.jpg"
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("/Users/vitorlevy/Documents/asdf.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Success!", file)
}
