package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/top", HandleTop())
	log.Fatal(http.ListenAndServe(":9000", mux))
}

type NewsID = int

type IDPayload struct {
	IDs []NewsID
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}
type Stories struct {
	AllStories []Story `json:"top_stories"`
}

type SimpleHandler struct{}

func (sh *SimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Only methods are allowed", http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}
func HandleTop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		if r.Method != http.MethodGet {
			http.Error(w, "Only methods are allowed", http.StatusBadRequest)
			return
		}
		log.Println("HandleTop called")
		res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")

		if err != nil {
			panic(err)
		}
		var payload IDPayload
		json.NewDecoder(res.Body).Decode(&payload.IDs)
		var URLs [10]string
		for i := 0; i < 10; i++ {
			URLs[i] = "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(payload.IDs[i]) + ".json?print=pretty"
		}
		var topStories Stories
		for i := 0; i < 10; i++ {
			res, err := http.Get(URLs[i])
			if err != nil {
				panic(err)
			}
			topStories.AllStories = append(topStories.AllStories, Story{})
			json.NewDecoder(res.Body).Decode(&topStories.AllStories[i])
		}

		storiesJSON, err := json.Marshal(topStories)
		if err != nil {
			panic(err)
		}
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, storiesJSON, "", "    "); err != nil {
			panic(err)
		}
		w.Write(prettyJSON.Bytes())

	}

}
