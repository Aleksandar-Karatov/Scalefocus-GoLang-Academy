package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandleTop(t *testing.T) {
	var allStories Stories

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("Server was called!")
		t.Log(r.URL.Path)
		if r.URL.Path == "/IDs" && r.Method == "GET" {
			IDsToSend := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			IDsJson, err := json.Marshal(IDsToSend)
			if err != nil {
				t.Log(err)
			}
			httptest.NewRequest("POST", "http://localhost:9000/api/top", bytes.NewBuffer(IDsJson))
		} else {
			var storyToSend Story = Story{}
			switch r.URL.Path {
			case "/1":
				storyToSend = Story{Title: "title1", Score: 1}
			case "/2":
				storyToSend = Story{Title: "title2", Score: 2}
			case "/3":
				storyToSend = Story{Title: "title13", Score: 13}
			case "/4":
				storyToSend = Story{Title: "title4", Score: 4}
			case "/5":
				storyToSend = Story{Title: "title5", Score: 5}
			case "/6":
				storyToSend = Story{Title: "title6", Score: 6}
			case "/7":
				storyToSend = Story{Title: "title7", Score: 7}
			case "/8":
				storyToSend = Story{Title: "title8", Score: 8}
			case "/9":
				storyToSend = Story{Title: "title9", Score: 9}
			case "/10":
				storyToSend = Story{Title: "title10", Score: 10}
			}
			if storyToSend.Title != "" {
				allStories.AllStories = append(allStories.AllStories, storyToSend)
				storyJson, err := json.Marshal(storyToSend)
				if err != nil {
					t.Log(err)
				}
				httptest.NewRequest("POST", "http://localhost:9000/api/top", bytes.NewBuffer(storyJson))
			}
		}
	}))

	mux := http.NewServeMux()
	mux.Handle("/api/top", HandleTop(mockServer.URL))
	resp, err := http.Get("http://localhost:9000/api/top")
	if err != nil {
		t.Log(err)
	}
	var toCheck Stories
	json.NewDecoder(resp.Body).Decode(&toCheck)
	reflect.DeepEqual(toCheck, allStories)

}
