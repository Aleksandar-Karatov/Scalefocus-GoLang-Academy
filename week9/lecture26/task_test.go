package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"week9Lecture26Task/week9/lecture26/dataStructs"
	"week9Lecture26Task/week9/lecture26/repository"
)

func TestHandleTop(t *testing.T) {
	var allStories dataStructs.Stories

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
			var storyToSend dataStructs.Story = dataStructs.Story{}
			switch r.URL.Path {
			case "/1":
				storyToSend = dataStructs.Story{Title: "title1", Score: 1}
			case "/2":
				storyToSend = dataStructs.Story{Title: "title2", Score: 2}
			case "/3":
				storyToSend = dataStructs.Story{Title: "title3", Score: 3}
			case "/4":
				storyToSend = dataStructs.Story{Title: "title4", Score: 4}
			case "/5":
				storyToSend = dataStructs.Story{Title: "title5", Score: 5}
			case "/6":
				storyToSend = dataStructs.Story{Title: "title6", Score: 6}
			case "/7":
				storyToSend = dataStructs.Story{Title: "title7", Score: 7}
			case "/8":
				storyToSend = dataStructs.Story{Title: "title8", Score: 8}
			case "/9":
				storyToSend = dataStructs.Story{Title: "title9", Score: 9}
			case "/10":
				storyToSend = dataStructs.Story{Title: "title10", Score: 10}
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
	db, err := sql.Open("sqlite", "DBL26_TEST.db")
	if err != nil {
		log.Fatal(err)
	}
	var dbc repository.DBcontext
	dbc.DB = db

	mux := http.NewServeMux()
	mux.Handle("/api/top", HandleTop(dbc, mockServer.URL))
	resp, err := http.Get("http://localhost:9000/api/top")
	if err != nil {
		t.Log(err)
	}
	var toCheck dataStructs.Stories
	json.NewDecoder(resp.Body).Decode(&toCheck)
	fmt.Println(toCheck)
	fmt.Println(allStories)
	reflect.DeepEqual(toCheck, allStories)
	dbc.DB.Close()
}
