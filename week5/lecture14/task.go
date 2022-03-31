package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s!\n", name)
}
func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}
