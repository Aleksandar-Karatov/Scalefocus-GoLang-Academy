package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/api/top", HandleTop("https://hacker-news.firebaseio.com/v0/"))

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

func HandleTop(dest string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		if r.Method != http.MethodGet {
			http.Error(w, "Only methods are allowed", http.StatusBadRequest)
			return
		}
		log.Println("HandleTop called")
		var res *http.Response
		var err error
		if dest == "https://hacker-news.firebaseio.com/v0/" {
			res, err = http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
		} else {
			res, err = http.Get(dest + "/IDs")
		}

		if err != nil {
			panic(err)
		}
		var payload IDPayload
		json.NewDecoder(res.Body).Decode(&payload.IDs)
		var URLs [10]string
		for i := 0; i < 10; i++ {
			if dest == "https://hacker-news.firebaseio.com/v0/" {
				URLs[i] = "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(payload.IDs[i]) + ".json?print=pretty"
			} else {
				URLs[i] = dest + "/" + strconv.Itoa(payload.IDs[i])
			}

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
