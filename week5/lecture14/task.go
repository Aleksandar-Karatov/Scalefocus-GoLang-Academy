package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"net/http"
)

func main() {
	var connections int64
	flag.Int64Var(&connections, "c", 2, "maximum number of concurrent connections ")
	flag.Parse()
	commands := os.Args
	Run(connections, commands)
}

func Run(connectionsNumber int64, commands []string) {
	ch := make(chan error, connectionsNumber)
	tempConnectionsNumber := connectionsNumber
	go func() {
		for i := 3; i < len(commands); i++ {
			if tempConnectionsNumber > 0 {
				ch <- pingURL(commands[i])
				time.Sleep(10 * time.Millisecond)
				tempConnectionsNumber--
			} else {
				fmt.Println("Buffer full. Exiting")
				close(ch)
				return
			}

		}
		close(ch)
	}()
	for err := range ch {
		if len(ch) == cap(ch) {
			close(ch)
			return
		}
		if err != nil {
			fmt.Println(err)
		}
	}
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
