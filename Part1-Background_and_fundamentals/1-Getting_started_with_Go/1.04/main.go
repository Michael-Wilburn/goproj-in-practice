package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal("could not retrieve example.com\n", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("could not read body", err)
	}
	log.Println(string(body))
}
